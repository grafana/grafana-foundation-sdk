<?php

namespace Grafana\Foundation\Cloudwatch;

class SQLExpression implements \JsonSerializable
{
    /**
     * SELECT part of the SQL expression
     */
    public ?\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression $select;

    /**
     * FROM part of the SQL expression
     * @var \Grafana\Foundation\Cloudwatch\QueryEditorPropertyExpression|\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression|null
     */
    public $from;

    /**
     * WHERE part of the SQL expression
     */
    public ?\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression $where;

    /**
     * GROUP BY part of the SQL expression
     */
    public ?\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression $groupBy;

    /**
     * ORDER BY part of the SQL expression
     */
    public ?\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression $orderBy;

    /**
     * The sort order of the SQL expression, `ASC` or `DESC`
     */
    public ?string $orderByDirection;

    /**
     * LIMIT part of the SQL expression
     */
    public ?int $limit;

    /**
     * @param \Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression|null $select
     * @param \Grafana\Foundation\Cloudwatch\QueryEditorPropertyExpression|\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression|null $from
     * @param \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression|null $where
     * @param \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression|null $groupBy
     * @param \Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression|null $orderBy
     * @param string|null $orderByDirection
     * @param int|null $limit
     */
    public function __construct(?\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression $select = null,  $from = null, ?\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression $where = null, ?\Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression $groupBy = null, ?\Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression $orderBy = null, ?string $orderByDirection = null, ?int $limit = null)
    {
        $this->select = $select;
        $this->from = $from;
        $this->where = $where;
        $this->groupBy = $groupBy;
        $this->orderBy = $orderBy;
        $this->orderByDirection = $orderByDirection;
        $this->limit = $limit;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{select?: mixed, from?: mixed|mixed, where?: mixed, groupBy?: mixed, orderBy?: mixed, orderByDirection?: string, limit?: int} $inputData */
        $data = $inputData;
        return new self(
            select: isset($data["select"]) ? (function($input) {
    	/** @var array{type?: "function", name?: string, parameters?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression::fromArray($val);
    })($data["select"]) : null,
            from: isset($data["from"]) ? (function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
        /** @var array<string, mixed> $input */
        switch ($input["type"]) {
        case "function":
            return QueryEditorFunctionExpression::fromArray($input);
        case "property":
            return QueryEditorPropertyExpression::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    })($data["from"]) : null,
            where: isset($data["where"]) ? (function($input) {
    	/** @var array{type?: string, expressions?: array<mixed|mixed|mixed|mixed|mixed|mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression::fromArray($val);
    })($data["where"]) : null,
            groupBy: isset($data["groupBy"]) ? (function($input) {
    	/** @var array{type?: string, expressions?: array<mixed|mixed|mixed|mixed|mixed|mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Cloudwatch\QueryEditorArrayExpression::fromArray($val);
    })($data["groupBy"]) : null,
            orderBy: isset($data["orderBy"]) ? (function($input) {
    	/** @var array{type?: "function", name?: string, parameters?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Cloudwatch\QueryEditorFunctionExpression::fromArray($val);
    })($data["orderBy"]) : null,
            orderByDirection: $data["orderByDirection"] ?? null,
            limit: $data["limit"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->select)) {
            $data->select = $this->select;
        }
        if (isset($this->from)) {
            $data->from = $this->from;
        }
        if (isset($this->where)) {
            $data->where = $this->where;
        }
        if (isset($this->groupBy)) {
            $data->groupBy = $this->groupBy;
        }
        if (isset($this->orderBy)) {
            $data->orderBy = $this->orderBy;
        }
        if (isset($this->orderByDirection)) {
            $data->orderByDirection = $this->orderByDirection;
        }
        if (isset($this->limit)) {
            $data->limit = $this->limit;
        }
        return $data;
    }
}
