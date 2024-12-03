<?php

namespace Grafana\Foundation\Athena;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: "grafana-athena-datasource",
            fromArray: [\Grafana\Foundation\Athena\Dataquery::class, 'fromArray'],
            convert: [\Grafana\Foundation\Athena\DataqueryConverter::class, 'convert'],
        );
    }
}