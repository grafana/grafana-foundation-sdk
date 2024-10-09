<?php

namespace Grafana\Foundation\Barchart;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'barchart',
            optionsFromArray: [\Grafana\Foundation\Barchart\Options::class, 'fromArray'],
            fieldConfigFromArray: [\Grafana\Foundation\Barchart\FieldConfig::class, 'fromArray'],
            convert: [\Grafana\Foundation\Barchart\PanelConverter::class, 'convert'],

        );
    }
}