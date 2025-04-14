<?php

namespace Grafana\Foundation\Azuremonitor;

class BuilderQueryExpression implements \JsonSerializable
{
    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyExpression $from;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorColumnsExpression $columns;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray $where;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorReduceExpressionArray $reduce;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorGroupByExpressionArray $groupBy;

    public ?int $limit;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorOrderByExpressionArray $orderBy;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray $fuzzySearch;

    public ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray $timeFilter;

    /**
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyExpression|null $from
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorColumnsExpression|null $columns
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray|null $where
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorReduceExpressionArray|null $reduce
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorGroupByExpressionArray|null $groupBy
     * @param int|null $limit
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorOrderByExpressionArray|null $orderBy
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray|null $fuzzySearch
     * @param \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray|null $timeFilter
     */
    public function __construct(?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyExpression $from = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorColumnsExpression $columns = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray $where = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorReduceExpressionArray $reduce = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorGroupByExpressionArray $groupBy = null, ?int $limit = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorOrderByExpressionArray $orderBy = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray $fuzzySearch = null, ?\Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray $timeFilter = null)
    {
        $this->from = $from;
        $this->columns = $columns;
        $this->where = $where;
        $this->reduce = $reduce;
        $this->groupBy = $groupBy;
        $this->limit = $limit;
        $this->orderBy = $orderBy;
        $this->fuzzySearch = $fuzzySearch;
        $this->timeFilter = $timeFilter;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{from?: mixed, columns?: mixed, where?: mixed, reduce?: mixed, groupBy?: mixed, limit?: int, orderBy?: mixed, fuzzySearch?: mixed, timeFilter?: mixed} $inputData */
        $data = $inputData;
        return new self(
            from: isset($data["from"]) ? (function($input) {
    	/** @var array{property?: mixed, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorPropertyExpression::fromArray($val);
    })($data["from"]) : null,
            columns: isset($data["columns"]) ? (function($input) {
    	/** @var array{columns?: array<string>, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorColumnsExpression::fromArray($val);
    })($data["columns"]) : null,
            where: isset($data["where"]) ? (function($input) {
    	/** @var array{expressions?: array<mixed>, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray::fromArray($val);
    })($data["where"]) : null,
            reduce: isset($data["reduce"]) ? (function($input) {
    	/** @var array{expressions?: array<mixed>, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorReduceExpressionArray::fromArray($val);
    })($data["reduce"]) : null,
            groupBy: isset($data["groupBy"]) ? (function($input) {
    	/** @var array{expressions?: array<mixed>, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorGroupByExpressionArray::fromArray($val);
    })($data["groupBy"]) : null,
            limit: $data["limit"] ?? null,
            orderBy: isset($data["orderBy"]) ? (function($input) {
    	/** @var array{expressions?: array<mixed>, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorOrderByExpressionArray::fromArray($val);
    })($data["orderBy"]) : null,
            fuzzySearch: isset($data["fuzzySearch"]) ? (function($input) {
    	/** @var array{expressions?: array<mixed>, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray::fromArray($val);
    })($data["fuzzySearch"]) : null,
            timeFilter: isset($data["timeFilter"]) ? (function($input) {
    	/** @var array{expressions?: array<mixed>, type?: string} */
    $val = $input;
    	return \Grafana\Foundation\Azuremonitor\BuilderQueryEditorWhereExpressionArray::fromArray($val);
    })($data["timeFilter"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->from)) {
            $data["from"] = $this->from;
        }
        if (isset($this->columns)) {
            $data["columns"] = $this->columns;
        }
        if (isset($this->where)) {
            $data["where"] = $this->where;
        }
        if (isset($this->reduce)) {
            $data["reduce"] = $this->reduce;
        }
        if (isset($this->groupBy)) {
            $data["groupBy"] = $this->groupBy;
        }
        if (isset($this->limit)) {
            $data["limit"] = $this->limit;
        }
        if (isset($this->orderBy)) {
            $data["orderBy"] = $this->orderBy;
        }
        if (isset($this->fuzzySearch)) {
            $data["fuzzySearch"] = $this->fuzzySearch;
        }
        if (isset($this->timeFilter)) {
            $data["timeFilter"] = $this->timeFilter;
        }
        return $data;
    }
}
