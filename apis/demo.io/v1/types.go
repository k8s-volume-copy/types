// +kubebuilder:object:generate=true
package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RsyncSource contains list of configurations for a rsync daemon.
type RsyncSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Spec contains list of configurations for a rsync daemon.
	Spec RsyncSourceSpec `json:"spec"`
	// +optional
	Status RsyncSourceStatus `json:"status"`
}

// RsyncSourceSpec contains the information of rsync source
type RsyncSourceSpec struct {
	// Image name for rsync source pod
	Image string `json:"image"`
	// +optional
	// NodeName to schedule source pod to a node
	NodeName string `json:"nodeName"`
	// +optional
	// HostName is kubernetes.io/hostname node label value for the given node name
	HostName string `json:"hostName"`
	// Username to access this source by a client
	Username string `json:"username"`
	// Password to access this source by a client
	Password string `json:"password"`
	// Volume that we want to make available to the client
	Volume corev1.Volume `json:"volume"`
}

// RsyncSourceStatus contains the information of rsync source
type RsyncSourceStatus struct {
	Phase corev1.PodPhase `json:"phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RsyncSourceList is a list of RsyncSource objects
type RsyncSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// List of RsyncSources
	Items []RsyncSource `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RsyncPopulator is a volume populator that helps
// to create a volume from any rsync source.
type RsyncPopulator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Spec contains details of rsync daemon. Rsync client will use these
	// information to connect and get the data for the volume.
	Spec RsyncPopulatorSpec `json:"spec"`
}

// RsyncPopulatorSpec contains the information of rsync daemon.
type RsyncPopulatorSpec struct {
	// Username is used as credential to access rsync daemon by the client.
	Username string `json:"username"`
	// Password is used as credential to access rsync daemon by the client.
	Password string `json:"password"`
	// URL is rsync daemon url it can be dns:port or ip:port.
	URL string `json:"url"`
	// Path is mount path of the volume which we want to sync by the clinet.
	Path string `json:"path"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RsyncPopulatorList is a list of RsyncPopulator objects
type RsyncPopulatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// List of RsyncPopulators
	Items []RsyncPopulator `json:"items" protobuf:"bytes,2,rep,name=items"`
}
