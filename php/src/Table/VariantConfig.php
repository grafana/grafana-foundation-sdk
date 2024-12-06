<?php

namespace Grafana\Foundation\Table;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'table',
            optionsFromArray: [\Grafana\Foundation\Table\Options::class, 'fromArray'],
            fieldConfigFromArray: [\Grafana\Foundation\Table\FieldConfig::class, 'fromArray'],
            convert: [\Grafana\Foundation\Table\PanelConverter::class, 'convert'],

        );
    }
}