<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class QueryGroupKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\QueryGroupSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\QueryGroupSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\QueryGroupSpec $spec = null)
    {
        $this->kind = "QueryGroup";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\QueryGroupSpec();
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
    	/** @var array{queries?: array<mixed>, transformations?: array<mixed>, queryOptions?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\QueryGroupSpec::fromArray($val);
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
