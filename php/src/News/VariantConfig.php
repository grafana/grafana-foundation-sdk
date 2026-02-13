<?php

namespace Grafana\Foundation\News;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'news',
            optionsFromArray: [Options::class, 'fromArray'],
            fieldConfigFromArray: null,
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
