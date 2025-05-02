<?php

namespace App\Linux;

use Grafana\Foundation\Common as SDKCommon;
use Grafana\Foundation\Dashboard as SDKDashboard;
use Grafana\Foundation\Stat;
use Grafana\Foundation\Table;
use Grafana\Foundation\Timeseries;
use Grafana\Foundation\Units\Constants as Units;

class Disk
{
    public static function rootMountSizeStat(): Stat\PanelBuilder
    {
        return Common::defaultStat()
            ->title('Root mount size')
            ->description('Total capacity on the primary mount point /.')
            ->WithTarget(
                Common::prometheusQuery('node_filesystem_size_bytes{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", mountpoint="/", fstype!="rootfs"}', '')
            )
            ->unit(Units::BYTES_IEC)
            ->colorMode(SDKCommon\BigValueColorMode::none())
        ;
    }

    public static function readWriteTimeseries(): Timeseries\PanelBuilder
    {
        return Common::defaultTimeseries()
            ->title('Disk reads/writes')
            ->description('Disk read/writes in bytes per second.')
            ->WithTarget(Common::prometheusQuery('irate(node_disk_read_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", device!=""}[$__rate_interval])', '{{ device }} read'))
            ->WithTarget(Common::prometheusQuery('irate(node_disk_written_bytes_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", device!=""}[$__rate_interval])', '{{ device }} written'))
            ->WithTarget(Common::prometheusQuery('irate(node_disk_io_time_seconds_total{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance", device!=""}[$__rate_interval])', '{{ device }} io util'))
            ->decimals(0)
            ->unit(Units::BYTES_PER_SECOND_SI)
            ->fillOpacity(1)
            ->gradientMode(SDKCommon\GraphGradientMode::opacity())
            ->overrideByRegexp('/time|used|busy|util/', [
                (new SDKDashboard\DynamicConfigValue(id: 'custom.axisSoftMax', value: 100)),
                (new SDKDashboard\DynamicConfigValue(id: 'custom.drawStyle', value: 'points')),
                (new SDKDashboard\DynamicConfigValue(id: 'unit', value: Units::PERCENT)),
            ])
        ;
    }

    public static function spaceUsageTable(): Table\PanelBuilder
    {
        return (new Table\PanelBuilder())
            ->title('Disk space usage')
            ->description('Disk utilisation in percent, by mountpoint. Some duplication can occur if the same filesystem is mounted in multiple locations.')
            ->footer(
                (new SDKCommon\TableFooterOptionsBuilder())
                    ->countRows(false)
                    ->reducer(['sum'])
            )
            ->datasource(new SDKDashboard\DataSourceRef(uid: '$datasource'))

            ->targets([
                Common::tablePrometheusQuery('node_filesystem_size_bytes{fstype!="", mountpoint!="", job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', 'TOTAL')
                    ->instant(),
                Common::tablePrometheusQuery('node_filesystem_avail_bytes{fstype!="", mountpoint!="", job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', 'FREE')
                    ->legendFormat('{{ mountpoint }} free')
                    ->instant(),
            ])
            ->unit(Units::BYTES_IEC)
            ->thresholds(
                (new SDKDashboard\ThresholdsConfigBuilder())
                    ->mode(SDKDashboard\ThresholdsMode::absolute())
                    ->steps([
                        new SDKDashboard\Threshold(value: null, color: 'light-blue'),
                        new SDKDashboard\Threshold(value: 0.8, color: 'light-yellow'),
                        new SDKDashboard\Threshold(value: 0.9, color: 'light-red'),
                    ])
            )
            // Transformations
            ->withTransformation(
                new SDKDashboard\DataTransformerConfig(
                    id: 'groupBy',
                    options: [
                        "fields" => [
                            "Value #FREE" => [
                                "aggregations" => ["lastNotNull"],
                                "operation" => "aggregate",
                            ],
                            "Value #TOTAL" => [
                                "aggregations" => ["lastNotNull"],
                                "operation" => "aggregate",
                            ],
                            "mountpoint" => [
                                "aggregations" => [],
                                "operation" => "groupby",
                            ],
                        ]
                    ],
                )
            )
            ->withTransformation(new SDKDashboard\DataTransformerConfig(id: 'merge', options: []))
            ->withTransformation(
                new SDKDashboard\DataTransformerConfig(
                    id: 'calculateField',
                    options: [
                        "alias" => "Used",
                        "binary" => [
                            "left" => "Value #TOTAL (lastNotNull)",
                            "operator" => "-",
                            "reducer" => "sum",
                            "right" => "Value #FREE (lastNotNull)",
                        ],
                        "mode" => "binary",
                        "reduce" => ["reducer" => "sum"],
                    ],
                )
            )
            ->withTransformation(
                new SDKDashboard\DataTransformerConfig(
                    id: 'calculateField',
                    options: [
                        "alias" => "Used, %",
                        "binary" => [
                            "left" => "Used",
                            "operator" => "/",
                            "reducer" => "sum",
                            "right" => "Value #TOTAL (lastNotNull)",
                        ],
                        "mode" => "binary",
                        "reduce" => ["reducer" => "sum"],
                    ],
                )
            )
            ->withTransformation(
                new SDKDashboard\DataTransformerConfig(
                    id: 'organize',
                    options: [
                        "excludeByName" => [],
                        "indexByName" => [
                            "Used" => 3,
                            "Used, %" => 4,
                            "Value #FREE (lastNotNull)" => 2,
                            "Value #TOTAL (lastNotNull)" => 1,
                            "mountpoint" => 0,
                        ],
                        "renameByName" => [
                            "Value #TOTAL (lastNotNull)" => "Size",
                            "Value #FREE (lastNotNull)" => "Available",
                            "mountpoint" => "Mounted on",
                        ],
                    ],
                )
            )
            ->withTransformation(new SDKDashboard\DataTransformerConfig(
                id: 'sortBy',
                options: [
                    "fields" => [],
                    "sort" => [["field" => "Mounted on", "desc" => false]],
                ],
            ))
            // Overrides
            ->overrideByName('Mounted on', [
                new SDKDashboard\DynamicConfigValue(id: 'custom.width', value: 260),
            ])
            ->overrideByName('Size', [
                new SDKDashboard\DynamicConfigValue(id: 'custom.width', value: 80),
            ])
            ->overrideByName('Used', [
                new SDKDashboard\DynamicConfigValue(id: 'custom.width', value: 80),
            ])
            ->overrideByName('Available', [
                new SDKDashboard\DynamicConfigValue(id: 'custom.width', value: 80),
            ])
            ->overrideByName('Used, %', [
                new SDKDashboard\DynamicConfigValue(id: 'custom.cellOptions', value: [
                    'mode' => 'basic',
                    'type' => 'gauge',
                ]),
                new SDKDashboard\DynamicConfigValue(id: 'min', value: 0),
                new SDKDashboard\DynamicConfigValue(id: 'max', value: 1),
                new SDKDashboard\DynamicConfigValue(id: 'unit', value: Units::PERCENT_UNIT),
            ])
        ;
    }
}
