<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Controls tooltip options
 */
class HeatmapTooltip implements \JsonSerializable
{
    /**
     * Controls if the tooltip is shown
     */
    public bool $show;

    /**
     * Controls if the tooltip shows a histogram of the y-axis values
     */
    public ?bool $yHistogram;

    /**
     * Controls if the tooltip shows a color scale in header
     */
    public ?bool $showColorScale;

    /**
     * @param bool|null $show
     * @param bool|null $yHistogram
     * @param bool|null $showColorScale
     */
    public function __construct(?bool $show = null, ?bool $yHistogram = null, ?bool $showColorScale = null)
    {
        $this->show = $show ?: false;
        $this->yHistogram = $yHistogram;
        $this->showColorScale = $showColorScale;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{show?: bool, yHistogram?: bool, showColorScale?: bool} $inputData */
        $data = $inputData;
        return new self(
            show: $data["show"] ?? null,
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
            "show" => $this->show,
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
