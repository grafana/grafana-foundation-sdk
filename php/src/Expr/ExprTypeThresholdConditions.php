<?php

namespace Grafana\Foundation\Expr;

class ExprTypeThresholdConditions implements \JsonSerializable
{
    public \Grafana\Foundation\Expr\ExprTypeThresholdConditionsEvaluator $evaluator;

    /**
     * @var mixed|null
     */
    public $loadedDimensions;

    public ?\Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluator $unloadEvaluator;

    /**
     * @param \Grafana\Foundation\Expr\ExprTypeThresholdConditionsEvaluator|null $evaluator
     * @param mixed|null $loadedDimensions
     * @param \Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluator|null $unloadEvaluator
     */
    public function __construct(?\Grafana\Foundation\Expr\ExprTypeThresholdConditionsEvaluator $evaluator = null,  $loadedDimensions = null, ?\Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluator $unloadEvaluator = null)
    {
        $this->evaluator = $evaluator ?: new \Grafana\Foundation\Expr\ExprTypeThresholdConditionsEvaluator();
        $this->loadedDimensions = $loadedDimensions;
        $this->unloadEvaluator = $unloadEvaluator;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{evaluator?: mixed, loadedDimensions?: mixed, unloadEvaluator?: mixed} $inputData */
        $data = $inputData;
        return new self(
            evaluator: isset($data["evaluator"]) ? (function($input) {
    	/** @var array{params?: array<float>, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Expr\ExprTypeThresholdConditionsEvaluator::fromArray($val);
    })($data["evaluator"]) : null,
            loadedDimensions: $data["loadedDimensions"] ?? null,
            unloadEvaluator: isset($data["unloadEvaluator"]) ? (function($input) {
    	/** @var array{params?: array<float>, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Expr\ExprTypeThresholdConditionsUnloadEvaluator::fromArray($val);
    })($data["unloadEvaluator"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->evaluator = $this->evaluator;
        if (isset($this->loadedDimensions)) {
            $data->loadedDimensions = $this->loadedDimensions;
        }
        if (isset($this->unloadEvaluator)) {
            $data->unloadEvaluator = $this->unloadEvaluator;
        }
        return $data;
    }
}
