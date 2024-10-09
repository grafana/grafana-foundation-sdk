<?php

namespace Grafana\Foundation\Alertgroups;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'alertGroups',
            optionsFromArray: [\Grafana\Foundation\Alertgroups\Options::class, 'fromArray'],
            fieldConfigFromArray: null,
            convert: [\Grafana\Foundation\Alertgroups\PanelConverter::class, 'convert'],

        );
    }
}