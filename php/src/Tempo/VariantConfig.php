<?php

namespace Grafana\Foundation\Tempo;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: "tempo",
            fromArray: [\Grafana\Foundation\Tempo\TempoQuery::class, 'fromArray'],
            convert: [\Grafana\Foundation\Tempo\TempoQueryConverter::class, 'convert'],
        );
    }
}