<?php

namespace Grafana\Foundation\Bigquery;

class SQLExpression implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Bigquery\QueryEditorFunctionExpression>|null
     */
    public ?array $columns;

    public ?string $from;

    /**
     * whereJsonTree?: _
     */
    public ?string $whereString;

    /**
     * @var array<\Grafana\Foundation\Bigquery\QueryEditorGroupByExpression>|null
     */
    public ?array $groupBy;

    public ?\Grafana\Foundation\Bigquery\QueryEditorPropertyExpression $orderBy;

    public ?\Grafana\Foundation\Bigquery\OrderByDirection $orderByDirection;

    public ?int $limit;

    public ?int $offset;

    /**
     * @param array<\Grafana\Foundation\Bigquery\QueryEditorFunctionExpression>|null $columns
     * @param string|null $from
     * @param string|null $whereString
     * @param array<\Grafana\Foundation\Bigquery\QueryEditorGroupByExpression>|null $groupBy
     * @param \Grafana\Foundation\Bigquery\QueryEditorPropertyExpression|null $orderBy
     * @param \Grafana\Foundation\Bigquery\OrderByDirection|null $orderByDirection
     * @param int|null $limit
     * @param int|null $offset
     */
    public function __construct(?array $columns = null, ?string $from = null, ?string $whereString = null, ?array $groupBy = null, ?\Grafana\Foundation\Bigquery\QueryEditorPropertyExpression $orderBy = null, ?\Grafana\Foundation\Bigquery\OrderByDirection $orderByDirection = null, ?int $limit = null, ?int $offset = null)
    {
        $this->columns = $columns;
        $this->from = $from;
        $this->whereString = $whereString;
        $this->groupBy = $groupBy;
        $this->orderBy = $orderBy;
        $this->orderByDirection = $orderByDirection;
        $this->limit = $limit;
        $this->offset = $offset;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{columns?: array<mixed>, from?: string, whereString?: string, groupBy?: array<mixed>, orderBy?: mixed, orderByDirection?: string, limit?: int, offset?: int} $inputData */
        $data = $inputData;
        return new self(
            columns: array_filter(array_map((function($input) {
    	/** @var array{type?: string, name?: string, parameters?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Bigquery\QueryEditorFunctionExpression::fromArray($val);
    }), $data["columns"] ?? [])),
            from: $data["from"] ?? null,
            whereString: $data["whereString"] ?? null,
            groupBy: array_filter(array_map((function($input) {
    	/** @var array{type?: string, property?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Bigquery\QueryEditorGroupByExpression::fromArray($val);
    }), $data["groupBy"] ?? [])),
            orderBy: isset($data["orderBy"]) ? (function($input) {
    	/** @var array{type?: string, property?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Bigquery\QueryEditorPropertyExpression::fromArray($val);
    })($data["orderBy"]) : null,
            orderByDirection: isset($data["orderByDirection"]) ? (function($input) { return \Grafana\Foundation\Bigquery\OrderByDirection::fromValue($input); })($data["orderByDirection"]) : null,
            limit: $data["limit"] ?? null,
            offset: $data["offset"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->columns)) {
            $data["columns"] = $this->columns;
        }
        if (isset($this->from)) {
            $data["from"] = $this->from;
        }
        if (isset($this->whereString)) {
            $data["whereString"] = $this->whereString;
        }
        if (isset($this->groupBy)) {
            $data["groupBy"] = $this->groupBy;
        }
        if (isset($this->orderBy)) {
            $data["orderBy"] = $this->orderBy;
        }
        if (isset($this->orderByDirection)) {
            $data["orderByDirection"] = $this->orderByDirection;
        }
        if (isset($this->limit)) {
            $data["limit"] = $this->limit;
        }
        if (isset($this->offset)) {
            $data["offset"] = $this->offset;
        }
        return $data;
    }
}
