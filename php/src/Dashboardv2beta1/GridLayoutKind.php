<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class GridLayoutKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\GridLayoutSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\GridLayoutSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\GridLayoutSpec $spec = null)
    {
        $this->kind = "GridLayout";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\GridLayoutSpec();
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
    	/** @var array{items?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\GridLayoutSpec::fromArray($val);
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
