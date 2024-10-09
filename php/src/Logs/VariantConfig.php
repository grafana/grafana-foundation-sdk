<?php

namespace Grafana\Foundation\Logs;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'logs',
            optionsFromArray: [\Grafana\Foundation\Logs\Options::class, 'fromArray'],
            fieldConfigFromArray: null,
            convert: [\Grafana\Foundation\Logs\PanelConverter::class, 'convert'],

        );
    }
}