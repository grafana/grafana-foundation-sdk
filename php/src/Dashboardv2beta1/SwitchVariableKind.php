<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class SwitchVariableKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\SwitchVariableSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\SwitchVariableSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\SwitchVariableSpec $spec = null)
    {
        $this->kind = "SwitchVariable";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\SwitchVariableSpec();
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
    	/** @var array{name?: string, current?: string, enabledValue?: string, disabledValue?: string, label?: string, hide?: string, skipUrlSync?: bool, description?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\SwitchVariableSpec::fromArray($val);
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
