<?php

namespace Grafana\Foundation\Histogram;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'histogram',
            optionsFromArray: [\Grafana\Foundation\Histogram\Options::class, 'fromArray'],
            fieldConfigFromArray: [\Grafana\Foundation\Histogram\FieldConfig::class, 'fromArray'],
            convert: [\Grafana\Foundation\Histogram\PanelConverter::class, 'convert'],

        );
    }
}