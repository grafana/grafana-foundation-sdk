package main

import (
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func avgScrapeIntervalDurationTimeseries() *timeseries.PanelBuilder {
	return defaultTimeseries().
		Title("Average Scrape Interval Duration").
		Description("Actual intervals between scrapes.").
		Datasource(datasourceRef("$prometheus_datasource")).
		Unit(units.Seconds).
		WithTarget(prometheusQuery(
			`rate(prometheus_target_interval_length_seconds_sum{job=~"$job", instance=~"$instance"}[$__rate_interval]) / rate(prometheus_target_interval_length_seconds_count{job=~"$job", instance=~"$instance"}[$__rate_interval])`,
			"{{instance}} {{interval}} configured",
		).IntervalFactor(2))
}

func scrapeFailuresTimeseries() *timeseries.PanelBuilder {
	return defaultTimeseries().
		Title("Scrape failures").
		Description("Shows all scrape failures (sample limit exceeded, duplicate, out of bounds, out of order).").
		Datasource(datasourceRef("$prometheus_datasource")).
		Unit(units.Short).
		WithTarget(prometheusQuery(
			`sum by (job) (rate(prometheus_target_scrapes_exceeded_sample_limit_total{job=~"$job", instance=~"$instance"}[$__rate_interval]))`,
			"exceeded sample limit: {{job}}",
		)).
		WithTarget(prometheusQuery(
			`sum by (job) (rate(prometheus_target_scrapes_sample_duplicate_timestamp_total{job=~"$job", instance=~"$instance"}[$__rate_interval]))`,
			"duplicate timestamp: {{job}}",
		)).
		WithTarget(prometheusQuery(
			`sum by (job) (rate(prometheus_target_scrapes_sample_out_of_bounds_total{job=~"$job", instance=~"$instance"}[$__rate_interval]))`,
			"out of bounds: {{job}}",
		)).
		WithTarget(prometheusQuery(
			`sum by (job) (rate(prometheus_target_scrapes_sample_out_of_order_total{job=~"$job", instance=~"$instance"}[$__rate_interval]))`,
			"out of order: {{job}}",
		))
}

func appendedSamplesTimeseries() *timeseries.PanelBuilder {
	return defaultTimeseries().
		Title("Appended Samples").
		Description("Total number of samples appended to the WAL.").
		Datasource(datasourceRef("$prometheus_datasource")).
		Unit(units.Short).
		WithTarget(prometheusQuery(
			`sum by (job, instance_group_name) (rate(agent_wal_samples_appended_total{job=~"$job", instance=~"$instance"}[$__rate_interval]))`,
			"{{job}} {{instance_group_name}}",
		))
}
