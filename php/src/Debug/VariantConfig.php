<?php

namespace Grafana\Foundation\Debug;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'debug',
            optionsFromArray: [\Grafana\Foundation\Debug\Options::class, 'fromArray'],
            fieldConfigFromArray: null,
            convert: [\Grafana\Foundation\Debug\PanelConverter::class, 'convert'],

        );
    }
}