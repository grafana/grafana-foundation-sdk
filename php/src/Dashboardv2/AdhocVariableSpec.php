<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * Adhoc variable specification
 */
class AdhocVariableSpec implements \JsonSerializable
{
    public string $name;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2\AdHocFilterWithLabels>
     */
    public array $baseFilters;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2\AdHocFilterWithLabels>
     */
    public array $filters;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2\MetricFindValue>
     */
    public array $defaultKeys;

    public ?string $label;

    public \Grafana\Foundation\Dashboardv2\VariableHide $hide;

    public bool $skipUrlSync;

    public ?string $description;

    public bool $allowCustomValue;

    /**
     * Whether the group-by operator is enabled in the ad hoc filter combobox.
     */
    public ?bool $enableGroupBy;

    public ?\Grafana\Foundation\Dashboardv2\ControlSourceRef $origin;

    /**
     * @param string|null $name
     * @param array<\Grafana\Foundation\Dashboardv2\AdHocFilterWithLabels>|null $baseFilters
     * @param array<\Grafana\Foundation\Dashboardv2\AdHocFilterWithLabels>|null $filters
     * @param array<\Grafana\Foundation\Dashboardv2\MetricFindValue>|null $defaultKeys
     * @param string|null $label
     * @param \Grafana\Foundation\Dashboardv2\VariableHide|null $hide
     * @param bool|null $skipUrlSync
     * @param string|null $description
     * @param bool|null $allowCustomValue
     * @param bool|null $enableGroupBy
     * @param \Grafana\Foundation\Dashboardv2\ControlSourceRef|null $origin
     */
    public function __construct(?string $name = null, ?array $baseFilters = null, ?array $filters = null, ?array $defaultKeys = null, ?string $label = null, ?\Grafana\Foundation\Dashboardv2\VariableHide $hide = null, ?bool $skipUrlSync = null, ?string $description = null, ?bool $allowCustomValue = null, ?bool $enableGroupBy = null, ?\Grafana\Foundation\Dashboardv2\ControlSourceRef $origin = null)
    {
        $this->name = $name ?: "";
        $this->baseFilters = $baseFilters ?: [];
        $this->filters = $filters ?: [];
        $this->defaultKeys = $defaultKeys ?: [];
        $this->label = $label;
        $this->hide = $hide ?: \Grafana\Foundation\Dashboardv2\VariableHide::dontHide();
        $this->skipUrlSync = $skipUrlSync ?: false;
        $this->description = $description;
        $this->allowCustomValue = $allowCustomValue ?: true;
        $this->enableGroupBy = $enableGroupBy;
        $this->origin = $origin;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, baseFilters?: array<mixed>, filters?: array<mixed>, defaultKeys?: array<mixed>, label?: string, hide?: string, skipUrlSync?: bool, description?: string, allowCustomValue?: bool, enableGroupBy?: bool, origin?: mixed} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            baseFilters: array_filter(array_map((function($input) {
    	/** @var array{key?: string, operator?: string, value?: string, values?: array<string>, keyLabel?: string, valueLabels?: array<string>, forceEdit?: bool, origin?: "dashboard", condition?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\AdHocFilterWithLabels::fromArray($val);
    }), $data["baseFilters"] ?? [])),
            filters: array_filter(array_map((function($input) {
    	/** @var array{key?: string, operator?: string, value?: string, values?: array<string>, keyLabel?: string, valueLabels?: array<string>, forceEdit?: bool, origin?: "dashboard", condition?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\AdHocFilterWithLabels::fromArray($val);
    }), $data["filters"] ?? [])),
            defaultKeys: array_filter(array_map((function($input) {
    	/** @var array{text?: string, value?: string|float, group?: string, expandable?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\MetricFindValue::fromArray($val);
    }), $data["defaultKeys"] ?? [])),
            label: $data["label"] ?? null,
            hide: isset($data["hide"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2\VariableHide::fromValue($input); })($data["hide"]) : null,
            skipUrlSync: $data["skipUrlSync"] ?? null,
            description: $data["description"] ?? null,
            allowCustomValue: $data["allowCustomValue"] ?? null,
            enableGroupBy: $data["enableGroupBy"] ?? null,
            origin: isset($data["origin"]) ? (function($input) {
    	/** @var array{type?: string, group?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\ControlSourceRef::fromArray($val);
    })($data["origin"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->name = $this->name;
        $data->baseFilters = $this->baseFilters;
        $data->filters = $this->filters;
        $data->defaultKeys = $this->defaultKeys;
        $data->hide = $this->hide;
        $data->skipUrlSync = $this->skipUrlSync;
        $data->allowCustomValue = $this->allowCustomValue;
        if (isset($this->label)) {
            $data->label = $this->label;
        }
        if (isset($this->description)) {
            $data->description = $this->description;
        }
        if (isset($this->enableGroupBy)) {
            $data->enableGroupBy = $this->enableGroupBy;
        }
        if (isset($this->origin)) {
            $data->origin = $this->origin;
        }
        return $data;
    }
}
