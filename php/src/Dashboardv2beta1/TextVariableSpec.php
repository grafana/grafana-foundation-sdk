<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Text variable specification
 */
class TextVariableSpec implements \JsonSerializable
{
    public string $name;

    public \Grafana\Foundation\Dashboardv2beta1\VariableOption $current;

    public string $query;

    public ?string $label;

    public \Grafana\Foundation\Dashboardv2beta1\VariableHide $hide;

    public bool $skipUrlSync;

    public ?string $description;

    /**
     * @param string|null $name
     * @param \Grafana\Foundation\Dashboardv2beta1\VariableOption|null $current
     * @param string|null $query
     * @param string|null $label
     * @param \Grafana\Foundation\Dashboardv2beta1\VariableHide|null $hide
     * @param bool|null $skipUrlSync
     * @param string|null $description
     */
    public function __construct(?string $name = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableOption $current = null, ?string $query = null, ?string $label = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableHide $hide = null, ?bool $skipUrlSync = null, ?string $description = null)
    {
        $this->name = $name ?: "";
        $this->current = $current ?: new \Grafana\Foundation\Dashboardv2beta1\VariableOption(text: "", value: "");
        $this->query = $query ?: "";
        $this->label = $label;
        $this->hide = $hide ?: \Grafana\Foundation\Dashboardv2beta1\VariableHide::dontHide();
        $this->skipUrlSync = $skipUrlSync ?: false;
        $this->description = $description;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, current?: mixed, query?: string, label?: string, hide?: string, skipUrlSync?: bool, description?: string} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            current: isset($data["current"]) ? (function($input) {
    	/** @var array{selected?: bool, text?: string|array<string>, value?: string|array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\VariableOption::fromArray($val);
    })($data["current"]) : null,
            query: $data["query"] ?? null,
            label: $data["label"] ?? null,
            hide: isset($data["hide"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\VariableHide::fromValue($input); })($data["hide"]) : null,
            skipUrlSync: $data["skipUrlSync"] ?? null,
            description: $data["description"] ?? null,
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
        $data->query = $this->query;
        $data->hide = $this->hide;
        $data->skipUrlSync = $this->skipUrlSync;
        if (isset($this->label)) {
            $data->label = $this->label;
        }
        if (isset($this->description)) {
            $data->description = $this->description;
        }
        return $data;
    }
}
