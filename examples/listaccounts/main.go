package main

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/CDNA-Technologies/kbcli/v3/kbclient"
	"github.com/CDNA-Technologies/kbcli/v3/kbclient/account"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewClient creates new kill bill client
func NewClient() *kbclient.KillBill {
	trp := httptransport.New("arise-dev-killbill.gonuclei.com", "", []string{"https"})
	// Add text/xml producer which is not handled by openapi runtime.
	trp.Producers["text/xml"] = runtime.TextProducer()
	// Set this to true to dump http messages
	trp.Debug = false
	apiKey := "golang-integration-subscription2"
	apiSecret := "golang-integration-subscription2"
	authWriter := runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		encoded := base64.StdEncoding.EncodeToString([]byte("admin" /*username*/ + ":" + "password" /**password*/))
		if err := r.SetHeaderParam("Authorization", "Basic "+encoded); err != nil {
			return err
		}
		if err := r.SetHeaderParam("X-KillBill-ApiKey", apiKey); err != nil {
			return err
		}
		if err := r.SetHeaderParam("X-KillBill-ApiSecret", apiSecret); err != nil {
			return err
		}
		return nil
	})
	client := kbclient.New(trp, strfmt.Default, authWriter, kbclient.KillbillDefaults{})
	return client
}

func main() {
	c := NewClient()

	resp, err := c.Account.GetAccounts(context.Background(), &account.GetAccountsParams{})
	if err != nil {
		panic(err)
	}
	for _, acc := range resp.Payload {
		fmt.Printf("%s %s %s\n", acc.AccountID, acc.ExternalKey, acc.Email)
	}
}
