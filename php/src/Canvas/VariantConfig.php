<?php

namespace Grafana\Foundation\Canvas;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'canvas',
            optionsFromArray: [\Grafana\Foundation\Canvas\Options::class, 'fromArray'],
            fieldConfigFromArray: null,
            convert: [\Grafana\Foundation\Canvas\PanelConverter::class, 'convert'],

        );
    }
}