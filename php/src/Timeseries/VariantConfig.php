<?php

namespace Grafana\Foundation\Timeseries;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'timeseries',
            optionsFromArray: [\Grafana\Foundation\Timeseries\Options::class, 'fromArray'],
            fieldConfigFromArray: [\Grafana\Foundation\Timeseries\FieldConfig::class, 'fromArray'],
            convert: [\Grafana\Foundation\Timeseries\PanelConverter::class, 'convert'],

        );
    }
}