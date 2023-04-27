package storkctl

import (
	"fmt"
	"strings"
	"time"

	storkv1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	storkops "github.com/portworx/sched-ops/k8s/stork"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/cmd/util"
)

const (
	failoverCommand                    = "failover"
	nameTimeSuffixFormat string        = "2006-01-02-150405"
	actionWaitTimeout    time.Duration = 10 * time.Minute
	actionWaitInterval   time.Duration = 10 * time.Second
)

func newDoCommand(cmdFactory Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	doCommands := &cobra.Command{
		Use:    "do",
		Short:  "do actions",
		Hidden: true,
	}

	doCommands.AddCommand(
		newFailoverCommand(cmdFactory, ioStreams),
	)
	return doCommands
}

func newFailoverCommand(cmdFactory Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	getClusterPairCommand := &cobra.Command{
		Use:   failoverCommand,
		Short: "Initiate failover for the given namespaces",
		Run: func(c *cobra.Command, args []string) {
			namespaces, err := cmdFactory.GetAllNamespaces()
			if err != nil {
				util.CheckErr(err)
				return
			}
			for _, namespace := range namespaces {
				if incompleteAction := getAnyIncompleteAction(namespace); incompleteAction != nil {
					printMsg(
						fmt.Sprintf(
							"Failed to start failover for namespace %v as action %v is already %v",
							namespace, incompleteAction.Name, incompleteAction.Status),
						ioStreams.Out)
					continue
				}
				action := storkv1.Action{
					ObjectMeta: metav1.ObjectMeta{
						Name:      newActionName(storkv1.ActionTypeFailover),
						Namespace: namespace,
					},
					Spec: storkv1.ActionSpec{
						ActionType: storkv1.ActionTypeFailover,
					},
					Status: storkv1.ActionStatusScheduled,
				}
				_, err = storkops.Instance().CreateAction(&action)
				if err != nil {
					printMsg(
						fmt.Sprintf(
							"Failed to start failover for namespace %v due to error %v",
							namespace, err),
						ioStreams.ErrOut)
					continue
				}
				printMsg(fmt.Sprintf("Started failover for namespace %v", namespace), ioStreams.Out)
				printMsg(getDescribeActionMessage(&action), ioStreams.Out)
			}
		},
	}
	return getClusterPairCommand
}

func isActionIncomplete(action *storkv1.Action) bool {
	return action.Status == storkv1.ActionStatusScheduled || action.Status == storkv1.ActionStatusInProgress
}

// check if there is already an Action scheduled or in-progress
func getAnyIncompleteAction(namespace string) *storkv1.Action {
	actionList, err := storkops.Instance().ListActions(namespace)
	if err != nil {
		util.CheckErr(err)
		return nil
	}
	for _, action := range actionList.Items {
		if isActionIncomplete(&action) {
			return &action
		}
	}
	return nil
}

func newActionName(action storkv1.ActionType) string {
	return strings.Join([]string{string(action), time.Now().Format(nameTimeSuffixFormat)}, "-")
}

func getDescribeActionMessage(action *storkv1.Action) string {
	return "To check Action status use: " +
		fmt.Sprintf("kubectl describe action %v -n %v", action.Name, action.Namespace)
}
