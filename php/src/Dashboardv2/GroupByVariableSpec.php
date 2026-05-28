<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * GroupBy variable specification
 */
class GroupByVariableSpec implements \JsonSerializable
{
    public string $name;

    public ?\Grafana\Foundation\Dashboardv2\VariableOption $defaultValue;

    public \Grafana\Foundation\Dashboardv2\VariableOption $current;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2\VariableOption>
     */
    public array $options;

    public bool $multi;

    public ?string $label;

    public \Grafana\Foundation\Dashboardv2\VariableHide $hide;

    public bool $skipUrlSync;

    public ?string $description;

    public ?\Grafana\Foundation\Dashboardv2\ControlSourceRef $origin;

    /**
     * @param string|null $name
     * @param \Grafana\Foundation\Dashboardv2\VariableOption|null $defaultValue
     * @param \Grafana\Foundation\Dashboardv2\VariableOption|null $current
     * @param array<\Grafana\Foundation\Dashboardv2\VariableOption>|null $options
     * @param bool|null $multi
     * @param string|null $label
     * @param \Grafana\Foundation\Dashboardv2\VariableHide|null $hide
     * @param bool|null $skipUrlSync
     * @param string|null $description
     * @param \Grafana\Foundation\Dashboardv2\ControlSourceRef|null $origin
     */
    public function __construct(?string $name = null, ?\Grafana\Foundation\Dashboardv2\VariableOption $defaultValue = null, ?\Grafana\Foundation\Dashboardv2\VariableOption $current = null, ?array $options = null, ?bool $multi = null, ?string $label = null, ?\Grafana\Foundation\Dashboardv2\VariableHide $hide = null, ?bool $skipUrlSync = null, ?string $description = null, ?\Grafana\Foundation\Dashboardv2\ControlSourceRef $origin = null)
    {
        $this->name = $name ?: "";
        $this->defaultValue = $defaultValue;
        $this->current = $current ?: new \Grafana\Foundation\Dashboardv2\VariableOption(text: "", value: "");
        $this->options = $options ?: [];
        $this->multi = $multi ?: false;
        $this->label = $label;
        $this->hide = $hide ?: \Grafana\Foundation\Dashboardv2\VariableHide::dontHide();
        $this->skipUrlSync = $skipUrlSync ?: false;
        $this->description = $description;
        $this->origin = $origin;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, defaultValue?: mixed, current?: mixed, options?: array<mixed>, multi?: bool, label?: string, hide?: string, skipUrlSync?: bool, description?: string, origin?: mixed} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            defaultValue: isset($data["defaultValue"]) ? (function($input) {
    	/** @var array{selected?: bool, text?: string|array<string>, value?: string|array<string>, properties?: array<string, string>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\VariableOption::fromArray($val);
    })($data["defaultValue"]) : null,
            current: isset($data["current"]) ? (function($input) {
    	/** @var array{selected?: bool, text?: string|array<string>, value?: string|array<string>, properties?: array<string, string>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\VariableOption::fromArray($val);
    })($data["current"]) : null,
            options: array_filter(array_map((function($input) {
    	/** @var array{selected?: bool, text?: string|array<string>, value?: string|array<string>, properties?: array<string, string>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\VariableOption::fromArray($val);
    }), $data["options"] ?? [])),
            multi: $data["multi"] ?? null,
            label: $data["label"] ?? null,
            hide: isset($data["hide"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2\VariableHide::fromValue($input); })($data["hide"]) : null,
            skipUrlSync: $data["skipUrlSync"] ?? null,
            description: $data["description"] ?? null,
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
        $data->current = $this->current;
        $data->options = $this->options;
        $data->multi = $this->multi;
        $data->hide = $this->hide;
        $data->skipUrlSync = $this->skipUrlSync;
        if (isset($this->defaultValue)) {
            $data->defaultValue = $this->defaultValue;
        }
        if (isset($this->label)) {
            $data->label = $this->label;
        }
        if (isset($this->description)) {
            $data->description = $this->description;
        }
        if (isset($this->origin)) {
            $data->origin = $this->origin;
        }
        return $data;
    }
}
