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
    case "Annotations":
        return CloudWatchAnnotationQuery::fromArray($input);
    case "Logs":
        return CloudWatchLogsQuery::fromArray($input);
    case "Metrics":
        return CloudWatchMetricsQuery::fromArray($input);
    default:
        throw new \ValueError('can not parse disjunction from array');
}
}),
            convert: (function(\Grafana\Foundation\Cog\Dataquery $input) {

    switch (true) {
    case $input instanceof CloudWatchAnnotationQuery:
        return CloudWatchAnnotationQueryConverter::convert($input);
    case $input instanceof CloudWatchLogsQuery:
        return CloudWatchLogsQueryConverter::convert($input);
    case $input instanceof CloudWatchMetricsQuery:
        return CloudWatchMetricsQueryConverter::convert($input);
    default:
        throw new \ValueError('can not convert unknown disjunction branch');
}
}),
        );
    }
}