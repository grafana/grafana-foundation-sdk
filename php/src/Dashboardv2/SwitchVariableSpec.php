<?php

namespace Grafana\Foundation\Dashboardv2;

class SwitchVariableSpec implements \JsonSerializable
{
    public string $name;

    public string $current;

    public string $enabledValue;

    public string $disabledValue;

    public ?string $label;

    public \Grafana\Foundation\Dashboardv2\VariableHide $hide;

    public bool $skipUrlSync;

    public ?string $description;

    public ?\Grafana\Foundation\Dashboardv2\ControlSourceRef $origin;

    /**
     * @param string|null $name
     * @param string|null $current
     * @param string|null $enabledValue
     * @param string|null $disabledValue
     * @param string|null $label
     * @param \Grafana\Foundation\Dashboardv2\VariableHide|null $hide
     * @param bool|null $skipUrlSync
     * @param string|null $description
     * @param \Grafana\Foundation\Dashboardv2\ControlSourceRef|null $origin
     */
    public function __construct(?string $name = null, ?string $current = null, ?string $enabledValue = null, ?string $disabledValue = null, ?string $label = null, ?\Grafana\Foundation\Dashboardv2\VariableHide $hide = null, ?bool $skipUrlSync = null, ?string $description = null, ?\Grafana\Foundation\Dashboardv2\ControlSourceRef $origin = null)
    {
        $this->name = $name ?: "";
        $this->current = $current ?: "false";
        $this->enabledValue = $enabledValue ?: "true";
        $this->disabledValue = $disabledValue ?: "false";
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
        /** @var array{name?: string, current?: string, enabledValue?: string, disabledValue?: string, label?: string, hide?: string, skipUrlSync?: bool, description?: string, origin?: mixed} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            current: $data["current"] ?? null,
            enabledValue: $data["enabledValue"] ?? null,
            disabledValue: $data["disabledValue"] ?? null,
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
        $data->enabledValue = $this->enabledValue;
        $data->disabledValue = $this->disabledValue;
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
