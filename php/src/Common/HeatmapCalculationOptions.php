<?php

namespace Grafana\Foundation\Common;

class HeatmapCalculationOptions implements \JsonSerializable
{
    /**
     * The number of buckets to use for the xAxis in the heatmap
     */
    public ?\Grafana\Foundation\Common\HeatmapCalculationBucketConfig $xBuckets;

    /**
     * The number of buckets to use for the yAxis in the heatmap
     */
    public ?\Grafana\Foundation\Common\HeatmapCalculationBucketConfig $yBuckets;

    /**
     * @param \Grafana\Foundation\Common\HeatmapCalculationBucketConfig|null $xBuckets
     * @param \Grafana\Foundation\Common\HeatmapCalculationBucketConfig|null $yBuckets
     */
    public function __construct(?\Grafana\Foundation\Common\HeatmapCalculationBucketConfig $xBuckets = null, ?\Grafana\Foundation\Common\HeatmapCalculationBucketConfig $yBuckets = null)
    {
        $this->xBuckets = $xBuckets;
        $this->yBuckets = $yBuckets;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{xBuckets?: mixed, yBuckets?: mixed} $inputData */
        $data = $inputData;
        return new self(
            xBuckets: isset($data["xBuckets"]) ? (function($input) {
    	/** @var array{mode?: string, value?: string, scale?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Common\HeatmapCalculationBucketConfig::fromArray($val);
    })($data["xBuckets"]) : null,
            yBuckets: isset($data["yBuckets"]) ? (function($input) {
    	/** @var array{mode?: string, value?: string, scale?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Common\HeatmapCalculationBucketConfig::fromArray($val);
    })($data["yBuckets"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->xBuckets)) {
            $data["xBuckets"] = $this->xBuckets;
        }
        if (isset($this->yBuckets)) {
            $data["yBuckets"] = $this->yBuckets;
        }
        return $data;
    }
}
