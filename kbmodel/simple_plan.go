// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SimplePlan simple plan
//
// swagger:model SimplePlan
type SimplePlan struct {

	// amount
	Amount float64 `json:"amount,omitempty"`

	// available base products
	AvailableBaseProducts []string `json:"availableBaseProducts"`

	// billing period
	// Enum: [DAILY WEEKLY BIWEEKLY THIRTY_DAYS SIXTY_DAYS NINETY_DAYS MONTHLY BIMESTRIAL QUARTERLY TRIANNUAL BIANNUAL ANNUAL BIENNIAL NO_BILLING_PERIOD]
	BillingPeriod string `json:"billingPeriod,omitempty"`

	// currency
	// Enum: [AED AFN ALL AMD ANG AOA ARS AUD AWG AZN BAM BBD BDT BGN BHD BIF BMD BND BOB BRL BSD BTN BWP BYR BZD CAD CDF CHF CLP CNY COP CRC CUC CUP CVE CZK DJF DKK DOP DZD EGP ERN ETB EUR FJD FKP GBP GEL GGP GHS GIP GMD GNF GTQ GYD HKD HNL HRK HTG HUF IDR ILS IMP INR IQD IRR ISK JEP JMD JOD JPY KES KGS KHR KMF KPW KRW KWD KYD KZT LAK LBP LKR LRD LSL LTL LVL LYD MAD MDL MGA MKD MMK MNT MOP MRO MUR MVR MWK MXN MYR MZN NAD NGN NIO NOK NPR NZD OMR PAB PEN PGK PHP PKR PLN PYG QAR RON RSD RUB RWF SAR SBD SCR SDG SEK SGD SHP SLL SOS SPL SRD STD SVC SYP SZL THB TJS TMT TND TOP TRY TTD TVD TWD TZS UAH UGX USD UYU UZS VEF VND VUV WST XAF XCD XDR XOF XPF YER ZAR ZMW ZWD BTC]
	Currency string `json:"currency,omitempty"`

	// plan Id
	PlanID string `json:"planId,omitempty"`

	// product category
	// Enum: [BASE ADD_ON STANDALONE]
	ProductCategory string `json:"productCategory,omitempty"`

	// product name
	ProductName string `json:"productName,omitempty"`

	// trial length
	TrialLength int32 `json:"trialLength,omitempty"`

	// trial time unit
	// Enum: [DAYS WEEKS MONTHS YEARS UNLIMITED]
	TrialTimeUnit string `json:"trialTimeUnit,omitempty"`
}

// Validate validates this simple plan
func (m *SimplePlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrialTimeUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var simplePlanTypeBillingPeriodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DAILY","WEEKLY","BIWEEKLY","THIRTY_DAYS","SIXTY_DAYS","NINETY_DAYS","MONTHLY","BIMESTRIAL","QUARTERLY","TRIANNUAL","BIANNUAL","ANNUAL","BIENNIAL","NO_BILLING_PERIOD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simplePlanTypeBillingPeriodPropEnum = append(simplePlanTypeBillingPeriodPropEnum, v)
	}
}

const (

	// SimplePlanBillingPeriodDAILY captures enum value "DAILY"
	SimplePlanBillingPeriodDAILY string = "DAILY"

	// SimplePlanBillingPeriodWEEKLY captures enum value "WEEKLY"
	SimplePlanBillingPeriodWEEKLY string = "WEEKLY"

	// SimplePlanBillingPeriodBIWEEKLY captures enum value "BIWEEKLY"
	SimplePlanBillingPeriodBIWEEKLY string = "BIWEEKLY"

	// SimplePlanBillingPeriodTHIRTYDAYS captures enum value "THIRTY_DAYS"
	SimplePlanBillingPeriodTHIRTYDAYS string = "THIRTY_DAYS"

	// SimplePlanBillingPeriodSIXTYDAYS captures enum value "SIXTY_DAYS"
	SimplePlanBillingPeriodSIXTYDAYS string = "SIXTY_DAYS"

	// SimplePlanBillingPeriodNINETYDAYS captures enum value "NINETY_DAYS"
	SimplePlanBillingPeriodNINETYDAYS string = "NINETY_DAYS"

	// SimplePlanBillingPeriodMONTHLY captures enum value "MONTHLY"
	SimplePlanBillingPeriodMONTHLY string = "MONTHLY"

	// SimplePlanBillingPeriodBIMESTRIAL captures enum value "BIMESTRIAL"
	SimplePlanBillingPeriodBIMESTRIAL string = "BIMESTRIAL"

	// SimplePlanBillingPeriodQUARTERLY captures enum value "QUARTERLY"
	SimplePlanBillingPeriodQUARTERLY string = "QUARTERLY"

	// SimplePlanBillingPeriodTRIANNUAL captures enum value "TRIANNUAL"
	SimplePlanBillingPeriodTRIANNUAL string = "TRIANNUAL"

	// SimplePlanBillingPeriodBIANNUAL captures enum value "BIANNUAL"
	SimplePlanBillingPeriodBIANNUAL string = "BIANNUAL"

	// SimplePlanBillingPeriodANNUAL captures enum value "ANNUAL"
	SimplePlanBillingPeriodANNUAL string = "ANNUAL"

	// SimplePlanBillingPeriodBIENNIAL captures enum value "BIENNIAL"
	SimplePlanBillingPeriodBIENNIAL string = "BIENNIAL"

	// SimplePlanBillingPeriodNOBILLINGPERIOD captures enum value "NO_BILLING_PERIOD"
	SimplePlanBillingPeriodNOBILLINGPERIOD string = "NO_BILLING_PERIOD"
)

// prop value enum
func (m *SimplePlan) validateBillingPeriodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, simplePlanTypeBillingPeriodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SimplePlan) validateBillingPeriod(formats strfmt.Registry) error {
	if swag.IsZero(m.BillingPeriod) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingPeriodEnum("billingPeriod", "body", m.BillingPeriod); err != nil {
		return err
	}

	return nil
}

var simplePlanTypeCurrencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYR","BZD","CAD","CDF","CHF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","EUR","FJD","FKP","GBP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SOS","SPL","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TVD","TWD","TZS","UAH","UGX","USD","UYU","UZS","VEF","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMW","ZWD","BTC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simplePlanTypeCurrencyPropEnum = append(simplePlanTypeCurrencyPropEnum, v)
	}
}

const (

	// SimplePlanCurrencyAED captures enum value "AED"
	SimplePlanCurrencyAED string = "AED"

	// SimplePlanCurrencyAFN captures enum value "AFN"
	SimplePlanCurrencyAFN string = "AFN"

	// SimplePlanCurrencyALL captures enum value "ALL"
	SimplePlanCurrencyALL string = "ALL"

	// SimplePlanCurrencyAMD captures enum value "AMD"
	SimplePlanCurrencyAMD string = "AMD"

	// SimplePlanCurrencyANG captures enum value "ANG"
	SimplePlanCurrencyANG string = "ANG"

	// SimplePlanCurrencyAOA captures enum value "AOA"
	SimplePlanCurrencyAOA string = "AOA"

	// SimplePlanCurrencyARS captures enum value "ARS"
	SimplePlanCurrencyARS string = "ARS"

	// SimplePlanCurrencyAUD captures enum value "AUD"
	SimplePlanCurrencyAUD string = "AUD"

	// SimplePlanCurrencyAWG captures enum value "AWG"
	SimplePlanCurrencyAWG string = "AWG"

	// SimplePlanCurrencyAZN captures enum value "AZN"
	SimplePlanCurrencyAZN string = "AZN"

	// SimplePlanCurrencyBAM captures enum value "BAM"
	SimplePlanCurrencyBAM string = "BAM"

	// SimplePlanCurrencyBBD captures enum value "BBD"
	SimplePlanCurrencyBBD string = "BBD"

	// SimplePlanCurrencyBDT captures enum value "BDT"
	SimplePlanCurrencyBDT string = "BDT"

	// SimplePlanCurrencyBGN captures enum value "BGN"
	SimplePlanCurrencyBGN string = "BGN"

	// SimplePlanCurrencyBHD captures enum value "BHD"
	SimplePlanCurrencyBHD string = "BHD"

	// SimplePlanCurrencyBIF captures enum value "BIF"
	SimplePlanCurrencyBIF string = "BIF"

	// SimplePlanCurrencyBMD captures enum value "BMD"
	SimplePlanCurrencyBMD string = "BMD"

	// SimplePlanCurrencyBND captures enum value "BND"
	SimplePlanCurrencyBND string = "BND"

	// SimplePlanCurrencyBOB captures enum value "BOB"
	SimplePlanCurrencyBOB string = "BOB"

	// SimplePlanCurrencyBRL captures enum value "BRL"
	SimplePlanCurrencyBRL string = "BRL"

	// SimplePlanCurrencyBSD captures enum value "BSD"
	SimplePlanCurrencyBSD string = "BSD"

	// SimplePlanCurrencyBTN captures enum value "BTN"
	SimplePlanCurrencyBTN string = "BTN"

	// SimplePlanCurrencyBWP captures enum value "BWP"
	SimplePlanCurrencyBWP string = "BWP"

	// SimplePlanCurrencyBYR captures enum value "BYR"
	SimplePlanCurrencyBYR string = "BYR"

	// SimplePlanCurrencyBZD captures enum value "BZD"
	SimplePlanCurrencyBZD string = "BZD"

	// SimplePlanCurrencyCAD captures enum value "CAD"
	SimplePlanCurrencyCAD string = "CAD"

	// SimplePlanCurrencyCDF captures enum value "CDF"
	SimplePlanCurrencyCDF string = "CDF"

	// SimplePlanCurrencyCHF captures enum value "CHF"
	SimplePlanCurrencyCHF string = "CHF"

	// SimplePlanCurrencyCLP captures enum value "CLP"
	SimplePlanCurrencyCLP string = "CLP"

	// SimplePlanCurrencyCNY captures enum value "CNY"
	SimplePlanCurrencyCNY string = "CNY"

	// SimplePlanCurrencyCOP captures enum value "COP"
	SimplePlanCurrencyCOP string = "COP"

	// SimplePlanCurrencyCRC captures enum value "CRC"
	SimplePlanCurrencyCRC string = "CRC"

	// SimplePlanCurrencyCUC captures enum value "CUC"
	SimplePlanCurrencyCUC string = "CUC"

	// SimplePlanCurrencyCUP captures enum value "CUP"
	SimplePlanCurrencyCUP string = "CUP"

	// SimplePlanCurrencyCVE captures enum value "CVE"
	SimplePlanCurrencyCVE string = "CVE"

	// SimplePlanCurrencyCZK captures enum value "CZK"
	SimplePlanCurrencyCZK string = "CZK"

	// SimplePlanCurrencyDJF captures enum value "DJF"
	SimplePlanCurrencyDJF string = "DJF"

	// SimplePlanCurrencyDKK captures enum value "DKK"
	SimplePlanCurrencyDKK string = "DKK"

	// SimplePlanCurrencyDOP captures enum value "DOP"
	SimplePlanCurrencyDOP string = "DOP"

	// SimplePlanCurrencyDZD captures enum value "DZD"
	SimplePlanCurrencyDZD string = "DZD"

	// SimplePlanCurrencyEGP captures enum value "EGP"
	SimplePlanCurrencyEGP string = "EGP"

	// SimplePlanCurrencyERN captures enum value "ERN"
	SimplePlanCurrencyERN string = "ERN"

	// SimplePlanCurrencyETB captures enum value "ETB"
	SimplePlanCurrencyETB string = "ETB"

	// SimplePlanCurrencyEUR captures enum value "EUR"
	SimplePlanCurrencyEUR string = "EUR"

	// SimplePlanCurrencyFJD captures enum value "FJD"
	SimplePlanCurrencyFJD string = "FJD"

	// SimplePlanCurrencyFKP captures enum value "FKP"
	SimplePlanCurrencyFKP string = "FKP"

	// SimplePlanCurrencyGBP captures enum value "GBP"
	SimplePlanCurrencyGBP string = "GBP"

	// SimplePlanCurrencyGEL captures enum value "GEL"
	SimplePlanCurrencyGEL string = "GEL"

	// SimplePlanCurrencyGGP captures enum value "GGP"
	SimplePlanCurrencyGGP string = "GGP"

	// SimplePlanCurrencyGHS captures enum value "GHS"
	SimplePlanCurrencyGHS string = "GHS"

	// SimplePlanCurrencyGIP captures enum value "GIP"
	SimplePlanCurrencyGIP string = "GIP"

	// SimplePlanCurrencyGMD captures enum value "GMD"
	SimplePlanCurrencyGMD string = "GMD"

	// SimplePlanCurrencyGNF captures enum value "GNF"
	SimplePlanCurrencyGNF string = "GNF"

	// SimplePlanCurrencyGTQ captures enum value "GTQ"
	SimplePlanCurrencyGTQ string = "GTQ"

	// SimplePlanCurrencyGYD captures enum value "GYD"
	SimplePlanCurrencyGYD string = "GYD"

	// SimplePlanCurrencyHKD captures enum value "HKD"
	SimplePlanCurrencyHKD string = "HKD"

	// SimplePlanCurrencyHNL captures enum value "HNL"
	SimplePlanCurrencyHNL string = "HNL"

	// SimplePlanCurrencyHRK captures enum value "HRK"
	SimplePlanCurrencyHRK string = "HRK"

	// SimplePlanCurrencyHTG captures enum value "HTG"
	SimplePlanCurrencyHTG string = "HTG"

	// SimplePlanCurrencyHUF captures enum value "HUF"
	SimplePlanCurrencyHUF string = "HUF"

	// SimplePlanCurrencyIDR captures enum value "IDR"
	SimplePlanCurrencyIDR string = "IDR"

	// SimplePlanCurrencyILS captures enum value "ILS"
	SimplePlanCurrencyILS string = "ILS"

	// SimplePlanCurrencyIMP captures enum value "IMP"
	SimplePlanCurrencyIMP string = "IMP"

	// SimplePlanCurrencyINR captures enum value "INR"
	SimplePlanCurrencyINR string = "INR"

	// SimplePlanCurrencyIQD captures enum value "IQD"
	SimplePlanCurrencyIQD string = "IQD"

	// SimplePlanCurrencyIRR captures enum value "IRR"
	SimplePlanCurrencyIRR string = "IRR"

	// SimplePlanCurrencyISK captures enum value "ISK"
	SimplePlanCurrencyISK string = "ISK"

	// SimplePlanCurrencyJEP captures enum value "JEP"
	SimplePlanCurrencyJEP string = "JEP"

	// SimplePlanCurrencyJMD captures enum value "JMD"
	SimplePlanCurrencyJMD string = "JMD"

	// SimplePlanCurrencyJOD captures enum value "JOD"
	SimplePlanCurrencyJOD string = "JOD"

	// SimplePlanCurrencyJPY captures enum value "JPY"
	SimplePlanCurrencyJPY string = "JPY"

	// SimplePlanCurrencyKES captures enum value "KES"
	SimplePlanCurrencyKES string = "KES"

	// SimplePlanCurrencyKGS captures enum value "KGS"
	SimplePlanCurrencyKGS string = "KGS"

	// SimplePlanCurrencyKHR captures enum value "KHR"
	SimplePlanCurrencyKHR string = "KHR"

	// SimplePlanCurrencyKMF captures enum value "KMF"
	SimplePlanCurrencyKMF string = "KMF"

	// SimplePlanCurrencyKPW captures enum value "KPW"
	SimplePlanCurrencyKPW string = "KPW"

	// SimplePlanCurrencyKRW captures enum value "KRW"
	SimplePlanCurrencyKRW string = "KRW"

	// SimplePlanCurrencyKWD captures enum value "KWD"
	SimplePlanCurrencyKWD string = "KWD"

	// SimplePlanCurrencyKYD captures enum value "KYD"
	SimplePlanCurrencyKYD string = "KYD"

	// SimplePlanCurrencyKZT captures enum value "KZT"
	SimplePlanCurrencyKZT string = "KZT"

	// SimplePlanCurrencyLAK captures enum value "LAK"
	SimplePlanCurrencyLAK string = "LAK"

	// SimplePlanCurrencyLBP captures enum value "LBP"
	SimplePlanCurrencyLBP string = "LBP"

	// SimplePlanCurrencyLKR captures enum value "LKR"
	SimplePlanCurrencyLKR string = "LKR"

	// SimplePlanCurrencyLRD captures enum value "LRD"
	SimplePlanCurrencyLRD string = "LRD"

	// SimplePlanCurrencyLSL captures enum value "LSL"
	SimplePlanCurrencyLSL string = "LSL"

	// SimplePlanCurrencyLTL captures enum value "LTL"
	SimplePlanCurrencyLTL string = "LTL"

	// SimplePlanCurrencyLVL captures enum value "LVL"
	SimplePlanCurrencyLVL string = "LVL"

	// SimplePlanCurrencyLYD captures enum value "LYD"
	SimplePlanCurrencyLYD string = "LYD"

	// SimplePlanCurrencyMAD captures enum value "MAD"
	SimplePlanCurrencyMAD string = "MAD"

	// SimplePlanCurrencyMDL captures enum value "MDL"
	SimplePlanCurrencyMDL string = "MDL"

	// SimplePlanCurrencyMGA captures enum value "MGA"
	SimplePlanCurrencyMGA string = "MGA"

	// SimplePlanCurrencyMKD captures enum value "MKD"
	SimplePlanCurrencyMKD string = "MKD"

	// SimplePlanCurrencyMMK captures enum value "MMK"
	SimplePlanCurrencyMMK string = "MMK"

	// SimplePlanCurrencyMNT captures enum value "MNT"
	SimplePlanCurrencyMNT string = "MNT"

	// SimplePlanCurrencyMOP captures enum value "MOP"
	SimplePlanCurrencyMOP string = "MOP"

	// SimplePlanCurrencyMRO captures enum value "MRO"
	SimplePlanCurrencyMRO string = "MRO"

	// SimplePlanCurrencyMUR captures enum value "MUR"
	SimplePlanCurrencyMUR string = "MUR"

	// SimplePlanCurrencyMVR captures enum value "MVR"
	SimplePlanCurrencyMVR string = "MVR"

	// SimplePlanCurrencyMWK captures enum value "MWK"
	SimplePlanCurrencyMWK string = "MWK"

	// SimplePlanCurrencyMXN captures enum value "MXN"
	SimplePlanCurrencyMXN string = "MXN"

	// SimplePlanCurrencyMYR captures enum value "MYR"
	SimplePlanCurrencyMYR string = "MYR"

	// SimplePlanCurrencyMZN captures enum value "MZN"
	SimplePlanCurrencyMZN string = "MZN"

	// SimplePlanCurrencyNAD captures enum value "NAD"
	SimplePlanCurrencyNAD string = "NAD"

	// SimplePlanCurrencyNGN captures enum value "NGN"
	SimplePlanCurrencyNGN string = "NGN"

	// SimplePlanCurrencyNIO captures enum value "NIO"
	SimplePlanCurrencyNIO string = "NIO"

	// SimplePlanCurrencyNOK captures enum value "NOK"
	SimplePlanCurrencyNOK string = "NOK"

	// SimplePlanCurrencyNPR captures enum value "NPR"
	SimplePlanCurrencyNPR string = "NPR"

	// SimplePlanCurrencyNZD captures enum value "NZD"
	SimplePlanCurrencyNZD string = "NZD"

	// SimplePlanCurrencyOMR captures enum value "OMR"
	SimplePlanCurrencyOMR string = "OMR"

	// SimplePlanCurrencyPAB captures enum value "PAB"
	SimplePlanCurrencyPAB string = "PAB"

	// SimplePlanCurrencyPEN captures enum value "PEN"
	SimplePlanCurrencyPEN string = "PEN"

	// SimplePlanCurrencyPGK captures enum value "PGK"
	SimplePlanCurrencyPGK string = "PGK"

	// SimplePlanCurrencyPHP captures enum value "PHP"
	SimplePlanCurrencyPHP string = "PHP"

	// SimplePlanCurrencyPKR captures enum value "PKR"
	SimplePlanCurrencyPKR string = "PKR"

	// SimplePlanCurrencyPLN captures enum value "PLN"
	SimplePlanCurrencyPLN string = "PLN"

	// SimplePlanCurrencyPYG captures enum value "PYG"
	SimplePlanCurrencyPYG string = "PYG"

	// SimplePlanCurrencyQAR captures enum value "QAR"
	SimplePlanCurrencyQAR string = "QAR"

	// SimplePlanCurrencyRON captures enum value "RON"
	SimplePlanCurrencyRON string = "RON"

	// SimplePlanCurrencyRSD captures enum value "RSD"
	SimplePlanCurrencyRSD string = "RSD"

	// SimplePlanCurrencyRUB captures enum value "RUB"
	SimplePlanCurrencyRUB string = "RUB"

	// SimplePlanCurrencyRWF captures enum value "RWF"
	SimplePlanCurrencyRWF string = "RWF"

	// SimplePlanCurrencySAR captures enum value "SAR"
	SimplePlanCurrencySAR string = "SAR"

	// SimplePlanCurrencySBD captures enum value "SBD"
	SimplePlanCurrencySBD string = "SBD"

	// SimplePlanCurrencySCR captures enum value "SCR"
	SimplePlanCurrencySCR string = "SCR"

	// SimplePlanCurrencySDG captures enum value "SDG"
	SimplePlanCurrencySDG string = "SDG"

	// SimplePlanCurrencySEK captures enum value "SEK"
	SimplePlanCurrencySEK string = "SEK"

	// SimplePlanCurrencySGD captures enum value "SGD"
	SimplePlanCurrencySGD string = "SGD"

	// SimplePlanCurrencySHP captures enum value "SHP"
	SimplePlanCurrencySHP string = "SHP"

	// SimplePlanCurrencySLL captures enum value "SLL"
	SimplePlanCurrencySLL string = "SLL"

	// SimplePlanCurrencySOS captures enum value "SOS"
	SimplePlanCurrencySOS string = "SOS"

	// SimplePlanCurrencySPL captures enum value "SPL"
	SimplePlanCurrencySPL string = "SPL"

	// SimplePlanCurrencySRD captures enum value "SRD"
	SimplePlanCurrencySRD string = "SRD"

	// SimplePlanCurrencySTD captures enum value "STD"
	SimplePlanCurrencySTD string = "STD"

	// SimplePlanCurrencySVC captures enum value "SVC"
	SimplePlanCurrencySVC string = "SVC"

	// SimplePlanCurrencySYP captures enum value "SYP"
	SimplePlanCurrencySYP string = "SYP"

	// SimplePlanCurrencySZL captures enum value "SZL"
	SimplePlanCurrencySZL string = "SZL"

	// SimplePlanCurrencyTHB captures enum value "THB"
	SimplePlanCurrencyTHB string = "THB"

	// SimplePlanCurrencyTJS captures enum value "TJS"
	SimplePlanCurrencyTJS string = "TJS"

	// SimplePlanCurrencyTMT captures enum value "TMT"
	SimplePlanCurrencyTMT string = "TMT"

	// SimplePlanCurrencyTND captures enum value "TND"
	SimplePlanCurrencyTND string = "TND"

	// SimplePlanCurrencyTOP captures enum value "TOP"
	SimplePlanCurrencyTOP string = "TOP"

	// SimplePlanCurrencyTRY captures enum value "TRY"
	SimplePlanCurrencyTRY string = "TRY"

	// SimplePlanCurrencyTTD captures enum value "TTD"
	SimplePlanCurrencyTTD string = "TTD"

	// SimplePlanCurrencyTVD captures enum value "TVD"
	SimplePlanCurrencyTVD string = "TVD"

	// SimplePlanCurrencyTWD captures enum value "TWD"
	SimplePlanCurrencyTWD string = "TWD"

	// SimplePlanCurrencyTZS captures enum value "TZS"
	SimplePlanCurrencyTZS string = "TZS"

	// SimplePlanCurrencyUAH captures enum value "UAH"
	SimplePlanCurrencyUAH string = "UAH"

	// SimplePlanCurrencyUGX captures enum value "UGX"
	SimplePlanCurrencyUGX string = "UGX"

	// SimplePlanCurrencyUSD captures enum value "USD"
	SimplePlanCurrencyUSD string = "USD"

	// SimplePlanCurrencyUYU captures enum value "UYU"
	SimplePlanCurrencyUYU string = "UYU"

	// SimplePlanCurrencyUZS captures enum value "UZS"
	SimplePlanCurrencyUZS string = "UZS"

	// SimplePlanCurrencyVEF captures enum value "VEF"
	SimplePlanCurrencyVEF string = "VEF"

	// SimplePlanCurrencyVND captures enum value "VND"
	SimplePlanCurrencyVND string = "VND"

	// SimplePlanCurrencyVUV captures enum value "VUV"
	SimplePlanCurrencyVUV string = "VUV"

	// SimplePlanCurrencyWST captures enum value "WST"
	SimplePlanCurrencyWST string = "WST"

	// SimplePlanCurrencyXAF captures enum value "XAF"
	SimplePlanCurrencyXAF string = "XAF"

	// SimplePlanCurrencyXCD captures enum value "XCD"
	SimplePlanCurrencyXCD string = "XCD"

	// SimplePlanCurrencyXDR captures enum value "XDR"
	SimplePlanCurrencyXDR string = "XDR"

	// SimplePlanCurrencyXOF captures enum value "XOF"
	SimplePlanCurrencyXOF string = "XOF"

	// SimplePlanCurrencyXPF captures enum value "XPF"
	SimplePlanCurrencyXPF string = "XPF"

	// SimplePlanCurrencyYER captures enum value "YER"
	SimplePlanCurrencyYER string = "YER"

	// SimplePlanCurrencyZAR captures enum value "ZAR"
	SimplePlanCurrencyZAR string = "ZAR"

	// SimplePlanCurrencyZMW captures enum value "ZMW"
	SimplePlanCurrencyZMW string = "ZMW"

	// SimplePlanCurrencyZWD captures enum value "ZWD"
	SimplePlanCurrencyZWD string = "ZWD"

	// SimplePlanCurrencyBTC captures enum value "BTC"
	SimplePlanCurrencyBTC string = "BTC"
)

// prop value enum
func (m *SimplePlan) validateCurrencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, simplePlanTypeCurrencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SimplePlan) validateCurrency(formats strfmt.Registry) error {
	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	// value enum
	if err := m.validateCurrencyEnum("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

var simplePlanTypeProductCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BASE","ADD_ON","STANDALONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simplePlanTypeProductCategoryPropEnum = append(simplePlanTypeProductCategoryPropEnum, v)
	}
}

const (

	// SimplePlanProductCategoryBASE captures enum value "BASE"
	SimplePlanProductCategoryBASE string = "BASE"

	// SimplePlanProductCategoryADDON captures enum value "ADD_ON"
	SimplePlanProductCategoryADDON string = "ADD_ON"

	// SimplePlanProductCategorySTANDALONE captures enum value "STANDALONE"
	SimplePlanProductCategorySTANDALONE string = "STANDALONE"
)

// prop value enum
func (m *SimplePlan) validateProductCategoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, simplePlanTypeProductCategoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SimplePlan) validateProductCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.ProductCategory) { // not required
		return nil
	}

	// value enum
	if err := m.validateProductCategoryEnum("productCategory", "body", m.ProductCategory); err != nil {
		return err
	}

	return nil
}

var simplePlanTypeTrialTimeUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DAYS","WEEKS","MONTHS","YEARS","UNLIMITED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simplePlanTypeTrialTimeUnitPropEnum = append(simplePlanTypeTrialTimeUnitPropEnum, v)
	}
}

const (

	// SimplePlanTrialTimeUnitDAYS captures enum value "DAYS"
	SimplePlanTrialTimeUnitDAYS string = "DAYS"

	// SimplePlanTrialTimeUnitWEEKS captures enum value "WEEKS"
	SimplePlanTrialTimeUnitWEEKS string = "WEEKS"

	// SimplePlanTrialTimeUnitMONTHS captures enum value "MONTHS"
	SimplePlanTrialTimeUnitMONTHS string = "MONTHS"

	// SimplePlanTrialTimeUnitYEARS captures enum value "YEARS"
	SimplePlanTrialTimeUnitYEARS string = "YEARS"

	// SimplePlanTrialTimeUnitUNLIMITED captures enum value "UNLIMITED"
	SimplePlanTrialTimeUnitUNLIMITED string = "UNLIMITED"
)

// prop value enum
func (m *SimplePlan) validateTrialTimeUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, simplePlanTypeTrialTimeUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SimplePlan) validateTrialTimeUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.TrialTimeUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validateTrialTimeUnitEnum("trialTimeUnit", "body", m.TrialTimeUnit); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this simple plan based on context it is used
func (m *SimplePlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SimplePlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimplePlan) UnmarshalBinary(b []byte) error {
	var res SimplePlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
