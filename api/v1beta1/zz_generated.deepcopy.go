//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	"github.com/openstack-k8s-operators/lib-common/modules/common/service"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIOverrideSpec) DeepCopyInto(out *APIOverrideSpec) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = make(map[service.Endpoint]service.RoutedOverrideSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIOverrideSpec.
func (in *APIOverrideSpec) DeepCopy() *APIOverrideSpec {
	if in == nil {
		return nil
	}
	out := new(APIOverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Octavia) DeepCopyInto(out *Octavia) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Octavia.
func (in *Octavia) DeepCopy() *Octavia {
	if in == nil {
		return nil
	}
	out := new(Octavia)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Octavia) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaAPI) DeepCopyInto(out *OctaviaAPI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaAPI.
func (in *OctaviaAPI) DeepCopy() *OctaviaAPI {
	if in == nil {
		return nil
	}
	out := new(OctaviaAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OctaviaAPI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaAPIList) DeepCopyInto(out *OctaviaAPIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OctaviaAPI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaAPIList.
func (in *OctaviaAPIList) DeepCopy() *OctaviaAPIList {
	if in == nil {
		return nil
	}
	out := new(OctaviaAPIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OctaviaAPIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaAPISpec) DeepCopyInto(out *OctaviaAPISpec) {
	*out = *in
	in.OctaviaAPISpecCore.DeepCopyInto(&out.OctaviaAPISpecCore)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaAPISpec.
func (in *OctaviaAPISpec) DeepCopy() *OctaviaAPISpec {
	if in == nil {
		return nil
	}
	out := new(OctaviaAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaAPISpecCore) DeepCopyInto(out *OctaviaAPISpecCore) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	out.PasswordSelectors = in.PasswordSelectors
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DefaultConfigOverwrite != nil {
		in, out := &in.DefaultConfigOverwrite, &out.DefaultConfigOverwrite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	in.Override.DeepCopyInto(&out.Override)
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.TLS.DeepCopyInto(&out.TLS)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaAPISpecCore.
func (in *OctaviaAPISpecCore) DeepCopy() *OctaviaAPISpecCore {
	if in == nil {
		return nil
	}
	out := new(OctaviaAPISpecCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaAPIStatus) DeepCopyInto(out *OctaviaAPIStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaAPIStatus.
func (in *OctaviaAPIStatus) DeepCopy() *OctaviaAPIStatus {
	if in == nil {
		return nil
	}
	out := new(OctaviaAPIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaAmphoraController) DeepCopyInto(out *OctaviaAmphoraController) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaAmphoraController.
func (in *OctaviaAmphoraController) DeepCopy() *OctaviaAmphoraController {
	if in == nil {
		return nil
	}
	out := new(OctaviaAmphoraController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OctaviaAmphoraController) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaAmphoraControllerList) DeepCopyInto(out *OctaviaAmphoraControllerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OctaviaAmphoraController, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaAmphoraControllerList.
func (in *OctaviaAmphoraControllerList) DeepCopy() *OctaviaAmphoraControllerList {
	if in == nil {
		return nil
	}
	out := new(OctaviaAmphoraControllerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OctaviaAmphoraControllerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaAmphoraControllerSpec) DeepCopyInto(out *OctaviaAmphoraControllerSpec) {
	*out = *in
	in.OctaviaAmphoraControllerSpecCore.DeepCopyInto(&out.OctaviaAmphoraControllerSpecCore)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaAmphoraControllerSpec.
func (in *OctaviaAmphoraControllerSpec) DeepCopy() *OctaviaAmphoraControllerSpec {
	if in == nil {
		return nil
	}
	out := new(OctaviaAmphoraControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaAmphoraControllerSpecCore) DeepCopyInto(out *OctaviaAmphoraControllerSpecCore) {
	*out = *in
	out.PasswordSelectors = in.PasswordSelectors
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DefaultConfigOverwrite != nil {
		in, out := &in.DefaultConfigOverwrite, &out.DefaultConfigOverwrite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AmphoraCustomFlavors != nil {
		in, out := &in.AmphoraCustomFlavors, &out.AmphoraCustomFlavors
		*out = make([]OctaviaAmphoraFlavor, len(*in))
		copy(*out, *in)
	}
	out.TLS = in.TLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaAmphoraControllerSpecCore.
func (in *OctaviaAmphoraControllerSpecCore) DeepCopy() *OctaviaAmphoraControllerSpecCore {
	if in == nil {
		return nil
	}
	out := new(OctaviaAmphoraControllerSpecCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaAmphoraControllerStatus) DeepCopyInto(out *OctaviaAmphoraControllerStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaAmphoraControllerStatus.
func (in *OctaviaAmphoraControllerStatus) DeepCopy() *OctaviaAmphoraControllerStatus {
	if in == nil {
		return nil
	}
	out := new(OctaviaAmphoraControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaAmphoraFlavor) DeepCopyInto(out *OctaviaAmphoraFlavor) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaAmphoraFlavor.
func (in *OctaviaAmphoraFlavor) DeepCopy() *OctaviaAmphoraFlavor {
	if in == nil {
		return nil
	}
	out := new(OctaviaAmphoraFlavor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaApiTLS) DeepCopyInto(out *OctaviaApiTLS) {
	*out = *in
	in.API.DeepCopyInto(&out.API)
	out.Ca = in.Ca
	in.Ovn.DeepCopyInto(&out.Ovn)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaApiTLS.
func (in *OctaviaApiTLS) DeepCopy() *OctaviaApiTLS {
	if in == nil {
		return nil
	}
	out := new(OctaviaApiTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaDefaults) DeepCopyInto(out *OctaviaDefaults) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaDefaults.
func (in *OctaviaDefaults) DeepCopy() *OctaviaDefaults {
	if in == nil {
		return nil
	}
	out := new(OctaviaDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaLbMgmtNetworks) DeepCopyInto(out *OctaviaLbMgmtNetworks) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaLbMgmtNetworks.
func (in *OctaviaLbMgmtNetworks) DeepCopy() *OctaviaLbMgmtNetworks {
	if in == nil {
		return nil
	}
	out := new(OctaviaLbMgmtNetworks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaList) DeepCopyInto(out *OctaviaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Octavia, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaList.
func (in *OctaviaList) DeepCopy() *OctaviaList {
	if in == nil {
		return nil
	}
	out := new(OctaviaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OctaviaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaRsyslog) DeepCopyInto(out *OctaviaRsyslog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaRsyslog.
func (in *OctaviaRsyslog) DeepCopy() *OctaviaRsyslog {
	if in == nil {
		return nil
	}
	out := new(OctaviaRsyslog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OctaviaRsyslog) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaRsyslogList) DeepCopyInto(out *OctaviaRsyslogList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OctaviaRsyslog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaRsyslogList.
func (in *OctaviaRsyslogList) DeepCopy() *OctaviaRsyslogList {
	if in == nil {
		return nil
	}
	out := new(OctaviaRsyslogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OctaviaRsyslogList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaRsyslogSpec) DeepCopyInto(out *OctaviaRsyslogSpec) {
	*out = *in
	in.OctaviaRsyslogSpecCore.DeepCopyInto(&out.OctaviaRsyslogSpecCore)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaRsyslogSpec.
func (in *OctaviaRsyslogSpec) DeepCopy() *OctaviaRsyslogSpec {
	if in == nil {
		return nil
	}
	out := new(OctaviaRsyslogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaRsyslogSpecCore) DeepCopyInto(out *OctaviaRsyslogSpecCore) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DefaultConfigOverwrite != nil {
		in, out := &in.DefaultConfigOverwrite, &out.DefaultConfigOverwrite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaRsyslogSpecCore.
func (in *OctaviaRsyslogSpecCore) DeepCopy() *OctaviaRsyslogSpecCore {
	if in == nil {
		return nil
	}
	out := new(OctaviaRsyslogSpecCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaRsyslogStatus) DeepCopyInto(out *OctaviaRsyslogStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaRsyslogStatus.
func (in *OctaviaRsyslogStatus) DeepCopy() *OctaviaRsyslogStatus {
	if in == nil {
		return nil
	}
	out := new(OctaviaRsyslogStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaSpec) DeepCopyInto(out *OctaviaSpec) {
	*out = *in
	in.OctaviaSpecBase.DeepCopyInto(&out.OctaviaSpecBase)
	in.OctaviaAPI.DeepCopyInto(&out.OctaviaAPI)
	in.OctaviaHousekeeping.DeepCopyInto(&out.OctaviaHousekeeping)
	in.OctaviaHealthManager.DeepCopyInto(&out.OctaviaHealthManager)
	in.OctaviaWorker.DeepCopyInto(&out.OctaviaWorker)
	in.OctaviaRsyslog.DeepCopyInto(&out.OctaviaRsyslog)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaSpec.
func (in *OctaviaSpec) DeepCopy() *OctaviaSpec {
	if in == nil {
		return nil
	}
	out := new(OctaviaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaSpecBase) DeepCopyInto(out *OctaviaSpecBase) {
	*out = *in
	out.PasswordSelectors = in.PasswordSelectors
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DefaultConfigOverwrite != nil {
		in, out := &in.DefaultConfigOverwrite, &out.DefaultConfigOverwrite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.LbMgmtNetworks.DeepCopyInto(&out.LbMgmtNetworks)
	if in.AmphoraCustomFlavors != nil {
		in, out := &in.AmphoraCustomFlavors, &out.AmphoraCustomFlavors
		*out = make([]OctaviaAmphoraFlavor, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaSpecBase.
func (in *OctaviaSpecBase) DeepCopy() *OctaviaSpecBase {
	if in == nil {
		return nil
	}
	out := new(OctaviaSpecBase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaSpecCore) DeepCopyInto(out *OctaviaSpecCore) {
	*out = *in
	in.OctaviaSpecBase.DeepCopyInto(&out.OctaviaSpecBase)
	in.OctaviaAPI.DeepCopyInto(&out.OctaviaAPI)
	in.OctaviaHousekeeping.DeepCopyInto(&out.OctaviaHousekeeping)
	in.OctaviaHealthManager.DeepCopyInto(&out.OctaviaHealthManager)
	in.OctaviaWorker.DeepCopyInto(&out.OctaviaWorker)
	in.OctaviaRsyslog.DeepCopyInto(&out.OctaviaRsyslog)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaSpecCore.
func (in *OctaviaSpecCore) DeepCopy() *OctaviaSpecCore {
	if in == nil {
		return nil
	}
	out := new(OctaviaSpecCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaStatus) DeepCopyInto(out *OctaviaStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaStatus.
func (in *OctaviaStatus) DeepCopy() *OctaviaStatus {
	if in == nil {
		return nil
	}
	out := new(OctaviaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PasswordSelector) DeepCopyInto(out *PasswordSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordSelector.
func (in *PasswordSelector) DeepCopy() *PasswordSelector {
	if in == nil {
		return nil
	}
	out := new(PasswordSelector)
	in.DeepCopyInto(out)
	return out
}
