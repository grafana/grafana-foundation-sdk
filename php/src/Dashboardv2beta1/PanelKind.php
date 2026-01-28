<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class PanelKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\PanelSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\PanelSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\PanelSpec $spec = null)
    {
        $this->kind = "Panel";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\PanelSpec();
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
    	/** @var array{id?: float, title?: string, description?: string, links?: array<mixed>, data?: mixed, vizConfig?: mixed, transparent?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\PanelSpec::fromArray($val);
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
