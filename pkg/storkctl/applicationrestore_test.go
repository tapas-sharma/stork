//go:build unittest
// +build unittest

package storkctl

import (
	"strings"
	"testing"
	"time"

	storkv1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	core "github.com/portworx/sched-ops/k8s/core"
	storkops "github.com/portworx/sched-ops/k8s/stork"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetRestoresNoRestore(t *testing.T) {
	cmdArgs := []string{"get", "applicationrestores"}

	var restoreList storkv1.ApplicationRestoreList
	expected := "No resources found.\n"
	testCommon(t, cmdArgs, &restoreList, expected, false)
}

func createApplicationRestoreAndVerify(
	t *testing.T,
	name string,
	namespace string,
	namespaces []string,
	backupLocation string,
	backupName string,
	resources string,
	createBackup bool,
) {

	if createBackup {
		createApplicationBackupAndVerify(t, backupName, namespace, namespaces, backupLocation, "", "", "")
	}

	cmdArgs := []string{"create", "apprestores", "-n", namespace, name, "--backupLocation", backupLocation, "--backupName", backupName}

	if len(resources) > 0 {
		cmdArgs = append(cmdArgs, "--resources", resources)
	}

	expected := "ApplicationRestore " + name + " started successfully\n"
	testCommon(t, cmdArgs, nil, expected, false)

	// Make sure it was created correctly
	restore, err := storkops.Instance().GetApplicationRestore(name, namespace)
	require.NoError(t, err, "Error getting restore")
	require.Equal(t, name, restore.Name, "ApplicationRestore name mismatch")
	require.Equal(t, namespace, restore.Namespace, "ApplicationRestore namespace mismatch")
	require.Equal(t, backupLocation, restore.Spec.BackupLocation, "ApplicationRestore backupLocation mismatch")
	require.Equal(t, backupName, restore.Spec.BackupName, "ApplicationRestore backupName mismatch")
}

func TestGetApplicationRestoresOneApplicationRestore(t *testing.T) {
	defer resetTest()
	createApplicationRestoreAndVerify(t, "getrestoretest", "test", []string{"namespace1"}, "backuplocation", "backupname", "", true)

	expected := "NAME             STAGE   STATUS   VOLUMES   RESOURCES   CREATED   ELAPSED\n" +
		"getrestoretest                    0/0       0                     \n"

	cmdArgs := []string{"get", "apprestores", "-n", "test"}
	testCommon(t, cmdArgs, nil, expected, false)
}

func TestGetApplicationRestoresMultiple(t *testing.T) {
	defer resetTest()
	_, err := core.Instance().CreateNamespace(&v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "default"}})
	require.NoError(t, err, "Error creating default namespace")

	createApplicationRestoreAndVerify(t, "getrestoretest1", "default", []string{"namespace1"}, "backuplocation", "backupname", "", true)
	createApplicationRestoreAndVerify(t, "getrestoretest2", "default", []string{"namespace1"}, "backuplocation", "backupname", "", false)

	expected := "NAME              STAGE   STATUS   VOLUMES   RESOURCES   CREATED   ELAPSED\n" +
		"getrestoretest1                    0/0       0                     \n" +
		"getrestoretest2                    0/0       0                     \n"

	cmdArgs := []string{"get", "apprestores", "getrestoretest1", "getrestoretest2"}
	testCommon(t, cmdArgs, nil, expected, false)

	// Should get all restores if no name given
	cmdArgs = []string{"get", "apprestores"}
	testCommon(t, cmdArgs, nil, expected, false)

	expected = "NAME              STAGE   STATUS   VOLUMES   RESOURCES   CREATED   ELAPSED\n" +
		"getrestoretest1                    0/0       0                     \n"
	// Should get only one restore if name given
	cmdArgs = []string{"get", "apprestores", "getrestoretest1"}
	testCommon(t, cmdArgs, nil, expected, false)

	_, err = core.Instance().CreateNamespace(&v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "ns1"}})
	require.NoError(t, err, "Error creating ns1 namespace")
	createApplicationRestoreAndVerify(t, "getrestoretest21", "ns1", []string{"namespace1"}, "backuplocation", "backupname", "", true)
	cmdArgs = []string{"get", "apprestores", "--all-namespaces"}
	expected = "NAMESPACE   NAME               STAGE   STATUS   VOLUMES   RESOURCES   CREATED   ELAPSED\n" +
		"default     getrestoretest1                     0/0       0                     \n" +
		"default     getrestoretest2                     0/0       0                     \n" +
		"ns1         getrestoretest21                    0/0       0                     \n"
	testCommon(t, cmdArgs, nil, expected, false)
}

func TestGetApplicationRestoresWithStatusAndProgress(t *testing.T) {
	defer resetTest()
	createApplicationRestoreAndVerify(t, "getrestorestatustest", "default", []string{"namespace1"}, "backuplocation", "backupname", "", true)
	restore, err := storkops.Instance().GetApplicationRestore("getrestorestatustest", "default")
	require.NoError(t, err, "Error getting restore")

	// Update the status of the restore
	restore.Status.FinishTimestamp = metav1.Now()
	restore.CreationTimestamp = metav1.NewTime(restore.Status.FinishTimestamp.Add(-5 * time.Minute))
	restore.Status.Stage = storkv1.ApplicationRestoreStageFinal
	restore.Status.Status = storkv1.ApplicationRestoreStatusSuccessful
	restore.Status.Volumes = []*storkv1.ApplicationRestoreVolumeInfo{}
	_, err = storkops.Instance().UpdateApplicationRestore(restore)
	require.NoError(t, err, "Error updating restore")

	expected := "NAME                   STAGE   STATUS       VOLUMES   RESOURCES   CREATED               ELAPSED\n" +
		"getrestorestatustest   Final   Successful   0/0       0           " + toTimeString(restore.CreationTimestamp.Time) + "   5m0s\n"
	cmdArgs := []string{"get", "apprestores", "getrestorestatustest"}
	testCommon(t, cmdArgs, nil, expected, false)
}

func TestSpecificObjectRestore(t *testing.T) {
	createApplicationRestoreAndVerify(t, "specificrestoretest1", "test", []string{"namespace1"}, "backuplocation", "backupname", "pvc/namespace1/pvc1", true)

	expected := "NAME                   STAGE   STATUS   VOLUMES   RESOURCES   CREATED   ELAPSED\n" +
		"specificrestoretest1                    0/0       0                     \n"

	cmdArgs := []string{"get", "apprestores", "-n", "test"}
	testCommon(t, cmdArgs, nil, expected, false)

	restore, err := storkops.Instance().GetApplicationRestore("specificrestoretest1", "test")
	require.NoError(t, err, "Error getting restore specificrestoretest1")
	require.Equal(t, "PersistentVolumeClaim", restore.Spec.IncludeResources[0].Kind, "ApplicationRestore Resource Kind Mismatch")
	require.Equal(t, "pvc1", restore.Spec.IncludeResources[0].Name, "ApplicationRestore Resource Name Mismatch")
	require.Equal(t, "namespace1", restore.Spec.IncludeResources[0].Namespace, "ApplicationRestore Resource Namespace Mismatch")
	require.Equal(t, "", restore.Spec.IncludeResources[0].Group, "ApplicationRestore Resource Group Mismatch")
	require.Equal(t, "v1", restore.Spec.IncludeResources[0].Version, "ApplicationRestore Resource Version Mismatch")

	resourceList := []string{
		"Deployment/namespace1/dep1",
		"deploy/namespace1/dep2",
		"PersistentVolumeClaim/namespace1/pvc1",
		"ServiceAccount/namespace1/sa1",
		"Service/namespace1/svc1",
		"StatefulSet/namespace1/sts1",
		"sts/namespace1/sts1",
		"Ingress/namespace1/ingress1",
		"ConfigMap/namespace1/cm1",
		"NetworkPolicy/namespace1/np1",
		"ClusterRole/cr1",
		"ClusterRoleBinding/crb1",
		"roles/namespace1/role1",
		"RoleBinding/namespace1/role1",
		"pv/pv1",
	}
	resources := strings.Join(resourceList, ",")
	createApplicationRestoreAndVerify(t, "specificrestoretest2", "test", []string{"namespace1"}, "backuplocation", "backupname", resources, false)

	expected = "NAME                   STAGE   STATUS   VOLUMES   RESOURCES   CREATED   ELAPSED\n" +
		"specificrestoretest1                    0/0       0                     \n" +
		"specificrestoretest2                    0/0       0                     \n"

	cmdArgs = []string{"get", "apprestores", "-n", "test"}
	testCommon(t, cmdArgs, nil, expected, false)

	cmdArgs = []string{"create", "apprestores", "-n", "test", "restoreName", "--backupLocation", "backupLocation", "--backupName", "backupname", "--resources", "ConfMapps/namespace1/cm1"}
	expected = "Error getting resource type for input resourcetype: ConfMapps\n" +
		"ApplicationRestore restoreName started successfully\n"
	testCommon(t, cmdArgs, nil, expected, false)

}

func TestCreateApplicationRestoreNoBackup(t *testing.T) {
	cmdArgs := []string{"create", "apprestores", "-n", "test", "restoreName", "--backupLocation", "backupLocation", "--backupName", "backupName"}
	expected := "error: applicationbackup backupName does not exist in namespace test"
	testCommon(t, cmdArgs, nil, expected, true)
}

func TestCreateApplicationRestoresNoName(t *testing.T) {
	cmdArgs := []string{"create", "apprestores"}

	expected := "error: exactly one name needs to be provided for applicationrestore name"
	testCommon(t, cmdArgs, nil, expected, true)
}

func TestCreateApplicationRestores(t *testing.T) {
	defer resetTest()
	createApplicationRestoreAndVerify(t, "createrestore", "default", []string{"namespace1"}, "backuplocation", "backupname", "", true)
}

func TestCreateApplicationRestoresMissingParameters(t *testing.T) {
	defer resetTest()
	cmdArgs := []string{"create", "apprestores", "createrestore", "--backupName", "backupname"}
	expected := "error: need to provide BackupLocation to use for restore"
	testCommon(t, cmdArgs, nil, expected, true)

	cmdArgs = []string{"create", "apprestores", "createrestore", "--backupLocation", "backuplocation"}
	expected = "error: need to provide BackupName to restore"
	testCommon(t, cmdArgs, nil, expected, true)
}

func TestCreateDuplicateApplicationRestores(t *testing.T) {
	defer resetTest()
	createApplicationRestoreAndVerify(t, "createrestore", "default", []string{"namespace1"}, "backuplocation", "backupname", "", true)
	cmdArgs := []string{"create", "apprestores", "createrestore", "--backupLocation", "backuplocation", "--backupName", "backupname"}

	expected := "Error from server (AlreadyExists): applicationrestores.stork.libopenstorage.org \"createrestore\" already exists"
	testCommon(t, cmdArgs, nil, expected, true)
}

func TestDeleteApplicationRestoresNoApplicationRestoreName(t *testing.T) {
	cmdArgs := []string{"delete", "apprestores"}

	var restoreList storkv1.ApplicationRestoreList
	expected := "error: at least one argument needs to be provided for applicationrestore name"
	testCommon(t, cmdArgs, &restoreList, expected, true)
}

func TestDeleteApplicationRestores(t *testing.T) {
	defer resetTest()
	createApplicationRestoreAndVerify(t, "deleterestore", "default", []string{"namespace1"}, "backuplocation", "backupname", "", true)

	cmdArgs := []string{"delete", "apprestores", "deleterestore"}
	expected := "ApplicationRestore deleterestore deleted successfully\n"
	testCommon(t, cmdArgs, nil, expected, false)

	cmdArgs = []string{"delete", "apprestores", "deleterestore"}
	expected = "Error from server (NotFound): applicationrestores.stork.libopenstorage.org \"deleterestore\" not found"
	testCommon(t, cmdArgs, nil, expected, true)

	createApplicationRestoreAndVerify(t, "deleterestore1", "default", []string{"namespace1"}, "backuplocation", "backupname1", "", true)
	createApplicationRestoreAndVerify(t, "deleterestore2", "default", []string{"namespace1"}, "backuplocation", "backupname2", "", true)

	cmdArgs = []string{"delete", "apprestores", "deleterestore1", "deleterestore2"}
	expected = "ApplicationRestore deleterestore1 deleted successfully\n"
	expected += "ApplicationRestore deleterestore2 deleted successfully\n"
	testCommon(t, cmdArgs, nil, expected, false)

	createApplicationRestoreAndVerify(t, "deleterestore1", "default", []string{"namespace1"}, "backuplocation", "backupname1", "", false)
	createApplicationRestoreAndVerify(t, "deleterestore2", "default", []string{"namespace1"}, "backuplocation", "backupname2", "", false)
}

func TestCreateApplicationRestoreWaitSuccess(t *testing.T) {
	restoreStatusRetryInterval = 10 * time.Second
	defer resetTest()

	namespace := "dummy-namespace"
	name := "dummy-name"

	createApplicationBackupAndVerify(t, "backupname", namespace, []string{namespace}, "backuplocation", "", "", "")

	cmdArgs := []string{"create", "apprestores", "-n", namespace, name, "--backupLocation", "backuplocation", "--backupName", "backupname", "--wait"}

	expected := "ApplicationRestore dummy-name started successfully\n" +
		"STAGE\t\tSTATUS              \n" +
		"\t\t                    \n" +
		"Volumes\t\tSuccessful          \n" +
		"ApplicationRestore dummy-name completed successfully\n"
	go setApplicationRestoreStatus(name, namespace, false, t)
	testCommon(t, cmdArgs, nil, expected, false)
}

func TestCreateApplicationRestoreWaitFailed(t *testing.T) {
	restoreStatusRetryInterval = 10 * time.Second
	defer resetTest()

	namespace := "dummy-namespace"
	name := "dummy-name"

	createApplicationBackupAndVerify(t, "backupname", namespace, []string{namespace}, "backuplocation", "", "", "")

	cmdArgs := []string{"create", "applicationrestore", "-n", namespace, name, "--backupLocation", "backuplocation", "--backupName", "backupname", "--wait"}

	expected := "ApplicationRestore dummy-name started successfully\n" +
		"STAGE\t\tSTATUS              \n" +
		"\t\t                    \n" +
		"Volumes\t\tFailed              \n" +
		"ApplicationRestore dummy-name failed\n"
	go setApplicationRestoreStatus(name, namespace, true, t)
	testCommon(t, cmdArgs, nil, expected, false)
}

func setApplicationRestoreStatus(name, namespace string, isFail bool, t *testing.T) {
	time.Sleep(10 * time.Second)
	restore, err := storkops.Instance().GetApplicationRestore(name, namespace)
	require.NoError(t, err, "Error getting ApplicationRestore details")
	require.Equal(t, restore.Status.Status, storkv1.ApplicationRestoreStatusInitial)
	require.Equal(t, restore.Status.Stage, storkv1.ApplicationRestoreStageInitial)
	restore.Status.Status = storkv1.ApplicationRestoreStatusSuccessful
	restore.Status.Stage = storkv1.ApplicationRestoreStageVolumes
	if isFail {
		restore.Status.Status = storkv1.ApplicationRestoreStatusFailed
	}

	_, err = storkops.Instance().UpdateApplicationRestore(restore)
	require.NoError(t, err, "Error updating ApplicationRestores")
}
