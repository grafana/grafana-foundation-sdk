<?php

namespace Grafana\Foundation\Azuremonitor;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: 'grafana-azure-monitor-datasource',
            fromArray: [AzureMonitorQuery::class, 'fromArray'],
            convert: [AzureMonitorQueryConverter::class, 'convert'],
        );
    }

}
