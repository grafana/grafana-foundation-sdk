<?php

namespace Grafana\Foundation\Stat;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'stat',
            optionsFromArray: [Options::class, 'fromArray'],
            fieldConfigFromArray: null,
            convert: [PanelConverter::class, 'convert'],
        );
    }
}
