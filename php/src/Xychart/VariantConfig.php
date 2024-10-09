<?php

namespace Grafana\Foundation\Xychart;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'xychart',
            optionsFromArray: [\Grafana\Foundation\Xychart\Options::class, 'fromArray'],
            fieldConfigFromArray: [\Grafana\Foundation\Xychart\FieldConfig::class, 'fromArray'],
            convert: [\Grafana\Foundation\Xychart\PanelConverter::class, 'convert'],

        );
    }
}