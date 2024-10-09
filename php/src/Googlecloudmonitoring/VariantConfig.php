<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: "cloud-monitoring",
            fromArray: [\Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQuery::class, 'fromArray'],
            convert: [\Grafana\Foundation\Googlecloudmonitoring\CloudMonitoringQueryConverter::class, 'convert'],
        );
    }
}