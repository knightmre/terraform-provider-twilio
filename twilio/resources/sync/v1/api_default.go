/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.1
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
	. "github.com/twilio/twilio-go/rest/sync/v1"
)

func ResourceServicesDocuments() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesDocuments,
		ReadContext:   readServicesDocuments,
		UpdateContext: updateServicesDocuments,
		DeleteContext: deleteServicesDocuments,
		Schema: map[string]*schema.Schema{
			"service_sid": AsString(SchemaRequired),
			"data":        AsString(SchemaComputedOptional),
			"ttl":         AsInt(SchemaComputedOptional),
			"unique_name": AsString(SchemaComputedOptional),
			"sid":         AsString(SchemaComputed),
			"if_match":    AsString(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesDocumentsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateDocumentParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.CreateDocument(serviceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.SyncV1.DeleteDocument(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.FetchDocument(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesDocumentsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateServicesDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateDocumentParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.UpdateDocument(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServices() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServices,
		ReadContext:   readServices,
		UpdateContext: updateServices,
		DeleteContext: deleteServices,
		Schema: map[string]*schema.Schema{
			"acl_enabled":                     AsBool(SchemaComputedOptional),
			"friendly_name":                   AsString(SchemaComputedOptional),
			"reachability_debouncing_enabled": AsBool(SchemaComputedOptional),
			"reachability_debouncing_window":  AsInt(SchemaComputedOptional),
			"reachability_webhooks_enabled":   AsBool(SchemaComputedOptional),
			"webhook_url":                     AsString(SchemaComputedOptional),
			"webhooks_from_rest_enabled":      AsBool(SchemaComputedOptional),
			"sid":                             AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateServiceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.SyncV1.CreateService(&params)
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

func deleteServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.SyncV1.DeleteService(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.FetchService(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateServiceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.UpdateService(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesLists() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesLists,
		ReadContext:   readServicesLists,
		UpdateContext: updateServicesLists,
		DeleteContext: deleteServicesLists,
		Schema: map[string]*schema.Schema{
			"service_sid":    AsString(SchemaRequired),
			"collection_ttl": AsInt(SchemaComputedOptional),
			"ttl":            AsInt(SchemaComputedOptional),
			"unique_name":    AsString(SchemaComputedOptional),
			"sid":            AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesListsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesLists(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSyncListParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.CreateSyncList(serviceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesLists(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.SyncV1.DeleteSyncList(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesLists(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.FetchSyncList(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesListsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateServicesLists(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSyncListParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.UpdateSyncList(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesListsItems() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesListsItems,
		ReadContext:   readServicesListsItems,
		UpdateContext: updateServicesListsItems,
		DeleteContext: deleteServicesListsItems,
		Schema: map[string]*schema.Schema{
			"service_sid":    AsString(SchemaRequired),
			"list_sid":       AsString(SchemaRequired),
			"data":           AsString(SchemaRequired),
			"collection_ttl": AsInt(SchemaComputedOptional),
			"item_ttl":       AsInt(SchemaComputedOptional),
			"ttl":            AsInt(SchemaComputedOptional),
			"index":          AsInt(SchemaComputed),
			"if_match":       AsString(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesListsItemsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesListsItems(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSyncListItemParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	listSid := d.Get("list_sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.CreateSyncListItem(serviceSid, listSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid, listSid}
	idParts = append(idParts, IntToString(*r.Index))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesListsItems(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := DeleteSyncListItemParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	listSid := d.Get("list_sid").(string)
	index := d.Get("index").(int)

	err := m.(*client.Config).Client.SyncV1.DeleteSyncListItem(serviceSid, listSid, index, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesListsItems(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	listSid := d.Get("list_sid").(string)
	index := d.Get("index").(int)

	r, err := m.(*client.Config).Client.SyncV1.FetchSyncListItem(serviceSid, listSid, index)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesListsItemsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/list_sid/index"

	if len(importParts) != 3 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("list_sid", importParts[1])
	d.Set("index", importParts[2])

	return nil
}
func updateServicesListsItems(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSyncListItemParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	listSid := d.Get("list_sid").(string)
	index := d.Get("index").(int)

	r, err := m.(*client.Config).Client.SyncV1.UpdateSyncListItem(serviceSid, listSid, index, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesMaps() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesMaps,
		ReadContext:   readServicesMaps,
		UpdateContext: updateServicesMaps,
		DeleteContext: deleteServicesMaps,
		Schema: map[string]*schema.Schema{
			"service_sid":    AsString(SchemaRequired),
			"collection_ttl": AsInt(SchemaComputedOptional),
			"ttl":            AsInt(SchemaComputedOptional),
			"unique_name":    AsString(SchemaComputedOptional),
			"sid":            AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesMapsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesMaps(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSyncMapParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.CreateSyncMap(serviceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesMaps(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.SyncV1.DeleteSyncMap(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesMaps(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.FetchSyncMap(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesMapsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateServicesMaps(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSyncMapParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.UpdateSyncMap(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesMapsItems() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesMapsItems,
		ReadContext:   readServicesMapsItems,
		UpdateContext: updateServicesMapsItems,
		DeleteContext: deleteServicesMapsItems,
		Schema: map[string]*schema.Schema{
			"service_sid":    AsString(SchemaRequired),
			"map_sid":        AsString(SchemaRequired),
			"data":           AsString(SchemaRequired),
			"key":            AsString(SchemaRequired),
			"collection_ttl": AsInt(SchemaComputedOptional),
			"item_ttl":       AsInt(SchemaComputedOptional),
			"ttl":            AsInt(SchemaComputedOptional),
			"if_match":       AsString(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesMapsItemsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesMapsItems(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSyncMapItemParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	mapSid := d.Get("map_sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.CreateSyncMapItem(serviceSid, mapSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid, mapSid}
	idParts = append(idParts, (*r.Key))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesMapsItems(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := DeleteSyncMapItemParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	mapSid := d.Get("map_sid").(string)
	key := d.Get("key").(string)

	err := m.(*client.Config).Client.SyncV1.DeleteSyncMapItem(serviceSid, mapSid, key, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesMapsItems(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	mapSid := d.Get("map_sid").(string)
	key := d.Get("key").(string)

	r, err := m.(*client.Config).Client.SyncV1.FetchSyncMapItem(serviceSid, mapSid, key)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesMapsItemsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/map_sid/key"

	if len(importParts) != 3 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("map_sid", importParts[1])
	d.Set("key", importParts[2])

	return nil
}
func updateServicesMapsItems(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSyncMapItemParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	mapSid := d.Get("map_sid").(string)
	key := d.Get("key").(string)

	r, err := m.(*client.Config).Client.SyncV1.UpdateSyncMapItem(serviceSid, mapSid, key, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesStreams() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesStreams,
		ReadContext:   readServicesStreams,
		UpdateContext: updateServicesStreams,
		DeleteContext: deleteServicesStreams,
		Schema: map[string]*schema.Schema{
			"service_sid": AsString(SchemaRequired),
			"ttl":         AsInt(SchemaComputedOptional),
			"unique_name": AsString(SchemaComputedOptional),
			"sid":         AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesStreamsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesStreams(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSyncStreamParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.CreateSyncStream(serviceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesStreams(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.SyncV1.DeleteSyncStream(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesStreams(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.FetchSyncStream(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesStreamsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateServicesStreams(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSyncStreamParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.UpdateSyncStream(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
