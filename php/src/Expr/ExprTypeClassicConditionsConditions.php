<?php

namespace Grafana\Foundation\Expr;

class ExprTypeClassicConditionsConditions implements \JsonSerializable
{
    public \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsEvaluator $evaluator;

    public \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperator $operator;

    public \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsQuery $query;

    public \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsReducer $reducer;

    /**
     * @param \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsEvaluator|null $evaluator
     * @param \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperator|null $operator
     * @param \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsQuery|null $query
     * @param \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsReducer|null $reducer
     */
    public function __construct(?\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsEvaluator $evaluator = null, ?\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperator $operator = null, ?\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsQuery $query = null, ?\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsReducer $reducer = null)
    {
        $this->evaluator = $evaluator ?: new \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsEvaluator();
        $this->operator = $operator ?: new \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperator();
        $this->query = $query ?: new \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsQuery();
        $this->reducer = $reducer ?: new \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsReducer();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{evaluator?: mixed, operator?: mixed, query?: mixed, reducer?: mixed} $inputData */
        $data = $inputData;
        return new self(
            evaluator: isset($data["evaluator"]) ? (function($input) {
    	/** @var array{params?: array<float>, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsEvaluator::fromArray($val);
    })($data["evaluator"]) : null,
            operator: isset($data["operator"]) ? (function($input) {
    	/** @var array{type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsOperator::fromArray($val);
    })($data["operator"]) : null,
            query: isset($data["query"]) ? (function($input) {
    	/** @var array{params?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsQuery::fromArray($val);
    })($data["query"]) : null,
            reducer: isset($data["reducer"]) ? (function($input) {
    	/** @var array{type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Expr\ExprTypeClassicConditionsConditionsReducer::fromArray($val);
    })($data["reducer"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->evaluator = $this->evaluator;
        $data->operator = $this->operator;
        $data->query = $this->query;
        $data->reducer = $this->reducer;
        return $data;
    }
}
