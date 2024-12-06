<?php

namespace Grafana\Foundation\Datagrid;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'datagrid',
            optionsFromArray: [\Grafana\Foundation\Datagrid\Options::class, 'fromArray'],
            fieldConfigFromArray: null,
            convert: [\Grafana\Foundation\Datagrid\PanelConverter::class, 'convert'],

        );
    }
}