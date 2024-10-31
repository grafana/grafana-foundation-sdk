<?php

namespace Grafana\Foundation\Statetimeline;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'state-timeline',
            optionsFromArray: [\Grafana\Foundation\Statetimeline\Options::class, 'fromArray'],
            fieldConfigFromArray: [\Grafana\Foundation\Statetimeline\FieldConfig::class, 'fromArray'],
            convert: [\Grafana\Foundation\Statetimeline\PanelConverter::class, 'convert'],

        );
    }
}