<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * A cell holds non-panel narrative content (markdown text, code) in a notebook layout.
 * Panel cells are not represented here — they reuse V2PanelKind.
 */
class CellKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\CellSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\CellSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\CellSpec $spec = null)
    {
        $this->kind = "Cell";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\CellSpec();
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
    	/** @var array{content?: mixed|mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\CellSpec::fromArray($val);
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
