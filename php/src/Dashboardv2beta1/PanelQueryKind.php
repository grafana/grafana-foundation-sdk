<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class PanelQueryKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\PanelQuerySpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\PanelQuerySpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\PanelQuerySpec $spec = null)
    {
        $this->kind = "PanelQuery";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\PanelQuerySpec();
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
    	/** @var array{query?: mixed, refId?: string, hidden?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\PanelQuerySpec::fromArray($val);
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
