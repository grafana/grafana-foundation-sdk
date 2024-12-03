<?php

namespace Grafana\Foundation\Bigquery;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: "grafana-bigquery-datasource",
            fromArray: [\Grafana\Foundation\Bigquery\Dataquery::class, 'fromArray'],
            convert: [\Grafana\Foundation\Bigquery\DataqueryConverter::class, 'convert'],
        );
    }
}