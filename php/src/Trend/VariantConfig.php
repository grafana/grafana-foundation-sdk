<?php

namespace Grafana\Foundation\Trend;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'trend',
            optionsFromArray: [\Grafana\Foundation\Trend\Options::class, 'fromArray'],
            fieldConfigFromArray: [\Grafana\Foundation\Trend\FieldConfig::class, 'fromArray'],
            convert: [\Grafana\Foundation\Trend\PanelConverter::class, 'convert'],

        );
    }
}