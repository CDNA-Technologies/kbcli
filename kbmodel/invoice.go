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

// Invoice invoice
//
// swagger:model Invoice
type Invoice struct {

	// account Id
	// Format: uuid
	AccountID strfmt.UUID `json:"accountId,omitempty"`

	// amount
	Amount float64 `json:"amount,omitempty"`

	// audit logs
	AuditLogs []*AuditLog `json:"auditLogs"`

	// balance
	Balance float64 `json:"balance,omitempty"`

	// bundle keys
	BundleKeys string `json:"bundleKeys,omitempty"`

	// credit adj
	CreditAdj float64 `json:"creditAdj,omitempty"`

	// credits
	Credits []*InvoiceItem `json:"credits"`

	// currency
	// Enum: [AED AFN ALL AMD ANG AOA ARS AUD AWG AZN BAM BBD BDT BGN BHD BIF BMD BND BOB BRL BSD BTN BWP BYR BZD CAD CDF CHF CLP CNY COP CRC CUC CUP CVE CZK DJF DKK DOP DZD EGP ERN ETB EUR FJD FKP GBP GEL GGP GHS GIP GMD GNF GTQ GYD HKD HNL HRK HTG HUF IDR ILS IMP INR IQD IRR ISK JEP JMD JOD JPY KES KGS KHR KMF KPW KRW KWD KYD KZT LAK LBP LKR LRD LSL LTL LVL LYD MAD MDL MGA MKD MMK MNT MOP MRO MUR MVR MWK MXN MYR MZN NAD NGN NIO NOK NPR NZD OMR PAB PEN PGK PHP PKR PLN PYG QAR RON RSD RUB RWF SAR SBD SCR SDG SEK SGD SHP SLL SOS SPL SRD STD SVC SYP SZL THB TJS TMT TND TOP TRY TTD TVD TWD TZS UAH UGX USD UYU UZS VEF VND VUV WST XAF XCD XDR XOF XPF YER ZAR ZMW ZWD BTC]
	Currency string `json:"currency,omitempty"`

	// invoice date
	// Format: date
	InvoiceDate strfmt.Date `json:"invoiceDate,omitempty"`

	// invoice Id
	// Format: uuid
	InvoiceID strfmt.UUID `json:"invoiceId,omitempty"`

	// invoice number
	InvoiceNumber string `json:"invoiceNumber,omitempty"`

	// is parent invoice
	IsParentInvoice bool `json:"isParentInvoice,omitempty"`

	// items
	Items []*InvoiceItem `json:"items"`

	// parent account Id
	// Format: uuid
	ParentAccountID strfmt.UUID `json:"parentAccountId,omitempty"`

	// parent invoice Id
	// Format: uuid
	ParentInvoiceID strfmt.UUID `json:"parentInvoiceId,omitempty"`

	// refund adj
	RefundAdj float64 `json:"refundAdj,omitempty"`

	// status
	// Enum: [DRAFT COMMITTED VOID]
	Status string `json:"status,omitempty"`

	// target date
	// Format: date
	TargetDate strfmt.Date `json:"targetDate,omitempty"`

	// tracking ids
	TrackingIds []string `json:"trackingIds"`
}

// Validate validates this invoice
func (m *Invoice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuditLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentInvoiceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Invoice) validateAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountID) { // not required
		return nil
	}

	if err := validate.FormatOf("accountId", "body", "uuid", m.AccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateAuditLogs(formats strfmt.Registry) error {
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

func (m *Invoice) validateCredits(formats strfmt.Registry) error {
	if swag.IsZero(m.Credits) { // not required
		return nil
	}

	for i := 0; i < len(m.Credits); i++ {
		if swag.IsZero(m.Credits[i]) { // not required
			continue
		}

		if m.Credits[i] != nil {
			if err := m.Credits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("credits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var invoiceTypeCurrencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYR","BZD","CAD","CDF","CHF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","EUR","FJD","FKP","GBP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SOS","SPL","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TVD","TWD","TZS","UAH","UGX","USD","UYU","UZS","VEF","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMW","ZWD","BTC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceTypeCurrencyPropEnum = append(invoiceTypeCurrencyPropEnum, v)
	}
}

const (

	// InvoiceCurrencyAED captures enum value "AED"
	InvoiceCurrencyAED string = "AED"

	// InvoiceCurrencyAFN captures enum value "AFN"
	InvoiceCurrencyAFN string = "AFN"

	// InvoiceCurrencyALL captures enum value "ALL"
	InvoiceCurrencyALL string = "ALL"

	// InvoiceCurrencyAMD captures enum value "AMD"
	InvoiceCurrencyAMD string = "AMD"

	// InvoiceCurrencyANG captures enum value "ANG"
	InvoiceCurrencyANG string = "ANG"

	// InvoiceCurrencyAOA captures enum value "AOA"
	InvoiceCurrencyAOA string = "AOA"

	// InvoiceCurrencyARS captures enum value "ARS"
	InvoiceCurrencyARS string = "ARS"

	// InvoiceCurrencyAUD captures enum value "AUD"
	InvoiceCurrencyAUD string = "AUD"

	// InvoiceCurrencyAWG captures enum value "AWG"
	InvoiceCurrencyAWG string = "AWG"

	// InvoiceCurrencyAZN captures enum value "AZN"
	InvoiceCurrencyAZN string = "AZN"

	// InvoiceCurrencyBAM captures enum value "BAM"
	InvoiceCurrencyBAM string = "BAM"

	// InvoiceCurrencyBBD captures enum value "BBD"
	InvoiceCurrencyBBD string = "BBD"

	// InvoiceCurrencyBDT captures enum value "BDT"
	InvoiceCurrencyBDT string = "BDT"

	// InvoiceCurrencyBGN captures enum value "BGN"
	InvoiceCurrencyBGN string = "BGN"

	// InvoiceCurrencyBHD captures enum value "BHD"
	InvoiceCurrencyBHD string = "BHD"

	// InvoiceCurrencyBIF captures enum value "BIF"
	InvoiceCurrencyBIF string = "BIF"

	// InvoiceCurrencyBMD captures enum value "BMD"
	InvoiceCurrencyBMD string = "BMD"

	// InvoiceCurrencyBND captures enum value "BND"
	InvoiceCurrencyBND string = "BND"

	// InvoiceCurrencyBOB captures enum value "BOB"
	InvoiceCurrencyBOB string = "BOB"

	// InvoiceCurrencyBRL captures enum value "BRL"
	InvoiceCurrencyBRL string = "BRL"

	// InvoiceCurrencyBSD captures enum value "BSD"
	InvoiceCurrencyBSD string = "BSD"

	// InvoiceCurrencyBTN captures enum value "BTN"
	InvoiceCurrencyBTN string = "BTN"

	// InvoiceCurrencyBWP captures enum value "BWP"
	InvoiceCurrencyBWP string = "BWP"

	// InvoiceCurrencyBYR captures enum value "BYR"
	InvoiceCurrencyBYR string = "BYR"

	// InvoiceCurrencyBZD captures enum value "BZD"
	InvoiceCurrencyBZD string = "BZD"

	// InvoiceCurrencyCAD captures enum value "CAD"
	InvoiceCurrencyCAD string = "CAD"

	// InvoiceCurrencyCDF captures enum value "CDF"
	InvoiceCurrencyCDF string = "CDF"

	// InvoiceCurrencyCHF captures enum value "CHF"
	InvoiceCurrencyCHF string = "CHF"

	// InvoiceCurrencyCLP captures enum value "CLP"
	InvoiceCurrencyCLP string = "CLP"

	// InvoiceCurrencyCNY captures enum value "CNY"
	InvoiceCurrencyCNY string = "CNY"

	// InvoiceCurrencyCOP captures enum value "COP"
	InvoiceCurrencyCOP string = "COP"

	// InvoiceCurrencyCRC captures enum value "CRC"
	InvoiceCurrencyCRC string = "CRC"

	// InvoiceCurrencyCUC captures enum value "CUC"
	InvoiceCurrencyCUC string = "CUC"

	// InvoiceCurrencyCUP captures enum value "CUP"
	InvoiceCurrencyCUP string = "CUP"

	// InvoiceCurrencyCVE captures enum value "CVE"
	InvoiceCurrencyCVE string = "CVE"

	// InvoiceCurrencyCZK captures enum value "CZK"
	InvoiceCurrencyCZK string = "CZK"

	// InvoiceCurrencyDJF captures enum value "DJF"
	InvoiceCurrencyDJF string = "DJF"

	// InvoiceCurrencyDKK captures enum value "DKK"
	InvoiceCurrencyDKK string = "DKK"

	// InvoiceCurrencyDOP captures enum value "DOP"
	InvoiceCurrencyDOP string = "DOP"

	// InvoiceCurrencyDZD captures enum value "DZD"
	InvoiceCurrencyDZD string = "DZD"

	// InvoiceCurrencyEGP captures enum value "EGP"
	InvoiceCurrencyEGP string = "EGP"

	// InvoiceCurrencyERN captures enum value "ERN"
	InvoiceCurrencyERN string = "ERN"

	// InvoiceCurrencyETB captures enum value "ETB"
	InvoiceCurrencyETB string = "ETB"

	// InvoiceCurrencyEUR captures enum value "EUR"
	InvoiceCurrencyEUR string = "EUR"

	// InvoiceCurrencyFJD captures enum value "FJD"
	InvoiceCurrencyFJD string = "FJD"

	// InvoiceCurrencyFKP captures enum value "FKP"
	InvoiceCurrencyFKP string = "FKP"

	// InvoiceCurrencyGBP captures enum value "GBP"
	InvoiceCurrencyGBP string = "GBP"

	// InvoiceCurrencyGEL captures enum value "GEL"
	InvoiceCurrencyGEL string = "GEL"

	// InvoiceCurrencyGGP captures enum value "GGP"
	InvoiceCurrencyGGP string = "GGP"

	// InvoiceCurrencyGHS captures enum value "GHS"
	InvoiceCurrencyGHS string = "GHS"

	// InvoiceCurrencyGIP captures enum value "GIP"
	InvoiceCurrencyGIP string = "GIP"

	// InvoiceCurrencyGMD captures enum value "GMD"
	InvoiceCurrencyGMD string = "GMD"

	// InvoiceCurrencyGNF captures enum value "GNF"
	InvoiceCurrencyGNF string = "GNF"

	// InvoiceCurrencyGTQ captures enum value "GTQ"
	InvoiceCurrencyGTQ string = "GTQ"

	// InvoiceCurrencyGYD captures enum value "GYD"
	InvoiceCurrencyGYD string = "GYD"

	// InvoiceCurrencyHKD captures enum value "HKD"
	InvoiceCurrencyHKD string = "HKD"

	// InvoiceCurrencyHNL captures enum value "HNL"
	InvoiceCurrencyHNL string = "HNL"

	// InvoiceCurrencyHRK captures enum value "HRK"
	InvoiceCurrencyHRK string = "HRK"

	// InvoiceCurrencyHTG captures enum value "HTG"
	InvoiceCurrencyHTG string = "HTG"

	// InvoiceCurrencyHUF captures enum value "HUF"
	InvoiceCurrencyHUF string = "HUF"

	// InvoiceCurrencyIDR captures enum value "IDR"
	InvoiceCurrencyIDR string = "IDR"

	// InvoiceCurrencyILS captures enum value "ILS"
	InvoiceCurrencyILS string = "ILS"

	// InvoiceCurrencyIMP captures enum value "IMP"
	InvoiceCurrencyIMP string = "IMP"

	// InvoiceCurrencyINR captures enum value "INR"
	InvoiceCurrencyINR string = "INR"

	// InvoiceCurrencyIQD captures enum value "IQD"
	InvoiceCurrencyIQD string = "IQD"

	// InvoiceCurrencyIRR captures enum value "IRR"
	InvoiceCurrencyIRR string = "IRR"

	// InvoiceCurrencyISK captures enum value "ISK"
	InvoiceCurrencyISK string = "ISK"

	// InvoiceCurrencyJEP captures enum value "JEP"
	InvoiceCurrencyJEP string = "JEP"

	// InvoiceCurrencyJMD captures enum value "JMD"
	InvoiceCurrencyJMD string = "JMD"

	// InvoiceCurrencyJOD captures enum value "JOD"
	InvoiceCurrencyJOD string = "JOD"

	// InvoiceCurrencyJPY captures enum value "JPY"
	InvoiceCurrencyJPY string = "JPY"

	// InvoiceCurrencyKES captures enum value "KES"
	InvoiceCurrencyKES string = "KES"

	// InvoiceCurrencyKGS captures enum value "KGS"
	InvoiceCurrencyKGS string = "KGS"

	// InvoiceCurrencyKHR captures enum value "KHR"
	InvoiceCurrencyKHR string = "KHR"

	// InvoiceCurrencyKMF captures enum value "KMF"
	InvoiceCurrencyKMF string = "KMF"

	// InvoiceCurrencyKPW captures enum value "KPW"
	InvoiceCurrencyKPW string = "KPW"

	// InvoiceCurrencyKRW captures enum value "KRW"
	InvoiceCurrencyKRW string = "KRW"

	// InvoiceCurrencyKWD captures enum value "KWD"
	InvoiceCurrencyKWD string = "KWD"

	// InvoiceCurrencyKYD captures enum value "KYD"
	InvoiceCurrencyKYD string = "KYD"

	// InvoiceCurrencyKZT captures enum value "KZT"
	InvoiceCurrencyKZT string = "KZT"

	// InvoiceCurrencyLAK captures enum value "LAK"
	InvoiceCurrencyLAK string = "LAK"

	// InvoiceCurrencyLBP captures enum value "LBP"
	InvoiceCurrencyLBP string = "LBP"

	// InvoiceCurrencyLKR captures enum value "LKR"
	InvoiceCurrencyLKR string = "LKR"

	// InvoiceCurrencyLRD captures enum value "LRD"
	InvoiceCurrencyLRD string = "LRD"

	// InvoiceCurrencyLSL captures enum value "LSL"
	InvoiceCurrencyLSL string = "LSL"

	// InvoiceCurrencyLTL captures enum value "LTL"
	InvoiceCurrencyLTL string = "LTL"

	// InvoiceCurrencyLVL captures enum value "LVL"
	InvoiceCurrencyLVL string = "LVL"

	// InvoiceCurrencyLYD captures enum value "LYD"
	InvoiceCurrencyLYD string = "LYD"

	// InvoiceCurrencyMAD captures enum value "MAD"
	InvoiceCurrencyMAD string = "MAD"

	// InvoiceCurrencyMDL captures enum value "MDL"
	InvoiceCurrencyMDL string = "MDL"

	// InvoiceCurrencyMGA captures enum value "MGA"
	InvoiceCurrencyMGA string = "MGA"

	// InvoiceCurrencyMKD captures enum value "MKD"
	InvoiceCurrencyMKD string = "MKD"

	// InvoiceCurrencyMMK captures enum value "MMK"
	InvoiceCurrencyMMK string = "MMK"

	// InvoiceCurrencyMNT captures enum value "MNT"
	InvoiceCurrencyMNT string = "MNT"

	// InvoiceCurrencyMOP captures enum value "MOP"
	InvoiceCurrencyMOP string = "MOP"

	// InvoiceCurrencyMRO captures enum value "MRO"
	InvoiceCurrencyMRO string = "MRO"

	// InvoiceCurrencyMUR captures enum value "MUR"
	InvoiceCurrencyMUR string = "MUR"

	// InvoiceCurrencyMVR captures enum value "MVR"
	InvoiceCurrencyMVR string = "MVR"

	// InvoiceCurrencyMWK captures enum value "MWK"
	InvoiceCurrencyMWK string = "MWK"

	// InvoiceCurrencyMXN captures enum value "MXN"
	InvoiceCurrencyMXN string = "MXN"

	// InvoiceCurrencyMYR captures enum value "MYR"
	InvoiceCurrencyMYR string = "MYR"

	// InvoiceCurrencyMZN captures enum value "MZN"
	InvoiceCurrencyMZN string = "MZN"

	// InvoiceCurrencyNAD captures enum value "NAD"
	InvoiceCurrencyNAD string = "NAD"

	// InvoiceCurrencyNGN captures enum value "NGN"
	InvoiceCurrencyNGN string = "NGN"

	// InvoiceCurrencyNIO captures enum value "NIO"
	InvoiceCurrencyNIO string = "NIO"

	// InvoiceCurrencyNOK captures enum value "NOK"
	InvoiceCurrencyNOK string = "NOK"

	// InvoiceCurrencyNPR captures enum value "NPR"
	InvoiceCurrencyNPR string = "NPR"

	// InvoiceCurrencyNZD captures enum value "NZD"
	InvoiceCurrencyNZD string = "NZD"

	// InvoiceCurrencyOMR captures enum value "OMR"
	InvoiceCurrencyOMR string = "OMR"

	// InvoiceCurrencyPAB captures enum value "PAB"
	InvoiceCurrencyPAB string = "PAB"

	// InvoiceCurrencyPEN captures enum value "PEN"
	InvoiceCurrencyPEN string = "PEN"

	// InvoiceCurrencyPGK captures enum value "PGK"
	InvoiceCurrencyPGK string = "PGK"

	// InvoiceCurrencyPHP captures enum value "PHP"
	InvoiceCurrencyPHP string = "PHP"

	// InvoiceCurrencyPKR captures enum value "PKR"
	InvoiceCurrencyPKR string = "PKR"

	// InvoiceCurrencyPLN captures enum value "PLN"
	InvoiceCurrencyPLN string = "PLN"

	// InvoiceCurrencyPYG captures enum value "PYG"
	InvoiceCurrencyPYG string = "PYG"

	// InvoiceCurrencyQAR captures enum value "QAR"
	InvoiceCurrencyQAR string = "QAR"

	// InvoiceCurrencyRON captures enum value "RON"
	InvoiceCurrencyRON string = "RON"

	// InvoiceCurrencyRSD captures enum value "RSD"
	InvoiceCurrencyRSD string = "RSD"

	// InvoiceCurrencyRUB captures enum value "RUB"
	InvoiceCurrencyRUB string = "RUB"

	// InvoiceCurrencyRWF captures enum value "RWF"
	InvoiceCurrencyRWF string = "RWF"

	// InvoiceCurrencySAR captures enum value "SAR"
	InvoiceCurrencySAR string = "SAR"

	// InvoiceCurrencySBD captures enum value "SBD"
	InvoiceCurrencySBD string = "SBD"

	// InvoiceCurrencySCR captures enum value "SCR"
	InvoiceCurrencySCR string = "SCR"

	// InvoiceCurrencySDG captures enum value "SDG"
	InvoiceCurrencySDG string = "SDG"

	// InvoiceCurrencySEK captures enum value "SEK"
	InvoiceCurrencySEK string = "SEK"

	// InvoiceCurrencySGD captures enum value "SGD"
	InvoiceCurrencySGD string = "SGD"

	// InvoiceCurrencySHP captures enum value "SHP"
	InvoiceCurrencySHP string = "SHP"

	// InvoiceCurrencySLL captures enum value "SLL"
	InvoiceCurrencySLL string = "SLL"

	// InvoiceCurrencySOS captures enum value "SOS"
	InvoiceCurrencySOS string = "SOS"

	// InvoiceCurrencySPL captures enum value "SPL"
	InvoiceCurrencySPL string = "SPL"

	// InvoiceCurrencySRD captures enum value "SRD"
	InvoiceCurrencySRD string = "SRD"

	// InvoiceCurrencySTD captures enum value "STD"
	InvoiceCurrencySTD string = "STD"

	// InvoiceCurrencySVC captures enum value "SVC"
	InvoiceCurrencySVC string = "SVC"

	// InvoiceCurrencySYP captures enum value "SYP"
	InvoiceCurrencySYP string = "SYP"

	// InvoiceCurrencySZL captures enum value "SZL"
	InvoiceCurrencySZL string = "SZL"

	// InvoiceCurrencyTHB captures enum value "THB"
	InvoiceCurrencyTHB string = "THB"

	// InvoiceCurrencyTJS captures enum value "TJS"
	InvoiceCurrencyTJS string = "TJS"

	// InvoiceCurrencyTMT captures enum value "TMT"
	InvoiceCurrencyTMT string = "TMT"

	// InvoiceCurrencyTND captures enum value "TND"
	InvoiceCurrencyTND string = "TND"

	// InvoiceCurrencyTOP captures enum value "TOP"
	InvoiceCurrencyTOP string = "TOP"

	// InvoiceCurrencyTRY captures enum value "TRY"
	InvoiceCurrencyTRY string = "TRY"

	// InvoiceCurrencyTTD captures enum value "TTD"
	InvoiceCurrencyTTD string = "TTD"

	// InvoiceCurrencyTVD captures enum value "TVD"
	InvoiceCurrencyTVD string = "TVD"

	// InvoiceCurrencyTWD captures enum value "TWD"
	InvoiceCurrencyTWD string = "TWD"

	// InvoiceCurrencyTZS captures enum value "TZS"
	InvoiceCurrencyTZS string = "TZS"

	// InvoiceCurrencyUAH captures enum value "UAH"
	InvoiceCurrencyUAH string = "UAH"

	// InvoiceCurrencyUGX captures enum value "UGX"
	InvoiceCurrencyUGX string = "UGX"

	// InvoiceCurrencyUSD captures enum value "USD"
	InvoiceCurrencyUSD string = "USD"

	// InvoiceCurrencyUYU captures enum value "UYU"
	InvoiceCurrencyUYU string = "UYU"

	// InvoiceCurrencyUZS captures enum value "UZS"
	InvoiceCurrencyUZS string = "UZS"

	// InvoiceCurrencyVEF captures enum value "VEF"
	InvoiceCurrencyVEF string = "VEF"

	// InvoiceCurrencyVND captures enum value "VND"
	InvoiceCurrencyVND string = "VND"

	// InvoiceCurrencyVUV captures enum value "VUV"
	InvoiceCurrencyVUV string = "VUV"

	// InvoiceCurrencyWST captures enum value "WST"
	InvoiceCurrencyWST string = "WST"

	// InvoiceCurrencyXAF captures enum value "XAF"
	InvoiceCurrencyXAF string = "XAF"

	// InvoiceCurrencyXCD captures enum value "XCD"
	InvoiceCurrencyXCD string = "XCD"

	// InvoiceCurrencyXDR captures enum value "XDR"
	InvoiceCurrencyXDR string = "XDR"

	// InvoiceCurrencyXOF captures enum value "XOF"
	InvoiceCurrencyXOF string = "XOF"

	// InvoiceCurrencyXPF captures enum value "XPF"
	InvoiceCurrencyXPF string = "XPF"

	// InvoiceCurrencyYER captures enum value "YER"
	InvoiceCurrencyYER string = "YER"

	// InvoiceCurrencyZAR captures enum value "ZAR"
	InvoiceCurrencyZAR string = "ZAR"

	// InvoiceCurrencyZMW captures enum value "ZMW"
	InvoiceCurrencyZMW string = "ZMW"

	// InvoiceCurrencyZWD captures enum value "ZWD"
	InvoiceCurrencyZWD string = "ZWD"

	// InvoiceCurrencyBTC captures enum value "BTC"
	InvoiceCurrencyBTC string = "BTC"
)

// prop value enum
func (m *Invoice) validateCurrencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceTypeCurrencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Invoice) validateCurrency(formats strfmt.Registry) error {
	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	// value enum
	if err := m.validateCurrencyEnum("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateInvoiceDate(formats strfmt.Registry) error {
	if swag.IsZero(m.InvoiceDate) { // not required
		return nil
	}

	if err := validate.FormatOf("invoiceDate", "body", "date", m.InvoiceDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateInvoiceID(formats strfmt.Registry) error {
	if swag.IsZero(m.InvoiceID) { // not required
		return nil
	}

	if err := validate.FormatOf("invoiceId", "body", "uuid", m.InvoiceID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Invoice) validateParentAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentAccountID) { // not required
		return nil
	}

	if err := validate.FormatOf("parentAccountId", "body", "uuid", m.ParentAccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateParentInvoiceID(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentInvoiceID) { // not required
		return nil
	}

	if err := validate.FormatOf("parentInvoiceId", "body", "uuid", m.ParentInvoiceID.String(), formats); err != nil {
		return err
	}

	return nil
}

var invoiceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DRAFT","COMMITTED","VOID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceTypeStatusPropEnum = append(invoiceTypeStatusPropEnum, v)
	}
}

const (

	// InvoiceStatusDRAFT captures enum value "DRAFT"
	InvoiceStatusDRAFT string = "DRAFT"

	// InvoiceStatusCOMMITTED captures enum value "COMMITTED"
	InvoiceStatusCOMMITTED string = "COMMITTED"

	// InvoiceStatusVOID captures enum value "VOID"
	InvoiceStatusVOID string = "VOID"
)

// prop value enum
func (m *Invoice) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Invoice) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateTargetDate(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetDate) { // not required
		return nil
	}

	if err := validate.FormatOf("targetDate", "body", "date", m.TargetDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this invoice based on the context it is used
func (m *Invoice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Invoice) contextValidateAuditLogs(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Invoice) contextValidateCredits(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Credits); i++ {

		if m.Credits[i] != nil {
			if err := m.Credits[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("credits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Invoice) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Invoice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Invoice) UnmarshalBinary(b []byte) error {
	var res Invoice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
