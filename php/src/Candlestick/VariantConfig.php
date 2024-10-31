<?php

namespace Grafana\Foundation\Candlestick;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\PanelcfgConfig
    {
        return new \Grafana\Foundation\Cog\PanelcfgConfig(
            identifier: 'candlestick',
            optionsFromArray: [\Grafana\Foundation\Candlestick\Options::class, 'fromArray'],
            fieldConfigFromArray: [\Grafana\Foundation\Candlestick\FieldConfig::class, 'fromArray'],
            convert: [\Grafana\Foundation\Candlestick\PanelConverter::class, 'convert'],

        );
    }
}