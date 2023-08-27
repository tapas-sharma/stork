package storkctl

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	storkv1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	apiextensionops "github.com/portworx/sched-ops/k8s/apiextensions"
	storkops "github.com/portworx/sched-ops/k8s/stork"
	"github.com/portworx/sched-ops/task"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1beta1 "k8s.io/apimachinery/pkg/apis/meta/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubernetes/pkg/printers"
)

var (
	restoreStatusRetryInterval = 30 * time.Second
	restoreStatusRetryTimeout  = 6 * time.Hour
)

var applicationRestoreColumns = []string{"NAME", "STAGE", "STATUS", "VOLUMES", "RESOURCES", "CREATED", "ELAPSED"}
var applicationRestoreSubcommand = "applicationrestores"
var applicationRestoreAliases = []string{"applicationrestore", "apprestore", "apprestores"}

func newCreateApplicationRestoreCommand(cmdFactory Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var applicationRestoreName string
	var backupLocation string
	var waitForCompletion bool
	var backupName string
	var replacePolicy string
	var resources string

	createApplicationRestoreCommand := &cobra.Command{
		Use:     applicationRestoreSubcommand,
		Aliases: applicationRestoreAliases,
		Short:   "Start an applicationRestore",
		Run: func(c *cobra.Command, args []string) {
			if len(args) != 1 {
				util.CheckErr(fmt.Errorf("exactly one name needs to be provided for applicationrestore name"))
				return
			}
			if backupLocation == "" {
				util.CheckErr(fmt.Errorf("need to provide BackupLocation to use for restore"))
				return
			}
			if backupName == "" {
				util.CheckErr(fmt.Errorf("need to provide BackupName to restore"))
				return
			}

			backup, err := storkops.Instance().GetApplicationBackup(backupName, cmdFactory.GetNamespace())
			if err != nil {
				util.CheckErr(fmt.Errorf("applicationbackup %s does not exist in namespace %s", backupName, cmdFactory.GetNamespace()))
				return
			}

			applicationRestoreName = args[0]
			applicationRestore := &storkv1.ApplicationRestore{
				Spec: storkv1.ApplicationRestoreSpec{
					BackupLocation: backupLocation,
					BackupName:     backupName,
					ReplacePolicy:  storkv1.ApplicationRestoreReplacePolicyType(replacePolicy),
				},
			}
			if len(resources) > 0 {
				objects := getObjectInfos(resources, ioStreams)
				applicationRestore.Spec.IncludeResources = objects
			}

			applicationRestore.Spec.NamespaceMapping = getDefaultNamespaceMapping(backup)
			applicationRestore.Name = applicationRestoreName
			applicationRestore.Namespace = cmdFactory.GetNamespace()
			_, err = storkops.Instance().CreateApplicationRestore(applicationRestore)
			if err != nil {
				util.CheckErr(err)
				return
			}

			msg := "ApplicationRestore " + applicationRestoreName + " started successfully"
			printMsg(msg, ioStreams.Out)

			if waitForCompletion {
				msg, err := waitForApplicationRestore(applicationRestore.Name, applicationRestore.Namespace, ioStreams)
				if err != nil {
					util.CheckErr(err)
					return
				}
				printMsg(msg, ioStreams.Out)
			}
		},
	}
	createApplicationRestoreCommand.Flags().BoolVarP(&waitForCompletion, "wait", "", false, "Wait for applicationrestore to complete")
	createApplicationRestoreCommand.Flags().StringVarP(&backupLocation, "backupLocation", "l", "", "BackupLocation to use for the restore")
	createApplicationRestoreCommand.Flags().StringVarP(&backupName, "backupName", "b", "", "Backup to restore from")
	createApplicationRestoreCommand.Flags().StringVarP(&replacePolicy, "replacePolicy", "r", "Retain", "Policy to use if resources being restored already exist (Retain or Delete).")
	createApplicationRestoreCommand.Flags().StringVarP(&resources, "resources", "", "",
		"Specific resources for restoring, should be given in format \"<kind>/<namespace>/<name>,<kind>/<namespace>/<name>,<kind>/<name>\", ex: \"<Deployment>/<ns1>/<dep1>,<PersistentVolumeClaim>/<ns1>/<pvc1>,<ClusterRole>/<clusterrole1>\"")

	return createApplicationRestoreCommand
}

func newGetApplicationRestoreCommand(cmdFactory Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	getApplicationRestoreCommand := &cobra.Command{
		Use:     applicationRestoreSubcommand,
		Aliases: applicationRestoreAliases,
		Short:   "Get applicationrestore resources",
		Run: func(c *cobra.Command, args []string) {
			var applicationRestores *storkv1.ApplicationRestoreList
			var err error

			namespaces, err := cmdFactory.GetAllNamespaces()
			if err != nil {
				util.CheckErr(err)
				return
			}
			if len(args) > 0 {
				applicationRestores = new(storkv1.ApplicationRestoreList)
				for _, applicationRestoreName := range args {
					for _, ns := range namespaces {
						applicationRestore, err := storkops.Instance().GetApplicationRestore(applicationRestoreName, ns)
						if err != nil {
							util.CheckErr(err)
							return
						}
						applicationRestores.Items = append(applicationRestores.Items, *applicationRestore)
					}
				}
			} else {
				var tempApplicationRestores storkv1.ApplicationRestoreList
				for _, ns := range namespaces {
					applicationRestores, err = storkops.Instance().ListApplicationRestores(ns, metav1.ListOptions{})
					if err != nil {
						util.CheckErr(err)
						return
					}
					tempApplicationRestores.Items = append(tempApplicationRestores.Items, applicationRestores.Items...)
				}
				applicationRestores = &tempApplicationRestores
			}

			if len(applicationRestores.Items) == 0 {
				handleEmptyList(ioStreams.Out)
				return
			}
			if cmdFactory.IsWatchSet() {
				if err := printObjectsWithWatch(c, applicationRestores, cmdFactory, applicationRestoreColumns, applicationRestorePrinter, ioStreams.Out); err != nil {
					util.CheckErr(err)
					return
				}
				return
			}
			if err := printObjects(c, applicationRestores, cmdFactory, applicationRestoreColumns, applicationRestorePrinter, ioStreams.Out); err != nil {
				util.CheckErr(err)
				return
			}
		},
	}
	cmdFactory.BindGetFlags(getApplicationRestoreCommand.Flags())

	return getApplicationRestoreCommand
}

func newDeleteApplicationRestoreCommand(cmdFactory Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	deleteApplicationRestoreCommand := &cobra.Command{
		Use:     applicationRestoreSubcommand,
		Aliases: applicationRestoreAliases,
		Short:   "Delete applicationrestore resources",
		Run: func(c *cobra.Command, args []string) {
			var applicationRestores []string

			if len(args) == 0 {
				util.CheckErr(fmt.Errorf("at least one argument needs to be provided for applicationrestore name"))
				return
			}
			applicationRestores = args

			deleteApplicationRestores(applicationRestores, cmdFactory.GetNamespace(), ioStreams)
		},
	}

	return deleteApplicationRestoreCommand
}

func deleteApplicationRestores(applicationRestores []string, namespace string, ioStreams genericclioptions.IOStreams) {
	for _, applicationRestore := range applicationRestores {
		err := storkops.Instance().DeleteApplicationRestore(applicationRestore, namespace)
		if err != nil {
			util.CheckErr(err)
			return
		}
		msg := fmt.Sprintf("ApplicationRestore %v deleted successfully", applicationRestore)
		printMsg(msg, ioStreams.Out)
	}
}

func applicationRestorePrinter(
	applicationRestoreList *storkv1.ApplicationRestoreList,
	options printers.GenerateOptions,
) ([]metav1beta1.TableRow, error) {
	if applicationRestoreList == nil {
		return nil, nil
	}
	rows := make([]metav1beta1.TableRow, 0)
	for _, applicationRestore := range applicationRestoreList.Items {
		name := applicationRestore.Name

		totalVolumes := len(applicationRestore.Status.Volumes)
		doneVolumes := 0
		for _, volume := range applicationRestore.Status.Volumes {
			if volume.Status == storkv1.ApplicationRestoreStatusSuccessful {
				doneVolumes++
			}
		}
		volumeStatus := fmt.Sprintf("%v/%v", doneVolumes, totalVolumes)

		elapsed := ""
		if !applicationRestore.CreationTimestamp.IsZero() {
			if applicationRestore.Status.Stage == storkv1.ApplicationRestoreStageFinal {
				if !applicationRestore.Status.FinishTimestamp.IsZero() {
					elapsed = applicationRestore.Status.FinishTimestamp.Sub(applicationRestore.CreationTimestamp.Time).String()
				}
			} else {
				elapsed = time.Since(applicationRestore.CreationTimestamp.Time).String()
			}
		}

		creationTime := toTimeString(applicationRestore.CreationTimestamp.Time)
		row := getRow(&applicationRestore,
			[]interface{}{name,
				applicationRestore.Status.Stage,
				applicationRestore.Status.Status,
				volumeStatus,
				len(applicationRestore.Status.Resources),
				creationTime,
				elapsed},
		)
		rows = append(rows, row)
	}
	return rows, nil
}

func waitForApplicationRestore(name, namespace string, ioStreams genericclioptions.IOStreams) (string, error) {
	var msg string
	var err error

	log.SetFlags(0)
	log.SetOutput(io.Discard)
	heading := fmt.Sprintf("%s\t\t%-20s", stage, status)
	printMsg(heading, ioStreams.Out)
	t := func() (interface{}, bool, error) {
		restore, err := storkops.Instance().GetApplicationRestore(name, namespace)
		if err != nil {
			util.CheckErr(err)
			return "", false, err
		}
		stat := fmt.Sprintf("%s\t\t%-20s", restore.Status.Stage, restore.Status.Status)
		printMsg(stat, ioStreams.Out)
		if restore.Status.Status == storkv1.ApplicationRestoreStatusSuccessful ||
			restore.Status.Status == storkv1.ApplicationRestoreStatusPartialSuccess {
			msg = fmt.Sprintf("ApplicationRestore %v completed successfully", name)
			return "", false, nil
		}
		if restore.Status.Status == storkv1.ApplicationRestoreStatusFailed {
			msg = fmt.Sprintf("ApplicationRestore %v failed", name)
			return "", false, nil
		}
		return "", true, fmt.Errorf("%v", restore.Status.Status)
	}
	// sleep just so that instead of blank initial stage/status,
	// we have something at start
	time.Sleep(5 * time.Second)
	if _, err = task.DoRetryWithTimeout(t, restoreStatusRetryTimeout, restoreStatusRetryInterval); err != nil {
		msg = "Timed out performing task"
	}

	return msg, err
}

func getDefaultNamespaceMapping(backup *storkv1.ApplicationBackup) map[string]string {
	nsMapping := make(map[string]string)
	for _, ns := range backup.Spec.Namespaces {
		nsMapping[ns] = ns
	}
	return nsMapping
}

func getObjectInfos(resourceString string, ioStreams genericclioptions.IOStreams) []storkv1.ObjectInfo {
	objects := make([]storkv1.ObjectInfo, 0)

	aecClient, err := apiextensionops.Instance().GetExtensionClient()
	if err != nil {
		printMsg(fmt.Sprintf("Error getting api extension client: %v", err), ioStreams.Out)
		return objects
	}
	discoveryClient := aecClient.Discovery()

	// List the available API resources
	apiResourceList, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		msg := fmt.Sprintf("Error getting API resources: %v", err)
		printMsg(msg, ioStreams.Out)
		return objects
	}

	resources := strings.Split(resourceString, ",")
	for _, resource := range resources {
		// resources will be provided like "<type>/<ns>/<name>,<type>/<name>"
		resourceDetails := strings.Split(resource, "/")
		object := storkv1.ObjectInfo{}
		if len(resourceDetails) == 3 {
			object.Name = resourceDetails[2]
			object.Namespace = resourceDetails[1]
		} else if len(resourceDetails) == 2 {
			object.Name = resourceDetails[1]
		} else {
			msg := fmt.Sprintf("Unsupported resource info %s", resource)
			printMsg(msg, ioStreams.Out)
			continue
		}

		resourceType := resourceDetails[0]
		found := false
		for _, group := range apiResourceList {
			groupVersion, err := schema.ParseGroupVersion(group.GroupVersion)
			if err != nil {
				continue
			}
			for _, apiResource := range group.APIResources {
				if isValidResourceType(resourceType, apiResource) {
					// valid input
					object.Kind = apiResource.Kind
					object.Version = groupVersion.Version
					object.Group = groupVersion.Group
					objects = append(objects, object)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			msg := fmt.Sprintf("Error getting resource type for input resourcetype: %s", resourceType)
			printMsg(msg, ioStreams.Out)
		}
	}
	return objects
}
