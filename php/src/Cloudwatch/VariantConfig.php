<?php

namespace Grafana\Foundation\Cloudwatch;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: 'cloudwatch',
            fromArray: (function($input) {
    \assert(is_array($input), 'expected disjunction value to be an array');
    /** @var array<string, mixed> $input */
    switch ($input["queryMode"]) {
    case "Annotations":
        return AnnotationQuery::fromArray($input);
    case "Logs":
        return LogsQuery::fromArray($input);
    case "Metrics":
        return MetricsQuery::fromArray($input);
    default:
        throw new \ValueError('can not parse disjunction from array');
}
}),
            convert: (function(\Grafana\Foundation\Cog\Dataquery $input) {

    switch (true) {
    case $input instanceof AnnotationQuery:
        return AnnotationQueryConverter::convert($input);
    case $input instanceof LogsQuery:
        return LogsQueryConverter::convert($input);
    case $input instanceof MetricsQuery:
        return MetricsQueryConverter::convert($input);
    default:
        throw new \ValueError('can not convert unknown disjunction branch');
}
}),
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
