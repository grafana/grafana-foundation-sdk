<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Interval variable kind
 */
class IntervalVariableKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\IntervalVariableSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\IntervalVariableSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\IntervalVariableSpec $spec = null)
    {
        $this->kind = "IntervalVariable";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\IntervalVariableSpec();
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
    	/** @var array{name?: string, query?: string, current?: mixed, options?: array<mixed>, auto?: bool, auto_min?: string, auto_count?: int, refresh?: string, label?: string, hide?: string, skipUrlSync?: bool, description?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\IntervalVariableSpec::fromArray($val);
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
