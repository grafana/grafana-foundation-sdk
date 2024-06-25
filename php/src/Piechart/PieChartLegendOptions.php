<?php

namespace Grafana\Foundation\Piechart;

class PieChartLegendOptions implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Piechart\PieChartLegendValues>
     */
    public array $values;

    public \Grafana\Foundation\Common\LegendDisplayMode $displayMode;

    public \Grafana\Foundation\Common\LegendPlacement $placement;

    public bool $showLegend;

    public ?bool $asTable;

    public ?bool $isVisible;

    public ?string $sortBy;

    public ?bool $sortDesc;

    public ?float $width;

    /**
     * @var array<string>
     */
    public array $calcs;

    /**
     * @param array<\Grafana\Foundation\Piechart\PieChartLegendValues>|null $values
     * @param \Grafana\Foundation\Common\LegendDisplayMode|null $displayMode
     * @param \Grafana\Foundation\Common\LegendPlacement|null $placement
     * @param bool|null $showLegend
     * @param bool|null $asTable
     * @param bool|null $isVisible
     * @param string|null $sortBy
     * @param bool|null $sortDesc
     * @param float|null $width
     * @param array<string>|null $calcs
     */
    public function __construct(?array $values = null, ?\Grafana\Foundation\Common\LegendDisplayMode $displayMode = null, ?\Grafana\Foundation\Common\LegendPlacement $placement = null, ?bool $showLegend = null, ?bool $asTable = null, ?bool $isVisible = null, ?string $sortBy = null, ?bool $sortDesc = null, ?float $width = null, ?array $calcs = null)
    {
        $this->values = $values ?: [];
        $this->displayMode = $displayMode ?: \Grafana\Foundation\Common\LegendDisplayMode::List();
        $this->placement = $placement ?: \Grafana\Foundation\Common\LegendPlacement::Bottom();
        $this->showLegend = $showLegend ?: false;
        $this->asTable = $asTable;
        $this->isVisible = $isVisible;
        $this->sortBy = $sortBy;
        $this->sortDesc = $sortDesc;
        $this->width = $width;
        $this->calcs = $calcs ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{values?: array<string>, displayMode?: string, placement?: string, showLegend?: bool, asTable?: bool, isVisible?: bool, sortBy?: string, sortDesc?: bool, width?: float, calcs?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            values: array_filter(array_map((function($input) { return \Grafana\Foundation\Piechart\PieChartLegendValues::fromValue($input); }), $data["values"] ?? [])),
            displayMode: isset($data["displayMode"]) ? (function($input) { return \Grafana\Foundation\Common\LegendDisplayMode::fromValue($input); })($data["displayMode"]) : null,
            placement: isset($data["placement"]) ? (function($input) { return \Grafana\Foundation\Common\LegendPlacement::fromValue($input); })($data["placement"]) : null,
            showLegend: $data["showLegend"] ?? null,
            asTable: $data["asTable"] ?? null,
            isVisible: $data["isVisible"] ?? null,
            sortBy: $data["sortBy"] ?? null,
            sortDesc: $data["sortDesc"] ?? null,
            width: $data["width"] ?? null,
            calcs: $data["calcs"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "values" => $this->values,
            "displayMode" => $this->displayMode,
            "placement" => $this->placement,
            "showLegend" => $this->showLegend,
            "calcs" => $this->calcs,
        ];
        if (isset($this->asTable)) {
            $data["asTable"] = $this->asTable;
        }
        if (isset($this->isVisible)) {
            $data["isVisible"] = $this->isVisible;
        }
        if (isset($this->sortBy)) {
            $data["sortBy"] = $this->sortBy;
        }
        if (isset($this->sortDesc)) {
            $data["sortDesc"] = $this->sortDesc;
        }
        if (isset($this->width)) {
            $data["width"] = $this->width;
        }
        return $data;
    }
}
