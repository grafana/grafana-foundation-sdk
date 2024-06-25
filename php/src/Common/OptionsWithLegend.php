<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class OptionsWithLegend implements \JsonSerializable
{
    public \Grafana\Foundation\Common\VizLegendOptions $legend;

    /**
     * @param \Grafana\Foundation\Common\VizLegendOptions|null $legend
     */
    public function __construct(?\Grafana\Foundation\Common\VizLegendOptions $legend = null)
    {
        $this->legend = $legend ?: new \Grafana\Foundation\Common\VizLegendOptions();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{legend?: mixed} $inputData */
        $data = $inputData;
        return new self(
            legend: isset($data["legend"]) ? (function($input) {
    	/** @var array{displayMode?: string, placement?: string, showLegend?: bool, asTable?: bool, isVisible?: bool, sortBy?: string, sortDesc?: bool, width?: float, calcs?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizLegendOptions::fromArray($val);
    })($data["legend"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "legend" => $this->legend,
        ];
        return $data;
    }
}
