<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Datasource variable specification
 */
class DatasourceVariableSpec implements \JsonSerializable
{
    public string $name;

    public string $pluginId;

    public \Grafana\Foundation\Dashboardv2beta1\VariableRefresh $refresh;

    public string $regex;

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

    /**
     * @param string|null $name
     * @param string|null $pluginId
     * @param \Grafana\Foundation\Dashboardv2beta1\VariableRefresh|null $refresh
     * @param string|null $regex
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
     */
    public function __construct(?string $name = null, ?string $pluginId = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableRefresh $refresh = null, ?string $regex = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableOption $current = null, ?array $options = null, ?bool $multi = null, ?bool $includeAll = null, ?string $allValue = null, ?string $label = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableHide $hide = null, ?bool $skipUrlSync = null, ?string $description = null, ?bool $allowCustomValue = null)
    {
        $this->name = $name ?: "";
        $this->pluginId = $pluginId ?: "";
        $this->refresh = $refresh ?: \Grafana\Foundation\Dashboardv2beta1\VariableRefresh::never();
        $this->regex = $regex ?: "";
        $this->current = $current ?: new \Grafana\Foundation\Dashboardv2beta1\VariableOption(text: "", value: "");
        $this->options = $options ?: [];
        $this->multi = $multi ?: false;
        $this->includeAll = $includeAll ?: false;
        $this->allValue = $allValue;
        $this->label = $label;
        $this->hide = $hide ?: \Grafana\Foundation\Dashboardv2beta1\VariableHide::dontHide();
        $this->skipUrlSync = $skipUrlSync ?: false;
        $this->description = $description;
        $this->allowCustomValue = $allowCustomValue ?: true;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, pluginId?: string, refresh?: string, regex?: string, current?: mixed, options?: array<mixed>, multi?: bool, includeAll?: bool, allValue?: string, label?: string, hide?: string, skipUrlSync?: bool, description?: string, allowCustomValue?: bool} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            pluginId: $data["pluginId"] ?? null,
            refresh: isset($data["refresh"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\VariableRefresh::fromValue($input); })($data["refresh"]) : null,
            regex: $data["regex"] ?? null,
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
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->name = $this->name;
        $data->pluginId = $this->pluginId;
        $data->refresh = $this->refresh;
        $data->regex = $this->regex;
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
        return $data;
    }
}
