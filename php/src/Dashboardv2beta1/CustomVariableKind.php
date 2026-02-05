<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Custom variable kind
 */
class CustomVariableKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\CustomVariableSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\CustomVariableSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\CustomVariableSpec $spec = null)
    {
        $this->kind = "CustomVariable";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\CustomVariableSpec();
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
    	/** @var array{name?: string, query?: string, current?: mixed, options?: array<mixed>, multi?: bool, includeAll?: bool, allValue?: string, label?: string, hide?: string, skipUrlSync?: bool, description?: string, allowCustomValue?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\CustomVariableSpec::fromArray($val);
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
