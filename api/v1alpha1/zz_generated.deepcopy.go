// +build !ignore_autogenerated

//
// Copyright 2021 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudPakInfo) DeepCopyInto(out *CloudPakInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudPakInfo.
func (in *CloudPakInfo) DeepCopy() *CloudPakInfo {
	if in == nil {
		return nil
	}
	out := new(CloudPakInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonWebUI) DeepCopyInto(out *CommonWebUI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonWebUI.
func (in *CommonWebUI) DeepCopy() *CommonWebUI {
	if in == nil {
		return nil
	}
	out := new(CommonWebUI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CommonWebUI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonWebUIConfig) DeepCopyInto(out *CommonWebUIConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonWebUIConfig.
func (in *CommonWebUIConfig) DeepCopy() *CommonWebUIConfig {
	if in == nil {
		return nil
	}
	out := new(CommonWebUIConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonWebUIList) DeepCopyInto(out *CommonWebUIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CommonWebUI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonWebUIList.
func (in *CommonWebUIList) DeepCopy() *CommonWebUIList {
	if in == nil {
		return nil
	}
	out := new(CommonWebUIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CommonWebUIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonWebUISpec) DeepCopyInto(out *CommonWebUISpec) {
	*out = *in
	out.CommonWebUIConfig = in.CommonWebUIConfig
	out.GlobalUIConfig = in.GlobalUIConfig
	out.Resources = in.Resources
	out.License = in.License
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonWebUISpec.
func (in *CommonWebUISpec) DeepCopy() *CommonWebUISpec {
	if in == nil {
		return nil
	}
	out := new(CommonWebUISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonWebUIStatus) DeepCopyInto(out *CommonWebUIStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Versions = in.Versions
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonWebUIStatus.
func (in *CommonWebUIStatus) DeepCopy() *CommonWebUIStatus {
	if in == nil {
		return nil
	}
	out := new(CommonWebUIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalUIConfig) DeepCopyInto(out *GlobalUIConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalUIConfig.
func (in *GlobalUIConfig) DeepCopy() *GlobalUIConfig {
	if in == nil {
		return nil
	}
	out := new(GlobalUIConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *License) DeepCopyInto(out *License) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new License.
func (in *License) DeepCopy() *License {
	if in == nil {
		return nil
	}
	out := new(License)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Limits) DeepCopyInto(out *Limits) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Limits.
func (in *Limits) DeepCopy() *Limits {
	if in == nil {
		return nil
	}
	out := new(Limits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Requests) DeepCopyInto(out *Requests) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Requests.
func (in *Requests) DeepCopy() *Requests {
	if in == nil {
		return nil
	}
	out := new(Requests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	out.Requests = in.Requests
	out.Limits = in.Limits
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitcherItem) DeepCopyInto(out *SwitcherItem) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitcherItem.
func (in *SwitcherItem) DeepCopy() *SwitcherItem {
	if in == nil {
		return nil
	}
	out := new(SwitcherItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwitcherItem) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitcherItemList) DeepCopyInto(out *SwitcherItemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SwitcherItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitcherItemList.
func (in *SwitcherItemList) DeepCopy() *SwitcherItemList {
	if in == nil {
		return nil
	}
	out := new(SwitcherItemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwitcherItemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitcherItemSpec) DeepCopyInto(out *SwitcherItemSpec) {
	*out = *in
	out.CloudPakInfo = in.CloudPakInfo
	out.License = in.License
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitcherItemSpec.
func (in *SwitcherItemSpec) DeepCopy() *SwitcherItemSpec {
	if in == nil {
		return nil
	}
	out := new(SwitcherItemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitcherItemStatus) DeepCopyInto(out *SwitcherItemStatus) {
	*out = *in
	out.Versions = in.Versions
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitcherItemStatus.
func (in *SwitcherItemStatus) DeepCopy() *SwitcherItemStatus {
	if in == nil {
		return nil
	}
	out := new(SwitcherItemStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Versions) DeepCopyInto(out *Versions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Versions.
func (in *Versions) DeepCopy() *Versions {
	if in == nil {
		return nil
	}
	out := new(Versions)
	in.DeepCopyInto(out)
	return out
}
