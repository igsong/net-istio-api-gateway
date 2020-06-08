/*
Copyright 2019 The Knative Authors.

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
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/kmeta"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Route defines HTTP routes which specifies matching rules and destination.
// Basically most of parts are directly translated into VirtualService of Istio,
// but if `knative-serving-route` is specified,
// then all the traffic is routed to the specified Knative route.
type Route struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// RouteSpec holds the desired state of the Route (from the client).
	// +optional
	Spec RouteSpec `json:"spec,omitempty"`

	// RouteStatus communicates the observed state of the Route (from the controller).
	// +optional
	Status RouteStatus `json:"status,omitempty"`
}

// Check that Route can be validated and defaulted.
var _ apis.Validatable = (*Route)(nil)
var _ apis.Defaultable = (*Route)(nil)
var _ kmeta.OwnerRefable = (*Route)(nil)

// RouteSpec holds the desired state of the Route (from the client).
type RouteSpec struct {
	// Hosts specify a list of hostname of rules defined the current route
	Hosts []string `json:"hosts"`
	// Gateway specify a list of gateways of rules
	Gateways []string `json:"gateways"`
	// Http specify HTTP routing rules
	HTTPRoute []HTTPRoute `json:"http"`
}

// HTTPRoute defines a specific routing rule for HTTP
type HTTPRoute struct {
	Name       string             `json:"name,omitempty"`
	Match      []HTTPMatchRequest `json:"match,omitempty"`
	Redirect   HTTPRedirect       `json:"redirect,omitempty"`
	Rewrite    HTTPRewrite        `json:"rewrite,omitempty"`
	Timeout    metav1.Duration    `json:"timeout,omitempty"`
	Retries    HTTPRetry          `json:"retries,omitempty"`
	Headers    Headers            `json:"headers,omitempty"`
	CorsPolicy CorsPolicy         `json:"corsPolicy,omitempty"`
}

type HTTPMatchRequest struct {
	Name            string                 `json:"name,omitempty"`
	URI             StringMatch            `json:"uri,omitempty"`
	Scheme          StringMatch            `json:"scheme,omitempty"`
	Method          StringMatch            `json:"method,omitempty"`
	Authority       StringMatch            `json:"authority,omitempty"`
	Headers         map[string]StringMatch `json:"headers,omitempty"`
	Port            uint32                 `json:"port,omitempty"`
	SourceLabels    map[string]string      `json:"sourceLabels,omitempty"`
	Gateways        []string               `json:"gateways,omitempty"`
	QueryParams     []string               `json:"queryParams,omitempty"`
	IgnoreURICase   bool                   `json:"ignoreUriCase,omitempty"`
	WithoutHeaders  map[string]StringMatch `json:"withoutHeaders,omitempty"`
	SourceNamespace string                 `json:"sourceNamespace,omitempty"`
}

type StringMatch struct {
	Exact  string `json:"exact,omitempty"`
	Prefix string `json:"prefix,omitempty"`
	Regex  string `json:"regex,omitempty"`
}

type HTTPRedirect struct {
	URI          string `json:"uri,omitempty"`
	Authority    string `json:"authority,omitempty"`
	RedirectCode uint32 `json:"redirectCode,omitempty"`
}

type HTTPRewrite struct {
	URI       string `json:"uri,omitempty"`
	Authority string `json:"authority,omitempty"`
}

type HTTPRetry struct {
	Attempts              uint32          `json:"attempts"`
	PerTryTimeout         metav1.Duration `json:"perTryTimeout,omitempty"`
	RetryOn               string          `json:"retryOn,omitempty"`
	RetryRemoteLocalities bool            `json:"retryRemoteLocalities,omitempty"`
}

type HeaderOperations struct {
	Set    map[string]string `json:"set,omitempty"`
	Add    map[string]string `json:"add,omitempty"`
	Remove []string          `json:"remove,omitempty"`
}

type Headers struct {
	Request  HeaderOperations `json:"request,omitempty"`
	Response HeaderOperations `json:"response,omitempty"`
}

type CorsPolicy struct {
	AllowOrigins     []StringMatch   `json:"allowOrigins,omitempty"`
	AllowMethods     []string        `json:"allowMethods,omitempty"`
	AllowHeaders     []string        `json:"allowHeaders,omitempty"`
	ExposeHeaders    []string        `json:"exposeHeaders,omitempty"`
	MaxAge           metav1.Duration `json:"maxAge,omitempty"`
	AllowCredentials bool            `json:"allowCredentials,omitempty"`
}

const (
	// RouteConditionReady indicates if the corresponding route is ready or not.
	RouteConditionReady = apis.ConditionReady
)

// RouteStatus communicates the observed state of the Route (from the controller).
type RouteStatus struct {
	duckv1.Status `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RouteList is a list of Route resources
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Route `json:"items"`
}
