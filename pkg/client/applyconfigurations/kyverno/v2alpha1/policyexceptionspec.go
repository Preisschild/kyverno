/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v2alpha1

import (
	v2beta1 "github.com/kyverno/kyverno/pkg/client/applyconfigurations/kyverno/v2beta1"
)

// PolicyExceptionSpecApplyConfiguration represents an declarative configuration of the PolicyExceptionSpec type for use
// with apply.
type PolicyExceptionSpecApplyConfiguration struct {
	Background *bool                                     `json:"background,omitempty"`
	Match      *v2beta1.MatchResourcesApplyConfiguration `json:"match,omitempty"`
	Exceptions []ExceptionApplyConfiguration             `json:"exceptions,omitempty"`
}

// PolicyExceptionSpecApplyConfiguration constructs an declarative configuration of the PolicyExceptionSpec type for use with
// apply.
func PolicyExceptionSpec() *PolicyExceptionSpecApplyConfiguration {
	return &PolicyExceptionSpecApplyConfiguration{}
}

// WithBackground sets the Background field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Background field is set to the value of the last call.
func (b *PolicyExceptionSpecApplyConfiguration) WithBackground(value bool) *PolicyExceptionSpecApplyConfiguration {
	b.Background = &value
	return b
}

// WithMatch sets the Match field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Match field is set to the value of the last call.
func (b *PolicyExceptionSpecApplyConfiguration) WithMatch(value *v2beta1.MatchResourcesApplyConfiguration) *PolicyExceptionSpecApplyConfiguration {
	b.Match = value
	return b
}

// WithExceptions adds the given value to the Exceptions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Exceptions field.
func (b *PolicyExceptionSpecApplyConfiguration) WithExceptions(values ...*ExceptionApplyConfiguration) *PolicyExceptionSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithExceptions")
		}
		b.Exceptions = append(b.Exceptions, *values[i])
	}
	return b
}
