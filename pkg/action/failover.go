package action

import (
	"fmt"
	storkv1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	"github.com/libopenstorage/stork/pkg/log"
	migration "github.com/libopenstorage/stork/pkg/migration/controllers"
	"github.com/libopenstorage/stork/pkg/utils"
	storkops "github.com/portworx/sched-ops/k8s/stork"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

// validateBeforeFailover is called as part of the Initial stage of Failover
func (ac *ActionController) validateBeforeFailover(action *storkv1.Action) {
	if action.Status.Status == storkv1.ActionStatusSuccessful {
		action.Status.Stage = storkv1.ActionStageScaleDownSource
		action.Status.Status = storkv1.ActionStatusInitial
		action.Status.FinishTimestamp = metav1.Now()
		ac.updateAction(action)
		return
	} else if action.Status.Status == storkv1.ActionStatusFailed {
		action.Status.Stage = storkv1.ActionStageFinal
		action.Status.FinishTimestamp = metav1.Now()
		ac.updateAction(action)
		return
	}

	if action.Status.Status == storkv1.ActionStatusInitial {
		action.Status.Status = storkv1.ActionStatusScheduled
		ac.updateAction(action)
	}

	// first get the migrationSchedule's static copy in the destination cluster
	migrationSchedule, err := storkops.Instance().GetMigrationSchedule(action.Spec.ActionParameter.FailoverParameter.MigrationScheduleReference, action.Namespace)
	if err != nil {
		msg := fmt.Sprintf("Error fetching the MigrationSchedule %s/%s for Failover", action.Namespace, action.Spec.ActionParameter.FailoverParameter.MigrationScheduleReference)
		logEvents := ac.printFunc(action, string(storkv1.ActionStatusFailed))
		logEvents(msg, "err")
		action.Status.Status = storkv1.ActionStatusFailed
		action.Status.Reason = msg
		ac.updateAction(action)
		return
	}

	// ensure that the provided migrationSchedule is a static copy
	if migrationSchedule.GetAnnotations() == nil || migrationSchedule.GetAnnotations()[migration.StorkMigrationScheduleCopied] != "true" {
		msg := fmt.Sprintf("The provided MigrationSchedule %s/%s is not created by Stork. Please ensure that the failover is initiated in the cluster you want to failover to", migrationSchedule.Namespace, migrationSchedule.Name)
		logEvents := ac.printFunc(action, string(storkv1.ActionStatusFailed))
		logEvents(msg, "err")
		action.Status.Status = storkv1.ActionStatusFailed
		action.Status.Reason = msg
		ac.updateAction(action)
		return
	}

	drMode, err := ac.getDRMode(migrationSchedule.Spec.Template.Spec.ClusterPair, action.Namespace)
	if err != nil {
		msg := "Failed to determine if the DR plan's mode is async-dr or sync-dr"
		logEvents := ac.printFunc(action, string(storkv1.ActionStatusFailed))
		logEvents(msg, "err")
		action.Status.Status = storkv1.ActionStatusFailed
		action.Status.Reason = msg
		ac.updateAction(action)
		return
	}

	if drMode == asyncDR && *action.Spec.ActionParameter.FailoverParameter.SkipSourceOperations {
		// we always honour skip-source-operations flag in async-dr mode irrespective of the clusterPair presence in the destination cluster
		// while for sync-dr in case clusterpair is present in destination cluster, we will try to ScaleDownSource in case source cluster is accessible
		msg := "Skipping source cluster operations and moving to the Failover ScaleUpDestinationStage since SkipSourceOperations is set to true"
		ac.updateActionToSkipSourceOperations(action, msg)
		return
	}

	// get sourceConfig from clusterPair in the destination cluster
	sourceConfig, err := getClusterPairSchedulerConfig(migrationSchedule.Spec.Template.Spec.ClusterPair, migrationSchedule.Namespace)
	if err != nil {
		if *action.Spec.ActionParameter.FailoverParameter.SkipSourceOperations {
			// if clusterPair is absent in destination cluster and skipSourceOperations is provided always move to ScaleUpDestination
			msg := "Skipping source cluster operations and moving to the Failover ScaleUpDestinationStage since SkipSourceOperations is set to true"
			ac.updateActionToSkipSourceOperations(action, msg)
			return
		}
		msg := fmt.Sprintf("Error fetching the ClusterPair %s/%s for Failover", migrationSchedule.Namespace, migrationSchedule.Spec.Template.Spec.ClusterPair)
		logEvents := ac.printFunc(action, string(storkv1.ActionStatusFailed))
		logEvents(msg, "err")
		action.Status.Status = storkv1.ActionStatusFailed
		action.Status.Reason = msg
		ac.updateAction(action)
		return
	}
	// check accessibility of the source cluster, get the migrationSchedule in the source cluster to check for any ongoing migrations and to suspend the migrationSchedule
	if !ac.isClusterAccessible(action, sourceConfig) {
		msg := "Unable to access the remote cluster. Directly moving to the Failover ScaleUpDestinationStage"
		ac.updateActionToSkipSourceOperations(action, msg)
		return
	}

	remoteOps, err := storkops.NewForConfig(sourceConfig)
	if err != nil {
		msg := fmt.Sprintf("Failed to create the stork client to access cluster with the config Host : %v", sourceConfig.Host)
		logEvents := ac.printFunc(action, string(storkv1.ActionStatusFailed))
		logEvents(msg, "err")
		action.Status.Status = storkv1.ActionStatusFailed
		action.Status.Reason = msg
		ac.updateAction(action)
		return
	}
	srcMigrSched, err := remoteOps.GetMigrationSchedule(action.Spec.ActionParameter.FailoverParameter.MigrationScheduleReference, action.Namespace)
	if err != nil {
		msg := fmt.Sprintf("Error fetching the MigrationSchedule %s/%s for Failover in the remote cluster : %v", migrationSchedule.Namespace, migrationSchedule.Name, err)
		logEvents := ac.printFunc(action, string(storkv1.ActionStatusFailed))
		logEvents(msg, "err")
		action.Status.Status = storkv1.ActionStatusFailed
		action.Status.Reason = msg
		ac.updateAction(action)
		return
	}

	_, latestMigration := getLatestMigrationPolicyAndStatus(*srcMigrSched)
	// Suspend the migration schedule in the source cluster
	suspend := true
	if !*srcMigrSched.Spec.Suspend {
		srcMigrSched.Spec.Suspend = &suspend
		msg := fmt.Sprintf("Suspending the MigrationSchedule %s/%s before proceeding with the Failover operation", srcMigrSched.Namespace, srcMigrSched.Name)
		log.ActionLog(action).Infof(msg)
		_, err := remoteOps.UpdateMigrationSchedule(srcMigrSched)
		if err != nil {
			log.ActionLog(action).Errorf("Error suspending the MigrationSchedule %s/%s: %v", srcMigrSched.Namespace, srcMigrSched.Name, err)
			return
		}
	}
	if latestMigration != nil {
		if !isMigrationComplete(latestMigration.Status) {
			// a migration is still in progress. failover needs to wait
			msg := fmt.Sprintf("Waiting for completion of the migration %s/%s", srcMigrSched.Namespace, latestMigration.Name)
			log.ActionLog(action).Infof(msg)
			ac.recorder.Event(action,
				v1.EventTypeWarning,
				string(storkv1.ActionStatusScheduled),
				msg)
			ac.updateAction(action)
			return
		} else if latestMigration.Status == storkv1.MigrationStatusFailed {
			// log a warning event that last migration was unsuccessful. We will still proceed with the failover
			msg := fmt.Sprintf("The latest migration %s/%s status is %s. Proceeding with the Failover", srcMigrSched.Namespace, latestMigration.Name, latestMigration.Status)
			log.ActionLog(action).Infof(msg)
			ac.recorder.Event(action,
				v1.EventTypeWarning,
				string(storkv1.ActionStatusScheduled),
				msg)
			ac.updateAction(action)
		}
	}

	msg := fmt.Sprintf("Failover ActionStageInitial status %s", string(storkv1.ActionStatusSuccessful))
	log.ActionLog(action).Infof(msg)
	action.Status.Status = storkv1.ActionStatusSuccessful
	action.Status.Reason = ""
	ac.updateAction(action)
}

func (ac *ActionController) deactivateSourceDuringFailover(action *storkv1.Action) {
	if action.Status.Status == storkv1.ActionStatusSuccessful {
		action.Status.Stage = storkv1.ActionStageWaitAfterScaleDown
		action.Status.Status = storkv1.ActionStatusInitial
		action.Status.FinishTimestamp = metav1.Now()
		ac.updateAction(action)
		return
	} else if action.Status.Status == storkv1.ActionStatusFailed {
		action.Status.Stage = storkv1.ActionStageFinal
		action.Status.FinishTimestamp = metav1.Now()
		ac.updateAction(action)
		return
	}

	if action.Status.Status == storkv1.ActionStatusInitial {
		action.Status.Status = storkv1.ActionStatusInProgress
		ac.updateAction(action)
	}

	migrationScheduleName := action.Spec.ActionParameter.FailoverParameter.MigrationScheduleReference
	namespaces := action.Spec.ActionParameter.FailoverParameter.FailoverNamespaces
	migrationSchedule, err := storkops.Instance().GetMigrationSchedule(migrationScheduleName, action.Namespace)
	if err != nil {
		msg := fmt.Sprintf("Error fetching the MigrationSchedule %s/%s", action.Namespace, migrationScheduleName)
		logEvents := ac.printFunc(action, string(storkv1.ActionStatusFailed))
		logEvents(msg, "err")
		action.Status.Status = storkv1.ActionStatusFailed
		action.Status.Reason = msg
		ac.updateAction(action)
		return
	}
	migrationNamespaces, err := utils.GetMergedNamespacesWithLabelSelector(migrationSchedule.Spec.Template.Spec.Namespaces, migrationSchedule.Spec.Template.Spec.NamespaceSelectors)
	if err != nil {
		msg := fmt.Sprintf("Failed to fetch list of namespaces from the MigrationSchedule %s/%s", migrationSchedule.Namespace, migrationSchedule.Name)
		logEvents := ac.printFunc(action, string(storkv1.ActionStatusFailed))
		logEvents(msg, "err")
		action.Status.Status = storkv1.ActionStatusFailed
		action.Status.Reason = msg
		ac.updateAction(action)
		return
	}

	// only consider namespaces which are a part of both namespaces and migrationNamespaces
	// this means if there are some invalid namespaces provided for failover/failback we will ignore them
	_, actualNamespaces, _ := utils.IsSubList(namespaces, migrationNamespaces)

	// get sourceConfig from clusterPair in destination cluster
	remoteConfig, err := getClusterPairSchedulerConfig(migrationSchedule.Spec.Template.Spec.ClusterPair, migrationSchedule.Namespace)
	if err != nil {
		msg := fmt.Sprintf("Error fetching the ClusterPair %s/%s", migrationSchedule.Namespace, migrationSchedule.Spec.Template.Spec.ClusterPair)
		logEvents := ac.printFunc(action, string(storkv1.ActionStatusFailed))
		logEvents(msg, "err")
		action.Status.Status = storkv1.ActionStatusFailed
		action.Status.Reason = msg
		ac.updateAction(action)
		return
	}
	// get destination i.e. current cluster's config
	clusterConfig := ac.config
	ac.deactivateClusterDuringDR(action, actualNamespaces, migrationNamespaces, clusterConfig, remoteConfig)
}

// activateClusterDuringFailover is used for both activation of destination and reactivation of source during failover.
// If rollback true -> reactivate source, else activate destination
func (ac *ActionController) activateClusterDuringFailover(action *storkv1.Action, rollback bool) {
	// Always Move to Final stage whether this stage succeeds or fails
	if action.Status.Status == storkv1.ActionStatusSuccessful || action.Status.Status == storkv1.ActionStatusFailed {
		action.Status.Stage = storkv1.ActionStageFinal
		if rollback {
			// Irrespective of stage status, mark the failover operation failed
			action.Status.Status = storkv1.ActionStatusFailed
		}
		action.Status.FinishTimestamp = metav1.Now()
		ac.updateAction(action)
		return
	}

	if action.Status.Status == storkv1.ActionStatusInitial {
		action.Status.Status = storkv1.ActionStatusInProgress
		ac.updateAction(action)
	}

	var config *rest.Config
	namespaces := action.Spec.ActionParameter.FailoverParameter.FailoverNamespaces

	migrationSchedule, err := storkops.Instance().GetMigrationSchedule(action.Spec.ActionParameter.FailoverParameter.MigrationScheduleReference, action.Namespace)
	if err != nil {
		msg := fmt.Sprintf("Error fetching the MigrationSchedule %s/%s", action.Namespace, action.Spec.ActionParameter.FailoverParameter.MigrationScheduleReference)
		logEvents := ac.printFunc(action, string(storkv1.ActionStatusFailed))
		logEvents(msg, "err")
		action.Status.Status = storkv1.ActionStatusFailed
		action.Status.Reason = msg
		ac.updateAction(action)
		return
	}
	if rollback {
		// In case of reactivation/rollback you have to activate the apps which were scaled down in the source cluster i.e. remote cluster
		config, err = getClusterPairSchedulerConfig(migrationSchedule.Spec.Template.Spec.ClusterPair, migrationSchedule.Namespace)
		if err != nil {
			msg := fmt.Sprintf("Failed to get the remote config from the ClusterPair %s", migrationSchedule.Spec.Template.Spec.ClusterPair)
			logEvents := ac.printFunc(action, string(storkv1.ActionStatusFailed))
			logEvents(msg, "err")
			action.Status.Status = storkv1.ActionStatusFailed
			action.Status.Reason = msg
			ac.updateAction(action)
			return
		}
	} else {
		config = ac.config
	}
	ac.activateClusterDuringDR(action, namespaces, migrationSchedule, config, rollback)
}

func (ac *ActionController) performLastMileMigrationDuringFailover(action *storkv1.Action) {
	if action.Status.Status == storkv1.ActionStatusSuccessful {
		action.Status.Stage = storkv1.ActionStageScaleUpDestination
		action.Status.Status = storkv1.ActionStatusInitial
		action.Status.FinishTimestamp = metav1.Now()
		ac.updateAction(action)
		return
	} else if action.Status.Status == storkv1.ActionStatusFailed {
		// move to reactivate/rollback stage
		action.Status.Stage = storkv1.ActionStageScaleUpSource
		action.Status.Status = storkv1.ActionStatusInitial
		action.Status.FinishTimestamp = metav1.Now()
		ac.updateAction(action)
		return
	}

	if action.Status.Status == storkv1.ActionStatusInitial {
		action.Status.Status = storkv1.ActionStatusInProgress
		ac.updateAction(action)
	}

	migrationScheduleReference := action.Spec.ActionParameter.FailoverParameter.MigrationScheduleReference

	migrationSchedule, err := storkops.Instance().GetMigrationSchedule(migrationScheduleReference, action.Namespace)
	if err != nil {
		msg := fmt.Sprintf("Error fetching the MigrationSchedule %s/%s", action.Namespace, migrationScheduleReference)
		log.ActionLog(action).Errorf(msg)
		action.Status.Status = storkv1.ActionStatusFailed
		action.Status.Reason = msg
		ac.recorder.Event(action,
			v1.EventTypeWarning,
			string(storkv1.ActionStatusFailed),
			msg)
		ac.updateAction(action)
		return
	}
	// In failover last-mile-migration is created in source->destination direction
	config, err := getClusterPairSchedulerConfig(migrationSchedule.Spec.Template.Spec.ClusterPair, migrationSchedule.Namespace)
	if err != nil {
		msg := fmt.Sprintf("Error fetching the ClusterPair %s/%s", migrationSchedule.Namespace, migrationSchedule.Spec.Template.Spec.ClusterPair)
		log.ActionLog(action).Errorf(msg)
		action.Status.Status = storkv1.ActionStatusFailed
		action.Status.Reason = msg
		ac.recorder.Event(action,
			v1.EventTypeWarning,
			string(storkv1.ActionStatusFailed),
			msg)
		ac.updateAction(action)
		return
	}

	ac.createLastMileMigration(action, config, migrationSchedule)
	// once last mile migration is successful, we should deactivate source cluster-domain if sync-dr
	// I have placed this here since if deactivation of cluster domain fails, and we fail the failover action
	// we should roll back the scale down of source cluster
	if action.Status.Status == storkv1.ActionStatusSuccessful {
		drMode, err := ac.getDRMode(migrationSchedule.Spec.Template.Spec.ClusterPair, action.Namespace)
		if err != nil {
			msg := fmt.Sprintf("Failed to determine the mode of the DR plan: %v", err)
			logEvents := ac.printFunc(action, string(storkv1.ActionStatusFailed))
			logEvents(msg, "err")
			action.Status.Status = storkv1.ActionStatusFailed
			action.Status.Reason = msg
			ac.updateAction(action)
			return
		}
		if drMode == syncDR {
			// deactivate the source cluster domain
			err := ac.remoteClusterDomainUpdate(false, action)
			if err != nil {
				msg := fmt.Sprintf("Failed to deactivate the remote cluster domain: %v", err)
				logEvents := ac.printFunc(action, "DeactivateClusterDomain")
				logEvents(msg, "err")
				action.Status.Status = storkv1.ActionStatusFailed
				action.Status.Reason = msg
				ac.updateAction(action)
			}
		}
	}
}

// updateActionToSkipSourceOperations is used to update the action and log events when we directly move to ScaleUpDestinationStage
func (ac *ActionController) updateActionToSkipSourceOperations(action *storkv1.Action, msg string) {
	logEvents := ac.printFunc(action, string(storkv1.ActionStatusSuccessful))
	logEvents(msg, "out")
	action.Status.FinishTimestamp = metav1.Now()
	action.Status.Stage = storkv1.ActionStageScaleUpDestination
	action.Status.Status = storkv1.ActionStatusInitial
	ac.updateAction(action)
}
