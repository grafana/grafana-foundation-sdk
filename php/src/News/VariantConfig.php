<?php

namespace Grafana\Foundation\News;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'news',
            optionsFromArray: [Options::class, 'fromArray'],
            fieldConfigFromArray: null,
            convert: [PanelConverter::class, 'convert'],
        );
    }
}
