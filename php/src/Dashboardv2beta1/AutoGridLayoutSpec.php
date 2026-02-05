<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class AutoGridLayoutSpec implements \JsonSerializable
{
    public ?float $maxColumnCount;

    public \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecColumnWidthMode $columnWidthMode;

    public ?float $columnWidth;

    public \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecRowHeightMode $rowHeightMode;

    public ?float $rowHeight;

    public ?bool $fillScreen;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemKind>
     */
    public array $items;

    /**
     * @param float|null $maxColumnCount
     * @param \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecColumnWidthMode|null $columnWidthMode
     * @param float|null $columnWidth
     * @param \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecRowHeightMode|null $rowHeightMode
     * @param float|null $rowHeight
     * @param bool|null $fillScreen
     * @param array<\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemKind>|null $items
     */
    public function __construct(?float $maxColumnCount = null, ?\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecColumnWidthMode $columnWidthMode = null, ?float $columnWidth = null, ?\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecRowHeightMode $rowHeightMode = null, ?float $rowHeight = null, ?bool $fillScreen = null, ?array $items = null)
    {
        $this->maxColumnCount = $maxColumnCount;
        $this->columnWidthMode = $columnWidthMode ?: \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecColumnWidthMode::standard();
        $this->columnWidth = $columnWidth;
        $this->rowHeightMode = $rowHeightMode ?: \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecRowHeightMode::standard();
        $this->rowHeight = $rowHeight;
        $this->fillScreen = $fillScreen;
        $this->items = $items ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{maxColumnCount?: float, columnWidthMode?: string, columnWidth?: float, rowHeightMode?: string, rowHeight?: float, fillScreen?: bool, items?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            maxColumnCount: $data["maxColumnCount"] ?? null,
            columnWidthMode: isset($data["columnWidthMode"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecColumnWidthMode::fromValue($input); })($data["columnWidthMode"]) : null,
            columnWidth: $data["columnWidth"] ?? null,
            rowHeightMode: isset($data["rowHeightMode"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecRowHeightMode::fromValue($input); })($data["rowHeightMode"]) : null,
            rowHeight: $data["rowHeight"] ?? null,
            fillScreen: $data["fillScreen"] ?? null,
            items: array_filter(array_map((function($input) {
    	/** @var array{kind?: string, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemKind::fromArray($val);
    }), $data["items"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->columnWidthMode = $this->columnWidthMode;
        $data->rowHeightMode = $this->rowHeightMode;
        $data->items = $this->items;
        if (isset($this->maxColumnCount)) {
            $data->maxColumnCount = $this->maxColumnCount;
        }
        if (isset($this->columnWidth)) {
            $data->columnWidth = $this->columnWidth;
        }
        if (isset($this->rowHeight)) {
            $data->rowHeight = $this->rowHeight;
        }
        if (isset($this->fillScreen)) {
            $data->fillScreen = $this->fillScreen;
        }
        return $data;
    }
}
