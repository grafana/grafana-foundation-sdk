<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: 'cloud-monitoring',
            fromArray: [CloudMonitoringQuery::class, 'fromArray'],
            convert: [CloudMonitoringQueryConverter::class, 'convert'],
            convertv2: [QueryConverter::class, 'convert'],
        );
    }

}
