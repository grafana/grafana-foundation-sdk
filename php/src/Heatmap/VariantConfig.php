<?php

namespace Grafana\Foundation\Heatmap;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'heatmap',
            optionsFromArray: [Options::class, 'fromArray'],
            fieldConfigFromArray: [FieldConfig::class, 'fromArray'],
            convert: [self::class, 'convert'],
        );
    }
    public static function convert(mixed $input): string
    {
        if ($input instanceof \Grafana\Foundation\Dashboard\Panel) {
            return PanelConverter::convert($input);
        }
    
        if ($input instanceof \Grafana\Foundation\Dashboardv2beta1\VizConfigKind) {
            return VisualizationConverter::convert($input);
        }
    
        return '/* could not convert panel */';
    }
}
