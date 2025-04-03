<?php

namespace Grafana\Foundation\Histogram;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'histogram',
            optionsFromArray: [Options::class, 'fromArray'],
            fieldConfigFromArray: [FieldConfig::class, 'fromArray'],
            convert: [PanelConverter::class, 'convert'],
        );
    }
}
