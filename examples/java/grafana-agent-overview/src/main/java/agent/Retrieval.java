package agent;

import com.grafana.foundation.timeseries.PanelBuilder;
import com.grafana.foundation.units.Constants;

import static agent.Common.*;

public class Retrieval {
    static PanelBuilder avgScrapeIntervalDurationTimeSeries() {
        return defaultTimeSeries().
                title("Average Scrape Interval Duration").
                description("Actual intervals between scrapes.").
                datasource(datasourceRef()).
                unit(Constants.Seconds).
                withTarget(
                        prometheusQuery(
                                "rate(prometheus_target_interval_length_seconds_sum{job=~\"$job\", instance=~\"$instance\"}[$__rate_interval]) / rate(prometheus_target_interval_length_seconds_count{job=~\"$job\", instance=~\"$instance\"}[$__rate_interval])",
                                "{{instance}} {{interval}} configured"
                        ).
                        intervalFactor(2.0)
                );
    }

    static PanelBuilder scrapeFailuresTimeSeries() {
        return defaultTimeSeries().
                title("Scrape failures").
                description("Shows all scrape failures (sample limit exceeded, duplicate, out of bounds, out of order).").
                datasource(datasourceRef()).
                unit(Constants.Short).
                withTarget(
                        prometheusQuery(
                                "sum by (job) (rate(prometheus_target_scrapes_exceeded_sample_limit_total{job=~\"$job\", instance=~\"$instance\"}[$__rate_interval])",
                                "exceeded sample limit: {{job}}"
                        )
                ).
                withTarget(
                        prometheusQuery(
                                "sum by (job) (rate(prometheus_target_scrapes_sample_duplicate_timestamp_total{job=~\"$job\", instance=~\"$instance\"}[$__rate_interval]))",
                                "duplicate timestamp: {{job}}"
                        )
                ).
                withTarget(
                        prometheusQuery(
                                "sum by (job) (rate(prometheus_target_scrapes_sample_out_of_bounds_total{job=~\"$job\", instance=~\"$instance\"}[$__rate_interval]))",
                                "out of bounds: {{job}}"
                        )
                ).
                withTarget(
                        prometheusQuery(
                                "sum by (job) (rate(prometheus_target_scrapes_sample_out_of_order_total{job=~\"$job\", instance=~\"$instance\"}[$__rate_interval]))",
                                "out of order: {{job}}"
                        )
                );
    }

    static PanelBuilder appendedSamplesTimeSeries() {
        return defaultTimeSeries().
                title("Appended Samples").
                description("Total number of samples appended to the WAL.").
                datasource(datasourceRef()).
                unit(Constants.Short).
                withTarget(
                        prometheusQuery(
                                "sum by (job, instance_group_name) (rate(agent_wal_samples_appended_total{job=~\"$job\", instance=~\"$instance\"}[$__rate_interval]))",
                                "{{job}} {{instance_group_name}}"
                        )
                );
    }
}
