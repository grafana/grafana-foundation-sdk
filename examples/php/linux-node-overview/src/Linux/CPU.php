<?php

namespace App\Linux;

use Grafana\Foundation\Common as SDKCommon;
use Grafana\Foundation\Dashboard as SDKDashboard;
use Grafana\Foundation\Stat;
use Grafana\Foundation\Timeseries;
use Grafana\Foundation\Units\Constants as Units;

class CPU
{
    public static function countStat(): Stat\PanelBuilder
    {
        return Common::defaultStat()
            ->title('CPU count')
            ->description('CPU count is the number of processor cores or central processing units (CPUs) in a computer,\ndetermining its processing capability and ability to handle tasks concurrently.')
            ->WithTarget(
                Common::prometheusQuery('count without (cpu) (node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mode="idle"})', 'Cores')
            )
            ->decimals(0)
            ->colorMode(SDKCommon\BigValueColorMode::none())
        ;
    }

    public static function usageStat(): Stat\PanelBuilder
    {
        $query = <<<'PROMQL'
(((count by (instance) (count(node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}) by (cpu, instance))) 
- 
avg by (instance) (sum by (instance, mode)(irate(node_cpu_seconds_total{mode='idle',job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])))) * 100) 
/ 
count by(instance) (count(node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}) by (cpu, instance))
PROMQL;

        return Common::defaultStat()
            ->title('CPU usage')
            ->description("Total CPU utilization percent is a metric that indicates the overall level of central processing unit (CPU) usage in a computer system.\nIt represents the combined load placed on all CPU cores or processors.\n\nFor instance, if the total CPU utilization percent is 50%, it means that,\non average, half of the CPU's processing capacity is being used to execute tasks. A higher percentage indicates that the CPU is working more intensively, potentially leading to system slowdowns if it remains consistently high.")
            ->min(0)
            ->max(100)
            ->unit(Units::PERCENT)
            ->WithTarget(
                Common::prometheusQuery($query, '')
            )
            ->thresholds(
                (new SDKDashboard\ThresholdsConfigBuilder())
                    ->mode(SDKDashboard\ThresholdsMode::absolute())
                    ->steps([
                        new SDKDashboard\Threshold(value: null, color: 'green'),
                        new SDKDashboard\Threshold(value: 80, color: 'red'),
                    ])
            )
            ->graphMode(SDKCommon\BigValueGraphMode::area())
        ;
    }

    public static function usageTimeseries(): Timeseries\PanelBuilder
    {
        $query = <<<'PROMQL'
(
    (1 - sum without (mode) (rate(node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mode=~"idle|iowait|steal"}[$__rate_interval])))
/ ignoring(cpu) group_left
    count without (cpu, mode) (node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mode="idle"})
) * 100
PROMQL;

        return Common::defaultTimeseries()
            ->title('CPU usage')
            ->description("Total CPU utilization percent is a metric that indicates the overall level of central processing unit (CPU) usage in a computer system.\nIt represents the combined load placed on all CPU cores or processors.\n\nFor instance, if the total CPU utilization percent is 50%, it means that,\non average, half of the CPU's processing capacity is being used to execute tasks. A higher percentage indicates that the CPU is working more intensively, potentially leading to system slowdowns if it remains consistently high.")
            ->WithTarget(Common::prometheusQuery($query, 'CPU {{cpu}}'))
            ->min(0)
            ->max(100)
            ->unit(Units::PERCENT)
            ->thresholds(
                (new SDKDashboard\ThresholdsConfigBuilder())
                    ->mode(SDKDashboard\ThresholdsMode::absolute())
                    ->steps([
                        new SDKDashboard\Threshold(value: null, color: 'green'),
                        new SDKDashboard\Threshold(value: 80, color: 'red'),
                    ])
            )
        ;
    }

    public static function loadAverageTimeseries(): Timeseries\PanelBuilder
    {
        return Common::defaultTimeseries()
            ->title('Load average')
            ->description("System load average over the previous 1, 5, and 15 minute ranges.\n\nA measurement of how many processes are waiting for CPU cycles. The maximum number is the number of CPU cores for the node.")
            ->WithTarget(
                Common::prometheusQuery('node_load1{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', '1m'),
            )
            ->WithTarget(
                Common::prometheusQuery('node_load5{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', '5m'),
            )
            ->WithTarget(
                Common::prometheusQuery('node_load15{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', '15m'),
            )
            ->WithTarget(
                Common::prometheusQuery('count without (cpu) (node_cpu_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mode="idle"})', 'Cores'),
            )
            ->min(0)
            ->decimals(2)
            ->unit(Units::SHORT)
            ->thresholds(
                (new SDKDashboard\ThresholdsConfigBuilder())
                    ->mode(SDKDashboard\ThresholdsMode::absolute())
                    ->steps([
                        new SDKDashboard\Threshold(value: null, color: 'green'),
                        new SDKDashboard\Threshold(value: 80, color: 'red'),
                    ])
            )
            ->overrideByRegexp('Cores', [
                (new SDKDashboard\DynamicConfigValue(id: 'color', value: ['fixedColor' => 'light-orange', 'mode' => 'fixed'])),
                (new SDKDashboard\DynamicConfigValue(id: 'custom.fillOpacity', value: 0)),
                (new SDKDashboard\DynamicConfigValue(id: 'custom.lineStyle', value: ['dash' => [10, 10], 'fill' => 'dash'])),
            ])
        ;
    }
}
