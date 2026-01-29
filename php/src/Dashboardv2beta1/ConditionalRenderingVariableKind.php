<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class ConditionalRenderingVariableKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpec $spec = null)
    {
        $this->kind = "ConditionalRenderingVariable";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpec();
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
    	/** @var array{variable?: string, operator?: string, value?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableSpec::fromArray($val);
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
