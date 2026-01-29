<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Text variable kind
 */
class TextVariableKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\TextVariableSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\TextVariableSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\TextVariableSpec $spec = null)
    {
        $this->kind = "TextVariable";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\TextVariableSpec();
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
    	/** @var array{name?: string, current?: mixed, query?: string, label?: string, hide?: string, skipUrlSync?: bool, description?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\TextVariableSpec::fromArray($val);
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
