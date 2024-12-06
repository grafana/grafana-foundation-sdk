<?php

namespace Grafana\Foundation\Histogram;

class Options implements \JsonSerializable
{
    /**
     * Bucket count (approx)
     */
    public ?int $bucketCount;

    /**
     * Size of each bucket
     */
    public ?int $bucketSize;

    /**
     * Offset buckets by this amount
     */
    public ?float $bucketOffset;

    public \Grafana\Foundation\Common\VizLegendOptions $legend;

    public \Grafana\Foundation\Common\VizTooltipOptions $tooltip;

    /**
     * Combines multiple series into a single histogram
     */
    public ?bool $combine;

    /**
     * @param int|null $bucketCount
     * @param int|null $bucketSize
     * @param float|null $bucketOffset
     * @param \Grafana\Foundation\Common\VizLegendOptions|null $legend
     * @param \Grafana\Foundation\Common\VizTooltipOptions|null $tooltip
     * @param bool|null $combine
     */
    public function __construct(?int $bucketCount = null, ?int $bucketSize = null, ?float $bucketOffset = null, ?\Grafana\Foundation\Common\VizLegendOptions $legend = null, ?\Grafana\Foundation\Common\VizTooltipOptions $tooltip = null, ?bool $combine = null)
    {
        $this->bucketCount = $bucketCount;
        $this->bucketSize = $bucketSize;
        $this->bucketOffset = $bucketOffset;
        $this->legend = $legend ?: new \Grafana\Foundation\Common\VizLegendOptions();
        $this->tooltip = $tooltip ?: new \Grafana\Foundation\Common\VizTooltipOptions();
        $this->combine = $combine;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{bucketCount?: int, bucketSize?: int, bucketOffset?: float, legend?: mixed, tooltip?: mixed, combine?: bool} $inputData */
        $data = $inputData;
        return new self(
            bucketCount: $data["bucketCount"] ?? null,
            bucketSize: $data["bucketSize"] ?? null,
            bucketOffset: $data["bucketOffset"] ?? null,
            legend: isset($data["legend"]) ? (function($input) {
    	/** @var array{displayMode?: string, placement?: string, showLegend?: bool, asTable?: bool, isVisible?: bool, sortBy?: string, sortDesc?: bool, width?: float, calcs?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizLegendOptions::fromArray($val);
    })($data["legend"]) : null,
            tooltip: isset($data["tooltip"]) ? (function($input) {
    	/** @var array{mode?: string, sort?: string, maxWidth?: float, maxHeight?: float} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizTooltipOptions::fromArray($val);
    })($data["tooltip"]) : null,
            combine: $data["combine"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "legend" => $this->legend,
            "tooltip" => $this->tooltip,
        ];
        if (isset($this->bucketCount)) {
            $data["bucketCount"] = $this->bucketCount;
        }
        if (isset($this->bucketSize)) {
            $data["bucketSize"] = $this->bucketSize;
        }
        if (isset($this->bucketOffset)) {
            $data["bucketOffset"] = $this->bucketOffset;
        }
        if (isset($this->combine)) {
            $data["combine"] = $this->combine;
        }
        return $data;
    }
}
