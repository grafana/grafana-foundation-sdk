<?php

namespace App\Linux;

use Grafana\Foundation\Common as SDKCommon;
use Grafana\Foundation\Dashboard as SDKDashboard;
use Grafana\Foundation\Stat;
use Grafana\Foundation\Timeseries;
use Grafana\Foundation\Units\Constants as Units;

class Memory
{
    public static function totalStat(): Stat\PanelBuilder
    {
        return Common::defaultStat()
            ->title('Memory total')
            ->description("Amount of random-access memory (RAM) installed.\nIt represents the system's available working memory that applications and the operating system use to perform tasks.\nA higher memory total generally leads to better system performance and the ability to run more demanding applications and processes simultaneously.")
            ->WithTarget(
                Common::prometheusQuery('node_memory_MemTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', '')
            )
            ->decimals(2)
            ->unit(Units::BYTES_IEC)
            ->colorMode(SDKCommon\BigValueColorMode::none())
        ;
    }

    public static function totalSwapStat(): Stat\PanelBuilder
    {
        return Common::defaultStat()
            ->title('Total swap')
            ->description("Total swap available.\n\nSwap is a space on a storage device (usually a dedicated swap partition or a swap file) \nused as virtual memory when the physical RAM (random-access memory) is fully utilized.\nSwap space helps prevent memory-related performance issues by temporarily transferring less-used data from RAM to disk,\nfreeing up physical memory for active processes and applications.")
            ->WithTarget(
                Common::prometheusQuery('node_memory_SwapTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', '')
            )
            ->unit(Units::BYTES_IEC)
            ->colorMode(SDKCommon\BigValueColorMode::none())
        ;
    }

    public static function usageStat(): Stat\PanelBuilder
    {
        $query = <<<'PROMQL'
100 -
(
    avg by (instance) (node_memory_MemAvailable_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}) /
    avg by (instance) (node_memory_MemTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"})
* 100
)
PROMQL;
        
        return Common::defaultStat()
            ->title('Memory usage')
            ->description("RAM (random-access memory) currently in use by the operating system and running applications, in percent.")
            ->WithTarget(Common::prometheusQuery($query, ''))
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
            ->graphMode(SDKCommon\BigValueGraphMode::area())
        ;
    }

    public static function usageTimeseries(): Timeseries\PanelBuilder
    {
        $query = <<<'PROMQL'
(
node_memory_MemTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}
-
node_memory_MemFree_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}
-
node_memory_Buffers_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}
-
node_memory_Cached_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}
)
PROMQL;

        return Common::defaultTimeseries()
            ->title('Memory usage')
            ->description("- Used: The amount of physical memory currently in use by the system.\n- Cached: The amount of physical memory used for caching data from disk. The Linux kernel uses available memory to cache data that is read from or written to disk. This helps speed up disk access times.\n- Free: The amount of physical memory that is currently not in use.\n- Buffers: The amount of physical memory used for temporary storage of data being transferred between devices or applications.\n- Available: The amount of physical memory that is available for use by applications. This takes into account memory that is currently being used for caching but can be freed up if needed.")
            ->span(18)
            ->WithTarget(Common::prometheusQuery($query, 'Memory used'))
            ->WithTarget(Common::prometheusQuery('node_memory_Cached_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', 'Memory cached'))
            ->WithTarget(Common::prometheusQuery('node_memory_MemAvailable_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', 'Memory available'))
            ->WithTarget(Common::prometheusQuery('node_memory_Buffers_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', 'Memory buffers'))
            ->WithTarget(Common::prometheusQuery('node_memory_MemFree_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', 'Memory free'))
            ->WithTarget(Common::prometheusQuery('node_memory_MemTotal_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', 'Memory total'))
            ->decimals(2)
            ->unit(Units::BYTES_IEC)
            ->gradientMode(SDKCommon\GraphGradientMode::opacity())
            ->overrideByRegexp('.*(T|t)otal.*', [
                (new SDKDashboard\DynamicConfigValue(id: 'color', value: ['fixedColor' => 'light-orange', 'mode' => 'fixed'])),
                (new SDKDashboard\DynamicConfigValue(id: 'custom.fillOpacity', value: 0)),
                (new SDKDashboard\DynamicConfigValue(id: 'custom.lineStyle', value: ['dash' => [10, 10], 'fill' => 'dash'])),
            ])
        ;
    }
}
