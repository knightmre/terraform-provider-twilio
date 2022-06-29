/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Studio
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/twilio/terraform-provider-twilio/client"
	. "github.com/twilio/terraform-provider-twilio/core"
	. "github.com/twilio/twilio-go/rest/studio/v2"
)

func ResourceFlowsExecutions() *schema.Resource {
	return &schema.Resource{
		CreateContext: createFlowsExecutions,
		ReadContext:   readFlowsExecutions,
		UpdateContext: updateFlowsExecutions,
		DeleteContext: deleteFlowsExecutions,
		Schema: map[string]*schema.Schema{
			"flow_sid":   AsString(SchemaRequired),
			"from":       AsString(SchemaRequired),
			"to":         AsString(SchemaRequired),
			"parameters": AsString(SchemaComputedOptional),
			"sid":        AsString(SchemaComputed),
			"status":     AsString(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseFlowsExecutionsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createFlowsExecutions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateExecutionParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	flowSid := d.Get("flow_sid").(string)

	r, err := m.(*client.Config).Client.StudioV2.CreateExecution(flowSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{flowSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))
	d.Set("sid", *r.Sid)

	return updateFlowsExecutions(ctx, d, m)
}

func deleteFlowsExecutions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	flowSid := d.Get("flow_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.StudioV2.DeleteExecution(flowSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readFlowsExecutions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	flowSid := d.Get("flow_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.StudioV2.FetchExecution(flowSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseFlowsExecutionsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected flow_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("flow_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateFlowsExecutions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateExecutionParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	flowSid := d.Get("flow_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.StudioV2.UpdateExecution(flowSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceFlows() *schema.Resource {
	return &schema.Resource{
		CreateContext: createFlows,
		ReadContext:   readFlows,
		UpdateContext: updateFlows,
		DeleteContext: deleteFlows,
		Schema: map[string]*schema.Schema{
			"definition":     AsString(SchemaRequired),
			"friendly_name":  AsString(SchemaRequired),
			"status":         AsString(SchemaRequired),
			"commit_message": AsString(SchemaComputedOptional),
			"sid":            AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseFlowsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createFlows(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateFlowParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.StudioV2.CreateFlow(&params)
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

func deleteFlows(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.StudioV2.DeleteFlow(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readFlows(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.StudioV2.FetchFlow(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseFlowsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateFlows(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateFlowParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.StudioV2.UpdateFlow(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
