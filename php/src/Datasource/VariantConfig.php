<?php

namespace Grafana\Foundation\Datasource;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: 'datasource',
            fromArray: [Dataquery::class, 'fromArray'],
            convert: [DataqueryConverter::class, 'convert'],
        );
    }
}
