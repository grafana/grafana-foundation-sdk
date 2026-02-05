<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Query variable specification
 */
class QueryVariableSpec implements \JsonSerializable
{
    public string $name;

    public \Grafana\Foundation\Dashboardv2beta1\VariableOption $current;

    public ?string $label;

    public \Grafana\Foundation\Dashboardv2beta1\VariableHide $hide;

    public \Grafana\Foundation\Dashboardv2beta1\VariableRefresh $refresh;

    public bool $skipUrlSync;

    public ?string $description;

    public \Grafana\Foundation\Dashboardv2beta1\DataQueryKind $query;

    public string $regex;

    public \Grafana\Foundation\Dashboardv2beta1\VariableSort $sort;

    public ?string $definition;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\VariableOption>
     */
    public array $options;

    public bool $multi;

    public bool $includeAll;

    public ?string $allValue;

    public ?string $placeholder;

    public bool $allowCustomValue;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\VariableOption>|null
     */
    public ?array $staticOptions;

    public ?\Grafana\Foundation\Dashboardv2beta1\QueryVariableSpecStaticOptionsOrder $staticOptionsOrder;

    /**
     * @param string|null $name
     * @param \Grafana\Foundation\Dashboardv2beta1\VariableOption|null $current
     * @param string|null $label
     * @param \Grafana\Foundation\Dashboardv2beta1\VariableHide|null $hide
     * @param \Grafana\Foundation\Dashboardv2beta1\VariableRefresh|null $refresh
     * @param bool|null $skipUrlSync
     * @param string|null $description
     * @param \Grafana\Foundation\Dashboardv2beta1\DataQueryKind|null $query
     * @param string|null $regex
     * @param \Grafana\Foundation\Dashboardv2beta1\VariableSort|null $sort
     * @param string|null $definition
     * @param array<\Grafana\Foundation\Dashboardv2beta1\VariableOption>|null $options
     * @param bool|null $multi
     * @param bool|null $includeAll
     * @param string|null $allValue
     * @param string|null $placeholder
     * @param bool|null $allowCustomValue
     * @param array<\Grafana\Foundation\Dashboardv2beta1\VariableOption>|null $staticOptions
     * @param \Grafana\Foundation\Dashboardv2beta1\QueryVariableSpecStaticOptionsOrder|null $staticOptionsOrder
     */
    public function __construct(?string $name = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableOption $current = null, ?string $label = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableHide $hide = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableRefresh $refresh = null, ?bool $skipUrlSync = null, ?string $description = null, ?\Grafana\Foundation\Dashboardv2beta1\DataQueryKind $query = null, ?string $regex = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableSort $sort = null, ?string $definition = null, ?array $options = null, ?bool $multi = null, ?bool $includeAll = null, ?string $allValue = null, ?string $placeholder = null, ?bool $allowCustomValue = null, ?array $staticOptions = null, ?\Grafana\Foundation\Dashboardv2beta1\QueryVariableSpecStaticOptionsOrder $staticOptionsOrder = null)
    {
        $this->name = $name ?: "";
        $this->current = $current ?: new \Grafana\Foundation\Dashboardv2beta1\VariableOption(text: "", value: "");
        $this->label = $label;
        $this->hide = $hide ?: \Grafana\Foundation\Dashboardv2beta1\VariableHide::dontHide();
        $this->refresh = $refresh ?: \Grafana\Foundation\Dashboardv2beta1\VariableRefresh::never();
        $this->skipUrlSync = $skipUrlSync ?: false;
        $this->description = $description;
        $this->query = $query ?: new \Grafana\Foundation\Dashboardv2beta1\DataQueryKind();
        $this->regex = $regex ?: "";
        $this->sort = $sort ?: \Grafana\Foundation\Dashboardv2beta1\VariableSort::Disabled();
        $this->definition = $definition;
        $this->options = $options ?: [];
        $this->multi = $multi ?: false;
        $this->includeAll = $includeAll ?: false;
        $this->allValue = $allValue;
        $this->placeholder = $placeholder;
        $this->allowCustomValue = $allowCustomValue ?: true;
        $this->staticOptions = $staticOptions;
        $this->staticOptionsOrder = $staticOptionsOrder;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, current?: mixed, label?: string, hide?: string, refresh?: string, skipUrlSync?: bool, description?: string, query?: mixed, regex?: string, sort?: string, definition?: string, options?: array<mixed>, multi?: bool, includeAll?: bool, allValue?: string, placeholder?: string, allowCustomValue?: bool, staticOptions?: array<mixed>, staticOptionsOrder?: string} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            current: isset($data["current"]) ? (function($input) {
    	/** @var array{selected?: bool, text?: string|array<string>, value?: string|array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\VariableOption::fromArray($val);
    })($data["current"]) : null,
            label: $data["label"] ?? null,
            hide: isset($data["hide"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\VariableHide::fromValue($input); })($data["hide"]) : null,
            refresh: isset($data["refresh"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\VariableRefresh::fromValue($input); })($data["refresh"]) : null,
            skipUrlSync: $data["skipUrlSync"] ?? null,
            description: $data["description"] ?? null,
            query: isset($data["query"]) ? (function($input) {
    	/** @var array{kind?: string, group?: string, version?: string, datasource?: mixed, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\DataQueryKind::fromArray($val);
    })($data["query"]) : null,
            regex: $data["regex"] ?? null,
            sort: isset($data["sort"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\VariableSort::fromValue($input); })($data["sort"]) : null,
            definition: $data["definition"] ?? null,
            options: array_filter(array_map((function($input) {
    	/** @var array{selected?: bool, text?: string|array<string>, value?: string|array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\VariableOption::fromArray($val);
    }), $data["options"] ?? [])),
            multi: $data["multi"] ?? null,
            includeAll: $data["includeAll"] ?? null,
            allValue: $data["allValue"] ?? null,
            placeholder: $data["placeholder"] ?? null,
            allowCustomValue: $data["allowCustomValue"] ?? null,
            staticOptions: array_filter(array_map((function($input) {
    	/** @var array{selected?: bool, text?: string|array<string>, value?: string|array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\VariableOption::fromArray($val);
    }), $data["staticOptions"] ?? [])),
            staticOptionsOrder: isset($data["staticOptionsOrder"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\QueryVariableSpecStaticOptionsOrder::fromValue($input); })($data["staticOptionsOrder"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->name = $this->name;
        $data->current = $this->current;
        $data->hide = $this->hide;
        $data->refresh = $this->refresh;
        $data->skipUrlSync = $this->skipUrlSync;
        $data->query = $this->query;
        $data->regex = $this->regex;
        $data->sort = $this->sort;
        $data->options = $this->options;
        $data->multi = $this->multi;
        $data->includeAll = $this->includeAll;
        $data->allowCustomValue = $this->allowCustomValue;
        if (isset($this->label)) {
            $data->label = $this->label;
        }
        if (isset($this->description)) {
            $data->description = $this->description;
        }
        if (isset($this->definition)) {
            $data->definition = $this->definition;
        }
        if (isset($this->allValue)) {
            $data->allValue = $this->allValue;
        }
        if (isset($this->placeholder)) {
            $data->placeholder = $this->placeholder;
        }
        if (isset($this->staticOptions)) {
            $data->staticOptions = $this->staticOptions;
        }
        if (isset($this->staticOptionsOrder)) {
            $data->staticOptionsOrder = $this->staticOptionsOrder;
        }
        return $data;
    }
}
