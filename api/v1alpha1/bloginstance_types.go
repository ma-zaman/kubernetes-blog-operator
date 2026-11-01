/*
Copyright 2026.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BlogInstanceSpec defines the desired state of BlogInstance
type BlogInstanceSpec struct {
	// Title is the display title of the blog.
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=100
	Title string `json:"title"`

	// Replicas is the desired number of blog application instances.
	//
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=10
	// +kubebuilder:default=1
	Replicas int32 `json:"replicas,omitempty"`

	// Image is the container image used to run the blog application.
	//
	// +kubebuilder:validation:MinLength=1
	Image string `json:"image"`

	// Hostname is the optional hostname used to expose the blog
	//
	// +optional
	Hostname string `json:"hostname,omitempty"`
}

// BlogInstanceStatus defines the observed state of BlogInstance.
type BlogInstanceStatus struct {
	// ReadyReplicas is the number of the blog application instances currently ready to receive traffic
	Readyreplicas int32 `json:"readyReplicas,omitempty"`

	// URL is the address at wich the blog is expected to be available.
	URL string `json:"url,omitempty"`

	// ObservedGeneration is the most recent generation observed by the controller
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Title",type=string,JSONPath=".spec.title"
// +kubebuilder:printcolumn:name="Replicas",type=integer,JSONPath=".spec.replicas"
// +kubebuilder:printcolumn:name="Ready",type=integer,JSONPath=".spec.readyReplicas"
// +kubebuilder:printcolumn:name="URL",type=string,JSONPath=".spec.url"
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=".spec.creationTimestamp"

// BlogInstance is the Schema for the bloginstances API
type BlogInstance struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BlogInstanceSpec   `json:"spec,omitempty"`
	Status BlogInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlogInstanceList contains a list of BlogInstance
type BlogInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []BlogInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(func(s *runtime.Scheme) error {
		s.AddKnownTypes(SchemeGroupVersion, &BlogInstance{}, &BlogInstanceList{})
		return nil
	})
}
