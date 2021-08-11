// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OverdueStateConfig overdue state config
// swagger:model OverdueStateConfig
type OverdueStateConfig struct {

	// auto reevaluation interval days
	AutoReevaluationIntervalDays int32 `json:"autoReevaluationIntervalDays,omitempty"`

	// condition
	Condition *OverdueCondition `json:"condition,omitempty"`

	// external message
	ExternalMessage string `json:"externalMessage,omitempty"`

	// is block changes
	IsBlockChanges bool `json:"isBlockChanges,omitempty"`

	// is clear state
	IsClearState bool `json:"isClearState,omitempty"`

	// is disable entitlement
	IsDisableEntitlement bool `json:"isDisableEntitlement,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// subscription cancellation policy
	// Enum: [END_OF_TERM IMMEDIATE NONE]
	SubscriptionCancellationPolicy OverdueStateConfigSubscriptionCancellationPolicyEnum `json:"subscriptionCancellationPolicy,omitempty"`
}

// Validate validates this overdue state config
func (m *OverdueStateConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionCancellationPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OverdueStateConfig) validateCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.Condition) { // not required
		return nil
	}

	if m.Condition != nil {
		if err := m.Condition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("condition")
			}
			return err
		}
	}

	return nil
}

var overdueStateConfigTypeSubscriptionCancellationPolicyPropEnum []interface{}

func init() {
	var res []OverdueStateConfigSubscriptionCancellationPolicyEnum
	if err := json.Unmarshal([]byte(`["END_OF_TERM","IMMEDIATE","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		overdueStateConfigTypeSubscriptionCancellationPolicyPropEnum = append(overdueStateConfigTypeSubscriptionCancellationPolicyPropEnum, v)
	}
}

type OverdueStateConfigSubscriptionCancellationPolicyEnum string

const (

	// OverdueStateConfigSubscriptionCancellationPolicyENDOFTERM captures enum value "END_OF_TERM"
	OverdueStateConfigSubscriptionCancellationPolicyENDOFTERM OverdueStateConfigSubscriptionCancellationPolicyEnum = "END_OF_TERM"

	// OverdueStateConfigSubscriptionCancellationPolicyIMMEDIATE captures enum value "IMMEDIATE"
	OverdueStateConfigSubscriptionCancellationPolicyIMMEDIATE OverdueStateConfigSubscriptionCancellationPolicyEnum = "IMMEDIATE"

	// OverdueStateConfigSubscriptionCancellationPolicyNONE captures enum value "NONE"
	OverdueStateConfigSubscriptionCancellationPolicyNONE OverdueStateConfigSubscriptionCancellationPolicyEnum = "NONE"
)

var OverdueStateConfigSubscriptionCancellationPolicyEnumValues = []string{
	"END_OF_TERM",
	"IMMEDIATE",
	"NONE",
}

func (e OverdueStateConfigSubscriptionCancellationPolicyEnum) IsValid() bool {
	for _, v := range OverdueStateConfigSubscriptionCancellationPolicyEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *OverdueStateConfig) validateSubscriptionCancellationPolicyEnum(path, location string, value OverdueStateConfigSubscriptionCancellationPolicyEnum) error {
	if err := validate.Enum(path, location, value, overdueStateConfigTypeSubscriptionCancellationPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *OverdueStateConfig) validateSubscriptionCancellationPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionCancellationPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateSubscriptionCancellationPolicyEnum("subscriptionCancellationPolicy", "body", m.SubscriptionCancellationPolicy); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OverdueStateConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OverdueStateConfig) UnmarshalBinary(b []byte) error {
	var res OverdueStateConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
