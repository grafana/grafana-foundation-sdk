<?php

namespace Grafana\Foundation\Stat;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'stat',
            optionsFromArray: [\Grafana\Foundation\Stat\Options::class, 'fromArray'],
            fieldConfigFromArray: null,
            convert: [\Grafana\Foundation\Stat\PanelConverter::class, 'convert'],

        );
    }
}