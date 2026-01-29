<?php

namespace Grafana\Foundation\Tempo;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: 'tempo',
            fromArray: [TempoQuery::class, 'fromArray'],
            convert: [TempoQueryConverter::class, 'convert'],
        );
    }

}
