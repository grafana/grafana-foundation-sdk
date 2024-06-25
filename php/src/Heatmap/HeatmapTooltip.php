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
     * @param bool|null $yHistogram
     * @param bool|null $showColorScale
     */
    public function __construct(?\Grafana\Foundation\Common\TooltipDisplayMode $mode = null, ?bool $yHistogram = null, ?bool $showColorScale = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Common\TooltipDisplayMode::Single();
        $this->yHistogram = $yHistogram;
        $this->showColorScale = $showColorScale;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, yHistogram?: bool, showColorScale?: bool} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Common\TooltipDisplayMode::fromValue($input); })($data["mode"]) : null,
            yHistogram: $data["yHistogram"] ?? null,
            showColorScale: $data["showColorScale"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "mode" => $this->mode,
        ];
        if (isset($this->yHistogram)) {
            $data["yHistogram"] = $this->yHistogram;
        }
        if (isset($this->showColorScale)) {
            $data["showColorScale"] = $this->showColorScale;
        }
        return $data;
    }
}
