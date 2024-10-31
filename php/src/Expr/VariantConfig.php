<?php

namespace Grafana\Foundation\Expr;

final class VariantConfig
{
    public static function get(): \Grafana\Foundation\Cog\DataqueryConfig
    {
        return new \Grafana\Foundation\Cog\DataqueryConfig(
            identifier: "__expr__",
            fromArray: (function($input) {
    \assert(is_array($input), 'expected disjunction value to be an array');

    switch ($input["type"]) {
    case "classic_conditions":
        return TypeClassicConditions::fromArray($input);
    case "math":
        return TypeMath::fromArray($input);
    case "reduce":
        return TypeReduce::fromArray($input);
    case "resample":
        return TypeResample::fromArray($input);
    case "sql":
        return TypeSql::fromArray($input);
    case "threshold":
        return TypeThreshold::fromArray($input);
    default:
        throw new \ValueError('can not parse disjunction from array');
}
}),
            convert: (function(\Grafana\Foundation\Cog\Dataquery $input) {

    switch (true) {
    case $input instanceof TypeClassicConditions:
        return TypeClassicConditionsConverter::convert($input);
    case $input instanceof TypeMath:
        return TypeMathConverter::convert($input);
    case $input instanceof TypeReduce:
        return TypeReduceConverter::convert($input);
    case $input instanceof TypeResample:
        return TypeResampleConverter::convert($input);
    case $input instanceof TypeSql:
        return TypeSqlConverter::convert($input);
    case $input instanceof TypeThreshold:
        return TypeThresholdConverter::convert($input);
    default:
        throw new \ValueError('can not convert unknown disjunction branch');
}
}),
        );
    }
}