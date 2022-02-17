/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.0
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
	. "github.com/twilio/twilio-go/rest/video/v1"
)

func ResourceCompositions() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCompositions,
		ReadContext:   readCompositions,
		DeleteContext: deleteCompositions,
		Schema: map[string]*schema.Schema{
			"room_sid":               AsString(SchemaForceNewRequired),
			"audio_sources":          AsList(AsString(SchemaForceNewOptional), SchemaForceNewOptional),
			"audio_sources_excluded": AsList(AsString(SchemaForceNewOptional), SchemaForceNewOptional),
			"format":                 AsString(SchemaForceNewOptional),
			"resolution":             AsString(SchemaForceNewOptional),
			"status_callback":        AsString(SchemaForceNewOptional),
			"status_callback_method": AsString(SchemaForceNewOptional),
			"trim":                   AsBool(SchemaForceNewOptional),
			"video_layout":           AsString(SchemaForceNewOptional),
			"sid":                    AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseCompositionsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createCompositions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateCompositionParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.VideoV1.CreateComposition(&params)
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

func deleteCompositions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VideoV1.DeleteComposition(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readCompositions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VideoV1.FetchComposition(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseCompositionsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func ResourceCompositionHooks() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCompositionHooks,
		ReadContext:   readCompositionHooks,
		UpdateContext: updateCompositionHooks,
		DeleteContext: deleteCompositionHooks,
		Schema: map[string]*schema.Schema{
			"friendly_name":          AsString(SchemaRequired),
			"audio_sources":          AsList(AsString(SchemaComputedOptional), SchemaComputedOptional),
			"audio_sources_excluded": AsList(AsString(SchemaComputedOptional), SchemaComputedOptional),
			"enabled":                AsBool(SchemaComputedOptional),
			"format":                 AsString(SchemaComputedOptional),
			"resolution":             AsString(SchemaComputedOptional),
			"status_callback":        AsString(SchemaComputedOptional),
			"status_callback_method": AsString(SchemaComputedOptional),
			"trim":                   AsBool(SchemaComputedOptional),
			"video_layout":           AsString(SchemaComputedOptional),
			"sid":                    AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseCompositionHooksImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createCompositionHooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateCompositionHookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.VideoV1.CreateCompositionHook(&params)
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

func deleteCompositionHooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VideoV1.DeleteCompositionHook(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readCompositionHooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VideoV1.FetchCompositionHook(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseCompositionHooksImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateCompositionHooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateCompositionHookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VideoV1.UpdateCompositionHook(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
