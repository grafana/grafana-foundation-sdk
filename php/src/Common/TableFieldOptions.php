<?php

namespace Grafana\Foundation\Common;

/**
 * Field options for each field within a table (e.g 10, "The String", 64.20, etc.)
 * Generally defines alignment, filtering capabilties, display options, etc.
 */
class TableFieldOptions implements \JsonSerializable
{
    public ?float $width;

    public ?float $minWidth;

    public \Grafana\Foundation\Common\FieldTextAlignment $align;

    /**
     * This field is deprecated in favor of using cellOptions
     */
    public ?\Grafana\Foundation\Common\TableCellDisplayMode $displayMode;

    /**
     * @var \Grafana\Foundation\Common\TableAutoCellOptions|\Grafana\Foundation\Common\TableSparklineCellOptions|\Grafana\Foundation\Common\TableBarGaugeCellOptions|\Grafana\Foundation\Common\TableColoredBackgroundCellOptions|\Grafana\Foundation\Common\TableColorTextCellOptions|\Grafana\Foundation\Common\TableImageCellOptions|\Grafana\Foundation\Common\TableDataLinksCellOptions|\Grafana\Foundation\Common\TableActionsCellOptions|\Grafana\Foundation\Common\TableJsonViewCellOptions
     */
    public $cellOptions;

    /**
     * ?? default is missing or false ??
     */
    public ?bool $hidden;

    public bool $inspect;

    public ?bool $filterable;

    /**
     * Hides any header for a column, useful for columns that show some static content or buttons.
     */
    public ?bool $hideHeader;

    /**
     * @param float|null $width
     * @param float|null $minWidth
     * @param \Grafana\Foundation\Common\FieldTextAlignment|null $align
     * @param \Grafana\Foundation\Common\TableCellDisplayMode|null $displayMode
     * @param \Grafana\Foundation\Common\TableAutoCellOptions|\Grafana\Foundation\Common\TableSparklineCellOptions|\Grafana\Foundation\Common\TableBarGaugeCellOptions|\Grafana\Foundation\Common\TableColoredBackgroundCellOptions|\Grafana\Foundation\Common\TableColorTextCellOptions|\Grafana\Foundation\Common\TableImageCellOptions|\Grafana\Foundation\Common\TableDataLinksCellOptions|\Grafana\Foundation\Common\TableActionsCellOptions|\Grafana\Foundation\Common\TableJsonViewCellOptions|null $cellOptions
     * @param bool|null $hidden
     * @param bool|null $inspect
     * @param bool|null $filterable
     * @param bool|null $hideHeader
     */
    public function __construct(?float $width = null, ?float $minWidth = null, ?\Grafana\Foundation\Common\FieldTextAlignment $align = null, ?\Grafana\Foundation\Common\TableCellDisplayMode $displayMode = null,  $cellOptions = null, ?bool $hidden = null, ?bool $inspect = null, ?bool $filterable = null, ?bool $hideHeader = null)
    {
        $this->width = $width;
        $this->minWidth = $minWidth;
        $this->align = $align ?: \Grafana\Foundation\Common\FieldTextAlignment::auto();
        $this->displayMode = $displayMode;
        $this->cellOptions = $cellOptions ?: new \Grafana\Foundation\Common\TableAutoCellOptions();
        $this->hidden = $hidden;
        $this->inspect = $inspect ?: false;
        $this->filterable = $filterable;
        $this->hideHeader = $hideHeader;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{width?: float, minWidth?: float, align?: string, displayMode?: string, cellOptions?: mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed, hidden?: bool, inspect?: bool, filterable?: bool, hideHeader?: bool} $inputData */
        $data = $inputData;
        return new self(
            width: $data["width"] ?? null,
            minWidth: $data["minWidth"] ?? null,
            align: isset($data["align"]) ? (function($input) { return \Grafana\Foundation\Common\FieldTextAlignment::fromValue($input); })($data["align"]) : null,
            displayMode: isset($data["displayMode"]) ? (function($input) { return \Grafana\Foundation\Common\TableCellDisplayMode::fromValue($input); })($data["displayMode"]) : null,
            cellOptions: isset($data["cellOptions"]) ? (function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
    
        switch ($input["type"]) {
        case "actions":
            return TableActionsCellOptions::fromArray($input);
        case "auto":
            return TableAutoCellOptions::fromArray($input);
        case "color-background":
            return TableColoredBackgroundCellOptions::fromArray($input);
        case "color-text":
            return TableColorTextCellOptions::fromArray($input);
        case "data-links":
            return TableDataLinksCellOptions::fromArray($input);
        case "gauge":
            return TableBarGaugeCellOptions::fromArray($input);
        case "image":
            return TableImageCellOptions::fromArray($input);
        case "json-view":
            return TableJsonViewCellOptions::fromArray($input);
        case "sparkline":
            return TableSparklineCellOptions::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    })($data["cellOptions"]) : null,
            hidden: $data["hidden"] ?? null,
            inspect: $data["inspect"] ?? null,
            filterable: $data["filterable"] ?? null,
            hideHeader: $data["hideHeader"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "align" => $this->align,
            "cellOptions" => $this->cellOptions,
            "inspect" => $this->inspect,
        ];
        if (isset($this->width)) {
            $data["width"] = $this->width;
        }
        if (isset($this->minWidth)) {
            $data["minWidth"] = $this->minWidth;
        }
        if (isset($this->displayMode)) {
            $data["displayMode"] = $this->displayMode;
        }
        if (isset($this->hidden)) {
            $data["hidden"] = $this->hidden;
        }
        if (isset($this->filterable)) {
            $data["filterable"] = $this->filterable;
        }
        if (isset($this->hideHeader)) {
            $data["hideHeader"] = $this->hideHeader;
        }
        return $data;
    }
}
