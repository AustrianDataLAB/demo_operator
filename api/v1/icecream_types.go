/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IcecreamSpec defines the desired state of Icecream
type IcecreamSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Icecream. Edit icecream_types.go to remove/update
	Flavour []string `json:"flavour,omitempty"`
	Topping []string `json:"topping,omitempty"`
}

// IcecreamStatus defines the observed state of Icecream
type IcecreamStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// a randomly assigned boolean to indicate if the icecream is sold out
	SoldOut bool `json:"soldOut,omitempty"`
	// you can only have 3 toppings
	ToppingsCount int `json:"toppingsCount,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster
//+kubebuilder:printcolumn:name="SoldOut",type="string",JSONPath=".status.soldOut",description="SoldOut"
//+kubebuilder:printcolumn:name="ToppingsCount",type="integer",JSONPath=".status.toppingsCount",description="ToppingsCount"

// Icecream is the Schema for the icecreams API
type Icecream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IcecreamSpec   `json:"spec,omitempty"`
	Status IcecreamStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// IcecreamList contains a list of Icecream
type IcecreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Icecream `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Icecream{}, &IcecreamList{})
}
