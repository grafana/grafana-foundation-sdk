<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * Adhoc variable kind
 */
class AdhocVariableKind implements \JsonSerializable
{
    public string $kind;

    public string $group;

    /**
     * @var array<string, string>|null
     */
    public ?array $labels;

    public ?\Grafana\Foundation\Dashboardv2\Dashboardv2AdhocVariableKindDatasource $datasource;

    public \Grafana\Foundation\Dashboardv2\AdhocVariableSpec $spec;

    /**
     * @param string|null $group
     * @param array<string, string>|null $labels
     * @param \Grafana\Foundation\Dashboardv2\Dashboardv2AdhocVariableKindDatasource|null $datasource
     * @param \Grafana\Foundation\Dashboardv2\AdhocVariableSpec|null $spec
     */
    public function __construct(?string $group = null, ?array $labels = null, ?\Grafana\Foundation\Dashboardv2\Dashboardv2AdhocVariableKindDatasource $datasource = null, ?\Grafana\Foundation\Dashboardv2\AdhocVariableSpec $spec = null)
    {
        $this->kind = "AdhocVariable";
    
        $this->group = $group ?: "";
        $this->labels = $labels;
        $this->datasource = $datasource;
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2\AdhocVariableSpec();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{kind?: string, group?: string, labels?: array<string, string>, datasource?: mixed, spec?: mixed} $inputData */
        $data = $inputData;
        return new self(
            group: $data["group"] ?? null,
            labels: $data["labels"] ?? null,
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\Dashboardv2AdhocVariableKindDatasource::fromArray($val);
    })($data["datasource"]) : null,
            spec: isset($data["spec"]) ? (function($input) {
    	/** @var array{name?: string, baseFilters?: array<mixed>, filters?: array<mixed>, defaultKeys?: array<mixed>, label?: string, hide?: string, skipUrlSync?: bool, description?: string, allowCustomValue?: bool, enableGroupBy?: bool, origin?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\AdhocVariableSpec::fromArray($val);
    })($data["spec"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->kind = $this->kind;
        $data->group = $this->group;
        $data->spec = $this->spec;
        if (isset($this->labels)) {
            $data->labels = $this->labels;
        }
        if (isset($this->datasource)) {
            $data->datasource = $this->datasource;
        }
        return $data;
    }
}
