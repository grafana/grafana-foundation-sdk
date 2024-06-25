<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Controls legend options
 */
class HeatmapLegend implements \JsonSerializable
{
    /**
     * Controls if the legend is shown
     */
    public bool $show;

    /**
     * @param bool|null $show
     */
    public function __construct(?bool $show = null)
    {
        $this->show = $show ?: false;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{show?: bool} $inputData */
        $data = $inputData;
        return new self(
            show: $data["show"] ?? null,
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
        return $data;
    }
}
