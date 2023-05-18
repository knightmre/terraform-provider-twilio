/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
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
	. "github.com/twilio/twilio-go/rest/flex/v1"
)

func ResourceChannels() *schema.Resource {
	return &schema.Resource{
		CreateContext: createChannels,
		ReadContext:   readChannels,
		DeleteContext: deleteChannels,
		Schema: map[string]*schema.Schema{
			"flex_flow_sid":           AsString(SchemaForceNewRequired),
			"identity":                AsString(SchemaForceNewRequired),
			"chat_user_friendly_name": AsString(SchemaForceNewRequired),
			"chat_friendly_name":      AsString(SchemaForceNewRequired),
			"target":                  AsString(SchemaForceNewOptional),
			"chat_unique_name":        AsString(SchemaForceNewOptional),
			"pre_engagement_data":     AsString(SchemaForceNewOptional),
			"task_sid":                AsString(SchemaForceNewOptional),
			"task_attributes":         AsString(SchemaForceNewOptional),
			"long_lived":              AsBool(SchemaForceNewOptional),
			"sid":                     AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseChannelsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateChannelParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.FlexV1.CreateChannel(&params)
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

func deleteChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.FlexV1.DeleteChannel(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.FlexV1.FetchChannel(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseChannelsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func ResourceFlexFlows() *schema.Resource {
	return &schema.Resource{
		CreateContext: createFlexFlows,
		ReadContext:   readFlexFlows,
		UpdateContext: updateFlexFlows,
		DeleteContext: deleteFlexFlows,
		Schema: map[string]*schema.Schema{
			"friendly_name":                   AsString(SchemaRequired),
			"chat_service_sid":                AsString(SchemaRequired),
			"channel_type":                    AsString(SchemaRequired),
			"contact_identity":                AsString(SchemaComputedOptional),
			"enabled":                         AsBool(SchemaComputedOptional),
			"integration_type":                AsString(SchemaComputedOptional),
			"integration_flow_sid":            AsString(SchemaComputedOptional),
			"integration_url":                 AsString(SchemaComputedOptional),
			"integration_workspace_sid":       AsString(SchemaComputedOptional),
			"integration_workflow_sid":        AsString(SchemaComputedOptional),
			"integration_channel":             AsString(SchemaComputedOptional),
			"integration_timeout":             AsInt(SchemaComputedOptional),
			"integration_priority":            AsInt(SchemaComputedOptional),
			"integration_creation_on_message": AsBool(SchemaComputedOptional),
			"long_lived":                      AsBool(SchemaComputedOptional),
			"janitor_enabled":                 AsBool(SchemaComputedOptional),
			"integration_retry_count":         AsInt(SchemaComputedOptional),
			"sid":                             AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseFlexFlowsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createFlexFlows(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateFlexFlowParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.FlexV1.CreateFlexFlow(&params)
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

func deleteFlexFlows(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.FlexV1.DeleteFlexFlow(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readFlexFlows(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.FlexV1.FetchFlexFlow(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseFlexFlowsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateFlexFlows(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateFlexFlowParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.FlexV1.UpdateFlexFlow(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceInsightsQualityManagementQuestionnaires() *schema.Resource {
	return &schema.Resource{
		CreateContext: createInsightsQualityManagementQuestionnaires,
		ReadContext:   readInsightsQualityManagementQuestionnaires,
		UpdateContext: updateInsightsQualityManagementQuestionnaires,
		DeleteContext: deleteInsightsQualityManagementQuestionnaires,
		Schema: map[string]*schema.Schema{
			"name":              AsString(SchemaRequired),
			"authorization":     AsString(SchemaComputedOptional),
			"description":       AsString(SchemaComputedOptional),
			"active":            AsBool(SchemaComputedOptional),
			"question_sids":     AsList(AsString(SchemaComputedOptional), SchemaComputedOptional),
			"questionnaire_sid": AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseInsightsQualityManagementQuestionnairesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createInsightsQualityManagementQuestionnaires(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateInsightsQuestionnairesParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.FlexV1.CreateInsightsQuestionnaires(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.QuestionnaireSid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteInsightsQualityManagementQuestionnaires(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := DeleteInsightsQuestionnairesParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	questionnaireSid := d.Get("questionnaire_sid").(string)

	err := m.(*client.Config).Client.FlexV1.DeleteInsightsQuestionnaires(questionnaireSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readInsightsQualityManagementQuestionnaires(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := FetchInsightsQuestionnairesParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	questionnaireSid := d.Get("questionnaire_sid").(string)

	r, err := m.(*client.Config).Client.FlexV1.FetchInsightsQuestionnaires(questionnaireSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseInsightsQualityManagementQuestionnairesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected questionnaire_sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("questionnaire_sid", importParts[0])

	return nil
}
func updateInsightsQualityManagementQuestionnaires(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateInsightsQuestionnairesParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	questionnaireSid := d.Get("questionnaire_sid").(string)

	r, err := m.(*client.Config).Client.FlexV1.UpdateInsightsQuestionnaires(questionnaireSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceWebChannels() *schema.Resource {
	return &schema.Resource{
		CreateContext: createWebChannels,
		ReadContext:   readWebChannels,
		UpdateContext: updateWebChannels,
		DeleteContext: deleteWebChannels,
		Schema: map[string]*schema.Schema{
			"flex_flow_sid":          AsString(SchemaForceNewRequired),
			"identity":               AsString(SchemaForceNewRequired),
			"customer_friendly_name": AsString(SchemaForceNewRequired),
			"chat_friendly_name":     AsString(SchemaForceNewRequired),
			"chat_unique_name":       AsString(SchemaForceNewOptional),
			"pre_engagement_data":    AsString(SchemaForceNewOptional),
			"sid":                    AsString(SchemaComputed),
			"chat_status":            AsString(SchemaComputedOptional),
			"post_engagement_data":   AsString(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseWebChannelsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createWebChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateWebChannelParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.FlexV1.CreateWebChannel(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))
	d.Set("sid", *r.Sid)

	return updateWebChannels(ctx, d, m)
}

func deleteWebChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.FlexV1.DeleteWebChannel(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readWebChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.FlexV1.FetchWebChannel(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseWebChannelsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateWebChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateWebChannelParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.FlexV1.UpdateWebChannel(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
