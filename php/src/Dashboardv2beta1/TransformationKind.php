<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class TransformationKind implements \JsonSerializable
{
    /**
     * The kind of a TransformationKind is the transformation ID
     */
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\DataTransformerConfig $spec;

    /**
     * @param string|null $kind
     * @param \Grafana\Foundation\Dashboardv2beta1\DataTransformerConfig|null $spec
     */
    public function __construct(?string $kind = null, ?\Grafana\Foundation\Dashboardv2beta1\DataTransformerConfig $spec = null)
    {
        $this->kind = $kind ?: "";
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\DataTransformerConfig();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{kind?: string, spec?: mixed} $inputData */
        $data = $inputData;
        return new self(
            kind: $data["kind"] ?? null,
            spec: isset($data["spec"]) ? (function($input) {
    	/** @var array{id?: string, disabled?: bool, filter?: mixed, topic?: string, options?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\DataTransformerConfig::fromArray($val);
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
