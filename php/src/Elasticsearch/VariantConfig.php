<?php

namespace Grafana\Foundation\Elasticsearch;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: "elasticsearch",
            fromArray: [\Grafana\Foundation\Elasticsearch\Dataquery::class, 'fromArray'],
            convert: [\Grafana\Foundation\Elasticsearch\DataqueryConverter::class, 'convert'],
        );
    }
}