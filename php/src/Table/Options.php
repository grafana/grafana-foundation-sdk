<?php

namespace Grafana\Foundation\Table;

class Options implements \JsonSerializable
{
    /**
     * Represents the index of the selected frame
     */
    public float $frameIndex;

    /**
     * Controls whether the panel should show the header
     */
    public bool $showHeader;

    /**
     * Controls whether the header should show icons for the column types
     */
    public ?bool $showTypeIcons;

    /**
     * Used to control row sorting
     * @var array<\Grafana\Foundation\Common\TableSortByFieldState>|null
     */
    public ?array $sortBy;

    /**
     * Controls footer options
     */
    public ?\Grafana\Foundation\Common\TableFooterOptions $footer;

    /**
     * Controls the height of the rows
     */
    public ?\Grafana\Foundation\Common\TableCellHeight $cellHeight;

    /**
     * @param float|null $frameIndex
     * @param bool|null $showHeader
     * @param bool|null $showTypeIcons
     * @param array<\Grafana\Foundation\Common\TableSortByFieldState>|null $sortBy
     * @param \Grafana\Foundation\Common\TableFooterOptions|null $footer
     * @param \Grafana\Foundation\Common\TableCellHeight|null $cellHeight
     */
    public function __construct(?float $frameIndex = null, ?bool $showHeader = null, ?bool $showTypeIcons = null, ?array $sortBy = null, ?\Grafana\Foundation\Common\TableFooterOptions $footer = null, ?\Grafana\Foundation\Common\TableCellHeight $cellHeight = null)
    {
        $this->frameIndex = $frameIndex ?: 0;
        $this->showHeader = $showHeader ?: true;
        $this->showTypeIcons = $showTypeIcons;
        $this->sortBy = $sortBy;
        $this->footer = $footer;
        $this->cellHeight = $cellHeight;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{frameIndex?: float, showHeader?: bool, showTypeIcons?: bool, sortBy?: array<mixed>, footer?: mixed, cellHeight?: string} $inputData */
        $data = $inputData;
        return new self(
            frameIndex: $data["frameIndex"] ?? null,
            showHeader: $data["showHeader"] ?? null,
            showTypeIcons: $data["showTypeIcons"] ?? null,
            sortBy: array_filter(array_map((function($input) {
    	/** @var array{displayName?: string, desc?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Common\TableSortByFieldState::fromArray($val);
    }), $data["sortBy"] ?? [])),
            footer: isset($data["footer"]) ? (function($input) {
    	/** @var array{show?: bool, reducer?: array<string>, fields?: array<string>, enablePagination?: bool, countRows?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Common\TableFooterOptions::fromArray($val);
    })($data["footer"]) : null,
            cellHeight: isset($data["cellHeight"]) ? (function($input) { return \Grafana\Foundation\Common\TableCellHeight::fromValue($input); })($data["cellHeight"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "frameIndex" => $this->frameIndex,
            "showHeader" => $this->showHeader,
        ];
        if (isset($this->showTypeIcons)) {
            $data["showTypeIcons"] = $this->showTypeIcons;
        }
        if (isset($this->sortBy)) {
            $data["sortBy"] = $this->sortBy;
        }
        if (isset($this->footer)) {
            $data["footer"] = $this->footer;
        }
        if (isset($this->cellHeight)) {
            $data["cellHeight"] = $this->cellHeight;
        }
        return $data;
    }
}
