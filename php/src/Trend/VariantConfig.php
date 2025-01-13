<?php

namespace Grafana\Foundation\Trend;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'trend',
            optionsFromArray: [Options::class, 'fromArray'],
            fieldConfigFromArray: [FieldConfig::class, 'fromArray'],
            convert: [PanelConverter::class, 'convert'],
        );
    }
}
