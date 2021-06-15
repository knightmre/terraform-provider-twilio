/*
 * Twilio - Accounts
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/twilio/terraform-provider-twilio/client"
	. "github.com/twilio/terraform-provider-twilio/twilio/common"
	. "github.com/twilio/twilio-go/rest/accounts/v1"
)

func ResourceCredentialsAWS() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCredentialsAWS,
		ReadContext:   readCredentialsAWS,
		UpdateContext: updateCredentialsAWS,
		DeleteContext: deleteCredentialsAWS,
		Schema: map[string]*schema.Schema{
			"account_sid":   AsString(SchemaOptional),
			"credentials":   AsString(SchemaOptional),
			"friendly_name": AsString(SchemaOptional),
			"date_created":  AsString(SchemaOptional),
			"date_updated":  AsString(SchemaOptional),
			"sid":           AsString(SchemaOptional),
			"url":           AsString(SchemaOptional),
		},
	}
}

func createCredentialsAWS(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateCredentialAwsParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.AccountsV1.CreateCredentialAws(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId((*r.Sid))
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteCredentialsAWS(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.AccountsV1.DeleteCredentialAws(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readCredentialsAWS(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AccountsV1.FetchCredentialAws(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateCredentialsAWS(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateCredentialAwsParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AccountsV1.UpdateCredentialAws(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceCredentialsPublicKeys() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCredentialsPublicKeys,
		ReadContext:   readCredentialsPublicKeys,
		UpdateContext: updateCredentialsPublicKeys,
		DeleteContext: deleteCredentialsPublicKeys,
		Schema: map[string]*schema.Schema{
			"account_sid":   AsString(SchemaOptional),
			"friendly_name": AsString(SchemaOptional),
			"public_key":    AsString(SchemaOptional),
			"date_created":  AsString(SchemaOptional),
			"date_updated":  AsString(SchemaOptional),
			"sid":           AsString(SchemaOptional),
			"url":           AsString(SchemaOptional),
		},
	}
}

func createCredentialsPublicKeys(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateCredentialPublicKeyParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.AccountsV1.CreateCredentialPublicKey(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId((*r.Sid))
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteCredentialsPublicKeys(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.AccountsV1.DeleteCredentialPublicKey(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readCredentialsPublicKeys(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AccountsV1.FetchCredentialPublicKey(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateCredentialsPublicKeys(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateCredentialPublicKeyParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.AccountsV1.UpdateCredentialPublicKey(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
