<?php

namespace App\Linux;

use Grafana\Foundation\Common as SDKCommon;
use Grafana\Foundation\Dashboard as SDKDashboard;
use Grafana\Foundation\Dashboard\VariableHide;
use Grafana\Foundation\Prometheus;
use Grafana\Foundation\Table;
use Grafana\Foundation\Timeseries;

class Dashboard
{
    public static function generate(): SDKDashboard\Dashboard
    {
        return (new SDKDashboard\DashboardBuilder(title: '[Example] Linux node / overview'))
            ->uid('example-integration-linux-node')
            ->tags(['generated', 'grafana-foundation-sdk', 'inux-node-integration'])
            ->editable()
            ->tooltip(SDKDashboard\DashboardCursorSync::off())
            ->refresh('30s')
            ->time('now-30m', 'now')
            ->timezone(SDKCommon\Constants::TIME_ZONE_BROWSER)
            ->timepicker(
                (new SDKDashboard\TimePickerBuilder())
                    ->refreshIntervals(['5s', '10s', '30s', '1m', '5m', '15m', '30m', '1h', '2h', '1d'])
                    ->timeOptions(['5m', '15m', '1h', '6h', '12h', '24h', '2d', '7d', '30d'])
            )
            // "Back to fleet" link
            ->link(
                (new SDKDashboard\DashboardLinkBuilder('Back to Linux node / fleet'))
                    ->type(SDKDashboard\DashboardLinkType::link())
                    ->url('/d/node-fleet')
                    ->keepTime(true)
            )
            // Links to other linux-node-related dashboards
            ->link(
                (new SDKDashboard\DashboardLinkBuilder('All Linux node /  dashboards'))
                    ->type(SDKDashboard\DashboardLinkType::dashboards())
                    ->tags(['linux-node-integration'])
                    ->includeVars(true)
                    ->asDropdown(true)
                    ->keepTime(true)
            )
            // "Data source" variable
            ->withVariable(
                (new SDKDashboard\DatasourceVariableBuilder('datasource'))
                    ->label('Data source')
                    ->type('prometheus')
                    ->regex('(?!grafanacloud-usage|grafanacloud-ml-metrics).+')
                    ->multi(false)
            )
            // "Cluster" variable
            ->withVariable(
                Common::queryVariable('cluster', 'Cluster', 'label_values(node_uname_info{job=~"integrations/(node_exporter|unix)"}, cluster)')
                    ->allValue('.*')
            )
            // "Job" variable
            ->withVariable(
                Common::queryVariable('job', 'Job', 'label_values(node_uname_info{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster"}, job)')
                    ->allValue('.+')
            )
            // "Instance" variable
            ->withVariable(
                Common::queryVariable('instance', 'Instance', 'label_values(node_uname_info{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job"}, instance)')
                    ->multi(false)
                    ->includeAll(false)
            )
            // "Loki data source" variable
            ->withVariable(
                (new SDKDashboard\DatasourceVariableBuilder('loki_datasource'))
                    ->label('Loki data source')
                    ->type('loki')
                    ->regex('(?!grafanacloud.+usage-insights|grafanacloud.+alert-state-history).+')
                    ->multi(false)
                    ->hide(VariableHide::hideVariable())
            )
            // Panels
            ->withRow(new SDKDashboard\RowBuilder('Overview'))
            ->withPanel(Host::uptimeStat())
            ->withPanel(Host::hostnameStat())
            ->withPanel(Host::kernelVersionStat())
            ->withPanel(Host::osStat())
            ->withPanel(CPU::countStat())
            ->withPanel(Memory::totalStat())
            ->withPanel(Memory::totalSwapStat())
            ->withPanel(Disk::rootMountSizeStat())
            ->withRow(new SDKDashboard\RowBuilder('CPU'))
            ->withPanel(CPU::usageStat()->height(6))
            ->withPanel(CPU::usageTimeseries()->height(6))
            ->withPanel(CPU::loadAverageTimeseries()->height(6)->span(6))
            ->withRow(new SDKDashboard\RowBuilder('Memory'))
            ->withPanel(Memory::usageStat()->height(6))
            ->withPanel(Memory::usageTimeseries()->height(6))
            ->withRow(new SDKDashboard\RowBuilder('Disk'))
            ->withPanel(Disk::readWriteTimeseries()->height(8))
            ->withPanel(Disk::spaceUsageTable()->height(8))
            ->withRow(new SDKDashboard\RowBuilder('Network'))
            ->withPanel(Network::trafficTimeseries()->height(8))
            ->withPanel(Network::errorsTimeseries()->height(8))
            ->build();
    }

    private static function runningInstancesTable(): Table\PanelBuilder
    {
        return (new Table\PanelBuilder())
            ->title('Running Instances')
            ->description('General statistics of running grafana agent instances.')
            ->height(7)
            ->span(24)
            ->footer((new SDKCommon\TableFooterOptionsBuilder())->countRows(false)->reducer(['sum']))
            ->datasource(new SDKDashboard\DataSourceRef(uid: '$prometheus_datasource'))
            ->unit('s')
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
            ->withOverride(
                new SDKDashboard\MatcherConfig(id: 'byName', options: 'Value #B'),
                [new SDKDashboard\DynamicConfigValue(id: 'unit', value: 's'),],
            )
        ;
    }

    private static function targetSyncTimeseries(): Timeseries\PanelBuilder
    {
        return self::defaultTimeseries()
            ->title('Target Sync')
            ->description('Actual interval to sync the scrape pool.')
            ->datasource(new SDKDashboard\DataSourceRef(uid: '$prometheus_datasource'))
            ->unit('s')
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
            ->unit('short')
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
            ->unit('s')
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
            ->unit('s')
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
            ->unit('short')
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
            ->showPoints(SDKCommon\VisibilityMode::auto())
            ->drawStyle(SDKCommon\GraphDrawStyle::line())
            ->gradientMode(SDKCommon\GraphGradientMode::none())
            ->legend(
                (new SDKCommon\VizLegendOptionsBuilder())
                    ->showLegend(true)
                    ->placement(SDKCommon\LegendPlacement::bottom())
                    ->displayMode(SDKCommon\LegendDisplayMode::list())
            )
            ->tooltip(
                (new SDKCommon\VizTooltipOptionsBuilder())
                    ->mode(SDKCommon\TooltipDisplayMode::single())
                    ->sort(SDKCommon\SortOrder::none())
            )
            ->thresholdsStyle(
                (new SDKCommon\GraphThresholdsStyleConfigBuilder())
                    ->mode(SDKCommon\GraphThresholdsStyleMode::off())
            )
            ->spanNulls(false)
            ->axisBorderShow(false)
        ;
    }
}