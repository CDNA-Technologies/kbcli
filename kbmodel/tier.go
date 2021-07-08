// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Tier tier
//
// swagger:model Tier
type Tier struct {

	// blocks
	Blocks []*TieredBlock `json:"blocks"`

	// fixed price
	FixedPrice []*Price `json:"fixedPrice"`

	// limits
	Limits []*Limit `json:"limits"`

	// recurring price
	RecurringPrice []*Price `json:"recurringPrice"`
}

// Validate validates this tier
func (m *Tier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlocks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFixedPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecurringPrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tier) validateBlocks(formats strfmt.Registry) error {
	if swag.IsZero(m.Blocks) { // not required
		return nil
	}

	for i := 0; i < len(m.Blocks); i++ {
		if swag.IsZero(m.Blocks[i]) { // not required
			continue
		}

		if m.Blocks[i] != nil {
			if err := m.Blocks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("blocks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Tier) validateFixedPrice(formats strfmt.Registry) error {
	if swag.IsZero(m.FixedPrice) { // not required
		return nil
	}

	for i := 0; i < len(m.FixedPrice); i++ {
		if swag.IsZero(m.FixedPrice[i]) { // not required
			continue
		}

		if m.FixedPrice[i] != nil {
			if err := m.FixedPrice[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fixedPrice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Tier) validateLimits(formats strfmt.Registry) error {
	if swag.IsZero(m.Limits) { // not required
		return nil
	}

	for i := 0; i < len(m.Limits); i++ {
		if swag.IsZero(m.Limits[i]) { // not required
			continue
		}

		if m.Limits[i] != nil {
			if err := m.Limits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("limits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Tier) validateRecurringPrice(formats strfmt.Registry) error {
	if swag.IsZero(m.RecurringPrice) { // not required
		return nil
	}

	for i := 0; i < len(m.RecurringPrice); i++ {
		if swag.IsZero(m.RecurringPrice[i]) { // not required
			continue
		}

		if m.RecurringPrice[i] != nil {
			if err := m.RecurringPrice[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recurringPrice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tier based on the context it is used
func (m *Tier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBlocks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFixedPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecurringPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tier) contextValidateBlocks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Blocks); i++ {

		if m.Blocks[i] != nil {
			if err := m.Blocks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("blocks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Tier) contextValidateFixedPrice(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FixedPrice); i++ {

		if m.FixedPrice[i] != nil {
			if err := m.FixedPrice[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fixedPrice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Tier) contextValidateLimits(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Limits); i++ {

		if m.Limits[i] != nil {
			if err := m.Limits[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("limits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Tier) contextValidateRecurringPrice(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RecurringPrice); i++ {

		if m.RecurringPrice[i] != nil {
			if err := m.RecurringPrice[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recurringPrice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Tier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tier) UnmarshalBinary(b []byte) error {
	var res Tier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
