<?php

namespace Grafana\Foundation\Loki;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: "loki",
            fromArray: [\Grafana\Foundation\Loki\Dataquery::class, 'fromArray'],
            convert: [\Grafana\Foundation\Loki\DataqueryConverter::class, 'convert'],
        );
    }
}