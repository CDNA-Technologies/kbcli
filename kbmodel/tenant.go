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
	"github.com/go-openapi/validate"
)

// Tenant tenant
//
// swagger:model Tenant
type Tenant struct {

	// api key
	// Required: true
	APIKey *string `json:"apiKey"`

	// api secret
	// Required: true
	APISecret *string `json:"apiSecret"`

	// audit logs
	AuditLogs []*AuditLog `json:"auditLogs"`

	// external key
	ExternalKey string `json:"externalKey,omitempty"`

	// tenant Id
	// Format: uuid
	TenantID strfmt.UUID `json:"tenantId,omitempty"`
}

// Validate validates this tenant
func (m *Tenant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAPISecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuditLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tenant) validateAPIKey(formats strfmt.Registry) error {

	if err := validate.Required("apiKey", "body", m.APIKey); err != nil {
		return err
	}

	return nil
}

func (m *Tenant) validateAPISecret(formats strfmt.Registry) error {

	if err := validate.Required("apiSecret", "body", m.APISecret); err != nil {
		return err
	}

	return nil
}

func (m *Tenant) validateAuditLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.AuditLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditLogs); i++ {
		if swag.IsZero(m.AuditLogs[i]) { // not required
			continue
		}

		if m.AuditLogs[i] != nil {
			if err := m.AuditLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Tenant) validateTenantID(formats strfmt.Registry) error {
	if swag.IsZero(m.TenantID) { // not required
		return nil
	}

	if err := validate.FormatOf("tenantId", "body", "uuid", m.TenantID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this tenant based on the context it is used
func (m *Tenant) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tenant) contextValidateAuditLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AuditLogs); i++ {

		if m.AuditLogs[i] != nil {
			if err := m.AuditLogs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Tenant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tenant) UnmarshalBinary(b []byte) error {
	var res Tenant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
