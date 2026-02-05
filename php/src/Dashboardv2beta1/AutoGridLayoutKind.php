<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class AutoGridLayoutKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpec $spec;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpec|null $spec
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpec $spec = null)
    {
        $this->kind = "AutoGridLayout";
    
        $this->spec = $spec ?: new \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpec();
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
    	/** @var array{maxColumnCount?: float, columnWidthMode?: string, columnWidth?: float, rowHeightMode?: string, rowHeight?: float, fillScreen?: bool, items?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpec::fromArray($val);
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
