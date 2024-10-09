<?php

namespace Grafana\Foundation\Text;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'text',
            optionsFromArray: [\Grafana\Foundation\Text\Options::class, 'fromArray'],
            fieldConfigFromArray: null,
            convert: [\Grafana\Foundation\Text\PanelConverter::class, 'convert'],

        );
    }
}