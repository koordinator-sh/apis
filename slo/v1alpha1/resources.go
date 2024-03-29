/*
Copyright 2022 The Koordinator Authors.

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

	schedulingv1alpha1 "github.com/koordinator-sh/apis/scheduling/v1alpha1"
)

type ResourceMap struct {
	corev1.ResourceList `json:"resources,omitempty"`
	Devices             []schedulingv1alpha1.DeviceInfo `json:"devices,omitempty"`
}
