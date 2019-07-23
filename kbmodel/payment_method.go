// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PaymentMethod payment method
// swagger:model PaymentMethod
type PaymentMethod struct {

	// account Id
	// Format: uuid
	AccountID strfmt.UUID `json:"accountId,omitempty"`

	// audit logs
	AuditLogs []*AuditLog `json:"auditLogs"`

	// external key
	ExternalKey string `json:"externalKey,omitempty"`

	// is default
	IsDefault *bool `json:"isDefault,omitempty"`

	// payment method Id
	// Format: uuid
	PaymentMethodID strfmt.UUID `json:"paymentMethodId,omitempty"`

	// plugin info
	PluginInfo *PaymentMethodPluginDetail `json:"pluginInfo,omitempty"`

	// plugin name
	PluginName string `json:"pluginName,omitempty"`
}

// Validate validates this payment method
func (m *PaymentMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuditLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentMethodID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePluginInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentMethod) validateAccountID(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountID) { // not required
		return nil
	}

	if err := validate.FormatOf("accountId", "body", "uuid", m.AccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentMethod) validateAuditLogs(formats strfmt.Registry) error {

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

func (m *PaymentMethod) validatePaymentMethodID(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethodID) { // not required
		return nil
	}

	if err := validate.FormatOf("paymentMethodId", "body", "uuid", m.PaymentMethodID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentMethod) validatePluginInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.PluginInfo) { // not required
		return nil
	}

	if m.PluginInfo != nil {
		if err := m.PluginInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pluginInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentMethod) UnmarshalBinary(b []byte) error {
	var res PaymentMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
