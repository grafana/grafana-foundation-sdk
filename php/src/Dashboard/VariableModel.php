<?php

namespace Grafana\Foundation\Dashboard;

/**
 * A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
 */
class VariableModel implements \JsonSerializable
{
    /**
     * Unique numeric identifier for the variable.
     */
    public string $id;

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
    public \Grafana\Foundation\Dashboard\VariableHide $hide;

    /**
     * Whether the variable value should be managed by URL query params or not
     */
    public bool $skipUrlSync;

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
     * Format to use while fetching all values from data source, eg: wildcard, glob, regex, pipe, etc.
     */
    public ?string $allFormat;

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
     * @param string|null $id
     * @param \Grafana\Foundation\Dashboard\VariableType|null $type
     * @param string|null $name
     * @param string|null $label
     * @param \Grafana\Foundation\Dashboard\VariableHide|null $hide
     * @param bool|null $skipUrlSync
     * @param string|null $description
     * @param string|array<string, mixed>|null $query
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     * @param string|null $allFormat
     * @param \Grafana\Foundation\Dashboard\VariableOption|null $current
     * @param bool|null $multi
     * @param array<\Grafana\Foundation\Dashboard\VariableOption>|null $options
     * @param \Grafana\Foundation\Dashboard\VariableRefresh|null $refresh
     * @param \Grafana\Foundation\Dashboard\VariableSort|null $sort
     * @param bool|null $includeAll
     * @param string|null $allValue
     * @param string|null $regex
     */
    public function __construct(?string $id = null, ?\Grafana\Foundation\Dashboard\VariableType $type = null, ?string $name = null, ?string $label = null, ?\Grafana\Foundation\Dashboard\VariableHide $hide = null, ?bool $skipUrlSync = null, ?string $description = null,  $query = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null, ?string $allFormat = null, ?\Grafana\Foundation\Dashboard\VariableOption $current = null, ?bool $multi = null, ?array $options = null, ?\Grafana\Foundation\Dashboard\VariableRefresh $refresh = null, ?\Grafana\Foundation\Dashboard\VariableSort $sort = null, ?bool $includeAll = null, ?string $allValue = null, ?string $regex = null)
    {
        $this->id = $id ?: "00000000-0000-0000-0000-000000000000";
        $this->type = $type ?: \Grafana\Foundation\Dashboard\VariableType::Query();
        $this->name = $name ?: "";
        $this->label = $label;
        $this->hide = $hide ?: \Grafana\Foundation\Dashboard\VariableHide::DontHide();
        $this->skipUrlSync = $skipUrlSync ?: false;
        $this->description = $description;
        $this->query = $query;
        $this->datasource = $datasource;
        $this->allFormat = $allFormat;
        $this->current = $current;
        $this->multi = $multi;
        $this->options = $options;
        $this->refresh = $refresh;
        $this->sort = $sort;
        $this->includeAll = $includeAll;
        $this->allValue = $allValue;
        $this->regex = $regex;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{id?: string, type?: string, name?: string, label?: string, hide?: int, skipUrlSync?: bool, description?: string, query?: string|array<string, mixed>, datasource?: mixed, allFormat?: string, current?: mixed, multi?: bool, options?: array<mixed>, refresh?: int, sort?: int, includeAll?: bool, allValue?: string, regex?: string} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
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
            allFormat: $data["allFormat"] ?? null,
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
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "id" => $this->id,
            "type" => $this->type,
            "name" => $this->name,
            "hide" => $this->hide,
            "skipUrlSync" => $this->skipUrlSync,
        ];
        if (isset($this->label)) {
            $data["label"] = $this->label;
        }
        if (isset($this->description)) {
            $data["description"] = $this->description;
        }
        if (isset($this->query)) {
            $data["query"] = $this->query;
        }
        if (isset($this->datasource)) {
            $data["datasource"] = $this->datasource;
        }
        if (isset($this->allFormat)) {
            $data["allFormat"] = $this->allFormat;
        }
        if (isset($this->current)) {
            $data["current"] = $this->current;
        }
        if (isset($this->multi)) {
            $data["multi"] = $this->multi;
        }
        if (isset($this->options)) {
            $data["options"] = $this->options;
        }
        if (isset($this->refresh)) {
            $data["refresh"] = $this->refresh;
        }
        if (isset($this->sort)) {
            $data["sort"] = $this->sort;
        }
        if (isset($this->includeAll)) {
            $data["includeAll"] = $this->includeAll;
        }
        if (isset($this->allValue)) {
            $data["allValue"] = $this->allValue;
        }
        if (isset($this->regex)) {
            $data["regex"] = $this->regex;
        }
        return $data;
    }
}
