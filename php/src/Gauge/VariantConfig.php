<?php

namespace Grafana\Foundation\Gauge;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'gauge',
            optionsFromArray: [\Grafana\Foundation\Gauge\Options::class, 'fromArray'],
            fieldConfigFromArray: null,
            convert: [\Grafana\Foundation\Gauge\PanelConverter::class, 'convert'],

        );
    }
}