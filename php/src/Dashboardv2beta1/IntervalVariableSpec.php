<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Interval variable specification
 */
class IntervalVariableSpec implements \JsonSerializable
{
    public string $name;

    public string $query;

    public \Grafana\Foundation\Dashboardv2beta1\VariableOption $current;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\VariableOption>
     */
    public array $options;

    public bool $auto;

    public string $autoMin;

    public int $autoCount;

    public \Grafana\Foundation\Dashboardv2beta1\VariableRefresh $refresh;

    public ?string $label;

    public \Grafana\Foundation\Dashboardv2beta1\VariableHide $hide;

    public bool $skipUrlSync;

    public ?string $description;

    /**
     * @param string|null $name
     * @param string|null $query
     * @param \Grafana\Foundation\Dashboardv2beta1\VariableOption|null $current
     * @param array<\Grafana\Foundation\Dashboardv2beta1\VariableOption>|null $options
     * @param bool|null $auto
     * @param string|null $autoMin
     * @param int|null $autoCount
     * @param \Grafana\Foundation\Dashboardv2beta1\VariableRefresh|null $refresh
     * @param string|null $label
     * @param \Grafana\Foundation\Dashboardv2beta1\VariableHide|null $hide
     * @param bool|null $skipUrlSync
     * @param string|null $description
     */
    public function __construct(?string $name = null, ?string $query = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableOption $current = null, ?array $options = null, ?bool $auto = null, ?string $autoMin = null, ?int $autoCount = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableRefresh $refresh = null, ?string $label = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableHide $hide = null, ?bool $skipUrlSync = null, ?string $description = null)
    {
        $this->name = $name ?: "";
        $this->query = $query ?: "";
        $this->current = $current ?: new \Grafana\Foundation\Dashboardv2beta1\VariableOption(text: "", value: "");
        $this->options = $options ?: [];
        $this->auto = $auto ?: false;
        $this->autoMin = $autoMin ?: "";
        $this->autoCount = $autoCount ?: 0;
        $this->refresh = $refresh ?: \Grafana\Foundation\Dashboardv2beta1\VariableRefresh::never();
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
        /** @var array{name?: string, query?: string, current?: mixed, options?: array<mixed>, auto?: bool, auto_min?: string, auto_count?: int, refresh?: string, label?: string, hide?: string, skipUrlSync?: bool, description?: string} $inputData */
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
            auto: $data["auto"] ?? null,
            autoMin: $data["auto_min"] ?? null,
            autoCount: $data["auto_count"] ?? null,
            refresh: isset($data["refresh"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\VariableRefresh::fromValue($input); })($data["refresh"]) : null,
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
        $data->query = $this->query;
        $data->current = $this->current;
        $data->options = $this->options;
        $data->auto = $this->auto;
        $data->auto_min = $this->autoMin;
        $data->auto_count = $this->autoCount;
        $data->refresh = $this->refresh;
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
