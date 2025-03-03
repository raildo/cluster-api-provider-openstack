/*
Copyright 2020 The Kubernetes Authors.

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

package v1alpha3

import (
	corev1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	ctrlconversion "sigs.k8s.io/controller-runtime/pkg/conversion"

	"sigs.k8s.io/cluster-api-provider-openstack/api/v1beta1"
)

var _ ctrlconversion.Convertible = &OpenStackCluster{}

func (r *OpenStackCluster) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*v1beta1.OpenStackCluster)

	return Convert_v1alpha3_OpenStackCluster_To_v1beta1_OpenStackCluster(r, dst, nil)
}

func (r *OpenStackCluster) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*v1beta1.OpenStackCluster)

	return Convert_v1beta1_OpenStackCluster_To_v1alpha3_OpenStackCluster(src, r, nil)
}

var _ ctrlconversion.Convertible = &OpenStackClusterList{}

func (r *OpenStackClusterList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*v1beta1.OpenStackClusterList)

	return Convert_v1alpha3_OpenStackClusterList_To_v1beta1_OpenStackClusterList(r, dst, nil)
}

func (r *OpenStackClusterList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*v1beta1.OpenStackClusterList)

	return Convert_v1beta1_OpenStackClusterList_To_v1alpha3_OpenStackClusterList(src, r, nil)
}

var _ ctrlconversion.Convertible = &OpenStackMachine{}

func (r *OpenStackMachine) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*v1beta1.OpenStackMachine)

	return Convert_v1alpha3_OpenStackMachine_To_v1beta1_OpenStackMachine(r, dst, nil)
}

func (r *OpenStackMachine) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*v1beta1.OpenStackMachine)

	return Convert_v1beta1_OpenStackMachine_To_v1alpha3_OpenStackMachine(src, r, nil)
}

var _ ctrlconversion.Convertible = &OpenStackMachineList{}

func (r *OpenStackMachineList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*v1beta1.OpenStackMachineList)

	return Convert_v1alpha3_OpenStackMachineList_To_v1beta1_OpenStackMachineList(r, dst, nil)
}

func (r *OpenStackMachineList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*v1beta1.OpenStackMachineList)

	return Convert_v1beta1_OpenStackMachineList_To_v1alpha3_OpenStackMachineList(src, r, nil)
}

var _ ctrlconversion.Convertible = &OpenStackMachineTemplate{}

func (r *OpenStackMachineTemplate) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*v1beta1.OpenStackMachineTemplate)

	return Convert_v1alpha3_OpenStackMachineTemplate_To_v1beta1_OpenStackMachineTemplate(r, dst, nil)
}

func (r *OpenStackMachineTemplate) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*v1beta1.OpenStackMachineTemplate)

	return Convert_v1beta1_OpenStackMachineTemplate_To_v1alpha3_OpenStackMachineTemplate(src, r, nil)
}

var _ ctrlconversion.Convertible = &OpenStackMachineTemplateList{}

func (r *OpenStackMachineTemplateList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*v1beta1.OpenStackMachineTemplateList)

	return Convert_v1alpha3_OpenStackMachineTemplateList_To_v1beta1_OpenStackMachineTemplateList(r, dst, nil)
}

func (r *OpenStackMachineTemplateList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*v1beta1.OpenStackMachineTemplateList)

	return Convert_v1beta1_OpenStackMachineTemplateList_To_v1alpha3_OpenStackMachineTemplateList(src, r, nil)
}

// Convert_v1alpha3_OpenStackClusterSpec_To_v1beta1_OpenStackClusterSpec has to be added by us because we dropped
// the useOctavia parameter. We don't have to migrate this parameter to v1beta1 so there is nothing to do.
func Convert_v1alpha3_OpenStackClusterSpec_To_v1beta1_OpenStackClusterSpec(in *OpenStackClusterSpec, out *v1beta1.OpenStackClusterSpec, s conversion.Scope) error {
	if in.CloudsSecret != nil {
		out.IdentityRef = &v1beta1.OpenStackIdentityReference{
			Kind: "Secret",
			Name: in.CloudsSecret.Name,
		}
	}
	return autoConvert_v1alpha3_OpenStackClusterSpec_To_v1beta1_OpenStackClusterSpec(in, out, s)
}

// Convert_v1beta1_OpenStackClusterSpec_To_v1alpha3_OpenStackClusterSpec has to be added by us because we have to
// convert the Type of CloudsSecret from SecretReference to string.
func Convert_v1beta1_OpenStackClusterSpec_To_v1alpha3_OpenStackClusterSpec(in *v1beta1.OpenStackClusterSpec, out *OpenStackClusterSpec, s conversion.Scope) error {
	if in.IdentityRef != nil {
		out.CloudsSecret = &corev1.SecretReference{
			Name: in.IdentityRef.Name,
		}
	}

	if in.Bastion != nil && in.Bastion.Instance.IdentityRef != nil {
		outBastion := out.Bastion
		if outBastion == nil {
			outBastion = &Bastion{}
		}

		outBastion.Instance.CloudsSecret = &corev1.SecretReference{
			Name: in.Bastion.Instance.IdentityRef.Name,
		}
	}
	return autoConvert_v1beta1_OpenStackClusterSpec_To_v1alpha3_OpenStackClusterSpec(in, out, s)
}

// Convert_v1alpha3_OpenStackMachineSpec_To_v1beta1_OpenStackMachineSpec is an autogenerated conversion function.
// v1beta1 drops the field .UserDataSecret which is why we reuqire to define the function here.
func Convert_v1alpha3_OpenStackMachineSpec_To_v1beta1_OpenStackMachineSpec(in *OpenStackMachineSpec, out *v1beta1.OpenStackMachineSpec, s conversion.Scope) error {
	if in.CloudsSecret != nil {
		out.IdentityRef = &v1beta1.OpenStackIdentityReference{
			Name: in.CloudsSecret.Name,
			Kind: "Secret",
		}
	}
	return autoConvert_v1alpha3_OpenStackMachineSpec_To_v1beta1_OpenStackMachineSpec(in, out, s)
}

// Convert_v1beta1_Network_To_v1alpha3_Network has to be added by us for the new portOpts
// parameter in v1beta1. There is no intention to support this parameter in v1alpha3, so the field is just dropped.
func Convert_v1beta1_Network_To_v1alpha3_Network(in *v1beta1.Network, out *Network, s conversion.Scope) error {
	return autoConvert_v1beta1_Network_To_v1alpha3_Network(in, out, s)
}

// Convert_v1beta1_OpenStackMachineSpec_To_v1alpha3_OpenStackMachineSpec has to be added by us for the new ports
// parameter in v1beta1. There is no intention to support this parameter in v1alpha3, so the field is just dropped.
// Further, we want to convert the Type of CloudsSecret from SecretReference to string.
func Convert_v1beta1_OpenStackMachineSpec_To_v1alpha3_OpenStackMachineSpec(in *v1beta1.OpenStackMachineSpec, out *OpenStackMachineSpec, s conversion.Scope) error {
	if in.IdentityRef != nil {
		out.CloudsSecret = &corev1.SecretReference{
			Name: in.IdentityRef.Name,
		}
	}
	return autoConvert_v1beta1_OpenStackMachineSpec_To_v1alpha3_OpenStackMachineSpec(in, out, s)
}

// Convert_v1beta1_OpenStackClusterStatus_To_v1alpha3_OpenStackClusterStatus has to be added
// in order to drop the FailureReason and FailureMessage fields that are not present in v1alpha3.
func Convert_v1beta1_OpenStackClusterStatus_To_v1alpha3_OpenStackClusterStatus(in *v1beta1.OpenStackClusterStatus, out *OpenStackClusterStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_OpenStackClusterStatus_To_v1alpha3_OpenStackClusterStatus(in, out, s)
}

func Convert_Slice_v1beta1_Network_To_Slice_v1alpha3_Network(in *[]v1beta1.Network, out *[]Network, s conversion.Scope) error {
	for i := range *in {
		inNet := &(*in)[i]
		outNet := new(Network)
		if err := autoConvert_v1beta1_Network_To_v1alpha3_Network(inNet, outNet, s); err != nil {
			return err
		}
		*out = append(*out, *outNet)
	}
	return nil
}

func Convert_Slice_v1alpha3_Network_To_Slice_v1beta1_Network(in *[]Network, out *[]v1beta1.Network, s conversion.Scope) error {
	for i := range *in {
		inNet := &(*in)[i]
		outNet := new(v1beta1.Network)
		if err := autoConvert_v1alpha3_Network_To_v1beta1_Network(inNet, outNet, s); err != nil {
			return err
		}
		*out = append(*out, *outNet)
	}
	return nil
}

func Convert_v1alpha3_SubnetFilter_To_v1beta1_SubnetFilter(in *SubnetFilter, out *v1beta1.SubnetFilter, s conversion.Scope) error {
	out.Name = in.Name
	out.Description = in.Description
	if in.ProjectID != "" {
		out.ProjectID = in.ProjectID
	} else {
		out.ProjectID = in.TenantID
	}
	out.IPVersion = in.IPVersion
	out.GatewayIP = in.GatewayIP
	out.CIDR = in.CIDR
	out.IPv6AddressMode = in.IPv6AddressMode
	out.IPv6RAMode = in.IPv6RAMode
	out.ID = in.ID
	out.Tags = in.Tags
	out.TagsAny = in.TagsAny
	out.NotTags = in.NotTags
	out.NotTagsAny = in.NotTagsAny
	return nil
}

func Convert_v1beta1_SubnetFilter_To_v1alpha3_SubnetFilter(in *v1beta1.SubnetFilter, out *SubnetFilter, s conversion.Scope) error {
	out.TenantID = in.ProjectID
	return autoConvert_v1beta1_SubnetFilter_To_v1alpha3_SubnetFilter(in, out, s)
}

func Convert_v1alpha3_Filter_To_v1beta1_NetworkFilter(in *Filter, out *v1beta1.NetworkFilter, s conversion.Scope) error {
	out.Name = in.Name
	out.Description = in.Description
	if in.ProjectID != "" {
		out.ProjectID = in.ProjectID
	} else {
		out.ProjectID = in.TenantID
	}
	out.ID = in.ID
	out.Tags = in.Tags
	out.TagsAny = in.TagsAny
	out.NotTags = in.NotTags
	out.NotTagsAny = in.NotTagsAny
	return nil
}

func Convert_v1beta1_NetworkFilter_To_v1alpha3_Filter(in *v1beta1.NetworkFilter, out *Filter, s conversion.Scope) error {
	out.Name = in.Name
	out.Description = in.Description
	out.ProjectID = in.ProjectID
	out.TenantID = in.ProjectID
	out.ID = in.ID
	out.Tags = in.Tags
	out.TagsAny = in.TagsAny
	out.NotTags = in.NotTags
	out.NotTagsAny = in.NotTagsAny
	return nil
}
