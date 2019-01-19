/*
Copyright 2018 BlackRock, Inc.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NodePhase is the label for the condition of a node.
type NodePhase string

// possible types of node phases
const (
	NodePhaseRunning        NodePhase = "Running"        // the node is running
	NodePhaseError          NodePhase = "Error"          // the node has encountered an error in processing
	NodePhaseNew            NodePhase = ""               // the node is new
	NodePhaseCompleted      NodePhase = "Completed"      // node has completed running
	NodePhaseRemove         NodePhase = "Remove"         // stale node
	NodePhaseResourceUpdate NodePhase = "ResourceUpdate" // resource is updated
)

// DispatchProtocolType is type of the event dispatch protocol. Used for dispatching events from gateway to watchers
type DispatchProtocolType string

// possible types of event dispatch protocol
const (
	HTTPGateway DispatchProtocolType = "HTTP"
	NATSGateway DispatchProtocolType = "NATS"
)

// Gateway is the definition of a gateway resource
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type Gateway struct {
	// +k8s:openapi-gen=false
	metav1.TypeMeta `json:",inline"`
	// +k8s:openapi-gen=false
	metav1.ObjectMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`
	Status            GatewayStatus `json:"status" protobuf:"bytes,2,opt,name=status"`
	Spec              GatewaySpec   `json:"spec" protobuf:"bytes,3,opt,name=spec"`
}

// GatewayList is the list of Gateway resources
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type GatewayList struct {
	// +k8s:openapi-gen=false
	metav1.TypeMeta `json:",inline"`
	// +k8s:openapi-gen=false
	metav1.ListMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Gateway `json:"items" protobuf:"bytes,2,opt,name=items"`
}

// GatewaySpec represents gateway specifications
type GatewaySpec struct {
	// DeploySpec is the pod specification for the gateway
	// Refer https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.11/#pod-v1-core
	DeploySpec *corev1.Pod `json:"deploySpec" protobuf:"bytes,1,opt,name=deploySpec"`

	// ConfigMap is name of the configmap for gateway. This configmap contains event sources.
	ConfigMap string `json:"configMap,omitempty" protobuf:"bytes,2,opt,name=configmap"`

	// Type is the type of gateway. Used as metadata.
	Type string `json:"type" protobuf:"bytes,3,opt,name=type"`

	// Version is used for marking event version
	EventVersion string `json:"eventVersion" protobuf:"bytes,4,opt,name=eventVersion"`

	// ServiceSpec is the specifications of the service to expose the gateway
	// Refer https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.11/#service-v1-core
	ServiceSpec *corev1.Service `json:"serviceSpec,omitempty" protobuf:"bytes,5,opt,name=serviceSpec"`

	// Watchers are components which are interested listening to notifications from this gateway
	// These only need to be specified when gateway dispatch mechanism is through HTTP POST notifications.
	// In future, support for NATS, KAFKA will be added as a means to dispatch notifications in which case
	// specifying watchers would be unnecessary.
	Watchers *NotificationWatchers `json:"watchers,omitempty" protobuf:"bytes,6,opt,name=watchers"`

	// Port on which the gateway event source processor is running on.
	ProcessorPort string `json:"processorPort" protobuf:"bytes,7,opt,name=processorPort"`

	// EventProtocol is the underlying protocol used to send events from gateway to watchers(components interested in listening to event from this gateway)
	DispatchProtocol DispatchProtocol `json:"dispatchProtocol" protobuf:"bytes,8,opt,name=dispatchProtocol"`
}

// GatewayStatus contains information about the status of a gateway.
type GatewayStatus struct {
	// Phase is the high-level summary of the gateway
	Phase NodePhase `json:"phase" protobuf:"bytes,1,opt,name=phase"`

	// StartedAt is the time at which this gateway was initiated
	StartedAt metav1.Time `json:"startedAt,omitempty" protobuf:"bytes,2,opt,name=startedAt"`

	// Message is a human readable string indicating details about a gateway in its phase
	Message string `json:"message,omitempty" protobuf:"bytes,4,opt,name=message"`

	// Nodes is a mapping between a node ID and the node's status
	// it records the states for the configurations of gateway.
	Nodes map[string]NodeStatus `json:"nodes,omitempty" protobuf:"bytes,5,rep,name=nodes"`
}

// NodeStatus describes the status for an individual node in the gateway configurations.
// A single node can represent one configuration.
type NodeStatus struct {
	// ID is a unique identifier of a node within a sensor
	// It is a hash of the node name
	ID string `json:"id" protobuf:"bytes,1,opt,name=id"`

	// Name is a unique name in the node tree used to generate the node ID
	Name string `json:"name" protobuf:"bytes,3,opt,name=name"`

	// DisplayName is the human readable representation of the node
	DisplayName string `json:"displayName" protobuf:"bytes,5,opt,name=displayName"`

	// Phase of the node
	Phase NodePhase `json:"phase" protobuf:"bytes,6,opt,name=phase"`

	// StartedAt is the time at which this node started
	// +k8s:openapi-gen=false
	StartedAt metav1.MicroTime `json:"startedAt,omitempty" protobuf:"bytes,7,opt,name=startedAt"`

	// Message store data or something to save for configuration
	Message string `json:"message,omitempty" protobuf:"bytes,8,opt,name=message"`

	// UpdateTime is the time when node(gateway configuration) was updated
	// +k8s:openapi-gen=false
	UpdateTime metav1.MicroTime `json:"updateTime,omitempty" protobuf:"bytes,9,opt,name=updateTime"`
}

// NotificationWatchers are components which are interested listening to notifications from this gateway
type NotificationWatchers struct {
	// Gateways is the list of gateways interested in listening to notifications from this gateway
	Gateways []GatewayNotificationWatcher `json:"gateways,omitempty" protobuf:"bytes,1,opt,name=gateways"`

	// Sensors is the list of sensors interested in listening to notifications from this gateway
	Sensors []SensorNotificationWatcher `json:"sensors,omitempty" protobuf:"bytes,2,rep,name=sensors"`
}

// GatewayNotificationWatcher is the gateway interested in listening to notifications from this gateway
type GatewayNotificationWatcher struct {
	// Name is the gateway name
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`

	// Port is http server port on which gateway is running
	Port string `json:"port" protobuf:"bytes,2,opt,name=port"`

	// Endpoint is REST API endpoint to post event to.
	// Events are sent using HTTP POST method to this endpoint.
	Endpoint string `json:"endpoint" protobuf:"bytes,3,opt,name=endpoint"`
}

// SensorNotificationWatcher is the sensor interested in listening to notifications from this gateway
type SensorNotificationWatcher struct {
	// Name is name of the sensor
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
}

// Dispatch protocol contains configuration necessary to dispatch an event to sensor over different communication protocols
type DispatchProtocol struct {
	Type DispatchProtocolType `json:"type" protobuf:"bytes,1,opt,name=type"`

	Http Http `json:"http" protobuf:"bytes,2,opt,name=http"`

	Nats Nats `json:"nats" protobuf:"bytes,3,opt,name=nats"`
}

type Http struct {
	Port string `json:"port" protobuf:"bytes,1,opt,name=port"`
}

type Nats struct {
	URL string `json:"url" protobuf:"bytes,1,opt,name=url"`
}
