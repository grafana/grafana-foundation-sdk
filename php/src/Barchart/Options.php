<?php

namespace Grafana\Foundation\Barchart;

class Options implements \JsonSerializable
{
    /**
     * Manually select which field from the dataset to represent the x field.
     */
    public ?string $xField;

    /**
     * Use the color value for a sibling field to color each bar value.
     */
    public ?string $colorByField;

    /**
     * Controls the orientation of the bar chart, either vertical or horizontal.
     */
    public \Grafana\Foundation\Common\VizOrientation $orientation;

    /**
     * Controls the radius of each bar.
     */
    public ?float $barRadius;

    /**
     * Controls the rotation of the x axis labels.
     */
    public int $xTickLabelRotation;

    /**
     * Sets the max length that a label can have before it is truncated.
     */
    public int $xTickLabelMaxLength;

    /**
     * Controls the spacing between x axis labels.
     * negative values indicate backwards skipping behavior
     */
    public ?int $xTickLabelSpacing;

    /**
     * Controls whether bars are stacked or not, either normally or in percent mode.
     */
    public \Grafana\Foundation\Common\StackingMode $stacking;

    /**
     * This controls whether values are shown on top or to the left of bars.
     */
    public \Grafana\Foundation\Common\VisibilityMode $showValue;

    /**
     * Controls the width of bars. 1 = Max width, 0 = Min width.
     */
    public float $barWidth;

    /**
     * Controls the width of groups. 1 = max with, 0 = min width.
     */
    public float $groupWidth;

    public \Grafana\Foundation\Common\VizLegendOptions $legend;

    public \Grafana\Foundation\Common\VizTooltipOptions $tooltip;

    public ?\Grafana\Foundation\Common\VizTextDisplayOptions $text;

    /**
     * Enables mode which highlights the entire bar area and shows tooltip when cursor
     * hovers over highlighted area
     */
    public bool $fullHighlight;

    /**
     * @param string|null $xField
     * @param string|null $colorByField
     * @param \Grafana\Foundation\Common\VizOrientation|null $orientation
     * @param float|null $barRadius
     * @param int|null $xTickLabelRotation
     * @param int|null $xTickLabelMaxLength
     * @param int|null $xTickLabelSpacing
     * @param \Grafana\Foundation\Common\StackingMode|null $stacking
     * @param \Grafana\Foundation\Common\VisibilityMode|null $showValue
     * @param float|null $barWidth
     * @param float|null $groupWidth
     * @param \Grafana\Foundation\Common\VizLegendOptions|null $legend
     * @param \Grafana\Foundation\Common\VizTooltipOptions|null $tooltip
     * @param \Grafana\Foundation\Common\VizTextDisplayOptions|null $text
     * @param bool|null $fullHighlight
     */
    public function __construct(?string $xField = null, ?string $colorByField = null, ?\Grafana\Foundation\Common\VizOrientation $orientation = null, ?float $barRadius = null, ?int $xTickLabelRotation = null, ?int $xTickLabelMaxLength = null, ?int $xTickLabelSpacing = null, ?\Grafana\Foundation\Common\StackingMode $stacking = null, ?\Grafana\Foundation\Common\VisibilityMode $showValue = null, ?float $barWidth = null, ?float $groupWidth = null, ?\Grafana\Foundation\Common\VizLegendOptions $legend = null, ?\Grafana\Foundation\Common\VizTooltipOptions $tooltip = null, ?\Grafana\Foundation\Common\VizTextDisplayOptions $text = null, ?bool $fullHighlight = null)
    {
        $this->xField = $xField;
        $this->colorByField = $colorByField;
        $this->orientation = $orientation ?: \Grafana\Foundation\Common\VizOrientation::auto();
        $this->barRadius = $barRadius;
        $this->xTickLabelRotation = $xTickLabelRotation ?: 0;
        $this->xTickLabelMaxLength = $xTickLabelMaxLength ?: 0;
        $this->xTickLabelSpacing = $xTickLabelSpacing;
        $this->stacking = $stacking ?: \Grafana\Foundation\Common\StackingMode::none();
        $this->showValue = $showValue ?: \Grafana\Foundation\Common\VisibilityMode::auto();
        $this->barWidth = $barWidth ?: 0.97;
        $this->groupWidth = $groupWidth ?: 0.7;
        $this->legend = $legend ?: new \Grafana\Foundation\Common\VizLegendOptions();
        $this->tooltip = $tooltip ?: new \Grafana\Foundation\Common\VizTooltipOptions();
        $this->text = $text;
        $this->fullHighlight = $fullHighlight ?: false;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{xField?: string, colorByField?: string, orientation?: string, barRadius?: float, xTickLabelRotation?: int, xTickLabelMaxLength?: int, xTickLabelSpacing?: int, stacking?: string, showValue?: string, barWidth?: float, groupWidth?: float, legend?: mixed, tooltip?: mixed, text?: mixed, fullHighlight?: bool} $inputData */
        $data = $inputData;
        return new self(
            xField: $data["xField"] ?? null,
            colorByField: $data["colorByField"] ?? null,
            orientation: isset($data["orientation"]) ? (function($input) { return \Grafana\Foundation\Common\VizOrientation::fromValue($input); })($data["orientation"]) : null,
            barRadius: $data["barRadius"] ?? null,
            xTickLabelRotation: $data["xTickLabelRotation"] ?? null,
            xTickLabelMaxLength: $data["xTickLabelMaxLength"] ?? null,
            xTickLabelSpacing: $data["xTickLabelSpacing"] ?? null,
            stacking: isset($data["stacking"]) ? (function($input) { return \Grafana\Foundation\Common\StackingMode::fromValue($input); })($data["stacking"]) : null,
            showValue: isset($data["showValue"]) ? (function($input) { return \Grafana\Foundation\Common\VisibilityMode::fromValue($input); })($data["showValue"]) : null,
            barWidth: $data["barWidth"] ?? null,
            groupWidth: $data["groupWidth"] ?? null,
            legend: isset($data["legend"]) ? (function($input) {
    	/** @var array{displayMode?: string, placement?: string, showLegend?: bool, asTable?: bool, isVisible?: bool, sortBy?: string, sortDesc?: bool, width?: float, calcs?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizLegendOptions::fromArray($val);
    })($data["legend"]) : null,
            tooltip: isset($data["tooltip"]) ? (function($input) {
    	/** @var array{mode?: string, sort?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizTooltipOptions::fromArray($val);
    })($data["tooltip"]) : null,
            text: isset($data["text"]) ? (function($input) {
    	/** @var array{titleSize?: float, valueSize?: float} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizTextDisplayOptions::fromArray($val);
    })($data["text"]) : null,
            fullHighlight: $data["fullHighlight"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->orientation = $this->orientation;
        $data->xTickLabelRotation = $this->xTickLabelRotation;
        $data->xTickLabelMaxLength = $this->xTickLabelMaxLength;
        $data->stacking = $this->stacking;
        $data->showValue = $this->showValue;
        $data->barWidth = $this->barWidth;
        $data->groupWidth = $this->groupWidth;
        $data->legend = $this->legend;
        $data->tooltip = $this->tooltip;
        $data->fullHighlight = $this->fullHighlight;
        if (isset($this->xField)) {
            $data->xField = $this->xField;
        }
        if (isset($this->colorByField)) {
            $data->colorByField = $this->colorByField;
        }
        if (isset($this->barRadius)) {
            $data->barRadius = $this->barRadius;
        }
        if (isset($this->xTickLabelSpacing)) {
            $data->xTickLabelSpacing = $this->xTickLabelSpacing;
        }
        if (isset($this->text)) {
            $data->text = $this->text;
        }
        return $data;
    }
}
