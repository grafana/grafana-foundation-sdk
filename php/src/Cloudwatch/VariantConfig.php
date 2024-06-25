<?php

namespace Grafana\Foundation\Cloudwatch;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: "cloudwatch",
            fromArray: (function($input) {
    \assert(is_array($input), 'expected disjunction value to be an array');

    switch ($input["queryMode"]) {
    case "Metrics":
        return CloudWatchMetricsQuery::fromArray($input);
    case "Logs":
        return CloudWatchLogsQuery::fromArray($input);
    case "Annotations":
        return CloudWatchAnnotationQuery::fromArray($input);
    default:
        throw new \ValueError('can not parse disjunction from array');
}
}),
        );
    }
}