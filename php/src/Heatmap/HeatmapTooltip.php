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
     * @param bool|null $show
     * @param bool|null $yHistogram
     */
    public function __construct(?bool $show = null, ?bool $yHistogram = null)
    {
        $this->show = $show ?: false;
        $this->yHistogram = $yHistogram;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{show?: bool, yHistogram?: bool} $inputData */
        $data = $inputData;
        return new self(
            show: $data["show"] ?? null,
            yHistogram: $data["yHistogram"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->show = $this->show;
        if (isset($this->yHistogram)) {
            $data->yHistogram = $this->yHistogram;
        }
        return $data;
    }
}
