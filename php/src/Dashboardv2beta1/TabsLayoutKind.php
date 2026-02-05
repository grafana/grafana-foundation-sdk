<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class TabsLayoutKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\TabsLayoutSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\TabsLayoutSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\TabsLayoutSpec $spec = null)
    {
        $this->kind = "TabsLayout";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\TabsLayoutSpec();
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
    	/** @var array{tabs?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\TabsLayoutSpec::fromArray($val);
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
