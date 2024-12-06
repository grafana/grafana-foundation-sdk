<?php

namespace Grafana\Foundation\Grafanapyroscope;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: "grafanapyroscope",
            fromArray: [\Grafana\Foundation\Grafanapyroscope\Dataquery::class, 'fromArray'],
            convert: [\Grafana\Foundation\Grafanapyroscope\DataqueryConverter::class, 'convert'],
        );
    }
}