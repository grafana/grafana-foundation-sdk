<?php

namespace Grafana\Foundation\Nodegraph;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'nodeGraph',
            optionsFromArray: [\Grafana\Foundation\Nodegraph\Options::class, 'fromArray'],
            fieldConfigFromArray: null,
            convert: [\Grafana\Foundation\Nodegraph\PanelConverter::class, 'convert'],

        );
    }
}