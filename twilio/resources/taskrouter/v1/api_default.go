/*
 * Twilio - Taskrouter
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
	. "github.com/twilio/twilio-go/rest/taskrouter/v1"
)

func ResourceWorkspacesActivities() *schema.Resource {
	return &schema.Resource{
		CreateContext: createWorkspacesActivities,
		ReadContext:   readWorkspacesActivities,
		UpdateContext: updateWorkspacesActivities,
		DeleteContext: deleteWorkspacesActivities,
		Schema: map[string]*schema.Schema{
			"workspace_sid": AsString(SchemaRequired),
			"friendly_name": AsString(SchemaRequired),
			"available":     AsBool(SchemaComputedOptional),
			"sid":           AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseWorkspacesActivitiesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createWorkspacesActivities(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateActivityParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	workspaceSid := d.Get("workspace_sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.CreateActivity(workspaceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{workspaceSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteWorkspacesActivities(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TaskrouterV1.DeleteActivity(workspaceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readWorkspacesActivities(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.FetchActivity(workspaceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseWorkspacesActivitiesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected workspace_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("workspace_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateWorkspacesActivities(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateActivityParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.UpdateActivity(workspaceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceWorkspacesTasks() *schema.Resource {
	return &schema.Resource{
		CreateContext: createWorkspacesTasks,
		ReadContext:   readWorkspacesTasks,
		UpdateContext: updateWorkspacesTasks,
		DeleteContext: deleteWorkspacesTasks,
		Schema: map[string]*schema.Schema{
			"workspace_sid":     AsString(SchemaRequired),
			"attributes":        AsString(SchemaComputedOptional),
			"priority":          AsInt(SchemaComputedOptional),
			"task_channel":      AsString(SchemaComputedOptional),
			"timeout":           AsInt(SchemaComputedOptional),
			"workflow_sid":      AsString(SchemaComputedOptional),
			"sid":               AsString(SchemaComputed),
			"if_match":          AsString(SchemaComputedOptional),
			"assignment_status": AsString(SchemaComputedOptional),
			"reason":            AsString(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseWorkspacesTasksImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createWorkspacesTasks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateTaskParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	workspaceSid := d.Get("workspace_sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.CreateTask(workspaceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{workspaceSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))
	d.Set("sid", *r.Sid)

	return updateWorkspacesTasks(ctx, d, m)
}

func deleteWorkspacesTasks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := DeleteTaskParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TaskrouterV1.DeleteTask(workspaceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readWorkspacesTasks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.FetchTask(workspaceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseWorkspacesTasksImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected workspace_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("workspace_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateWorkspacesTasks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateTaskParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.UpdateTask(workspaceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceWorkspacesTaskChannels() *schema.Resource {
	return &schema.Resource{
		CreateContext: createWorkspacesTaskChannels,
		ReadContext:   readWorkspacesTaskChannels,
		UpdateContext: updateWorkspacesTaskChannels,
		DeleteContext: deleteWorkspacesTaskChannels,
		Schema: map[string]*schema.Schema{
			"workspace_sid":             AsString(SchemaRequired),
			"friendly_name":             AsString(SchemaRequired),
			"unique_name":               AsString(SchemaRequired),
			"channel_optimized_routing": AsBool(SchemaComputedOptional),
			"sid":                       AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseWorkspacesTaskChannelsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createWorkspacesTaskChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateTaskChannelParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	workspaceSid := d.Get("workspace_sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.CreateTaskChannel(workspaceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{workspaceSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteWorkspacesTaskChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TaskrouterV1.DeleteTaskChannel(workspaceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readWorkspacesTaskChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.FetchTaskChannel(workspaceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseWorkspacesTaskChannelsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected workspace_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("workspace_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateWorkspacesTaskChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateTaskChannelParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.UpdateTaskChannel(workspaceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceWorkspacesTaskQueues() *schema.Resource {
	return &schema.Resource{
		CreateContext: createWorkspacesTaskQueues,
		ReadContext:   readWorkspacesTaskQueues,
		UpdateContext: updateWorkspacesTaskQueues,
		DeleteContext: deleteWorkspacesTaskQueues,
		Schema: map[string]*schema.Schema{
			"workspace_sid":            AsString(SchemaRequired),
			"friendly_name":            AsString(SchemaRequired),
			"assignment_activity_sid":  AsString(SchemaComputedOptional),
			"max_reserved_workers":     AsInt(SchemaComputedOptional),
			"reservation_activity_sid": AsString(SchemaComputedOptional),
			"target_workers":           AsString(SchemaComputedOptional),
			"task_order":               AsString(SchemaComputedOptional),
			"sid":                      AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseWorkspacesTaskQueuesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createWorkspacesTaskQueues(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateTaskQueueParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	workspaceSid := d.Get("workspace_sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.CreateTaskQueue(workspaceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{workspaceSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteWorkspacesTaskQueues(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TaskrouterV1.DeleteTaskQueue(workspaceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readWorkspacesTaskQueues(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.FetchTaskQueue(workspaceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseWorkspacesTaskQueuesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected workspace_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("workspace_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateWorkspacesTaskQueues(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateTaskQueueParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.UpdateTaskQueue(workspaceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceWorkspacesWorkers() *schema.Resource {
	return &schema.Resource{
		CreateContext: createWorkspacesWorkers,
		ReadContext:   readWorkspacesWorkers,
		UpdateContext: updateWorkspacesWorkers,
		DeleteContext: deleteWorkspacesWorkers,
		Schema: map[string]*schema.Schema{
			"workspace_sid":               AsString(SchemaRequired),
			"friendly_name":               AsString(SchemaRequired),
			"activity_sid":                AsString(SchemaComputedOptional),
			"attributes":                  AsString(SchemaComputedOptional),
			"sid":                         AsString(SchemaComputed),
			"if_match":                    AsString(SchemaComputedOptional),
			"reject_pending_reservations": AsBool(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseWorkspacesWorkersImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createWorkspacesWorkers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateWorkerParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	workspaceSid := d.Get("workspace_sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.CreateWorker(workspaceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{workspaceSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))
	d.Set("sid", *r.Sid)

	return updateWorkspacesWorkers(ctx, d, m)
}

func deleteWorkspacesWorkers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := DeleteWorkerParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TaskrouterV1.DeleteWorker(workspaceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readWorkspacesWorkers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.FetchWorker(workspaceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseWorkspacesWorkersImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected workspace_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("workspace_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateWorkspacesWorkers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateWorkerParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.UpdateWorker(workspaceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceWorkspacesWorkflows() *schema.Resource {
	return &schema.Resource{
		CreateContext: createWorkspacesWorkflows,
		ReadContext:   readWorkspacesWorkflows,
		UpdateContext: updateWorkspacesWorkflows,
		DeleteContext: deleteWorkspacesWorkflows,
		Schema: map[string]*schema.Schema{
			"workspace_sid":                    AsString(SchemaRequired),
			"configuration":                    AsString(SchemaRequired),
			"friendly_name":                    AsString(SchemaRequired),
			"assignment_callback_url":          AsString(SchemaComputedOptional),
			"fallback_assignment_callback_url": AsString(SchemaComputedOptional),
			"task_reservation_timeout":         AsInt(SchemaComputedOptional),
			"sid":                              AsString(SchemaComputed),
			"re_evaluate_tasks":                AsString(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseWorkspacesWorkflowsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createWorkspacesWorkflows(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateWorkflowParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	workspaceSid := d.Get("workspace_sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.CreateWorkflow(workspaceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{workspaceSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))
	d.Set("sid", *r.Sid)

	return updateWorkspacesWorkflows(ctx, d, m)
}

func deleteWorkspacesWorkflows(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TaskrouterV1.DeleteWorkflow(workspaceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readWorkspacesWorkflows(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.FetchWorkflow(workspaceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseWorkspacesWorkflowsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected workspace_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("workspace_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateWorkspacesWorkflows(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateWorkflowParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.UpdateWorkflow(workspaceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceWorkspaces() *schema.Resource {
	return &schema.Resource{
		CreateContext: createWorkspaces,
		ReadContext:   readWorkspaces,
		UpdateContext: updateWorkspaces,
		DeleteContext: deleteWorkspaces,
		Schema: map[string]*schema.Schema{
			"friendly_name":          AsString(SchemaRequired),
			"event_callback_url":     AsString(SchemaComputedOptional),
			"events_filter":          AsString(SchemaComputedOptional),
			"multi_task_enabled":     AsBool(SchemaComputedOptional),
			"prioritize_queue_order": AsString(SchemaComputedOptional),
			"template":               AsString(SchemaComputedOptional),
			"sid":                    AsString(SchemaComputed),
			"default_activity_sid":   AsString(SchemaComputedOptional),
			"timeout_activity_sid":   AsString(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseWorkspacesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createWorkspaces(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateWorkspaceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.TaskrouterV1.CreateWorkspace(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))
	d.Set("sid", *r.Sid)

	return updateWorkspaces(ctx, d, m)
}

func deleteWorkspaces(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TaskrouterV1.DeleteWorkspace(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readWorkspaces(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.FetchWorkspace(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseWorkspacesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateWorkspaces(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateWorkspaceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TaskrouterV1.UpdateWorkspace(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
