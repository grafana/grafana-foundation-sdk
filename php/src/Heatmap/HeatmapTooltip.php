<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Controls tooltip options
 */
class HeatmapTooltip implements \JsonSerializable
{
    /**
     * Controls how the tooltip is shown
     */
    public \Grafana\Foundation\Common\TooltipDisplayMode $mode;

    public ?float $maxHeight;

    public ?float $maxWidth;

    /**
     * Controls if the tooltip shows a histogram of the y-axis values
     */
    public ?bool $yHistogram;

    /**
     * Controls if the tooltip shows a color scale in header
     */
    public ?bool $showColorScale;

    /**
     * @param \Grafana\Foundation\Common\TooltipDisplayMode|null $mode
     * @param float|null $maxHeight
     * @param float|null $maxWidth
     * @param bool|null $yHistogram
     * @param bool|null $showColorScale
     */
    public function __construct(?\Grafana\Foundation\Common\TooltipDisplayMode $mode = null, ?float $maxHeight = null, ?float $maxWidth = null, ?bool $yHistogram = null, ?bool $showColorScale = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Common\TooltipDisplayMode::Single();
        $this->maxHeight = $maxHeight;
        $this->maxWidth = $maxWidth;
        $this->yHistogram = $yHistogram;
        $this->showColorScale = $showColorScale;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, maxHeight?: float, maxWidth?: float, yHistogram?: bool, showColorScale?: bool} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Common\TooltipDisplayMode::fromValue($input); })($data["mode"]) : null,
            maxHeight: $data["maxHeight"] ?? null,
            maxWidth: $data["maxWidth"] ?? null,
            yHistogram: $data["yHistogram"] ?? null,
            showColorScale: $data["showColorScale"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->mode = $this->mode;
        if (isset($this->maxHeight)) {
            $data->maxHeight = $this->maxHeight;
        }
        if (isset($this->maxWidth)) {
            $data->maxWidth = $this->maxWidth;
        }
        if (isset($this->yHistogram)) {
            $data->yHistogram = $this->yHistogram;
        }
        if (isset($this->showColorScale)) {
            $data->showColorScale = $this->showColorScale;
        }
        return $data;
    }
}
