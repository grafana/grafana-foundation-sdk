<?php

namespace Grafana\Foundation\Geomap;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'geomap',
            optionsFromArray: [\Grafana\Foundation\Geomap\Options::class, 'fromArray'],
            fieldConfigFromArray: null,
            convert: [\Grafana\Foundation\Geomap\PanelConverter::class, 'convert'],

        );
    }
}