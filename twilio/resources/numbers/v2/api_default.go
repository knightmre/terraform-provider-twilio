/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.21.0
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
	. "github.com/twilio/twilio-go/rest/numbers/v2"
)

func ResourceRegulatoryComplianceBundles() *schema.Resource {
	return &schema.Resource{
		CreateContext: createRegulatoryComplianceBundles,
		ReadContext:   readRegulatoryComplianceBundles,
		UpdateContext: updateRegulatoryComplianceBundles,
		DeleteContext: deleteRegulatoryComplianceBundles,
		Schema: map[string]*schema.Schema{
			"email":           AsString(SchemaRequired),
			"friendly_name":   AsString(SchemaRequired),
			"end_user_type":   AsString(SchemaComputedOptional),
			"iso_country":     AsString(SchemaComputedOptional),
			"number_type":     AsString(SchemaComputedOptional),
			"regulation_sid":  AsString(SchemaComputedOptional),
			"status_callback": AsString(SchemaComputedOptional),
			"sid":             AsString(SchemaComputed),
			"status":          AsString(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseRegulatoryComplianceBundlesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createRegulatoryComplianceBundles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateBundleParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.NumbersV2.CreateBundle(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))
	d.Set("sid", *r.Sid)

	return updateRegulatoryComplianceBundles(ctx, d, m)
}

func deleteRegulatoryComplianceBundles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.NumbersV2.DeleteBundle(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readRegulatoryComplianceBundles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NumbersV2.FetchBundle(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseRegulatoryComplianceBundlesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateRegulatoryComplianceBundles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateBundleParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NumbersV2.UpdateBundle(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceRegulatoryComplianceEndUsers() *schema.Resource {
	return &schema.Resource{
		CreateContext: createRegulatoryComplianceEndUsers,
		ReadContext:   readRegulatoryComplianceEndUsers,
		UpdateContext: updateRegulatoryComplianceEndUsers,
		DeleteContext: deleteRegulatoryComplianceEndUsers,
		Schema: map[string]*schema.Schema{
			"friendly_name": AsString(SchemaRequired),
			"type":          AsString(SchemaRequired),
			"attributes":    AsString(SchemaComputedOptional),
			"sid":           AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseRegulatoryComplianceEndUsersImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createRegulatoryComplianceEndUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateEndUserParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.NumbersV2.CreateEndUser(&params)
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

func deleteRegulatoryComplianceEndUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.NumbersV2.DeleteEndUser(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readRegulatoryComplianceEndUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NumbersV2.FetchEndUser(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseRegulatoryComplianceEndUsersImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateRegulatoryComplianceEndUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateEndUserParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NumbersV2.UpdateEndUser(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceRegulatoryComplianceBundlesItemAssignments() *schema.Resource {
	return &schema.Resource{
		CreateContext: createRegulatoryComplianceBundlesItemAssignments,
		ReadContext:   readRegulatoryComplianceBundlesItemAssignments,
		DeleteContext: deleteRegulatoryComplianceBundlesItemAssignments,
		Schema: map[string]*schema.Schema{
			"bundle_sid": AsString(SchemaForceNewRequired),
			"object_sid": AsString(SchemaForceNewRequired),
			"sid":        AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseRegulatoryComplianceBundlesItemAssignmentsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createRegulatoryComplianceBundlesItemAssignments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateItemAssignmentParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	bundleSid := d.Get("bundle_sid").(string)

	r, err := m.(*client.Config).Client.NumbersV2.CreateItemAssignment(bundleSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{bundleSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteRegulatoryComplianceBundlesItemAssignments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	bundleSid := d.Get("bundle_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.NumbersV2.DeleteItemAssignment(bundleSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readRegulatoryComplianceBundlesItemAssignments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	bundleSid := d.Get("bundle_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NumbersV2.FetchItemAssignment(bundleSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseRegulatoryComplianceBundlesItemAssignmentsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected bundle_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("bundle_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func ResourceRegulatoryComplianceSupportingDocuments() *schema.Resource {
	return &schema.Resource{
		CreateContext: createRegulatoryComplianceSupportingDocuments,
		ReadContext:   readRegulatoryComplianceSupportingDocuments,
		UpdateContext: updateRegulatoryComplianceSupportingDocuments,
		DeleteContext: deleteRegulatoryComplianceSupportingDocuments,
		Schema: map[string]*schema.Schema{
			"friendly_name": AsString(SchemaRequired),
			"type":          AsString(SchemaRequired),
			"attributes":    AsString(SchemaComputedOptional),
			"sid":           AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseRegulatoryComplianceSupportingDocumentsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createRegulatoryComplianceSupportingDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSupportingDocumentParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.NumbersV2.CreateSupportingDocument(&params)
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

func deleteRegulatoryComplianceSupportingDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.NumbersV2.DeleteSupportingDocument(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readRegulatoryComplianceSupportingDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NumbersV2.FetchSupportingDocument(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseRegulatoryComplianceSupportingDocumentsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateRegulatoryComplianceSupportingDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSupportingDocumentParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NumbersV2.UpdateSupportingDocument(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
