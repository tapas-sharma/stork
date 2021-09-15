// +build !ignore_autogenerated

/*

LICENSE

*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupLocationMaintenance) DeepCopyInto(out *BackupLocationMaintenance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupLocationMaintenance.
func (in *BackupLocationMaintenance) DeepCopy() *BackupLocationMaintenance {
	if in == nil {
		return nil
	}
	out := new(BackupLocationMaintenance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupLocationMaintenance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupLocationMaintenanceList) DeepCopyInto(out *BackupLocationMaintenanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupLocationMaintenance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupLocationMaintenanceList.
func (in *BackupLocationMaintenanceList) DeepCopy() *BackupLocationMaintenanceList {
	if in == nil {
		return nil
	}
	out := new(BackupLocationMaintenanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupLocationMaintenanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupLocationMaintenanceSpec) DeepCopyInto(out *BackupLocationMaintenanceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupLocationMaintenanceSpec.
func (in *BackupLocationMaintenanceSpec) DeepCopy() *BackupLocationMaintenanceSpec {
	if in == nil {
		return nil
	}
	out := new(BackupLocationMaintenanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupLocationMaintenanceStatus) DeepCopyInto(out *BackupLocationMaintenanceStatus) {
	*out = *in
	if in.RepoStatus != nil {
		in, out := &in.RepoStatus, &out.RepoStatus
		*out = make([]*RepoMaintenanceStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(RepoMaintenanceStatus)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupLocationMaintenanceStatus.
func (in *BackupLocationMaintenanceStatus) DeepCopy() *BackupLocationMaintenanceStatus {
	if in == nil {
		return nil
	}
	out := new(BackupLocationMaintenanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataExport) DeepCopyInto(out *DataExport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataExport.
func (in *DataExport) DeepCopy() *DataExport {
	if in == nil {
		return nil
	}
	out := new(DataExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataExport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataExportList) DeepCopyInto(out *DataExportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataExport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataExportList.
func (in *DataExportList) DeepCopy() *DataExportList {
	if in == nil {
		return nil
	}
	out := new(DataExportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataExportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataExportObjectReference) DeepCopyInto(out *DataExportObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataExportObjectReference.
func (in *DataExportObjectReference) DeepCopy() *DataExportObjectReference {
	if in == nil {
		return nil
	}
	out := new(DataExportObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataExportSpec) DeepCopyInto(out *DataExportSpec) {
	*out = *in
	out.Source = in.Source
	out.Destination = in.Destination
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataExportSpec.
func (in *DataExportSpec) DeepCopy() *DataExportSpec {
	if in == nil {
		return nil
	}
	out := new(DataExportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportStatus) DeepCopyInto(out *ExportStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportStatus.
func (in *ExportStatus) DeepCopy() *ExportStatus {
	if in == nil {
		return nil
	}
	out := new(ExportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepoMaintenanceStatus) DeepCopyInto(out *RepoMaintenanceStatus) {
	*out = *in
	in.LastRunTimestamp.DeepCopyInto(&out.LastRunTimestamp)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepoMaintenanceStatus.
func (in *RepoMaintenanceStatus) DeepCopy() *RepoMaintenanceStatus {
	if in == nil {
		return nil
	}
	out := new(RepoMaintenanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBackup) DeepCopyInto(out *VolumeBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBackup.
func (in *VolumeBackup) DeepCopy() *VolumeBackup {
	if in == nil {
		return nil
	}
	out := new(VolumeBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBackupList) DeepCopyInto(out *VolumeBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VolumeBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBackupList.
func (in *VolumeBackupList) DeepCopy() *VolumeBackupList {
	if in == nil {
		return nil
	}
	out := new(VolumeBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBackupSpec) DeepCopyInto(out *VolumeBackupSpec) {
	*out = *in
	out.BackupLocation = in.BackupLocation
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBackupSpec.
func (in *VolumeBackupSpec) DeepCopy() *VolumeBackupSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBackupStatus) DeepCopyInto(out *VolumeBackupStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBackupStatus.
func (in *VolumeBackupStatus) DeepCopy() *VolumeBackupStatus {
	if in == nil {
		return nil
	}
	out := new(VolumeBackupStatus)
	in.DeepCopyInto(out)
	return out
}
