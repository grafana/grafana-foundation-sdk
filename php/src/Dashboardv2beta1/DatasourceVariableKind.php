<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Datasource variable kind
 */
class DatasourceVariableKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\DatasourceVariableSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\DatasourceVariableSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\DatasourceVariableSpec $spec = null)
    {
        $this->kind = "DatasourceVariable";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\DatasourceVariableSpec();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{kind?: string, spec?: mixed} $inputData */
        $data = $inputData;
        return new self(
            spec: isset($data["spec"]) ? (function($input) {
    	/** @var array{name?: string, pluginId?: string, refresh?: string, regex?: string, current?: mixed, options?: array<mixed>, multi?: bool, includeAll?: bool, allValue?: string, label?: string, hide?: string, skipUrlSync?: bool, description?: string, allowCustomValue?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\DatasourceVariableSpec::fromArray($val);
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
        $data->spec = $this->spec;
        return $data;
    }
}
