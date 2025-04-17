<?php

namespace Grafana\Foundation\Statetimeline;

class Options implements \JsonSerializable
{
    /**
     * Show timeline values on chart
     */
    public \Grafana\Foundation\Common\VisibilityMode $showValue;

    /**
     * Controls the row height
     */
    public float $rowHeight;

    /**
     * Merge equal consecutive values
     */
    public ?bool $mergeValues;

    public \Grafana\Foundation\Common\VizLegendOptions $legend;

    public \Grafana\Foundation\Common\VizTooltipOptions $tooltip;

    /**
     * @var array<string>|null
     */
    public ?array $timezone;

    /**
     * Controls value alignment on the timelines
     */
    public ?\Grafana\Foundation\Common\TimelineValueAlignment $alignValue;

    /**
     * @param \Grafana\Foundation\Common\VisibilityMode|null $showValue
     * @param float|null $rowHeight
     * @param bool|null $mergeValues
     * @param \Grafana\Foundation\Common\VizLegendOptions|null $legend
     * @param \Grafana\Foundation\Common\VizTooltipOptions|null $tooltip
     * @param array<string>|null $timezone
     * @param \Grafana\Foundation\Common\TimelineValueAlignment|null $alignValue
     */
    public function __construct(?\Grafana\Foundation\Common\VisibilityMode $showValue = null, ?float $rowHeight = null, ?bool $mergeValues = null, ?\Grafana\Foundation\Common\VizLegendOptions $legend = null, ?\Grafana\Foundation\Common\VizTooltipOptions $tooltip = null, ?array $timezone = null, ?\Grafana\Foundation\Common\TimelineValueAlignment $alignValue = null)
    {
        $this->showValue = $showValue ?: \Grafana\Foundation\Common\VisibilityMode::auto();
        $this->rowHeight = $rowHeight ?: 0.9;
        $this->mergeValues = $mergeValues;
        $this->legend = $legend ?: new \Grafana\Foundation\Common\VizLegendOptions();
        $this->tooltip = $tooltip ?: new \Grafana\Foundation\Common\VizTooltipOptions();
        $this->timezone = $timezone;
        $this->alignValue = $alignValue;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{showValue?: string, rowHeight?: float, mergeValues?: bool, legend?: mixed, tooltip?: mixed, timezone?: array<string>, alignValue?: string} $inputData */
        $data = $inputData;
        return new self(
            showValue: isset($data["showValue"]) ? (function($input) { return \Grafana\Foundation\Common\VisibilityMode::fromValue($input); })($data["showValue"]) : null,
            rowHeight: $data["rowHeight"] ?? null,
            mergeValues: $data["mergeValues"] ?? null,
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
            timezone: $data["timezone"] ?? null,
            alignValue: isset($data["alignValue"]) ? (function($input) { return \Grafana\Foundation\Common\TimelineValueAlignment::fromValue($input); })($data["alignValue"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->showValue = $this->showValue;
        $data->rowHeight = $this->rowHeight;
        $data->legend = $this->legend;
        $data->tooltip = $this->tooltip;
        if (isset($this->mergeValues)) {
            $data->mergeValues = $this->mergeValues;
        }
        if (isset($this->timezone)) {
            $data->timezone = $this->timezone;
        }
        if (isset($this->alignValue)) {
            $data->alignValue = $this->alignValue;
        }
        return $data;
    }
}
