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
    case "resample":
        return TypeResample::fromArray($input);
    case "classic_conditions":
        return TypeClassicConditions::fromArray($input);
    case "threshold":
        return TypeThreshold::fromArray($input);
    case "sql":
        return TypeSql::fromArray($input);
    case "math":
        return TypeMath::fromArray($input);
    case "reduce":
        return TypeReduce::fromArray($input);
    default:
        throw new \ValueError('can not parse disjunction from array');
}
}),
        );
    }
}