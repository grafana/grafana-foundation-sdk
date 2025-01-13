<?php

namespace Grafana\Foundation\Timeseries;

class Options implements \JsonSerializable
{
    /**
     * @var array<string>|null
     */
    public ?array $timezone;

    public \Grafana\Foundation\Common\VizLegendOptions $legend;

    public \Grafana\Foundation\Common\VizTooltipOptions $tooltip;

    public ?\Grafana\Foundation\Common\VizOrientation $orientation;

    /**
     * @param array<string>|null $timezone
     * @param \Grafana\Foundation\Common\VizLegendOptions|null $legend
     * @param \Grafana\Foundation\Common\VizTooltipOptions|null $tooltip
     * @param \Grafana\Foundation\Common\VizOrientation|null $orientation
     */
    public function __construct(?array $timezone = null, ?\Grafana\Foundation\Common\VizLegendOptions $legend = null, ?\Grafana\Foundation\Common\VizTooltipOptions $tooltip = null, ?\Grafana\Foundation\Common\VizOrientation $orientation = null)
    {
        $this->timezone = $timezone;
        $this->legend = $legend ?: new \Grafana\Foundation\Common\VizLegendOptions(calcs: [], displayMode: \Grafana\Foundation\Common\LegendDisplayMode::list(), placement: \Grafana\Foundation\Common\LegendPlacement::bottom());
        $this->tooltip = $tooltip ?: new \Grafana\Foundation\Common\VizTooltipOptions();
        $this->orientation = $orientation;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{timezone?: array<string>, legend?: mixed, tooltip?: mixed, orientation?: string} $inputData */
        $data = $inputData;
        return new self(
            timezone: $data["timezone"] ?? null,
            legend: isset($data["legend"]) ? (function($input) {
    	/** @var array{displayMode?: string, placement?: string, showLegend?: bool, asTable?: bool, isVisible?: bool, sortBy?: string, sortDesc?: bool, width?: float, calcs?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizLegendOptions::fromArray($val);
    })($data["legend"]) : null,
            tooltip: isset($data["tooltip"]) ? (function($input) {
    	/** @var array{mode?: string, sort?: string, maxWidth?: float, maxHeight?: float, hideZeros?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizTooltipOptions::fromArray($val);
    })($data["tooltip"]) : null,
            orientation: isset($data["orientation"]) ? (function($input) { return \Grafana\Foundation\Common\VizOrientation::fromValue($input); })($data["orientation"]) : null,
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
        if (isset($this->timezone)) {
            $data["timezone"] = $this->timezone;
        }
        if (isset($this->orientation)) {
            $data["orientation"] = $this->orientation;
        }
        return $data;
    }
}
