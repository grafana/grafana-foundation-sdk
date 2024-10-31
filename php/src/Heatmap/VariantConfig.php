<?php

namespace Grafana\Foundation\Heatmap;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'heatmap',
            optionsFromArray: [\Grafana\Foundation\Heatmap\Options::class, 'fromArray'],
            fieldConfigFromArray: [\Grafana\Foundation\Heatmap\FieldConfig::class, 'fromArray'],
            convert: [\Grafana\Foundation\Heatmap\PanelConverter::class, 'convert'],

        );
    }
}