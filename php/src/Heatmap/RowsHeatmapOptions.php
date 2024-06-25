<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Controls frame rows options
 */
class RowsHeatmapOptions implements \JsonSerializable
{
    /**
     * Sets the name of the cell when not calculating from data
     */
    public ?string $value;

    /**
     * Controls tick alignment when not calculating from data
     */
    public ?\Grafana\Foundation\Common\HeatmapCellLayout $layout;

    /**
     * @param string|null $value
     * @param \Grafana\Foundation\Common\HeatmapCellLayout|null $layout
     */
    public function __construct(?string $value = null, ?\Grafana\Foundation\Common\HeatmapCellLayout $layout = null)
    {
        $this->value = $value;
        $this->layout = $layout;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{value?: string, layout?: string} $inputData */
        $data = $inputData;
        return new self(
            value: $data["value"] ?? null,
            layout: isset($data["layout"]) ? (function($input) { return \Grafana\Foundation\Common\HeatmapCellLayout::fromValue($input); })($data["layout"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->value)) {
            $data["value"] = $this->value;
        }
        if (isset($this->layout)) {
            $data["layout"] = $this->layout;
        }
        return $data;
    }
}
