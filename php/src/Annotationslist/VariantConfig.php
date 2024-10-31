<?php

namespace Grafana\Foundation\Annotationslist;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'annolist',
            optionsFromArray: [\Grafana\Foundation\Annotationslist\Options::class, 'fromArray'],
            fieldConfigFromArray: null,
            convert: [\Grafana\Foundation\Annotationslist\PanelConverter::class, 'convert'],

        );
    }
}