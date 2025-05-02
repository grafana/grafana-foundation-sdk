<?php

namespace App\Agent;

use Grafana\Foundation\Common;
use Grafana\Foundation\Dashboard as SDKDashboard;
use Grafana\Foundation\Prometheus;
use Grafana\Foundation\Table;
use Grafana\Foundation\Timeseries;
use Grafana\Foundation\Units\Constants as Units;

class Dashboard
{
    public static function generate(): SDKDashboard\Dashboard
    {
        return (new SDKDashboard\DashboardBuilder(title: '[Example] Grafana Agent Overview'))
            ->uid('example-integration-agent')
            ->tags(['generated', 'grafana-foundation-sdk', 'grafana-agent-integration'])
            ->editable()
            ->tooltip(SDKDashboard\DashboardCursorSync::off())
            ->refresh('30s')
            ->time('now-30m', 'now')
            ->timezone(Common\Constants::TIME_ZONE_BROWSER)
            ->timepicker(
                (new SDKDashboard\TimePickerBuilder())
                    ->refreshIntervals(['5s', '10s', '30s', '1m', '5m', '15m', '30m', '1h', '2h', '1d'])
            )
            // Links to other agent-related dashboards
            ->link(
                (new SDKDashboard\DashboardLinkBuilder('Grafana Agent Dashboards'))
                    ->type(SDKDashboard\DashboardLinkType::dashboards())
                    ->icon('external link')
                    ->tags(['grafana-agent-integration'])
                    ->includeVars(true)
                    ->keepTime(true)
            )
            // "Data source" variable
            ->withVariable(
                (new SDKDashboard\DatasourceVariableBuilder('prometheus_datasource'))
                    ->label('Data source')
                    ->type('prometheus')
                    ->regex('(?!grafanacloud-usage|grafanacloud-ml-metrics).+')
                    ->multi(false)
            )
            // "Job" variable
            ->withVariable(
                (new SDKDashboard\QueryVariableBuilder('job'))
                    ->label('Job')
                    ->query('label_values(agent_build_info, job)')
                    ->datasource(new SDKDashboard\DataSourceRef(uid: '$prometheus_datasource'))
                    ->current(
                        new SDKDashboard\VariableOption(
                            selected: true,
                            text: 'All',
                            value: '$__all',
                        )
                    )
                    ->refresh(SDKDashboard\VariableRefresh::onTimeRangeChanged())
                    ->sort(SDKDashboard\VariableSort::alphabeticalAsc())
                    ->multi(true)
                    ->includeAll(true)
            )
            // "Instance" variable
            ->withVariable(
                (new SDKDashboard\QueryVariableBuilder('instance'))
                    ->label('Instance')
                    ->query('label_values(agent_build_info{job=~"$job"}, instance)')
                    ->datasource(new SDKDashboard\DataSourceRef(uid: '$prometheus_datasource'))
                    ->current(
                        new SDKDashboard\VariableOption(
                            selected: true,
                            text: 'All',
                            value: '$__all',
                        )
                    )
                    ->refresh(SDKDashboard\VariableRefresh::onTimeRangeChanged())
                    ->sort(SDKDashboard\VariableSort::alphabeticalAsc())
                    ->multi(true)
                    ->includeAll(true)
            )
            // Panels
            ->withRow(new SDKDashboard\RowBuilder('Overview'))
            ->withPanel(self::runningInstancesTable())
            ->withRow(new SDKDashboard\RowBuilder('Prometheus Discovery'))
            ->withPanel(self::targetSyncTimeseries())
            ->withPanel(self::targetsTimeseries())
            ->withRow(new SDKDashboard\RowBuilder('Prometheus Retrieval'))
            ->withPanel(self::avgScrapeIntervalDurationTimeseries())
            ->withPanel(self::scrapeFailuresTimeseries())
            ->withPanel(self::appendedSamplesTimeseries())
            ->build();
    }

    private static function runningInstancesTable(): Table\PanelBuilder
    {
        return (new Table\PanelBuilder())
            ->title('Running Instances')
            ->description('General statistics of running grafana agent instances.')
            ->height(7)
            ->span(24)
            ->footer((new Common\TableFooterOptionsBuilder())->countRows(false)->reducer(['sum']))
            ->datasource(new SDKDashboard\DataSourceRef(uid: '$prometheus_datasource'))
            ->unit(Units::SECONDS)
            ->withTarget(self::tablePrometheusQuery(
                'count by (instance, version) (agent_build_info{job=~"$job", instance=~"$instance"})',
                'A',
            ))
            ->withTarget(self::tablePrometheusQuery(
                'max by (instance) (time() - process_start_time_seconds{job=~"$job", instance=~"$instance"})',
                'B',
            ))
		    // Transformations
            ->withTransformation(new SDKDashboard\DataTransformerConfig(
                id: 'merge',
                options: [],
            ))
            ->withTransformation(new SDKDashboard\DataTransformerConfig(
                id: 'organize',
                options: [
                    'excludeByName' => [
                        'Time' => true,
                        'Value #A' => true,
                    ],
                    'renameByName' => [
                        'Value #B' => 'Uptime',
                    ],
                ],
            ))
            // Overrides
            ->overrideByName('Value #B', [
                new SDKDashboard\DynamicConfigValue(id: 'unit', value: 's'),
            ])
        ;
    }

    private static function targetSyncTimeseries(): Timeseries\PanelBuilder
    {
        return self::defaultTimeseries()
            ->title('Target Sync')
            ->description('Actual interval to sync the scrape pool.')
            ->datasource(new SDKDashboard\DataSourceRef(uid: '$prometheus_datasource'))
            ->unit(Units::SECONDS)
            ->withTarget(self::prometheusQuery(
                'sum(rate(prometheus_target_sync_length_seconds_sum{job=~"$job", instance=~"$instance"}[$__rate_interval])) by (instance, scrape_job)',
                '{{instance}}/{{scrape_job}}',
            ))
        ;
    }

    private static function targetsTimeseries(): Timeseries\PanelBuilder
    {
        return self::defaultTimeseries()
            ->title('Targets')
            ->description('Discovered targets by prometheus service discovery.')
            ->datasource(new SDKDashboard\DataSourceRef(uid: '$prometheus_datasource'))
            ->unit(Units::SHORT)
            ->withTarget(self::prometheusQuery(
                'sum by (instance) (prometheus_sd_discovered_targets{job=~"$job", instance=~"$instance"})',
                '{{instance}}',
            ))
        ;
    }

    private static function avgScrapeIntervalDurationTimeseries(): Timeseries\PanelBuilder
    {
        return self::defaultTimeseries()
            ->title('Average Scrape Interval Duration')
            ->description('Actual intervals between scrapes.')
            ->datasource(new SDKDashboard\DataSourceRef(uid: '$prometheus_datasource'))
            ->unit(Units::SECONDS)
            ->withTarget(self::prometheusQuery(
                'rate(prometheus_target_interval_length_seconds_sum{job=~"$job", instance=~"$instance"}[$__rate_interval]) / rate(prometheus_target_interval_length_seconds_count{job=~"$job", instance=~"$instance"}[$__rate_interval])',
                '{{instance}} {{interval}} configured',
            ))
        ;
    }

    private static function scrapeFailuresTimeseries(): Timeseries\PanelBuilder
    {
        return self::defaultTimeseries()
            ->title('Scrape failures')
            ->description('Shows all scrape failures (sample limit exceeded, duplicate, out of bounds, out of order).')
            ->datasource(new SDKDashboard\DataSourceRef(uid: '$prometheus_datasource'))
            ->unit(Units::SECONDS)
            ->withTarget(self::prometheusQuery(
                'sum by (job) (rate(prometheus_target_scrapes_exceeded_sample_limit_total{job=~"$job", instance=~"$instance"}[$__rate_interval]))',
                'exceeded sample limit: {{job}}',
            ))
            ->withTarget(self::prometheusQuery(
                'sum by (job) (rate(prometheus_target_scrapes_sample_duplicate_timestamp_total{job=~"$job", instance=~"$instance"}[$__rate_interval]))',
                'duplicate timestamp: {{job}}',
            ))
            ->withTarget(self::prometheusQuery(
                'sum by (job) (rate(prometheus_target_scrapes_sample_out_of_bounds_total{job=~"$job", instance=~"$instance"}[$__rate_interval]))',
                'out of bounds: {{job}}',
            ))
            ->withTarget(self::prometheusQuery(
                'sum by (job) (rate(prometheus_target_scrapes_sample_out_of_order_total{job=~"$job", instance=~"$instance"}[$__rate_interval]))',
                'out of order: {{job}}',
            ))
        ;
    }

    private static function appendedSamplesTimeseries(): Timeseries\PanelBuilder
    {
        return self::defaultTimeseries()
            ->title('Appended Samples')
            ->description('Total number of samples appended to the WAL.')
            ->datasource(new SDKDashboard\DataSourceRef(uid: '$prometheus_datasource'))
            ->unit(Units::SHORT)
            ->withTarget(self::prometheusQuery(
                'sum by (job, instance_group_name) (rate(agent_wal_samples_appended_total{job=~"$job", instance=~"$instance"}[$__rate_interval]))',
                '{{job}} {{instance_group_name}}',
            ))
        ;
    }

    private static function tablePrometheusQuery(string $query, string $refId): Prometheus\DataqueryBuilder
    {
        return (new Prometheus\DataqueryBuilder())
            ->expr($query)
            ->format(Prometheus\PromQueryFormat::table())
            ->instant()
            ->refId($refId)
        ;
    }

    private static function prometheusQuery(string $query, string $legend): Prometheus\DataqueryBuilder
    {
        return (new Prometheus\DataqueryBuilder())
            ->expr($query)
            ->legendFormat($legend)
        ;
    }

    private static function defaultTimeseries(): Timeseries\PanelBuilder
    {
        return (new Timeseries\PanelBuilder())
            ->height(7)
            ->span(12)
            ->lineWidth(1)
            ->fillOpacity(0)
            ->pointSize(5)
            ->showPoints(Common\VisibilityMode::auto())
            ->drawStyle(Common\GraphDrawStyle::line())
            ->gradientMode(Common\GraphGradientMode::none())
            ->legend(
                (new Common\VizLegendOptionsBuilder())
                    ->showLegend(true)
                    ->placement(Common\LegendPlacement::bottom())
                    ->displayMode(Common\LegendDisplayMode::list())
            )
            ->tooltip(
                (new Common\VizTooltipOptionsBuilder())
                    ->mode(Common\TooltipDisplayMode::single())
                    ->sort(Common\SortOrder::none())
            )
            ->thresholdsStyle(
                (new Common\GraphThresholdsStyleConfigBuilder())
                    ->mode(Common\GraphThresholdsStyleMode::off())
            )
            ->spanNulls(false)
            ->axisBorderShow(false)
        ;
    }
}
