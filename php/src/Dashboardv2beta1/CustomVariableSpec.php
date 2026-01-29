<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Custom variable specification
 */
class CustomVariableSpec implements \JsonSerializable
{
    public string $name;

    public string $query;

    public \Grafana\Foundation\Dashboardv2beta1\VariableOption $current;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\VariableOption>
     */
    public array $options;

    public bool $multi;

    public bool $includeAll;

    public ?string $allValue;

    public ?string $label;

    public \Grafana\Foundation\Dashboardv2beta1\VariableHide $hide;

    public bool $skipUrlSync;

    public ?string $description;

    public bool $allowCustomValue;

    public ?\Grafana\Foundation\Dashboardv2beta1\CustomVariableSpecValuesFormat $valuesFormat;

    /**
     * @param string|null $name
     * @param string|null $query
     * @param \Grafana\Foundation\Dashboardv2beta1\VariableOption|null $current
     * @param array<\Grafana\Foundation\Dashboardv2beta1\VariableOption>|null $options
     * @param bool|null $multi
     * @param bool|null $includeAll
     * @param string|null $allValue
     * @param string|null $label
     * @param \Grafana\Foundation\Dashboardv2beta1\VariableHide|null $hide
     * @param bool|null $skipUrlSync
     * @param string|null $description
     * @param bool|null $allowCustomValue
     * @param \Grafana\Foundation\Dashboardv2beta1\CustomVariableSpecValuesFormat|null $valuesFormat
     */
    public function __construct(?string $name = null, ?string $query = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableOption $current = null, ?array $options = null, ?bool $multi = null, ?bool $includeAll = null, ?string $allValue = null, ?string $label = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableHide $hide = null, ?bool $skipUrlSync = null, ?string $description = null, ?bool $allowCustomValue = null, ?\Grafana\Foundation\Dashboardv2beta1\CustomVariableSpecValuesFormat $valuesFormat = null)
    {
        $this->name = $name ?: "";
        $this->query = $query ?: "";
        $this->current = $current ?: new \Grafana\Foundation\Dashboardv2beta1\VariableOption();
        $this->options = $options ?: [];
        $this->multi = $multi ?: false;
        $this->includeAll = $includeAll ?: false;
        $this->allValue = $allValue;
        $this->label = $label;
        $this->hide = $hide ?: \Grafana\Foundation\Dashboardv2beta1\VariableHide::dontHide();
        $this->skipUrlSync = $skipUrlSync ?: false;
        $this->description = $description;
        $this->allowCustomValue = $allowCustomValue ?: true;
        $this->valuesFormat = $valuesFormat;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, query?: string, current?: mixed, options?: array<mixed>, multi?: bool, includeAll?: bool, allValue?: string, label?: string, hide?: string, skipUrlSync?: bool, description?: string, allowCustomValue?: bool, valuesFormat?: string} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            query: $data["query"] ?? null,
            current: isset($data["current"]) ? (function($input) {
    	/** @var array{selected?: bool, text?: string|array<string>, value?: string|array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\VariableOption::fromArray($val);
    })($data["current"]) : null,
            options: array_filter(array_map((function($input) {
    	/** @var array{selected?: bool, text?: string|array<string>, value?: string|array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\VariableOption::fromArray($val);
    }), $data["options"] ?? [])),
            multi: $data["multi"] ?? null,
            includeAll: $data["includeAll"] ?? null,
            allValue: $data["allValue"] ?? null,
            label: $data["label"] ?? null,
            hide: isset($data["hide"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\VariableHide::fromValue($input); })($data["hide"]) : null,
            skipUrlSync: $data["skipUrlSync"] ?? null,
            description: $data["description"] ?? null,
            allowCustomValue: $data["allowCustomValue"] ?? null,
            valuesFormat: isset($data["valuesFormat"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\CustomVariableSpecValuesFormat::fromValue($input); })($data["valuesFormat"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->name = $this->name;
        $data->query = $this->query;
        $data->current = $this->current;
        $data->options = $this->options;
        $data->multi = $this->multi;
        $data->includeAll = $this->includeAll;
        $data->hide = $this->hide;
        $data->skipUrlSync = $this->skipUrlSync;
        $data->allowCustomValue = $this->allowCustomValue;
        if (isset($this->allValue)) {
            $data->allValue = $this->allValue;
        }
        if (isset($this->label)) {
            $data->label = $this->label;
        }
        if (isset($this->description)) {
            $data->description = $this->description;
        }
        if (isset($this->valuesFormat)) {
            $data->valuesFormat = $this->valuesFormat;
        }
        return $data;
    }
}
