<?php

namespace App\Linux;

use Grafana\Foundation\Common as SDKCommon;
use Grafana\Foundation\Dashboard as SDKDashboard;
use Grafana\Foundation\Dashboard\VariableHide;

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
}
