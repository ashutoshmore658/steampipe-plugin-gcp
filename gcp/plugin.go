/*
Package gcp implements a steampipe plugin for gcp.

This plugin provides data that Steampipe uses to present foreign
tables that represent GCP resources.
*/
package gcp

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

const pluginName = "steampipe-plugin-gcp"

// Plugin creates this (gcp) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		TableMap: map[string]*plugin.Table{
			"gcp_iam_policy":                      tableGcpIAMPolicy(ctx),
			"gcp_iam_role":                        tableGcpIamRole(ctx),
			"gcp_logging_exclusion":               tableGcpLoggingExclusion(ctx),
			"gcp_logging_metric":                  tableGcpLoggingMetric(ctx),
			"gcp_logging_sink":                    tableGcpLoggingSink(ctx),
			"gcp_monitoring_group":                tableGcpMonitoringGroup(ctx),
			"gcp_monitoring_notification_channel": tableGcpMonitoringNotificationChannel(ctx),
			"gcp_project_service":                 tableGcpProjectService(ctx),
			"gcp_pubsub_snapshot":                 tableGcpPubSubSnapshot(ctx),
			"gcp_pubsub_subscription":             tableGcpPubSubSubscription(ctx),
			"gcp_pubsub_topic":                    tableGcpPubSubTopic(ctx),
			"gcp_service_account":                 tableGcpServiceAccount(ctx),
			"gcp_service_account_key":             tableGcpServiceAccountKey(ctx),
		},
	}

	return p
}