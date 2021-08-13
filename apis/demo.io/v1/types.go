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
	// Image name for rsync source or rsync daemon.
	Image string `json:"image"`
	// No of replicas required in the rsync daemon.
	Replicas *int32 `json:"replicas"`
	// +optional
	// HostName is kubernetes.io/hostname label value for node.
	// This is used to schedule rsync source to a node.
	HostName string `json:"hostName"`
	// Username added as valid user for rsync daemon. Rsync client that
	// has password can access the source uisng username and password.
	Username string `json:"username"`
	// Password to access rsync sourceby any client that has access to
	// username and pawword.
	Password string `json:"password"`
	// Volume that we want to make available to the client. This volume
	// is mounted to the source as a readonly filesystem. Clisnt can only
	// read that volume using credentials.
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
