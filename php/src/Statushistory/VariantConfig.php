<?php

namespace Grafana\Foundation\Statushistory;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'status-history',
            optionsFromArray: [Options::class, 'fromArray'],
            fieldConfigFromArray: [FieldConfig::class, 'fromArray'],
            convert: [PanelConverter::class, 'convert'],
        );
    }
}
