// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cloud) DeepCopyInto(out *Cloud) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cloud.
func (in *Cloud) DeepCopy() *Cloud {
	if in == nil {
		return nil
	}
	out := new(Cloud)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfo) DeepCopyInto(out *ClusterInfo) {
	*out = *in
	out.Cloud = in.Cloud
	in.Kubernetes.DeepCopyInto(&out.Kubernetes)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfo.
func (in *ClusterInfo) DeepCopy() *ClusterInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerDeployment) DeepCopyInto(out *ControllerDeployment) {
	*out = *in
	if in.ProviderConfig != nil {
		in, out := &in.ProviderConfig, &out.ProviderConfig
		*out = new(ProviderConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerDeployment.
func (in *ControllerDeployment) DeepCopy() *ControllerDeployment {
	if in == nil {
		return nil
	}
	out := new(ControllerDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerInstallation) DeepCopyInto(out *ControllerInstallation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerInstallation.
func (in *ControllerInstallation) DeepCopy() *ControllerInstallation {
	if in == nil {
		return nil
	}
	out := new(ControllerInstallation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerInstallation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerInstallationList) DeepCopyInto(out *ControllerInstallationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControllerInstallation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerInstallationList.
func (in *ControllerInstallationList) DeepCopy() *ControllerInstallationList {
	if in == nil {
		return nil
	}
	out := new(ControllerInstallationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerInstallationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerInstallationSpec) DeepCopyInto(out *ControllerInstallationSpec) {
	*out = *in
	out.RegistrationRef = in.RegistrationRef
	out.SeedRef = in.SeedRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerInstallationSpec.
func (in *ControllerInstallationSpec) DeepCopy() *ControllerInstallationSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerInstallationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerInstallationStatus) DeepCopyInto(out *ControllerInstallationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProviderStatus != nil {
		in, out := &in.ProviderStatus, &out.ProviderStatus
		*out = new(ProviderConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerInstallationStatus.
func (in *ControllerInstallationStatus) DeepCopy() *ControllerInstallationStatus {
	if in == nil {
		return nil
	}
	out := new(ControllerInstallationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRegistration) DeepCopyInto(out *ControllerRegistration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRegistration.
func (in *ControllerRegistration) DeepCopy() *ControllerRegistration {
	if in == nil {
		return nil
	}
	out := new(ControllerRegistration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerRegistration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRegistrationList) DeepCopyInto(out *ControllerRegistrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControllerRegistration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRegistrationList.
func (in *ControllerRegistrationList) DeepCopy() *ControllerRegistrationList {
	if in == nil {
		return nil
	}
	out := new(ControllerRegistrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerRegistrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRegistrationSpec) DeepCopyInto(out *ControllerRegistrationSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ControllerResource, len(*in))
		copy(*out, *in)
	}
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(ControllerDeployment)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRegistrationSpec.
func (in *ControllerRegistrationSpec) DeepCopy() *ControllerRegistrationSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerRegistrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerResource) DeepCopyInto(out *ControllerResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerResource.
func (in *ControllerResource) DeepCopy() *ControllerResource {
	if in == nil {
		return nil
	}
	out := new(ControllerResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8SNetworks) DeepCopyInto(out *K8SNetworks) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = new(CIDR)
		**out = **in
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = new(CIDR)
		**out = **in
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = new(CIDR)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8SNetworks.
func (in *K8SNetworks) DeepCopy() *K8SNetworks {
	if in == nil {
		return nil
	}
	out := new(K8SNetworks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastError) DeepCopyInto(out *LastError) {
	*out = *in
	if in.Codes != nil {
		in, out := &in.Codes, &out.Codes
		*out = make([]ErrorCode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastError.
func (in *LastError) DeepCopy() *LastError {
	if in == nil {
		return nil
	}
	out := new(LastError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastOperation) DeepCopyInto(out *LastOperation) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastOperation.
func (in *LastOperation) DeepCopy() *LastOperation {
	if in == nil {
		return nil
	}
	out := new(LastOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Logging) DeepCopyInto(out *Logging) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]Endpoint, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Logging.
func (in *Logging) DeepCopy() *Logging {
	if in == nil {
		return nil
	}
	out := new(Logging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitoring) DeepCopyInto(out *Monitoring) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]Endpoint, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitoring.
func (in *Monitoring) DeepCopy() *Monitoring {
	if in == nil {
		return nil
	}
	out := new(Monitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plant) DeepCopyInto(out *Plant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plant.
func (in *Plant) DeepCopy() *Plant {
	if in == nil {
		return nil
	}
	out := new(Plant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Plant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlantList) DeepCopyInto(out *PlantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Plant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlantList.
func (in *PlantList) DeepCopy() *PlantList {
	if in == nil {
		return nil
	}
	out := new(PlantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlantSpec) DeepCopyInto(out *PlantSpec) {
	*out = *in
	out.SecretRef = in.SecretRef
	in.Monitoring.DeepCopyInto(&out.Monitoring)
	in.Logging.DeepCopyInto(&out.Logging)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlantSpec.
func (in *PlantSpec) DeepCopy() *PlantSpec {
	if in == nil {
		return nil
	}
	out := new(PlantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlantStatus) DeepCopyInto(out *PlantStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	in.ClusterInfo.DeepCopyInto(&out.ClusterInfo)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlantStatus.
func (in *PlantStatus) DeepCopy() *PlantStatus {
	if in == nil {
		return nil
	}
	out := new(PlantStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderConfig) DeepCopyInto(out *ProviderConfig) {
	*out = *in
	in.RawExtension.DeepCopyInto(&out.RawExtension)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderConfig.
func (in *ProviderConfig) DeepCopy() *ProviderConfig {
	if in == nil {
		return nil
	}
	out := new(ProviderConfig)
	in.DeepCopyInto(out)
	return out
}
