package main

import (
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func targetSyncTimeseries() *timeseries.PanelBuilder {
	return defaultTimeseries().
		Title("Target Sync").
		Description("Actual interval to sync the scrape pool.").
		Datasource(datasourceRef("$prometheus_datasource")).
		Unit(units.Seconds).
		WithTarget(prometheusQuery(
			`sum(rate(prometheus_target_sync_length_seconds_sum{job=~"$job", instance=~"$instance"}[$__rate_interval])) by (instance, scrape_job)`,
			"{{instance}}/{{scrape_job}}",
		))
}

func targetsTimeseries() *timeseries.PanelBuilder {
	return defaultTimeseries().
		Title("Targets").
		Description("Discovered targets by prometheus service discovery.").
		Datasource(datasourceRef("$prometheus_datasource")).
		Unit(units.Short).
		WithTarget(prometheusQuery(
			`sum by (instance) (prometheus_sd_discovered_targets{job=~"$job", instance=~"$instance"})`,
			"{{instance}}",
		))
}
