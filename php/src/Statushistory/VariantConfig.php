<?php

namespace Grafana\Foundation\Statushistory;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'status-history',
            optionsFromArray: [\Grafana\Foundation\Statushistory\Options::class, 'fromArray'],
            fieldConfigFromArray: [\Grafana\Foundation\Statushistory\FieldConfig::class, 'fromArray'],
            convert: [\Grafana\Foundation\Statushistory\PanelConverter::class, 'convert'],

        );
    }
}