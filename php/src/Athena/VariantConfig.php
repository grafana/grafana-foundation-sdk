<?php

namespace Grafana\Foundation\Athena;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: 'grafana-athena-datasource',
            fromArray: [Dataquery::class, 'fromArray'],
            convert: [DataqueryConverter::class, 'convert'],
            convertv2: [self::class, 'convertV2'],
        );
    }

    public static function convertV2(\Grafana\Foundation\Dashboardv2\DataQueryKind|\Grafana\Foundation\Dashboardv2beta1\DataQueryKind $input): string
    {
        if ($input instanceof \Grafana\Foundation\Dashboardv2\DataQueryKind) {
            return QueryV2Converter::convert($input);
        }
        return QueryConverter::convert($input);
    }

}
