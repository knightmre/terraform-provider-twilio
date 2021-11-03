/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/twilio/terraform-provider-twilio/client"
	. "github.com/twilio/terraform-provider-twilio/core"
	. "github.com/twilio/twilio-go/rest/wireless/v1"
)

func ResourceCommands() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCommands,
		ReadContext:   readCommands,
		DeleteContext: deleteCommands,
		Schema: map[string]*schema.Schema{
			"command":                    AsString(SchemaForceNewRequired),
			"callback_method":            AsString(SchemaForceNewOptional),
			"callback_url":               AsString(SchemaForceNewOptional),
			"command_mode":               AsString(SchemaForceNewOptional),
			"delivery_receipt_requested": AsBool(SchemaForceNewOptional),
			"include_sid":                AsString(SchemaForceNewOptional),
			"sim":                        AsString(SchemaForceNewOptional),
			"sid":                        AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseCommandsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createCommands(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateCommandParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.WirelessV1.CreateCommand(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteCommands(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.WirelessV1.DeleteCommand(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readCommands(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.WirelessV1.FetchCommand(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseCommandsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func ResourceRatePlans() *schema.Resource {
	return &schema.Resource{
		CreateContext: createRatePlans,
		ReadContext:   readRatePlans,
		UpdateContext: updateRatePlans,
		DeleteContext: deleteRatePlans,
		Schema: map[string]*schema.Schema{
			"data_enabled":                     AsBool(SchemaComputedOptional),
			"data_limit":                       AsInt(SchemaComputedOptional),
			"data_metering":                    AsString(SchemaComputedOptional),
			"friendly_name":                    AsString(SchemaComputedOptional),
			"international_roaming":            AsList(AsString(SchemaComputedOptional), SchemaComputedOptional),
			"international_roaming_data_limit": AsInt(SchemaComputedOptional),
			"messaging_enabled":                AsBool(SchemaComputedOptional),
			"national_roaming_data_limit":      AsInt(SchemaComputedOptional),
			"national_roaming_enabled":         AsBool(SchemaComputedOptional),
			"unique_name":                      AsString(SchemaComputedOptional),
			"voice_enabled":                    AsBool(SchemaComputedOptional),
			"sid":                              AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseRatePlansImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createRatePlans(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateRatePlanParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.WirelessV1.CreateRatePlan(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteRatePlans(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.WirelessV1.DeleteRatePlan(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readRatePlans(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.WirelessV1.FetchRatePlan(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseRatePlansImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateRatePlans(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateRatePlanParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.WirelessV1.UpdateRatePlan(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
