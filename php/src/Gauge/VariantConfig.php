<?php

namespace Grafana\Foundation\Gauge;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'gauge',
            optionsFromArray: [Options::class, 'fromArray'],
            fieldConfigFromArray: null,
            convert: [self::class, 'convert'],
        );
    }
    private static function convert(mixed $input): string
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
