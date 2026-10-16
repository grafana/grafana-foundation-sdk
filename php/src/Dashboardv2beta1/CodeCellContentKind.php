<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class CodeCellContentKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\CodeCellContentSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\CodeCellContentSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\CodeCellContentSpec $spec = null)
    {
        $this->kind = "Code";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\CodeCellContentSpec();
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
    	/** @var array{language?: string, code?: string, highlight?: array<int>, annotation?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\CodeCellContentSpec::fromArray($val);
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
