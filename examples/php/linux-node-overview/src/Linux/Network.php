<?php

namespace App\Linux;

use Grafana\Foundation\Common as SDKCommon;
use Grafana\Foundation\Dashboard as SDKDashboard;
use Grafana\Foundation\Timeseries;
use Grafana\Foundation\Units\Constants as Units;

class Network
{
    public static function trafficTimeseries(): Timeseries\PanelBuilder
    {
        return Common::defaultTimeseries()
            ->title('Network traffic')
            ->description('Network traffic (bits per sec) measures data transmitted and received.')
            ->axisLabel('out(-) | in(+)')
            ->axisPlacement(SDKCommon\AxisPlacement::auto())
            ->WithTarget(Common::prometheusQuery('irate(node_network_receive_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])*8
            # only show interfaces that had traffic change at least once during selected dashboard interval:
            and
            increase(
                node_network_receive_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__range]
                ) > 0', '{{ device }} received'))
            ->WithTarget(Common::prometheusQuery('irate(node_network_transmit_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])*8
            # only show interfaces that had traffic change at least once during selected dashboard interval:
            and
            increase(
                node_network_transmit_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__range]
                ) > 0', '{{ device }} transmitted'))
            ->decimals(1)
            ->unit(Units::BITS_PER_SECOND_SI)
            ->gradientMode(SDKCommon\GraphGradientMode::opacity())
            ->overrideByRegexp('/transmit|tx|out/', [
                new SDKDashboard\DynamicConfigValue(id: 'custom.transform', value: 'negative-Y'),
            ])
        ;
    }

    public static function errorsTimeseries(): Timeseries\PanelBuilder
    {
        return Common::defaultTimeseries()
            ->title('Network errors and dropped packets')
            ->description('**Network errors**:

Network errors refer to issues that occur during the transmission of data across a network. 

These errors can result from various factors, including physical issues, jitter, collisions, noise and interference.

Monitoring network errors is essential for diagnosing and resolving issues, as they can indicate problems with network hardware or environmental factors affecting network quality.

**Dropped packets**:

Dropped packets occur when data packets traveling through a network are intentionally discarded or lost due to congestion, resource limitations, or network configuration issues. 

Common causes include network congestion, buffer overflows, QoS settings, and network errors, as corrupted or incomplete packets may be discarded by receiving devices.

Dropped packets can impact network performance and lead to issues such as degraded voice or video quality in real-time applications.
')
            ->axisLabel('out(-) | in(+)')
            ->axisPlacement(SDKCommon\AxisPlacement::auto())
            ->withTarget(Common::prometheusQuery('irate(node_network_transmit_errs_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])>0', '{{ device }} errors transmitted'))
            ->withTarget(Common::prometheusQuery('irate(node_network_receive_errs_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])>0', '{{ device }} errors received'))
            ->withTarget(Common::prometheusQuery('irate(node_network_transmit_drop_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])>0', '{{ device }} transmitted dropped'))
            ->withTarget(Common::prometheusQuery('irate(node_network_receive_drop_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}[$__rate_interval])>0', '{{ device }} received dropped'))
            ->decimals(1)
            ->unit(Units::PACKETS_PER_SECOND)
            ->gradientMode(SDKCommon\GraphGradientMode::opacity())
            ->overrideByRegexp('/transmit|tx|out/', [
                new SDKDashboard\DynamicConfigValue(id: 'custom.transform', value: 'negative-Y'),
            ])
        ;
    }
}
