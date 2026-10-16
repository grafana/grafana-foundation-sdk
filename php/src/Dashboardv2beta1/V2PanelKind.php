<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * The notebook's own panel chain. It is a copy of the dashboard one down to the transformation,
 * which follows the dashboard v2 shape rather than the v2beta1 shape in this package. The chain has
 * to be forked rather than shared because PanelKind reaches TransformationKind through
 * QueryGroupKind, and those three are what Dashboard v2beta1 serves. Everything the chain does not
 * change (DataLink, VizConfigKind, PanelQueryKind, QueryOptionsSpec) stays shared.
 */
class V2PanelKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\V2PanelSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\V2PanelSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\V2PanelSpec $spec = null)
    {
        $this->kind = "Panel";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\V2PanelSpec();
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
    	/** @var array{id?: float, title?: string, description?: string, subtitle?: string, links?: array<mixed>, data?: mixed, vizConfig?: mixed, transparent?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\V2PanelSpec::fromArray($val);
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
