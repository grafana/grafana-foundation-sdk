<?php

namespace Grafana\Foundation\Dashboardlist;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'dashlist',
            optionsFromArray: [\Grafana\Foundation\Dashboardlist\Options::class, 'fromArray'],
            fieldConfigFromArray: null,
            convert: [\Grafana\Foundation\Dashboardlist\PanelConverter::class, 'convert'],

        );
    }
}