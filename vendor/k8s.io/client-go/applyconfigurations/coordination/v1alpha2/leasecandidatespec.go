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

package v1alpha2

import (
	coordinationv1 "k8s.io/api/coordination/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LeaseCandidateSpecApplyConfiguration represents a declarative configuration of the LeaseCandidateSpec type for use
// with apply.
type LeaseCandidateSpecApplyConfiguration struct {
	LeaseName        *string                                  `json:"leaseName,omitempty"`
	PingTime         *v1.MicroTime                            `json:"pingTime,omitempty"`
	RenewTime        *v1.MicroTime                            `json:"renewTime,omitempty"`
	BinaryVersion    *string                                  `json:"binaryVersion,omitempty"`
	EmulationVersion *string                                  `json:"emulationVersion,omitempty"`
	Strategy         *coordinationv1.CoordinatedLeaseStrategy `json:"strategy,omitempty"`
}

// LeaseCandidateSpecApplyConfiguration constructs a declarative configuration of the LeaseCandidateSpec type for use with
// apply.
func LeaseCandidateSpec() *LeaseCandidateSpecApplyConfiguration {
	return &LeaseCandidateSpecApplyConfiguration{}
}

// WithLeaseName sets the LeaseName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LeaseName field is set to the value of the last call.
func (b *LeaseCandidateSpecApplyConfiguration) WithLeaseName(value string) *LeaseCandidateSpecApplyConfiguration {
	b.LeaseName = &value
	return b
}

// WithPingTime sets the PingTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PingTime field is set to the value of the last call.
func (b *LeaseCandidateSpecApplyConfiguration) WithPingTime(value v1.MicroTime) *LeaseCandidateSpecApplyConfiguration {
	b.PingTime = &value
	return b
}

// WithRenewTime sets the RenewTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RenewTime field is set to the value of the last call.
func (b *LeaseCandidateSpecApplyConfiguration) WithRenewTime(value v1.MicroTime) *LeaseCandidateSpecApplyConfiguration {
	b.RenewTime = &value
	return b
}

// WithBinaryVersion sets the BinaryVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BinaryVersion field is set to the value of the last call.
func (b *LeaseCandidateSpecApplyConfiguration) WithBinaryVersion(value string) *LeaseCandidateSpecApplyConfiguration {
	b.BinaryVersion = &value
	return b
}

// WithEmulationVersion sets the EmulationVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EmulationVersion field is set to the value of the last call.
func (b *LeaseCandidateSpecApplyConfiguration) WithEmulationVersion(value string) *LeaseCandidateSpecApplyConfiguration {
	b.EmulationVersion = &value
	return b
}

// WithStrategy sets the Strategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Strategy field is set to the value of the last call.
func (b *LeaseCandidateSpecApplyConfiguration) WithStrategy(value coordinationv1.CoordinatedLeaseStrategy) *LeaseCandidateSpecApplyConfiguration {
	b.Strategy = &value
	return b
}
