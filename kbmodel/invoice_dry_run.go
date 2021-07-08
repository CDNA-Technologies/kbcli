// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InvoiceDryRun invoice dry run
//
// swagger:model InvoiceDryRun
type InvoiceDryRun struct {

	// billing period
	// Enum: [DAILY WEEKLY BIWEEKLY THIRTY_DAYS SIXTY_DAYS NINETY_DAYS MONTHLY BIMESTRIAL QUARTERLY TRIANNUAL BIANNUAL ANNUAL BIENNIAL NO_BILLING_PERIOD]
	BillingPeriod string `json:"billingPeriod,omitempty"`

	// billing policy
	// Enum: [START_OF_TERM END_OF_TERM IMMEDIATE ILLEGAL]
	BillingPolicy string `json:"billingPolicy,omitempty"`

	// bundle Id
	// Format: uuid
	BundleID strfmt.UUID `json:"bundleId,omitempty"`

	// dry run action
	// Enum: [START_ENTITLEMENT START_BILLING PAUSE_ENTITLEMENT PAUSE_BILLING RESUME_ENTITLEMENT RESUME_BILLING PHASE CHANGE STOP_ENTITLEMENT STOP_BILLING SERVICE_STATE_CHANGE]
	DryRunAction string `json:"dryRunAction,omitempty"`

	// dry run type
	// Enum: [TARGET_DATE UPCOMING_INVOICE SUBSCRIPTION_ACTION]
	DryRunType string `json:"dryRunType,omitempty"`

	// effective date
	// Format: date
	EffectiveDate strfmt.Date `json:"effectiveDate,omitempty"`

	// phase type
	// Enum: [TRIAL DISCOUNT FIXEDTERM EVERGREEN]
	PhaseType string `json:"phaseType,omitempty"`

	// price list name
	PriceListName string `json:"priceListName,omitempty"`

	// price overrides
	PriceOverrides []*PhasePrice `json:"priceOverrides"`

	// product category
	// Enum: [BASE ADD_ON STANDALONE]
	ProductCategory string `json:"productCategory,omitempty"`

	// product name
	ProductName string `json:"productName,omitempty"`

	// subscription Id
	// Format: uuid
	SubscriptionID strfmt.UUID `json:"subscriptionId,omitempty"`
}

// Validate validates this invoice dry run
func (m *InvoiceDryRun) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBundleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDryRunAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDryRunType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhaseType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceOverrides(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var invoiceDryRunTypeBillingPeriodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DAILY","WEEKLY","BIWEEKLY","THIRTY_DAYS","SIXTY_DAYS","NINETY_DAYS","MONTHLY","BIMESTRIAL","QUARTERLY","TRIANNUAL","BIANNUAL","ANNUAL","BIENNIAL","NO_BILLING_PERIOD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceDryRunTypeBillingPeriodPropEnum = append(invoiceDryRunTypeBillingPeriodPropEnum, v)
	}
}

const (

	// InvoiceDryRunBillingPeriodDAILY captures enum value "DAILY"
	InvoiceDryRunBillingPeriodDAILY string = "DAILY"

	// InvoiceDryRunBillingPeriodWEEKLY captures enum value "WEEKLY"
	InvoiceDryRunBillingPeriodWEEKLY string = "WEEKLY"

	// InvoiceDryRunBillingPeriodBIWEEKLY captures enum value "BIWEEKLY"
	InvoiceDryRunBillingPeriodBIWEEKLY string = "BIWEEKLY"

	// InvoiceDryRunBillingPeriodTHIRTYDAYS captures enum value "THIRTY_DAYS"
	InvoiceDryRunBillingPeriodTHIRTYDAYS string = "THIRTY_DAYS"

	// InvoiceDryRunBillingPeriodSIXTYDAYS captures enum value "SIXTY_DAYS"
	InvoiceDryRunBillingPeriodSIXTYDAYS string = "SIXTY_DAYS"

	// InvoiceDryRunBillingPeriodNINETYDAYS captures enum value "NINETY_DAYS"
	InvoiceDryRunBillingPeriodNINETYDAYS string = "NINETY_DAYS"

	// InvoiceDryRunBillingPeriodMONTHLY captures enum value "MONTHLY"
	InvoiceDryRunBillingPeriodMONTHLY string = "MONTHLY"

	// InvoiceDryRunBillingPeriodBIMESTRIAL captures enum value "BIMESTRIAL"
	InvoiceDryRunBillingPeriodBIMESTRIAL string = "BIMESTRIAL"

	// InvoiceDryRunBillingPeriodQUARTERLY captures enum value "QUARTERLY"
	InvoiceDryRunBillingPeriodQUARTERLY string = "QUARTERLY"

	// InvoiceDryRunBillingPeriodTRIANNUAL captures enum value "TRIANNUAL"
	InvoiceDryRunBillingPeriodTRIANNUAL string = "TRIANNUAL"

	// InvoiceDryRunBillingPeriodBIANNUAL captures enum value "BIANNUAL"
	InvoiceDryRunBillingPeriodBIANNUAL string = "BIANNUAL"

	// InvoiceDryRunBillingPeriodANNUAL captures enum value "ANNUAL"
	InvoiceDryRunBillingPeriodANNUAL string = "ANNUAL"

	// InvoiceDryRunBillingPeriodBIENNIAL captures enum value "BIENNIAL"
	InvoiceDryRunBillingPeriodBIENNIAL string = "BIENNIAL"

	// InvoiceDryRunBillingPeriodNOBILLINGPERIOD captures enum value "NO_BILLING_PERIOD"
	InvoiceDryRunBillingPeriodNOBILLINGPERIOD string = "NO_BILLING_PERIOD"
)

// prop value enum
func (m *InvoiceDryRun) validateBillingPeriodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceDryRunTypeBillingPeriodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceDryRun) validateBillingPeriod(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingPeriod) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingPeriodEnum("billingPeriod", "body", m.BillingPeriod); err != nil {
		return err
	}

	return nil
}

var invoiceDryRunTypeBillingPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["START_OF_TERM","END_OF_TERM","IMMEDIATE","ILLEGAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceDryRunTypeBillingPolicyPropEnum = append(invoiceDryRunTypeBillingPolicyPropEnum, v)
	}
}

const (

	// InvoiceDryRunBillingPolicySTARTOFTERM captures enum value "START_OF_TERM"
	InvoiceDryRunBillingPolicySTARTOFTERM string = "START_OF_TERM"

	// InvoiceDryRunBillingPolicyENDOFTERM captures enum value "END_OF_TERM"
	InvoiceDryRunBillingPolicyENDOFTERM string = "END_OF_TERM"

	// InvoiceDryRunBillingPolicyIMMEDIATE captures enum value "IMMEDIATE"
	InvoiceDryRunBillingPolicyIMMEDIATE string = "IMMEDIATE"

	// InvoiceDryRunBillingPolicyILLEGAL captures enum value "ILLEGAL"
	InvoiceDryRunBillingPolicyILLEGAL string = "ILLEGAL"
)

// prop value enum
func (m *InvoiceDryRun) validateBillingPolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceDryRunTypeBillingPolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceDryRun) validateBillingPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingPolicyEnum("billingPolicy", "body", m.BillingPolicy); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceDryRun) validateBundleID(formats strfmt.Registry) error {
	if swag.IsZero(m.BundleID) { // not required
		return nil
	}

	if err := validate.FormatOf("bundleId", "body", "uuid", m.BundleID.String(), formats); err != nil {
		return err
	}

	return nil
}

var invoiceDryRunTypeDryRunActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["START_ENTITLEMENT","START_BILLING","PAUSE_ENTITLEMENT","PAUSE_BILLING","RESUME_ENTITLEMENT","RESUME_BILLING","PHASE","CHANGE","STOP_ENTITLEMENT","STOP_BILLING","SERVICE_STATE_CHANGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceDryRunTypeDryRunActionPropEnum = append(invoiceDryRunTypeDryRunActionPropEnum, v)
	}
}

const (

	// InvoiceDryRunDryRunActionSTARTENTITLEMENT captures enum value "START_ENTITLEMENT"
	InvoiceDryRunDryRunActionSTARTENTITLEMENT string = "START_ENTITLEMENT"

	// InvoiceDryRunDryRunActionSTARTBILLING captures enum value "START_BILLING"
	InvoiceDryRunDryRunActionSTARTBILLING string = "START_BILLING"

	// InvoiceDryRunDryRunActionPAUSEENTITLEMENT captures enum value "PAUSE_ENTITLEMENT"
	InvoiceDryRunDryRunActionPAUSEENTITLEMENT string = "PAUSE_ENTITLEMENT"

	// InvoiceDryRunDryRunActionPAUSEBILLING captures enum value "PAUSE_BILLING"
	InvoiceDryRunDryRunActionPAUSEBILLING string = "PAUSE_BILLING"

	// InvoiceDryRunDryRunActionRESUMEENTITLEMENT captures enum value "RESUME_ENTITLEMENT"
	InvoiceDryRunDryRunActionRESUMEENTITLEMENT string = "RESUME_ENTITLEMENT"

	// InvoiceDryRunDryRunActionRESUMEBILLING captures enum value "RESUME_BILLING"
	InvoiceDryRunDryRunActionRESUMEBILLING string = "RESUME_BILLING"

	// InvoiceDryRunDryRunActionPHASE captures enum value "PHASE"
	InvoiceDryRunDryRunActionPHASE string = "PHASE"

	// InvoiceDryRunDryRunActionCHANGE captures enum value "CHANGE"
	InvoiceDryRunDryRunActionCHANGE string = "CHANGE"

	// InvoiceDryRunDryRunActionSTOPENTITLEMENT captures enum value "STOP_ENTITLEMENT"
	InvoiceDryRunDryRunActionSTOPENTITLEMENT string = "STOP_ENTITLEMENT"

	// InvoiceDryRunDryRunActionSTOPBILLING captures enum value "STOP_BILLING"
	InvoiceDryRunDryRunActionSTOPBILLING string = "STOP_BILLING"

	// InvoiceDryRunDryRunActionSERVICESTATECHANGE captures enum value "SERVICE_STATE_CHANGE"
	InvoiceDryRunDryRunActionSERVICESTATECHANGE string = "SERVICE_STATE_CHANGE"
)

// prop value enum
func (m *InvoiceDryRun) validateDryRunActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceDryRunTypeDryRunActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceDryRun) validateDryRunAction(formats strfmt.Registry) error {
	if swag.IsZero(m.DryRunAction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDryRunActionEnum("dryRunAction", "body", m.DryRunAction); err != nil {
		return err
	}

	return nil
}

var invoiceDryRunTypeDryRunTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TARGET_DATE","UPCOMING_INVOICE","SUBSCRIPTION_ACTION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceDryRunTypeDryRunTypePropEnum = append(invoiceDryRunTypeDryRunTypePropEnum, v)
	}
}

const (

	// InvoiceDryRunDryRunTypeTARGETDATE captures enum value "TARGET_DATE"
	InvoiceDryRunDryRunTypeTARGETDATE string = "TARGET_DATE"

	// InvoiceDryRunDryRunTypeUPCOMINGINVOICE captures enum value "UPCOMING_INVOICE"
	InvoiceDryRunDryRunTypeUPCOMINGINVOICE string = "UPCOMING_INVOICE"

	// InvoiceDryRunDryRunTypeSUBSCRIPTIONACTION captures enum value "SUBSCRIPTION_ACTION"
	InvoiceDryRunDryRunTypeSUBSCRIPTIONACTION string = "SUBSCRIPTION_ACTION"
)

// prop value enum
func (m *InvoiceDryRun) validateDryRunTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceDryRunTypeDryRunTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceDryRun) validateDryRunType(formats strfmt.Registry) error {
	if swag.IsZero(m.DryRunType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDryRunTypeEnum("dryRunType", "body", m.DryRunType); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceDryRun) validateEffectiveDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EffectiveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effectiveDate", "body", "date", m.EffectiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var invoiceDryRunTypePhaseTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TRIAL","DISCOUNT","FIXEDTERM","EVERGREEN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceDryRunTypePhaseTypePropEnum = append(invoiceDryRunTypePhaseTypePropEnum, v)
	}
}

const (

	// InvoiceDryRunPhaseTypeTRIAL captures enum value "TRIAL"
	InvoiceDryRunPhaseTypeTRIAL string = "TRIAL"

	// InvoiceDryRunPhaseTypeDISCOUNT captures enum value "DISCOUNT"
	InvoiceDryRunPhaseTypeDISCOUNT string = "DISCOUNT"

	// InvoiceDryRunPhaseTypeFIXEDTERM captures enum value "FIXEDTERM"
	InvoiceDryRunPhaseTypeFIXEDTERM string = "FIXEDTERM"

	// InvoiceDryRunPhaseTypeEVERGREEN captures enum value "EVERGREEN"
	InvoiceDryRunPhaseTypeEVERGREEN string = "EVERGREEN"
)

// prop value enum
func (m *InvoiceDryRun) validatePhaseTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceDryRunTypePhaseTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceDryRun) validatePhaseType(formats strfmt.Registry) error {
	if swag.IsZero(m.PhaseType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePhaseTypeEnum("phaseType", "body", m.PhaseType); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceDryRun) validatePriceOverrides(formats strfmt.Registry) error {
	if swag.IsZero(m.PriceOverrides) { // not required
		return nil
	}

	for i := 0; i < len(m.PriceOverrides); i++ {
		if swag.IsZero(m.PriceOverrides[i]) { // not required
			continue
		}

		if m.PriceOverrides[i] != nil {
			if err := m.PriceOverrides[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("priceOverrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var invoiceDryRunTypeProductCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BASE","ADD_ON","STANDALONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceDryRunTypeProductCategoryPropEnum = append(invoiceDryRunTypeProductCategoryPropEnum, v)
	}
}

const (

	// InvoiceDryRunProductCategoryBASE captures enum value "BASE"
	InvoiceDryRunProductCategoryBASE string = "BASE"

	// InvoiceDryRunProductCategoryADDON captures enum value "ADD_ON"
	InvoiceDryRunProductCategoryADDON string = "ADD_ON"

	// InvoiceDryRunProductCategorySTANDALONE captures enum value "STANDALONE"
	InvoiceDryRunProductCategorySTANDALONE string = "STANDALONE"
)

// prop value enum
func (m *InvoiceDryRun) validateProductCategoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceDryRunTypeProductCategoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceDryRun) validateProductCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.ProductCategory) { // not required
		return nil
	}

	// value enum
	if err := m.validateProductCategoryEnum("productCategory", "body", m.ProductCategory); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceDryRun) validateSubscriptionID(formats strfmt.Registry) error {
	if swag.IsZero(m.SubscriptionID) { // not required
		return nil
	}

	if err := validate.FormatOf("subscriptionId", "body", "uuid", m.SubscriptionID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this invoice dry run based on the context it is used
func (m *InvoiceDryRun) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePriceOverrides(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceDryRun) contextValidatePriceOverrides(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PriceOverrides); i++ {

		if m.PriceOverrides[i] != nil {
			if err := m.PriceOverrides[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("priceOverrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceDryRun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceDryRun) UnmarshalBinary(b []byte) error {
	var res InvoiceDryRun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
