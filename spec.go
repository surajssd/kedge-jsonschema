/*
Copyright 2017 The Kedge Authors All rights reserved.

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

package spec

import (
	api_v1 "k8s.io/client-go/pkg/api/v1"
	ext_v1beta1 "k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

// Define Persistent Volume to use in the app
// kedgeSpec: io.kedge.PersistentVolume
type PersistentVolume struct {
	// Data from the kubernetes persistent volume claim spec
	// k8s: io.k8s.api.core.v1.PersistentVolumeClaimSpec
	api_v1.PersistentVolumeClaimSpec `json:",inline"`
	// Name of the persistent Volume Claim
	// +optional
	Name string `json:"name"`
	// Size of persistent volume
	Size string `json:"size"`
}

// Define ConfigMap to be created
// kedgeSpec: io.kedge.ConfigMap
type ConfigMapMod struct {
	// Name of the configMap
	// +optional
	Name string `json:"name,omitempty"`
	// Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'
	Data map[string]string `json:"data,omitempty"`
}

// Define Kubernetes service
// kedgeSpec: io.kedge.ServiceSpec
type ServiceSpecMod struct {
	// k8s: io.k8s.api.core.v1.ServiceSpec
	api_v1.ServiceSpec `json:",inline"`
	// Name of the service
	// +optional
	Name string `json:"name,omitempty"`
	// The list of ports that are exposed by this service. More info:
	// https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	// ref: io.kedge.ServicePort
	Ports []ServicePortMod `json:"ports"`
}

// Define service port
// kedgeSpec: io.kedge.ServicePort
type ServicePortMod struct {
	// k8s: io.k8s.api.core.v1.ServicePort
	api_v1.ServicePort `json:",inline"`
	// Host to create ingress automatically. Endpoint allows specifying an
	// ingress resource in the format '<Host>/<Path>'
	// +optional
	Endpoint string `json:"endpoint"`
}

// Create ingress object
// kedgeSpec: io.kedge.IngressSpec
type IngressSpecMod struct {
	// Name of the ingress
	// +optional
	Name string `json:"name"`
	// k8s: io.k8s.api.extensions.v1beta1.IngressSpec
	ext_v1beta1.IngressSpec `json:",inline"`
}

// EnvFromSource represents the source of a set of ConfigMaps
// kedgeSpec: io.kedge.EnvFromSource
type EnvFromSource struct {
	//The ConfigMap to select from
	// ref: io.kedge.ConfigMapEnvSource
	ConfigMapRef *ConfigMapEnvSource `json:"configMapRef,omitempty"`
}

// ConfigMapEnvSource selects a ConfigMap to populate the environment
// variables with. The contents of the target ConfigMap's Data field will
// represent the key-value pairs as environment variables.
// kedgeSpec: io.kedge.ConfigMapEnvSource
type ConfigMapEnvSource struct {
	// Name of the referent.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name string `json:"name,omitempty"`
}

// A single application container that you want to run within a pod.
// kedgeSpec: io.kedge.ContainerSpec
type Container struct {
	// One common definitions for 'livenessProbe' and 'readinessProbe'
	// this allows to have only one place to define both probes (if they are the same)
	// Periodic probe of container liveness and readiness. Container will be restarted
	// if the probe fails. Cannot be updated. More info:
	// https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// ref: io.k8s.api.core.v1.Probe
	// +optional
	Health *api_v1.Probe `json:"health,omitempty"`
	// EnvFrom defines the collection of EnvFromSource to inject into containers.
	// ref: io.kedge.EnvFromSource
	// +optional
	EnvFrom []EnvFromSource `json:"envFrom,omitempty"`
	// k8s: io.k8s.api.core.v1.Container
	api_v1.Container `json:",inline"`
}

type PodSpecMod struct {
	// List of containers belonging to the pod. Containers cannot currently be
	// added or removed. There must be at least one container in a Pod. Cannot be updated.
	// ref: io.kedge.ContainerSpec
	Containers []Container `json:"containers,omitempty"`
	// k8s: io.k8s.api.core.v1.PodSpec
	api_v1.PodSpec `json:",inline"`
}

// AppSpec is a description of an app
// kedgeSpec: io.kedge.AppSpec
type App struct {
	// Name of the micro-service
	Name string `json:"name"`
	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and services. More info: http://kubernetes.io/docs/user-guide/labels
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	// List of persistentVolumes that should be mounted on the pod.
	// ref: io.kedge.PersistentVolume
	// +optional
	PersistentVolumes []PersistentVolume `json:"persistentVolumes,omitempty"`
	// List of configMaps
	// ref: io.kedge.ConfigMap
	// +optional
	ConfigMaps []ConfigMapMod `json:"configMaps,omitempty"`
	// List of Kubernetes Services
	// ref: io.kedge.ServiceSpec
	// +optional
	Services []ServiceSpecMod `json:"services,omitempty"`
	// List of Kubernetes Ingress
	// ref: io.kedge.IngressSpec
	// +optional
	Ingresses []IngressSpecMod `json:"ingresses,omitempty"`

	PodSpecMod `json:",inline"`
	// k8s: io.k8s.api.extensions.v1beta1.DeploymentSpec
	ext_v1beta1.DeploymentSpec `json:",inline"`
}