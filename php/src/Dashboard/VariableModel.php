<?php

namespace Grafana\Foundation\Dashboard;

/**
 * A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
 */
class VariableModel implements \JsonSerializable
{
    /**
     * Type of variable
     */
    public \Grafana\Foundation\Dashboard\VariableType $type;

    /**
     * Name of variable
     */
    public string $name;

    /**
     * Optional display name
     */
    public ?string $label;

    /**
     * Visibility configuration for the variable
     */
    public ?\Grafana\Foundation\Dashboard\VariableHide $hide;

    /**
     * Whether the variable value should be managed by URL query params or not
     */
    public ?bool $skipUrlSync;

    /**
     * Description of variable. It can be defined but `null`.
     */
    public ?string $description;

    /**
     * Query used to fetch values for a variable
     * @var string|array<string, mixed>|null
     */
    public $query;

    /**
     * Data source used to fetch values for a variable. It can be defined but `null`.
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * Shows current selected variable text/value on the dashboard
     */
    public ?\Grafana\Foundation\Dashboard\VariableOption $current;

    /**
     * Whether multiple values can be selected or not from variable value list
     */
    public ?bool $multi;

    /**
     * Options that can be selected for a variable.
     * @var array<\Grafana\Foundation\Dashboard\VariableOption>|null
     */
    public ?array $options;

    /**
     * Options to config when to refresh a variable
     */
    public ?\Grafana\Foundation\Dashboard\VariableRefresh $refresh;

    /**
     * Options sort order
     */
    public ?\Grafana\Foundation\Dashboard\VariableSort $sort;

    /**
     * Whether all value option is available or not
     */
    public ?bool $includeAll;

    /**
     * Custom all value
     */
    public ?string $allValue;

    /**
     * Optional field, if you want to extract part of a series name or metric node segment.
     * Named capture groups can be used to separate the display text and value.
     */
    public ?string $regex;

    /**
     * Dynamically calculates interval by dividing time range by the count specified.
     */
    public ?bool $auto;

    /**
     * The minimum threshold below which the step count intervals will not divide the time.
     */
    public ?string $autoMin;

    /**
     * How many times the current time range should be divided to calculate the value, similar to the Max data points query option.
     * For example, if the current visible time range is 30 minutes, then the auto interval groups the data into 30 one-minute increments.
     */
    public ?int $autoCount;

    /**
     * @param \Grafana\Foundation\Dashboard\VariableType|null $type
     * @param string|null $name
     * @param string|null $label
     * @param \Grafana\Foundation\Dashboard\VariableHide|null $hide
     * @param bool|null $skipUrlSync
     * @param string|null $description
     * @param string|array<string, mixed>|null $query
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     * @param \Grafana\Foundation\Dashboard\VariableOption|null $current
     * @param bool|null $multi
     * @param array<\Grafana\Foundation\Dashboard\VariableOption>|null $options
     * @param \Grafana\Foundation\Dashboard\VariableRefresh|null $refresh
     * @param \Grafana\Foundation\Dashboard\VariableSort|null $sort
     * @param bool|null $includeAll
     * @param string|null $allValue
     * @param string|null $regex
     * @param bool|null $auto
     * @param string|null $autoMin
     * @param int|null $autoCount
     */
    public function __construct(?\Grafana\Foundation\Dashboard\VariableType $type = null, ?string $name = null, ?string $label = null, ?\Grafana\Foundation\Dashboard\VariableHide $hide = null, ?bool $skipUrlSync = null, ?string $description = null,  $query = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null, ?\Grafana\Foundation\Dashboard\VariableOption $current = null, ?bool $multi = null, ?array $options = null, ?\Grafana\Foundation\Dashboard\VariableRefresh $refresh = null, ?\Grafana\Foundation\Dashboard\VariableSort $sort = null, ?bool $includeAll = null, ?string $allValue = null, ?string $regex = null, ?bool $auto = null, ?string $autoMin = null, ?int $autoCount = null)
    {
        $this->type = $type ?: \Grafana\Foundation\Dashboard\VariableType::Query();
        $this->name = $name ?: "";
        $this->label = $label;
        $this->hide = $hide;
        $this->skipUrlSync = $skipUrlSync;
        $this->description = $description;
        $this->query = $query;
        $this->datasource = $datasource;
        $this->current = $current;
        $this->multi = $multi;
        $this->options = $options;
        $this->refresh = $refresh;
        $this->sort = $sort;
        $this->includeAll = $includeAll;
        $this->allValue = $allValue;
        $this->regex = $regex;
        $this->auto = $auto;
        $this->autoMin = $autoMin;
        $this->autoCount = $autoCount;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, name?: string, label?: string, hide?: int, skipUrlSync?: bool, description?: string, query?: string|array<string, mixed>, datasource?: mixed, current?: mixed, multi?: bool, options?: array<mixed>, refresh?: int, sort?: int, includeAll?: bool, allValue?: string, regex?: string, auto?: bool, auto_min?: string, auto_count?: int} $inputData */
        $data = $inputData;
        return new self(
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Dashboard\VariableType::fromValue($input); })($data["type"]) : null,
            name: $data["name"] ?? null,
            label: $data["label"] ?? null,
            hide: isset($data["hide"]) ? (function($input) { return \Grafana\Foundation\Dashboard\VariableHide::fromValue($input); })($data["hide"]) : null,
            skipUrlSync: $data["skipUrlSync"] ?? null,
            description: $data["description"] ?? null,
            query: isset($data["query"]) ? (function($input) {
        switch (true) {
        case is_string($input):
            return $input;
        default:
            return $input;
    }
    })($data["query"]) : null,
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DataSourceRef::fromArray($val);
    })($data["datasource"]) : null,
            current: isset($data["current"]) ? (function($input) {
    	/** @var array{selected?: bool, text?: string|array<string>, value?: string|array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\VariableOption::fromArray($val);
    })($data["current"]) : null,
            multi: $data["multi"] ?? null,
            options: array_filter(array_map((function($input) {
    	/** @var array{selected?: bool, text?: string|array<string>, value?: string|array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\VariableOption::fromArray($val);
    }), $data["options"] ?? [])),
            refresh: isset($data["refresh"]) ? (function($input) { return \Grafana\Foundation\Dashboard\VariableRefresh::fromValue($input); })($data["refresh"]) : null,
            sort: isset($data["sort"]) ? (function($input) { return \Grafana\Foundation\Dashboard\VariableSort::fromValue($input); })($data["sort"]) : null,
            includeAll: $data["includeAll"] ?? null,
            allValue: $data["allValue"] ?? null,
            regex: $data["regex"] ?? null,
            auto: $data["auto"] ?? null,
            autoMin: $data["auto_min"] ?? null,
            autoCount: $data["auto_count"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->type = $this->type;
        $data->name = $this->name;
        if (isset($this->label)) {
            $data->label = $this->label;
        }
        if (isset($this->hide)) {
            $data->hide = $this->hide;
        }
        if (isset($this->skipUrlSync)) {
            $data->skipUrlSync = $this->skipUrlSync;
        }
        if (isset($this->description)) {
            $data->description = $this->description;
        }
        if (isset($this->query)) {
            $data->query = $this->query;
        }
        if (isset($this->datasource)) {
            $data->datasource = $this->datasource;
        }
        if (isset($this->current)) {
            $data->current = $this->current;
        }
        if (isset($this->multi)) {
            $data->multi = $this->multi;
        }
        if (isset($this->options)) {
            $data->options = $this->options;
        }
        if (isset($this->refresh)) {
            $data->refresh = $this->refresh;
        }
        if (isset($this->sort)) {
            $data->sort = $this->sort;
        }
        if (isset($this->includeAll)) {
            $data->includeAll = $this->includeAll;
        }
        if (isset($this->allValue)) {
            $data->allValue = $this->allValue;
        }
        if (isset($this->regex)) {
            $data->regex = $this->regex;
        }
        if (isset($this->auto)) {
            $data->auto = $this->auto;
        }
        if (isset($this->autoMin)) {
            $data->auto_min = $this->autoMin;
        }
        if (isset($this->autoCount)) {
            $data->auto_count = $this->autoCount;
        }
        return $data;
    }
}
