<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Constant variable specification
 */
class ConstantVariableSpec implements \JsonSerializable
{
    public string $name;

    public string $query;

    public \Grafana\Foundation\Dashboardv2beta1\VariableOption $current;

    public ?string $label;

    public \Grafana\Foundation\Dashboardv2beta1\VariableHide $hide;

    public bool $skipUrlSync;

    public ?string $description;

    public ?\Grafana\Foundation\Dashboardv2beta1\ControlSourceRef $origin;

    /**
     * @param string|null $name
     * @param string|null $query
     * @param \Grafana\Foundation\Dashboardv2beta1\VariableOption|null $current
     * @param string|null $label
     * @param \Grafana\Foundation\Dashboardv2beta1\VariableHide|null $hide
     * @param bool|null $skipUrlSync
     * @param string|null $description
     * @param \Grafana\Foundation\Dashboardv2beta1\ControlSourceRef|null $origin
     */
    public function __construct(?string $name = null, ?string $query = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableOption $current = null, ?string $label = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableHide $hide = null, ?bool $skipUrlSync = null, ?string $description = null, ?\Grafana\Foundation\Dashboardv2beta1\ControlSourceRef $origin = null)
    {
        $this->name = $name ?: "";
        $this->query = $query ?: "";
        $this->current = $current ?: new \Grafana\Foundation\Dashboardv2beta1\VariableOption(text: "", value: "");
        $this->label = $label;
        $this->hide = $hide ?: \Grafana\Foundation\Dashboardv2beta1\VariableHide::dontHide();
        $this->skipUrlSync = $skipUrlSync ?: false;
        $this->description = $description;
        $this->origin = $origin;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, query?: string, current?: mixed, label?: string, hide?: string, skipUrlSync?: bool, description?: string, origin?: mixed} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            query: $data["query"] ?? null,
            current: isset($data["current"]) ? (function($input) {
    	/** @var array{selected?: bool, text?: string|array<string>, value?: string|array<string>, properties?: array<string, string>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\VariableOption::fromArray($val);
    })($data["current"]) : null,
            label: $data["label"] ?? null,
            hide: isset($data["hide"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\VariableHide::fromValue($input); })($data["hide"]) : null,
            skipUrlSync: $data["skipUrlSync"] ?? null,
            description: $data["description"] ?? null,
            origin: isset($data["origin"]) ? (function($input) {
    	/** @var array{type?: string, group?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\ControlSourceRef::fromArray($val);
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
        $data->query = $this->query;
        $data->current = $this->current;
        $data->hide = $this->hide;
        $data->skipUrlSync = $this->skipUrlSync;
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
