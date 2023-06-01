/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	accountsV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/accounts/v1"
	api "github.com/twilio/terraform-provider-twilio/twilio/resources/api/v2010"
	autopilotV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/autopilot/v1"
	chatV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/chat/v1"
	chatV2 "github.com/twilio/terraform-provider-twilio/twilio/resources/chat/v2"
	conversationsV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/conversations/v1"
	eventsV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/events/v1"
	flexV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/flex/v1"
	intelligenceV2 "github.com/twilio/terraform-provider-twilio/twilio/resources/intelligence/v2"
	ip_messagingV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/ip_messaging/v1"
	ip_messagingV2 "github.com/twilio/terraform-provider-twilio/twilio/resources/ip_messaging/v2"
	messagingV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/messaging/v1"
	microvisorV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/microvisor/v1"
	notifyV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/notify/v1"
	numbersV2 "github.com/twilio/terraform-provider-twilio/twilio/resources/numbers/v2"
	proxyV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/proxy/v1"
	serverlessV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/serverless/v1"
	studioV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/studio/v1"
	studioV2 "github.com/twilio/terraform-provider-twilio/twilio/resources/studio/v2"
	supersimV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/supersim/v1"
	syncV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/sync/v1"
	taskrouterV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/taskrouter/v1"
	trunkingV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/trunking/v1"
	trusthubV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/trusthub/v1"
	verifyV2 "github.com/twilio/terraform-provider-twilio/twilio/resources/verify/v2"
	videoV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/video/v1"
	voiceV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/voice/v1"
	wirelessV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/wireless/v1"
)

type TwilioResources struct {
	Map map[string]*schema.Resource
}

func NewTwilioResources() *TwilioResources {
	return &TwilioResources{
		map[string]*schema.Resource{
			"twilio_accounts_credentials_aws_v1":                                          accountsV1.ResourceCredentialsAWS(),
			"twilio_accounts_credentials_public_keys_v1":                                  accountsV1.ResourceCredentialsPublicKeys(),
			"twilio_api_accounts_addresses":                                               api.ResourceAccountsAddresses(),
			"twilio_api_accounts_applications":                                            api.ResourceAccountsApplications(),
			"twilio_api_accounts_calls":                                                   api.ResourceAccountsCalls(),
			"twilio_api_accounts_calls_feedback_summary":                                  api.ResourceAccountsCallsFeedbackSummary(),
			"twilio_api_accounts_calls_recordings":                                        api.ResourceAccountsCallsRecordings(),
			"twilio_api_accounts_incoming_phone_numbers":                                  api.ResourceAccountsIncomingPhoneNumbers(),
			"twilio_api_accounts_incoming_phone_numbers_assigned_add_ons":                 api.ResourceAccountsIncomingPhoneNumbersAssignedAddOns(),
			"twilio_api_accounts_messages":                                                api.ResourceAccountsMessages(),
			"twilio_api_accounts_keys":                                                    api.ResourceAccountsKeys(),
			"twilio_api_accounts_signing_keys":                                            api.ResourceAccountsSigningKeys(),
			"twilio_api_accounts_conferences_participants":                                api.ResourceAccountsConferencesParticipants(),
			"twilio_api_accounts_queues":                                                  api.ResourceAccountsQueues(),
			"twilio_api_accounts_sip_domains_auth_calls_credential_list_mappings":         api.ResourceAccountsSIPDomainsAuthCallsCredentialListMappings(),
			"twilio_api_accounts_sip_domains_auth_calls_ip_access_control_list_mappings":  api.ResourceAccountsSIPDomainsAuthCallsIpAccessControlListMappings(),
			"twilio_api_accounts_sip_domains_auth_registrations_credential_list_mappings": api.ResourceAccountsSIPDomainsAuthRegistrationsCredentialListMappings(),
			"twilio_api_accounts_sip_credential_lists_credentials":                        api.ResourceAccountsSIPCredentialListsCredentials(),
			"twilio_api_accounts_sip_credential_lists":                                    api.ResourceAccountsSIPCredentialLists(),
			"twilio_api_accounts_sip_domains_credential_list_mappings":                    api.ResourceAccountsSIPDomainsCredentialListMappings(),
			"twilio_api_accounts_sip_domains":                                             api.ResourceAccountsSIPDomains(),
			"twilio_api_accounts_sip_ip_access_control_lists":                             api.ResourceAccountsSIPIpAccessControlLists(),
			"twilio_api_accounts_sip_domains_ip_access_control_list_mappings":             api.ResourceAccountsSIPDomainsIpAccessControlListMappings(),
			"twilio_api_accounts_sip_ip_access_control_lists_ip_addresses":                api.ResourceAccountsSIPIpAccessControlListsIpAddresses(),
			"twilio_api_accounts_usage_triggers":                                          api.ResourceAccountsUsageTriggers(),
			"twilio_autopilot_assistants_v1":                                              autopilotV1.ResourceAssistants(),
			"twilio_autopilot_assistants_tasks_fields_v1":                                 autopilotV1.ResourceAssistantsTasksFields(),
			"twilio_autopilot_assistants_field_types_v1":                                  autopilotV1.ResourceAssistantsFieldTypes(),
			"twilio_autopilot_assistants_field_types_field_values_v1":                     autopilotV1.ResourceAssistantsFieldTypesFieldValues(),
			"twilio_autopilot_assistants_model_builds_v1":                                 autopilotV1.ResourceAssistantsModelBuilds(),
			"twilio_autopilot_assistants_queries_v1":                                      autopilotV1.ResourceAssistantsQueries(),
			"twilio_autopilot_assistants_tasks_samples_v1":                                autopilotV1.ResourceAssistantsTasksSamples(),
			"twilio_autopilot_assistants_tasks_v1":                                        autopilotV1.ResourceAssistantsTasks(),
			"twilio_autopilot_assistants_webhooks_v1":                                     autopilotV1.ResourceAssistantsWebhooks(),
			"twilio_chat_services_channels_v1":                                            chatV1.ResourceServicesChannels(),
			"twilio_chat_credentials_v1":                                                  chatV1.ResourceCredentials(),
			"twilio_chat_services_channels_invites_v1":                                    chatV1.ResourceServicesChannelsInvites(),
			"twilio_chat_services_channels_members_v1":                                    chatV1.ResourceServicesChannelsMembers(),
			"twilio_chat_services_channels_messages_v1":                                   chatV1.ResourceServicesChannelsMessages(),
			"twilio_chat_services_roles_v1":                                               chatV1.ResourceServicesRoles(),
			"twilio_chat_services_v1":                                                     chatV1.ResourceServices(),
			"twilio_chat_services_users_v1":                                               chatV1.ResourceServicesUsers(),
			"twilio_chat_services_channels_v2":                                            chatV2.ResourceServicesChannels(),
			"twilio_chat_services_channels_webhooks_v2":                                   chatV2.ResourceServicesChannelsWebhooks(),
			"twilio_chat_credentials_v2":                                                  chatV2.ResourceCredentials(),
			"twilio_chat_services_channels_invites_v2":                                    chatV2.ResourceServicesChannelsInvites(),
			"twilio_chat_services_channels_members_v2":                                    chatV2.ResourceServicesChannelsMembers(),
			"twilio_chat_services_channels_messages_v2":                                   chatV2.ResourceServicesChannelsMessages(),
			"twilio_chat_services_roles_v2":                                               chatV2.ResourceServicesRoles(),
			"twilio_chat_services_v2":                                                     chatV2.ResourceServices(),
			"twilio_chat_services_users_v2":                                               chatV2.ResourceServicesUsers(),
			"twilio_conversations_configuration_addresses_v1":                             conversationsV1.ResourceConfigurationAddresses(),
			"twilio_conversations_conversations_v1":                                       conversationsV1.ResourceConversations(),
			"twilio_conversations_conversations_messages_v1":                              conversationsV1.ResourceConversationsMessages(),
			"twilio_conversations_conversations_participants_v1":                          conversationsV1.ResourceConversationsParticipants(),
			"twilio_conversations_conversations_webhooks_v1":                              conversationsV1.ResourceConversationsWebhooks(),
			"twilio_conversations_credentials_v1":                                         conversationsV1.ResourceCredentials(),
			"twilio_conversations_roles_v1":                                               conversationsV1.ResourceRoles(),
			"twilio_conversations_services_v1":                                            conversationsV1.ResourceServices(),
			"twilio_conversations_services_conversations_v1":                              conversationsV1.ResourceServicesConversations(),
			"twilio_conversations_services_conversations_messages_v1":                     conversationsV1.ResourceServicesConversationsMessages(),
			"twilio_conversations_services_conversations_participants_v1":                 conversationsV1.ResourceServicesConversationsParticipants(),
			"twilio_conversations_services_conversations_webhooks_v1":                     conversationsV1.ResourceServicesConversationsWebhooks(),
			"twilio_conversations_services_roles_v1":                                      conversationsV1.ResourceServicesRoles(),
			"twilio_conversations_services_users_v1":                                      conversationsV1.ResourceServicesUsers(),
			"twilio_conversations_users_v1":                                               conversationsV1.ResourceUsers(),
			"twilio_events_sinks_v1":                                                      eventsV1.ResourceSinks(),
			"twilio_events_subscriptions_subscribed_events_v1":                            eventsV1.ResourceSubscriptionsSubscribedEvents(),
			"twilio_events_subscriptions_v1":                                              eventsV1.ResourceSubscriptions(),
			"twilio_flex_channels_v1":                                                     flexV1.ResourceChannels(),
			"twilio_flex_flex_flows_v1":                                                   flexV1.ResourceFlexFlows(),
			"twilio_flex_insights_quality_management_questionnaires_v1":                   flexV1.ResourceInsightsQualityManagementQuestionnaires(),
			"twilio_flex_web_channels_v1":                                                 flexV1.ResourceWebChannels(),
			"twilio_intelligence_services_v2":                                             intelligenceV2.ResourceServices(),
			"twilio_intelligence_transcripts_v2":                                          intelligenceV2.ResourceTranscripts(),
			"twilio_ip_messaging_services_channels_v1":                                    ip_messagingV1.ResourceServicesChannels(),
			"twilio_ip_messaging_credentials_v1":                                          ip_messagingV1.ResourceCredentials(),
			"twilio_ip_messaging_services_channels_invites_v1":                            ip_messagingV1.ResourceServicesChannelsInvites(),
			"twilio_ip_messaging_services_channels_members_v1":                            ip_messagingV1.ResourceServicesChannelsMembers(),
			"twilio_ip_messaging_services_channels_messages_v1":                           ip_messagingV1.ResourceServicesChannelsMessages(),
			"twilio_ip_messaging_services_roles_v1":                                       ip_messagingV1.ResourceServicesRoles(),
			"twilio_ip_messaging_services_v1":                                             ip_messagingV1.ResourceServices(),
			"twilio_ip_messaging_services_users_v1":                                       ip_messagingV1.ResourceServicesUsers(),
			"twilio_ip_messaging_services_channels_v2":                                    ip_messagingV2.ResourceServicesChannels(),
			"twilio_ip_messaging_services_channels_webhooks_v2":                           ip_messagingV2.ResourceServicesChannelsWebhooks(),
			"twilio_ip_messaging_credentials_v2":                                          ip_messagingV2.ResourceCredentials(),
			"twilio_ip_messaging_services_channels_invites_v2":                            ip_messagingV2.ResourceServicesChannelsInvites(),
			"twilio_ip_messaging_services_channels_members_v2":                            ip_messagingV2.ResourceServicesChannelsMembers(),
			"twilio_ip_messaging_services_channels_messages_v2":                           ip_messagingV2.ResourceServicesChannelsMessages(),
			"twilio_ip_messaging_services_roles_v2":                                       ip_messagingV2.ResourceServicesRoles(),
			"twilio_ip_messaging_services_v2":                                             ip_messagingV2.ResourceServices(),
			"twilio_ip_messaging_services_users_v2":                                       ip_messagingV2.ResourceServicesUsers(),
			"twilio_messaging_services_alpha_senders_v1":                                  messagingV1.ResourceServicesAlphaSenders(),
			"twilio_messaging_services_phone_numbers_v1":                                  messagingV1.ResourceServicesPhoneNumbers(),
			"twilio_messaging_services_v1":                                                messagingV1.ResourceServices(),
			"twilio_messaging_services_short_codes_v1":                                    messagingV1.ResourceServicesShortCodes(),
			"twilio_messaging_services_compliance_usa2p_v1":                               messagingV1.ResourceServicesComplianceUsa2p(),
			"twilio_microvisor_configs_v1":                                                microvisorV1.ResourceConfigs(),
			"twilio_microvisor_secrets_v1":                                                microvisorV1.ResourceSecrets(),
			"twilio_microvisor_devices_configs_v1":                                        microvisorV1.ResourceDevicesConfigs(),
			"twilio_microvisor_devices_secrets_v1":                                        microvisorV1.ResourceDevicesSecrets(),
			"twilio_notify_services_bindings_v1":                                          notifyV1.ResourceServicesBindings(),
			"twilio_notify_credentials_v1":                                                notifyV1.ResourceCredentials(),
			"twilio_notify_services_v1":                                                   notifyV1.ResourceServices(),
			"twilio_numbers_regulatory_compliance_bundles_v2":                             numbersV2.ResourceRegulatoryComplianceBundles(),
			"twilio_numbers_regulatory_compliance_end_users_v2":                           numbersV2.ResourceRegulatoryComplianceEndUsers(),
			"twilio_numbers_regulatory_compliance_bundles_item_assignments_v2":            numbersV2.ResourceRegulatoryComplianceBundlesItemAssignments(),
			"twilio_numbers_regulatory_compliance_supporting_documents_v2":                numbersV2.ResourceRegulatoryComplianceSupportingDocuments(),
			"twilio_proxy_services_sessions_participants_v1":                              proxyV1.ResourceServicesSessionsParticipants(),
			"twilio_proxy_services_phone_numbers_v1":                                      proxyV1.ResourceServicesPhoneNumbers(),
			"twilio_proxy_services_v1":                                                    proxyV1.ResourceServices(),
			"twilio_proxy_services_sessions_v1":                                           proxyV1.ResourceServicesSessions(),
			"twilio_proxy_services_short_codes_v1":                                        proxyV1.ResourceServicesShortCodes(),
			"twilio_serverless_services_assets_v1":                                        serverlessV1.ResourceServicesAssets(),
			"twilio_serverless_services_builds_v1":                                        serverlessV1.ResourceServicesBuilds(),
			"twilio_serverless_services_environments_v1":                                  serverlessV1.ResourceServicesEnvironments(),
			"twilio_serverless_services_functions_v1":                                     serverlessV1.ResourceServicesFunctions(),
			"twilio_serverless_services_v1":                                               serverlessV1.ResourceServices(),
			"twilio_serverless_services_environments_variables_v1":                        serverlessV1.ResourceServicesEnvironmentsVariables(),
			"twilio_studio_flows_engagements_v1":                                          studioV1.ResourceFlowsEngagements(),
			"twilio_studio_flows_executions_v1":                                           studioV1.ResourceFlowsExecutions(),
			"twilio_studio_flows_executions_v2":                                           studioV2.ResourceFlowsExecutions(),
			"twilio_studio_flows_v2":                                                      studioV2.ResourceFlows(),
			"twilio_supersim_network_access_profiles_networks_v1":                         supersimV1.ResourceNetworkAccessProfilesNetworks(),
			"twilio_sync_services_documents_v1":                                           syncV1.ResourceServicesDocuments(),
			"twilio_sync_services_v1":                                                     syncV1.ResourceServices(),
			"twilio_sync_services_lists_v1":                                               syncV1.ResourceServicesLists(),
			"twilio_sync_services_lists_items_v1":                                         syncV1.ResourceServicesListsItems(),
			"twilio_sync_services_maps_v1":                                                syncV1.ResourceServicesMaps(),
			"twilio_sync_services_maps_items_v1":                                          syncV1.ResourceServicesMapsItems(),
			"twilio_sync_services_streams_v1":                                             syncV1.ResourceServicesStreams(),
			"twilio_taskrouter_workspaces_activities_v1":                                  taskrouterV1.ResourceWorkspacesActivities(),
			"twilio_taskrouter_workspaces_tasks_v1":                                       taskrouterV1.ResourceWorkspacesTasks(),
			"twilio_taskrouter_workspaces_task_channels_v1":                               taskrouterV1.ResourceWorkspacesTaskChannels(),
			"twilio_taskrouter_workspaces_task_queues_v1":                                 taskrouterV1.ResourceWorkspacesTaskQueues(),
			"twilio_taskrouter_workspaces_workers_v1":                                     taskrouterV1.ResourceWorkspacesWorkers(),
			"twilio_taskrouter_workspaces_workflows_v1":                                   taskrouterV1.ResourceWorkspacesWorkflows(),
			"twilio_taskrouter_workspaces_v1":                                             taskrouterV1.ResourceWorkspaces(),
			"twilio_trunking_trunks_credential_lists_v1":                                  trunkingV1.ResourceTrunksCredentialLists(),
			"twilio_trunking_trunks_ip_access_control_lists_v1":                           trunkingV1.ResourceTrunksIpAccessControlLists(),
			"twilio_trunking_trunks_origination_urls_v1":                                  trunkingV1.ResourceTrunksOriginationUrls(),
			"twilio_trunking_trunks_phone_numbers_v1":                                     trunkingV1.ResourceTrunksPhoneNumbers(),
			"twilio_trunking_trunks_v1":                                                   trunkingV1.ResourceTrunks(),
			"twilio_trusthub_customer_profiles_v1":                                        trusthubV1.ResourceCustomerProfiles(),
			"twilio_trusthub_customer_profiles_channel_endpoint_assignments_v1":           trusthubV1.ResourceCustomerProfilesChannelEndpointAssignments(),
			"twilio_trusthub_customer_profiles_entity_assignments_v1":                     trusthubV1.ResourceCustomerProfilesEntityAssignments(),
			"twilio_trusthub_end_users_v1":                                                trusthubV1.ResourceEndUsers(),
			"twilio_trusthub_supporting_documents_v1":                                     trusthubV1.ResourceSupportingDocuments(),
			"twilio_trusthub_trust_products_v1":                                           trusthubV1.ResourceTrustProducts(),
			"twilio_trusthub_trust_products_channel_endpoint_assignments_v1":              trusthubV1.ResourceTrustProductsChannelEndpointAssignments(),
			"twilio_trusthub_trust_products_entity_assignments_v1":                        trusthubV1.ResourceTrustProductsEntityAssignments(),
			"twilio_verify_services_rate_limits_buckets_v2":                               verifyV2.ResourceServicesRateLimitsBuckets(),
			"twilio_verify_services_entities_v2":                                          verifyV2.ResourceServicesEntities(),
			"twilio_verify_services_messaging_configurations_v2":                          verifyV2.ResourceServicesMessagingConfigurations(),
			"twilio_verify_services_entities_factors_v2":                                  verifyV2.ResourceServicesEntitiesFactors(),
			"twilio_verify_services_rate_limits_v2":                                       verifyV2.ResourceServicesRateLimits(),
			"twilio_verify_safe_list_numbers_v2":                                          verifyV2.ResourceSafeListNumbers(),
			"twilio_verify_services_v2":                                                   verifyV2.ResourceServices(),
			"twilio_verify_services_webhooks_v2":                                          verifyV2.ResourceServicesWebhooks(),
			"twilio_video_compositions_v1":                                                videoV1.ResourceCompositions(),
			"twilio_video_composition_hooks_v1":                                           videoV1.ResourceCompositionHooks(),
			"twilio_voice_byoc_trunks_v1":                                                 voiceV1.ResourceByocTrunks(),
			"twilio_voice_connection_policies_v1":                                         voiceV1.ResourceConnectionPolicies(),
			"twilio_voice_connection_policies_targets_v1":                                 voiceV1.ResourceConnectionPoliciesTargets(),
			"twilio_voice_ip_records_v1":                                                  voiceV1.ResourceIpRecords(),
			"twilio_voice_source_ip_mappings_v1":                                          voiceV1.ResourceSourceIpMappings(),
			"twilio_wireless_commands_v1":                                                 wirelessV1.ResourceCommands(),
			"twilio_wireless_rate_plans_v1":                                               wirelessV1.ResourceRatePlans(),
		},
	}
}
