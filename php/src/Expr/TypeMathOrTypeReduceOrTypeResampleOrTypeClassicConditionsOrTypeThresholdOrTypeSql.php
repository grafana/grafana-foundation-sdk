<?php

namespace Grafana\Foundation\Expr;

class TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql implements \JsonSerializable
{
    public ?\Grafana\Foundation\Expr\TypeMath $typeMath;

    public ?\Grafana\Foundation\Expr\TypeReduce $typeReduce;

    public ?\Grafana\Foundation\Expr\TypeResample $typeResample;

    public ?\Grafana\Foundation\Expr\TypeClassicConditions $typeClassicConditions;

    public ?\Grafana\Foundation\Expr\TypeThreshold $typeThreshold;

    public ?\Grafana\Foundation\Expr\TypeSql $typeSql;

    /**
     * @param \Grafana\Foundation\Expr\TypeMath|null $typeMath
     * @param \Grafana\Foundation\Expr\TypeReduce|null $typeReduce
     * @param \Grafana\Foundation\Expr\TypeResample|null $typeResample
     * @param \Grafana\Foundation\Expr\TypeClassicConditions|null $typeClassicConditions
     * @param \Grafana\Foundation\Expr\TypeThreshold|null $typeThreshold
     * @param \Grafana\Foundation\Expr\TypeSql|null $typeSql
     */
    public function __construct(?\Grafana\Foundation\Expr\TypeMath $typeMath = null, ?\Grafana\Foundation\Expr\TypeReduce $typeReduce = null, ?\Grafana\Foundation\Expr\TypeResample $typeResample = null, ?\Grafana\Foundation\Expr\TypeClassicConditions $typeClassicConditions = null, ?\Grafana\Foundation\Expr\TypeThreshold $typeThreshold = null, ?\Grafana\Foundation\Expr\TypeSql $typeSql = null)
    {
        $this->typeMath = $typeMath;
        $this->typeReduce = $typeReduce;
        $this->typeResample = $typeResample;
        $this->typeClassicConditions = $typeClassicConditions;
        $this->typeThreshold = $typeThreshold;
        $this->typeSql = $typeSql;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{TypeMath?: mixed, TypeReduce?: mixed, TypeResample?: mixed, TypeClassicConditions?: mixed, TypeThreshold?: mixed, TypeSql?: mixed} $inputData */
        $data = $inputData;
        return new self(
            typeMath: isset($data["TypeMath"]) ? (function($input) {
    	/** @var array{datasource?: mixed, expression?: string, hide?: bool, intervalMs?: float, maxDataPoints?: int, queryType?: string, refId?: string, resultAssertions?: mixed, timeRange?: mixed, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Expr\TypeMath::fromArray($val);
    })($data["TypeMath"]) : null,
            typeReduce: isset($data["TypeReduce"]) ? (function($input) {
    	/** @var array{datasource?: mixed, expression?: string, hide?: bool, intervalMs?: float, maxDataPoints?: int, queryType?: string, reducer?: string, refId?: string, resultAssertions?: mixed, settings?: mixed, timeRange?: mixed, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Expr\TypeReduce::fromArray($val);
    })($data["TypeReduce"]) : null,
            typeResample: isset($data["TypeResample"]) ? (function($input) {
    	/** @var array{datasource?: mixed, downsampler?: string, expression?: string, hide?: bool, intervalMs?: float, maxDataPoints?: int, queryType?: string, refId?: string, resultAssertions?: mixed, timeRange?: mixed, type?: string, upsampler?: string, window?: string} */
    $val = $input;
    	return \Grafana\Foundation\Expr\TypeResample::fromArray($val);
    })($data["TypeResample"]) : null,
            typeClassicConditions: isset($data["TypeClassicConditions"]) ? (function($input) {
    	/** @var array{conditions?: array<mixed>, datasource?: mixed, hide?: bool, intervalMs?: float, maxDataPoints?: int, queryType?: string, refId?: string, resultAssertions?: mixed, timeRange?: mixed, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Expr\TypeClassicConditions::fromArray($val);
    })($data["TypeClassicConditions"]) : null,
            typeThreshold: isset($data["TypeThreshold"]) ? (function($input) {
    	/** @var array{conditions?: array<mixed>, datasource?: mixed, expression?: string, hide?: bool, intervalMs?: float, maxDataPoints?: int, queryType?: string, refId?: string, resultAssertions?: mixed, timeRange?: mixed, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Expr\TypeThreshold::fromArray($val);
    })($data["TypeThreshold"]) : null,
            typeSql: isset($data["TypeSql"]) ? (function($input) {
    	/** @var array{datasource?: mixed, expression?: string, hide?: bool, intervalMs?: float, maxDataPoints?: int, queryType?: string, refId?: string, resultAssertions?: mixed, timeRange?: mixed, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Expr\TypeSql::fromArray($val);
    })($data["TypeSql"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->typeMath)) {
            $data["TypeMath"] = $this->typeMath;
        }
        if (isset($this->typeReduce)) {
            $data["TypeReduce"] = $this->typeReduce;
        }
        if (isset($this->typeResample)) {
            $data["TypeResample"] = $this->typeResample;
        }
        if (isset($this->typeClassicConditions)) {
            $data["TypeClassicConditions"] = $this->typeClassicConditions;
        }
        if (isset($this->typeThreshold)) {
            $data["TypeThreshold"] = $this->typeThreshold;
        }
        if (isset($this->typeSql)) {
            $data["TypeSql"] = $this->typeSql;
        }
        return $data;
    }
}
