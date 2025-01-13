<?php

namespace Grafana\Foundation\Testdata;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: '',
            fromArray: [Dataquery::class, 'fromArray'],
            convert: [DataqueryConverter::class, 'convert'],
        );
    }
}
