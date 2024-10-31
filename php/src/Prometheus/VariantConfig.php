<?php

namespace Grafana\Foundation\Prometheus;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: "prometheus",
            fromArray: [\Grafana\Foundation\Prometheus\Dataquery::class, 'fromArray'],
            convert: [\Grafana\Foundation\Prometheus\DataqueryConverter::class, 'convert'],
        );
    }
}