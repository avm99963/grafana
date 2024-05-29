// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v0alpha1

import (
	v0alpha1 "github.com/grafana/grafana/pkg/apis/alerting_notifications/v0alpha1"
)

// StatusOperatorStateApplyConfiguration represents an declarative configuration of the StatusOperatorState type for use
// with apply.
type StatusOperatorStateApplyConfiguration struct {
	DescriptiveState *string                            `json:"descriptiveState,omitempty"`
	Details          map[string]string                  `json:"details,omitempty"`
	LastEvaluation   *string                            `json:"lastEvaluation,omitempty"`
	State            *v0alpha1.StatusOperatorStateState `json:"state,omitempty"`
}

// StatusOperatorStateApplyConfiguration constructs an declarative configuration of the StatusOperatorState type for use with
// apply.
func StatusOperatorState() *StatusOperatorStateApplyConfiguration {
	return &StatusOperatorStateApplyConfiguration{}
}

// WithDescriptiveState sets the DescriptiveState field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DescriptiveState field is set to the value of the last call.
func (b *StatusOperatorStateApplyConfiguration) WithDescriptiveState(value string) *StatusOperatorStateApplyConfiguration {
	b.DescriptiveState = &value
	return b
}

// WithDetails puts the entries into the Details field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Details field,
// overwriting an existing map entries in Details field with the same key.
func (b *StatusOperatorStateApplyConfiguration) WithDetails(entries map[string]string) *StatusOperatorStateApplyConfiguration {
	if b.Details == nil && len(entries) > 0 {
		b.Details = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Details[k] = v
	}
	return b
}

// WithLastEvaluation sets the LastEvaluation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastEvaluation field is set to the value of the last call.
func (b *StatusOperatorStateApplyConfiguration) WithLastEvaluation(value string) *StatusOperatorStateApplyConfiguration {
	b.LastEvaluation = &value
	return b
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *StatusOperatorStateApplyConfiguration) WithState(value v0alpha1.StatusOperatorStateState) *StatusOperatorStateApplyConfiguration {
	b.State = &value
	return b
}