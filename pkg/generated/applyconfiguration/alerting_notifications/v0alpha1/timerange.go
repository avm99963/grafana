// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v0alpha1

// TimeRangeApplyConfiguration represents an declarative configuration of the TimeRange type for use
// with apply.
type TimeRangeApplyConfiguration struct {
	EndMinute   *string `json:"endMinute,omitempty"`
	StartMinute *string `json:"startMinute,omitempty"`
}

// TimeRangeApplyConfiguration constructs an declarative configuration of the TimeRange type for use with
// apply.
func TimeRange() *TimeRangeApplyConfiguration {
	return &TimeRangeApplyConfiguration{}
}

// WithEndMinute sets the EndMinute field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndMinute field is set to the value of the last call.
func (b *TimeRangeApplyConfiguration) WithEndMinute(value string) *TimeRangeApplyConfiguration {
	b.EndMinute = &value
	return b
}

// WithStartMinute sets the StartMinute field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StartMinute field is set to the value of the last call.
func (b *TimeRangeApplyConfiguration) WithStartMinute(value string) *TimeRangeApplyConfiguration {
	b.StartMinute = &value
	return b
}