<?php

namespace Grafana\Foundation\Testdata;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: "testdata",
            fromArray: [\Grafana\Foundation\Testdata\Dataquery::class, 'fromArray'],
            convert: [\Grafana\Foundation\Testdata\DataqueryConverter::class, 'convert'],
        );
    }
}