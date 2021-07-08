/*
Copyright 2021.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DeviceSpec defines the desired state of Device
type DeviceServiceSpec struct {
	config *corev1.ObjectReference `json:"config, omitempty"`
}

// DeviceStatus defines the observed state of Device
type DeviceServiceStatus struct {
	Condition DeviceServiceCondition `json:"condition, omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Device is the Schema for the devices API
type DeviceService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeviceServiceSpec   `json:"spec,omitempty"`
	Status DeviceServiceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DeviceServiceList contains a list of DeviceService
type DeviceServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeviceService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Device{}, &DeviceList{})
}

type DeviceServiceCondition struct {
	Type   DeviceServiceConditionType `json:"type, omitempty"`
	Status metav1.ConditionStatus     `json:"status, omitempty"`
}

type DeviceServiceConditionType string
