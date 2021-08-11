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

// AuditLog audit log
// swagger:model AuditLog
type AuditLog struct {

	// change date
	// Format: date-time
	ChangeDate strfmt.DateTime `json:"changeDate,omitempty"`

	// change type
	ChangeType string `json:"changeType,omitempty"`

	// changed by
	ChangedBy string `json:"changedBy,omitempty"`

	// comments
	Comments string `json:"comments,omitempty"`

	// history
	History *Entity `json:"history,omitempty"`

	// object Id
	// Format: uuid
	ObjectID strfmt.UUID `json:"objectId,omitempty"`

	// object type
	// Enum: [ACCOUNT ACCOUNT_EMAIL BLOCKING_STATES BUNDLE CUSTOM_FIELD INVOICE PAYMENT TRANSACTION INVOICE_ITEM INVOICE_PAYMENT SUBSCRIPTION SUBSCRIPTION_EVENT SERVICE_BROADCAST PAYMENT_ATTEMPT PAYMENT_METHOD TAG TAG_DEFINITION TENANT TENANT_KVS]
	ObjectType AuditLogObjectTypeEnum `json:"objectType,omitempty"`

	// reason code
	ReasonCode string `json:"reasonCode,omitempty"`

	// user token
	UserToken string `json:"userToken,omitempty"`
}

// Validate validates this audit log
func (m *AuditLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChangeDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditLog) validateChangeDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ChangeDate) { // not required
		return nil
	}

	if err := validate.FormatOf("changeDate", "body", "date-time", m.ChangeDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuditLog) validateHistory(formats strfmt.Registry) error {

	if swag.IsZero(m.History) { // not required
		return nil
	}

	if m.History != nil {
		if err := m.History.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("history")
			}
			return err
		}
	}

	return nil
}

func (m *AuditLog) validateObjectID(formats strfmt.Registry) error {

	if swag.IsZero(m.ObjectID) { // not required
		return nil
	}

	if err := validate.FormatOf("objectId", "body", "uuid", m.ObjectID.String(), formats); err != nil {
		return err
	}

	return nil
}

var auditLogTypeObjectTypePropEnum []interface{}

func init() {
	var res []AuditLogObjectTypeEnum
	if err := json.Unmarshal([]byte(`["ACCOUNT","ACCOUNT_EMAIL","BLOCKING_STATES","BUNDLE","CUSTOM_FIELD","INVOICE","PAYMENT","TRANSACTION","INVOICE_ITEM","INVOICE_PAYMENT","SUBSCRIPTION","SUBSCRIPTION_EVENT","SERVICE_BROADCAST","PAYMENT_ATTEMPT","PAYMENT_METHOD","TAG","TAG_DEFINITION","TENANT","TENANT_KVS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		auditLogTypeObjectTypePropEnum = append(auditLogTypeObjectTypePropEnum, v)
	}
}

type AuditLogObjectTypeEnum string

const (

	// AuditLogObjectTypeACCOUNT captures enum value "ACCOUNT"
	AuditLogObjectTypeACCOUNT AuditLogObjectTypeEnum = "ACCOUNT"

	// AuditLogObjectTypeACCOUNTEMAIL captures enum value "ACCOUNT_EMAIL"
	AuditLogObjectTypeACCOUNTEMAIL AuditLogObjectTypeEnum = "ACCOUNT_EMAIL"

	// AuditLogObjectTypeBLOCKINGSTATES captures enum value "BLOCKING_STATES"
	AuditLogObjectTypeBLOCKINGSTATES AuditLogObjectTypeEnum = "BLOCKING_STATES"

	// AuditLogObjectTypeBUNDLE captures enum value "BUNDLE"
	AuditLogObjectTypeBUNDLE AuditLogObjectTypeEnum = "BUNDLE"

	// AuditLogObjectTypeCUSTOMFIELD captures enum value "CUSTOM_FIELD"
	AuditLogObjectTypeCUSTOMFIELD AuditLogObjectTypeEnum = "CUSTOM_FIELD"

	// AuditLogObjectTypeINVOICE captures enum value "INVOICE"
	AuditLogObjectTypeINVOICE AuditLogObjectTypeEnum = "INVOICE"

	// AuditLogObjectTypePAYMENT captures enum value "PAYMENT"
	AuditLogObjectTypePAYMENT AuditLogObjectTypeEnum = "PAYMENT"

	// AuditLogObjectTypeTRANSACTION captures enum value "TRANSACTION"
	AuditLogObjectTypeTRANSACTION AuditLogObjectTypeEnum = "TRANSACTION"

	// AuditLogObjectTypeINVOICEITEM captures enum value "INVOICE_ITEM"
	AuditLogObjectTypeINVOICEITEM AuditLogObjectTypeEnum = "INVOICE_ITEM"

	// AuditLogObjectTypeINVOICEPAYMENT captures enum value "INVOICE_PAYMENT"
	AuditLogObjectTypeINVOICEPAYMENT AuditLogObjectTypeEnum = "INVOICE_PAYMENT"

	// AuditLogObjectTypeSUBSCRIPTION captures enum value "SUBSCRIPTION"
	AuditLogObjectTypeSUBSCRIPTION AuditLogObjectTypeEnum = "SUBSCRIPTION"

	// AuditLogObjectTypeSUBSCRIPTIONEVENT captures enum value "SUBSCRIPTION_EVENT"
	AuditLogObjectTypeSUBSCRIPTIONEVENT AuditLogObjectTypeEnum = "SUBSCRIPTION_EVENT"

	// AuditLogObjectTypeSERVICEBROADCAST captures enum value "SERVICE_BROADCAST"
	AuditLogObjectTypeSERVICEBROADCAST AuditLogObjectTypeEnum = "SERVICE_BROADCAST"

	// AuditLogObjectTypePAYMENTATTEMPT captures enum value "PAYMENT_ATTEMPT"
	AuditLogObjectTypePAYMENTATTEMPT AuditLogObjectTypeEnum = "PAYMENT_ATTEMPT"

	// AuditLogObjectTypePAYMENTMETHOD captures enum value "PAYMENT_METHOD"
	AuditLogObjectTypePAYMENTMETHOD AuditLogObjectTypeEnum = "PAYMENT_METHOD"

	// AuditLogObjectTypeTAG captures enum value "TAG"
	AuditLogObjectTypeTAG AuditLogObjectTypeEnum = "TAG"

	// AuditLogObjectTypeTAGDEFINITION captures enum value "TAG_DEFINITION"
	AuditLogObjectTypeTAGDEFINITION AuditLogObjectTypeEnum = "TAG_DEFINITION"

	// AuditLogObjectTypeTENANT captures enum value "TENANT"
	AuditLogObjectTypeTENANT AuditLogObjectTypeEnum = "TENANT"

	// AuditLogObjectTypeTENANTKVS captures enum value "TENANT_KVS"
	AuditLogObjectTypeTENANTKVS AuditLogObjectTypeEnum = "TENANT_KVS"
)

var AuditLogObjectTypeEnumValues = []string{
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

func (e AuditLogObjectTypeEnum) IsValid() bool {
	for _, v := range AuditLogObjectTypeEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *AuditLog) validateObjectTypeEnum(path, location string, value AuditLogObjectTypeEnum) error {
	if err := validate.Enum(path, location, value, auditLogTypeObjectTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AuditLog) validateObjectType(formats strfmt.Registry) error {

	if swag.IsZero(m.ObjectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateObjectTypeEnum("objectType", "body", m.ObjectType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLog) UnmarshalBinary(b []byte) error {
	var res AuditLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
