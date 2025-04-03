<?php

namespace Grafana\Foundation\Parca;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: 'parca',
            fromArray: [Dataquery::class, 'fromArray'],
            convert: [DataqueryConverter::class, 'convert'],
        );
    }
}
