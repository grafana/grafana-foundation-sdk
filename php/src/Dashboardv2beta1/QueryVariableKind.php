<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Query variable kind
 */
class QueryVariableKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\QueryVariableSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\QueryVariableSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\QueryVariableSpec $spec = null)
    {
        $this->kind = "QueryVariable";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\QueryVariableSpec();
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
    	/** @var array{name?: string, current?: mixed, label?: string, hide?: string, refresh?: string, skipUrlSync?: bool, description?: string, query?: mixed, regex?: string, sort?: string, definition?: string, options?: array<mixed>, multi?: bool, includeAll?: bool, allValue?: string, placeholder?: string, allowCustomValue?: bool, staticOptions?: array<mixed>, staticOptionsOrder?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\QueryVariableSpec::fromArray($val);
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
