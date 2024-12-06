<?php

namespace Grafana\Foundation\Piechart;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'piechart',
            optionsFromArray: [\Grafana\Foundation\Piechart\Options::class, 'fromArray'],
            fieldConfigFromArray: [\Grafana\Foundation\Piechart\FieldConfig::class, 'fromArray'],
            convert: [\Grafana\Foundation\Piechart\PanelConverter::class, 'convert'],

        );
    }
}