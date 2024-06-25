<?php

namespace Grafana\Foundation\Heatmap;

class FieldConfig implements \JsonSerializable
{
    public ?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution;

    public ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom;

    /**
     * @param \Grafana\Foundation\Common\ScaleDistributionConfig|null $scaleDistribution
     * @param \Grafana\Foundation\Common\HideSeriesConfig|null $hideFrom
     */
    public function __construct(?\Grafana\Foundation\Common\ScaleDistributionConfig $scaleDistribution = null, ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom = null)
    {
        $this->scaleDistribution = $scaleDistribution;
        $this->hideFrom = $hideFrom;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{scaleDistribution?: mixed, hideFrom?: mixed} $inputData */
        $data = $inputData;
        return new self(
            scaleDistribution: isset($data["scaleDistribution"]) ? (function($input) {
    	/** @var array{type?: string, log?: float, linearThreshold?: float} */
    $val = $input;
    	return \Grafana\Foundation\Common\ScaleDistributionConfig::fromArray($val);
    })($data["scaleDistribution"]) : null,
            hideFrom: isset($data["hideFrom"]) ? (function($input) {
    	/** @var array{tooltip?: bool, legend?: bool, viz?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Common\HideSeriesConfig::fromArray($val);
    })($data["hideFrom"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->scaleDistribution)) {
            $data["scaleDistribution"] = $this->scaleDistribution;
        }
        if (isset($this->hideFrom)) {
            $data["hideFrom"] = $this->hideFrom;
        }
        return $data;
    }
}
