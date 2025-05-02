<?php

namespace App\Linux;

use Grafana\Foundation\Common as SDKCommon;
use Grafana\Foundation\Dashboard as SDKDashboard;
use Grafana\Foundation\Stat;
use Grafana\Foundation\Units\Constants as Units;

class Host
{
    public static function uptimeStat(): Stat\PanelBuilder
    {
        return Common::defaultStat()
            ->title('Uptime')
            ->description('The duration of time that has passed since the last reboot or system start.')
            ->unit(Units::DURATION_SECONDS)
            ->WithTarget(
                Common::prometheusQuery('time() - node_boot_time_seconds{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', '')
            )
            ->thresholds(
                (new SDKDashboard\ThresholdsConfigBuilder())
                    ->mode(SDKDashboard\ThresholdsMode::absolute())
                    ->steps([
                        new SDKDashboard\Threshold(value: null, color: 'orange'),
                        new SDKDashboard\Threshold(value: 600, color: 'text'),
                    ])
            )
        ;
    }

    public static function hostnameStat(): Stat\PanelBuilder
    {
        return Common::unameStat(
            'Hostname',
            "System's hostname.",
            'nodename',
        );
    }

    public static function kernelVersionStat(): Stat\PanelBuilder
    {
        return Common::unameStat(
            'Kernel version',
            "Kernel version of linux host.",
            'release',
        );
    }

    public static function osStat(): Stat\PanelBuilder
    {
        return Common::defaultStat()
            ->title('OS')
            ->description('Operating system.')
            ->unit(Units::DURATION_SECONDS)
            ->WithTarget(
                Common::tablePrometheusQuery('node_os_info{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', 'A')
            )
            ->reduceOptions(
                (new SDKCommon\ReduceDataOptionsBuilder())
                    ->calcs(['lastNotNull'])
                    ->fields('pretty_name')
            )
            ->colorMode(SDKCommon\BigValueColorMode::none())
        ;
    }
}
