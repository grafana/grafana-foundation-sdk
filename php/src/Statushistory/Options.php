<?php

namespace Grafana\Foundation\Statushistory;

class Options implements \JsonSerializable
{
    /**
     * Set the height of the rows
     */
    public float $rowHeight;

    /**
     * Show values on the columns
     */
    public \Grafana\Foundation\Common\VisibilityMode $showValue;

    public \Grafana\Foundation\Common\VizLegendOptions $legend;

    public \Grafana\Foundation\Common\VizTooltipOptions $tooltip;

    /**
     * @var array<string>|null
     */
    public ?array $timezone;

    /**
     * Controls the column width
     */
    public ?float $colWidth;

    /**
     * @param float|null $rowHeight
     * @param \Grafana\Foundation\Common\VisibilityMode|null $showValue
     * @param \Grafana\Foundation\Common\VizLegendOptions|null $legend
     * @param \Grafana\Foundation\Common\VizTooltipOptions|null $tooltip
     * @param array<string>|null $timezone
     * @param float|null $colWidth
     */
    public function __construct(?float $rowHeight = null, ?\Grafana\Foundation\Common\VisibilityMode $showValue = null, ?\Grafana\Foundation\Common\VizLegendOptions $legend = null, ?\Grafana\Foundation\Common\VizTooltipOptions $tooltip = null, ?array $timezone = null, ?float $colWidth = null)
    {
        $this->rowHeight = $rowHeight ?: 0.9;
        $this->showValue = $showValue ?: \Grafana\Foundation\Common\VisibilityMode::auto();
        $this->legend = $legend ?: new \Grafana\Foundation\Common\VizLegendOptions();
        $this->tooltip = $tooltip ?: new \Grafana\Foundation\Common\VizTooltipOptions();
        $this->timezone = $timezone;
        $this->colWidth = $colWidth;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{rowHeight?: float, showValue?: string, legend?: mixed, tooltip?: mixed, timezone?: array<string>, colWidth?: float} $inputData */
        $data = $inputData;
        return new self(
            rowHeight: $data["rowHeight"] ?? null,
            showValue: isset($data["showValue"]) ? (function($input) { return \Grafana\Foundation\Common\VisibilityMode::fromValue($input); })($data["showValue"]) : null,
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
            colWidth: $data["colWidth"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->rowHeight = $this->rowHeight;
        $data->showValue = $this->showValue;
        $data->legend = $this->legend;
        $data->tooltip = $this->tooltip;
        if (isset($this->timezone)) {
            $data->timezone = $this->timezone;
        }
        if (isset($this->colWidth)) {
            $data->colWidth = $this->colWidth;
        }
        return $data;
    }
}
