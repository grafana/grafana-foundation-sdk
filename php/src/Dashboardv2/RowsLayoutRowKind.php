<?php

namespace Grafana\Foundation\Dashboardv2;

class RowsLayoutRowKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2\RowsLayoutRowSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2\RowsLayoutRowSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2\RowsLayoutRowSpec $spec = null)
    {
        $this->kind = "RowsLayoutRow";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2\RowsLayoutRowSpec();
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
    	/** @var array{title?: string, collapse?: bool, hideHeader?: bool, fillScreen?: bool, conditionalRendering?: mixed, repeat?: mixed, layout?: mixed|mixed|mixed|mixed, variables?: array<mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\RowsLayoutRowSpec::fromArray($val);
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
