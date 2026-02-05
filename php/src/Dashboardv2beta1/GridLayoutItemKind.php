<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class GridLayoutItemKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\GridLayoutItemSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\GridLayoutItemSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\GridLayoutItemSpec $spec = null)
    {
        $this->kind = "GridLayoutItem";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\GridLayoutItemSpec();
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
    	/** @var array{x?: int, y?: int, width?: int, height?: int, element?: mixed, repeat?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\GridLayoutItemSpec::fromArray($val);
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
