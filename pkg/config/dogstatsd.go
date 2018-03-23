// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

package config

var dsdGroup = Group{
	Name:        "Dogstatsd",
	Description: "Dogstatsd enables you to submit custom metrics, via UDP or Unix Sockets",
	Priority:    20,
	Options: []Option{
		{
			Name:        "use_dogstatsd",
			Default:     true,
			Description: "Run Dogstatsd for custom metrics + JMX",
		},
		{
			Name:        "dogstatsd_port",
			Default:     int64(8125),
			Description: "UDP port Dogstatsd will listen to (0 to disable)",
		},
		{
			Name:        "dogstatsd_non_local_traffic",
			Default:     false,
			Description: "Whether dogstatsd should listen to UDP traffic on all interfaces",
		},
		{
			Name:        "bind_host",
			Default:     "localhost",
			Description: "If dogstatsd_non_local_traffic is disabled, host to bind to",
		},
		{
			Name:        "dogstatsd_buffer_size",
			Default:     int64(8192), // 8 KiB
			Description: "The buffer size use to receive statsd packet, in bytes",
			YAMLDesc: `The buffer size use to receive statsd packet, in bytes.
# Adjust accordingly to your client-side buffering settings.`,
		},
		{
			Name:        "dogstatsd_socket",
			Default:     "",
			Description: "Path of the Unix Socket to listen to (empty string to disable)",
		},
		{
			Name:        "dogstatsd_origin_detection",
			Default:     false,
			Description: "Enable origin detection and container tagging for Unix Socket traffic",
		},
		{
			Name:        "statsd_metric_namespace",
			Default:     "",
			Description: "Prefix to add to all custom metric names",
		},
		{
			Name:        "statsd_forward_host",
			Default:     "",
			Description: "Hostname to forward all incomping statsd packets to",
			YAMLDesc: `Dogstatsd can forward every packet received to another statsd server.
# WARNING: Make sure that forwarded packets are regular statsd packets and not "dogstatsd" packets,
# as your other statsd server might not be able to handle them.`,
		},
		{
			Name:        "statsd_forward_port",
			Default:     int64(0),
			Description: "Port to forward all incomping statsd packets to",
			MergeDesc:   true,
		},
		{
			Name:        "dogstatsd_expiry_seconds",
			Default:     float64(300),
			Description: "Duration idle counters are kept before expiring",
		},
		{
			Name:        "dogstatsd_stats_enable",
			Default:     false,
			Description: "Publish internal stats as Go epxvars",
		},
		{
			Name:        "dogstatsd_expiry_seconds",
			Default:     int64(5000),
			Description: "The port to expose stats to",
		},
		{
			Name:        "dogstatsd_stats_buffer",
			Default:     int64(10),
			Description: "How many stat batches to keep in memory",
		},
	},
}

func init() {
	registerGroup(&dsdGroup)
}
