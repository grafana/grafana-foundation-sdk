<?php

namespace Grafana\Foundation\Loki;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: 'loki',
            fromArray: [Dataquery::class, 'fromArray'],
            convert: [DataqueryConverter::class, 'convert'],
            convertv2: [QueryConverter::class, 'convert'],
        );
    }

}
