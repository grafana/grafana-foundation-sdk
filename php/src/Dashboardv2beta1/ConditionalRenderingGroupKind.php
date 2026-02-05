<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class ConditionalRenderingGroupKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpec $spec = null)
    {
        $this->kind = "ConditionalRenderingGroup";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpec();
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
    	/** @var array{visibility?: string, condition?: string, items?: array<mixed|mixed|mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpec::fromArray($val);
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
