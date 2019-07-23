// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomField custom field
// swagger:model CustomField
type CustomField struct {

	// audit logs
	AuditLogs []*AuditLog `json:"auditLogs"`

	// custom field Id
	// Format: uuid
	CustomFieldID strfmt.UUID `json:"customFieldId,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// object Id
	// Format: uuid
	ObjectID strfmt.UUID `json:"objectId,omitempty"`

	// object type
	// Enum: [ACCOUNT ACCOUNT_EMAIL BLOCKING_STATES BUNDLE CUSTOM_FIELD INVOICE PAYMENT TRANSACTION INVOICE_ITEM INVOICE_PAYMENT SUBSCRIPTION SUBSCRIPTION_EVENT SERVICE_BROADCAST PAYMENT_ATTEMPT PAYMENT_METHOD TAG TAG_DEFINITION TENANT TENANT_KVS]
	ObjectType CustomFieldObjectTypeEnum `json:"objectType,omitempty"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this custom field
func (m *CustomField) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomFieldID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomField) validateAuditLogs(formats strfmt.Registry) error {

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

func (m *CustomField) validateCustomFieldID(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomFieldID) { // not required
		return nil
	}

	if err := validate.FormatOf("customFieldId", "body", "uuid", m.CustomFieldID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) validateObjectID(formats strfmt.Registry) error {

	if swag.IsZero(m.ObjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("objectId", "body", "uuid", m.ObjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

var customFieldTypeObjectTypePropEnum []interface{}

func init() {
	var res []CustomFieldObjectTypeEnum
	if err := json.Unmarshal([]byte(`["ACCOUNT","ACCOUNT_EMAIL","BLOCKING_STATES","BUNDLE","CUSTOM_FIELD","INVOICE","PAYMENT","TRANSACTION","INVOICE_ITEM","INVOICE_PAYMENT","SUBSCRIPTION","SUBSCRIPTION_EVENT","SERVICE_BROADCAST","PAYMENT_ATTEMPT","PAYMENT_METHOD","TAG","TAG_DEFINITION","TENANT","TENANT_KVS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customFieldTypeObjectTypePropEnum = append(customFieldTypeObjectTypePropEnum, v)
	}
}

type CustomFieldObjectTypeEnum string

const (

	// CustomFieldObjectTypeACCOUNT captures enum value "ACCOUNT"
	CustomFieldObjectTypeACCOUNT CustomFieldObjectTypeEnum = "ACCOUNT"

	// CustomFieldObjectTypeACCOUNTEMAIL captures enum value "ACCOUNT_EMAIL"
	CustomFieldObjectTypeACCOUNTEMAIL CustomFieldObjectTypeEnum = "ACCOUNT_EMAIL"

	// CustomFieldObjectTypeBLOCKINGSTATES captures enum value "BLOCKING_STATES"
	CustomFieldObjectTypeBLOCKINGSTATES CustomFieldObjectTypeEnum = "BLOCKING_STATES"

	// CustomFieldObjectTypeBUNDLE captures enum value "BUNDLE"
	CustomFieldObjectTypeBUNDLE CustomFieldObjectTypeEnum = "BUNDLE"

	// CustomFieldObjectTypeCUSTOMFIELD captures enum value "CUSTOM_FIELD"
	CustomFieldObjectTypeCUSTOMFIELD CustomFieldObjectTypeEnum = "CUSTOM_FIELD"

	// CustomFieldObjectTypeINVOICE captures enum value "INVOICE"
	CustomFieldObjectTypeINVOICE CustomFieldObjectTypeEnum = "INVOICE"

	// CustomFieldObjectTypePAYMENT captures enum value "PAYMENT"
	CustomFieldObjectTypePAYMENT CustomFieldObjectTypeEnum = "PAYMENT"

	// CustomFieldObjectTypeTRANSACTION captures enum value "TRANSACTION"
	CustomFieldObjectTypeTRANSACTION CustomFieldObjectTypeEnum = "TRANSACTION"

	// CustomFieldObjectTypeINVOICEITEM captures enum value "INVOICE_ITEM"
	CustomFieldObjectTypeINVOICEITEM CustomFieldObjectTypeEnum = "INVOICE_ITEM"

	// CustomFieldObjectTypeINVOICEPAYMENT captures enum value "INVOICE_PAYMENT"
	CustomFieldObjectTypeINVOICEPAYMENT CustomFieldObjectTypeEnum = "INVOICE_PAYMENT"

	// CustomFieldObjectTypeSUBSCRIPTION captures enum value "SUBSCRIPTION"
	CustomFieldObjectTypeSUBSCRIPTION CustomFieldObjectTypeEnum = "SUBSCRIPTION"

	// CustomFieldObjectTypeSUBSCRIPTIONEVENT captures enum value "SUBSCRIPTION_EVENT"
	CustomFieldObjectTypeSUBSCRIPTIONEVENT CustomFieldObjectTypeEnum = "SUBSCRIPTION_EVENT"

	// CustomFieldObjectTypeSERVICEBROADCAST captures enum value "SERVICE_BROADCAST"
	CustomFieldObjectTypeSERVICEBROADCAST CustomFieldObjectTypeEnum = "SERVICE_BROADCAST"

	// CustomFieldObjectTypePAYMENTATTEMPT captures enum value "PAYMENT_ATTEMPT"
	CustomFieldObjectTypePAYMENTATTEMPT CustomFieldObjectTypeEnum = "PAYMENT_ATTEMPT"

	// CustomFieldObjectTypePAYMENTMETHOD captures enum value "PAYMENT_METHOD"
	CustomFieldObjectTypePAYMENTMETHOD CustomFieldObjectTypeEnum = "PAYMENT_METHOD"

	// CustomFieldObjectTypeTAG captures enum value "TAG"
	CustomFieldObjectTypeTAG CustomFieldObjectTypeEnum = "TAG"

	// CustomFieldObjectTypeTAGDEFINITION captures enum value "TAG_DEFINITION"
	CustomFieldObjectTypeTAGDEFINITION CustomFieldObjectTypeEnum = "TAG_DEFINITION"

	// CustomFieldObjectTypeTENANT captures enum value "TENANT"
	CustomFieldObjectTypeTENANT CustomFieldObjectTypeEnum = "TENANT"

	// CustomFieldObjectTypeTENANTKVS captures enum value "TENANT_KVS"
	CustomFieldObjectTypeTENANTKVS CustomFieldObjectTypeEnum = "TENANT_KVS"
)

var CustomFieldObjectTypeEnumValues = []string{
	"ACCOUNT",
	"ACCOUNT_EMAIL",
	"BLOCKING_STATES",
	"BUNDLE",
	"CUSTOM_FIELD",
	"INVOICE",
	"PAYMENT",
	"TRANSACTION",
	"INVOICE_ITEM",
	"INVOICE_PAYMENT",
	"SUBSCRIPTION",
	"SUBSCRIPTION_EVENT",
	"SERVICE_BROADCAST",
	"PAYMENT_ATTEMPT",
	"PAYMENT_METHOD",
	"TAG",
	"TAG_DEFINITION",
	"TENANT",
	"TENANT_KVS",
}

func (e CustomFieldObjectTypeEnum) IsValid() bool {
	for _, v := range CustomFieldObjectTypeEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *CustomField) validateObjectTypeEnum(path, location string, value CustomFieldObjectTypeEnum) error {
	if err := validate.Enum(path, location, value, customFieldTypeObjectTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomField) validateObjectType(formats strfmt.Registry) error {

	if swag.IsZero(m.ObjectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateObjectTypeEnum("objectType", "body", m.ObjectType); err != nil {
		return err
	}

	return nil
}

func (m *CustomField) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomField) UnmarshalBinary(b []byte) error {
	var res CustomField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
