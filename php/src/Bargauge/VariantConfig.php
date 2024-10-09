<?php

namespace Grafana\Foundation\Bargauge;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'bargauge',
            optionsFromArray: [\Grafana\Foundation\Bargauge\Options::class, 'fromArray'],
            fieldConfigFromArray: null,
            convert: [\Grafana\Foundation\Bargauge\PanelConverter::class, 'convert'],

        );
    }
}