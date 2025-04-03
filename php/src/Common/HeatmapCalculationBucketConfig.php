<?php

namespace Grafana\Foundation\Common;

class HeatmapCalculationBucketConfig implements \JsonSerializable
{
    /**
     * Sets the bucket calculation mode
     */
    public ?\Grafana\Foundation\Common\HeatmapCalculationMode $mode;

    /**
     * The number of buckets to use for the axis in the heatmap
     */
    public ?string $value;

    /**
     * Controls the scale of the buckets
     */
    public ?\Grafana\Foundation\Common\ScaleDistributionConfig $scale;

    /**
     * @param \Grafana\Foundation\Common\HeatmapCalculationMode|null $mode
     * @param string|null $value
     * @param \Grafana\Foundation\Common\ScaleDistributionConfig|null $scale
     */
    public function __construct(?\Grafana\Foundation\Common\HeatmapCalculationMode $mode = null, ?string $value = null, ?\Grafana\Foundation\Common\ScaleDistributionConfig $scale = null)
    {
        $this->mode = $mode;
        $this->value = $value;
        $this->scale = $scale;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, value?: string, scale?: mixed} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Common\HeatmapCalculationMode::fromValue($input); })($data["mode"]) : null,
            value: $data["value"] ?? null,
            scale: isset($data["scale"]) ? (function($input) {
    	/** @var array{type?: string, log?: float, linearThreshold?: float} */
    $val = $input;
    	return \Grafana\Foundation\Common\ScaleDistributionConfig::fromArray($val);
    })($data["scale"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->mode)) {
            $data["mode"] = $this->mode;
        }
        if (isset($this->value)) {
            $data["value"] = $this->value;
        }
        if (isset($this->scale)) {
            $data["scale"] = $this->scale;
        }
        return $data;
    }
}
