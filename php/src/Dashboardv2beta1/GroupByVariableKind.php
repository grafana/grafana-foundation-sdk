<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Group variable kind
 */
class GroupByVariableKind implements \JsonSerializable
{
    public string $kind;

    public string $group;

    public ?\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1GroupByVariableKindDatasource $datasource;

    public \Grafana\Foundation\Dashboardv2beta1\GroupByVariableSpec $spec;

    /**
     * @param string|null $group
     * @param \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1GroupByVariableKindDatasource|null $datasource
     * @param \Grafana\Foundation\Dashboardv2beta1\GroupByVariableSpec|null $spec
     */
    public function __construct(?string $group = null, ?\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1GroupByVariableKindDatasource $datasource = null, ?\Grafana\Foundation\Dashboardv2beta1\GroupByVariableSpec $spec = null)
    {
        $this->kind = "GroupByVariable";
    
        $this->group = $group ?: "";
        $this->datasource = $datasource;
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\GroupByVariableSpec();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{kind?: string, group?: string, datasource?: mixed, spec?: mixed} $inputData */
        $data = $inputData;
        return new self(
            group: $data["group"] ?? null,
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1GroupByVariableKindDatasource::fromArray($val);
    })($data["datasource"]) : null,
            spec: isset($data["spec"]) ? (function($input) {
    	/** @var array{name?: string, defaultValue?: mixed, current?: mixed, options?: array<mixed>, multi?: bool, label?: string, hide?: string, skipUrlSync?: bool, description?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\GroupByVariableSpec::fromArray($val);
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
        if (isset($this->datasource)) {
            $data->datasource = $this->datasource;
        }
        return $data;
    }
}
