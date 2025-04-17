<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class VizLegendOptions implements \JsonSerializable
{
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
    public function __construct(?\Grafana\Foundation\Common\LegendDisplayMode $displayMode = null, ?\Grafana\Foundation\Common\LegendPlacement $placement = null, ?bool $showLegend = null, ?bool $asTable = null, ?bool $isVisible = null, ?string $sortBy = null, ?bool $sortDesc = null, ?float $width = null, ?array $calcs = null)
    {
        $this->displayMode = $displayMode ?: \Grafana\Foundation\Common\LegendDisplayMode::list();
        $this->placement = $placement ?: \Grafana\Foundation\Common\LegendPlacement::bottom();
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
        /** @var array{displayMode?: string, placement?: string, showLegend?: bool, asTable?: bool, isVisible?: bool, sortBy?: string, sortDesc?: bool, width?: float, calcs?: array<string>} $inputData */
        $data = $inputData;
        return new self(
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
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->displayMode = $this->displayMode;
        $data->placement = $this->placement;
        $data->showLegend = $this->showLegend;
        $data->calcs = $this->calcs;
        if (isset($this->asTable)) {
            $data->asTable = $this->asTable;
        }
        if (isset($this->isVisible)) {
            $data->isVisible = $this->isVisible;
        }
        if (isset($this->sortBy)) {
            $data->sortBy = $this->sortBy;
        }
        if (isset($this->sortDesc)) {
            $data->sortDesc = $this->sortDesc;
        }
        if (isset($this->width)) {
            $data->width = $this->width;
        }
        return $data;
    }
}
