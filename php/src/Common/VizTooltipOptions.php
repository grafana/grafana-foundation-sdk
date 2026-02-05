<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class VizTooltipOptions implements \JsonSerializable
{
    public \Grafana\Foundation\Common\TooltipDisplayMode $mode;

    public \Grafana\Foundation\Common\SortOrder $sort;

    public ?float $maxWidth;

    public ?float $maxHeight;

    public ?bool $hideZeros;

    /**
     * @param \Grafana\Foundation\Common\TooltipDisplayMode|null $mode
     * @param \Grafana\Foundation\Common\SortOrder|null $sort
     * @param float|null $maxWidth
     * @param float|null $maxHeight
     * @param bool|null $hideZeros
     */
    public function __construct(?\Grafana\Foundation\Common\TooltipDisplayMode $mode = null, ?\Grafana\Foundation\Common\SortOrder $sort = null, ?float $maxWidth = null, ?float $maxHeight = null, ?bool $hideZeros = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Common\TooltipDisplayMode::Single();
        $this->sort = $sort ?: \Grafana\Foundation\Common\SortOrder::Ascending();
        $this->maxWidth = $maxWidth;
        $this->maxHeight = $maxHeight;
        $this->hideZeros = $hideZeros;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, sort?: string, maxWidth?: float, maxHeight?: float, hideZeros?: bool} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Common\TooltipDisplayMode::fromValue($input); })($data["mode"]) : null,
            sort: isset($data["sort"]) ? (function($input) { return \Grafana\Foundation\Common\SortOrder::fromValue($input); })($data["sort"]) : null,
            maxWidth: $data["maxWidth"] ?? null,
            maxHeight: $data["maxHeight"] ?? null,
            hideZeros: $data["hideZeros"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->mode = $this->mode;
        $data->sort = $this->sort;
        if (isset($this->maxWidth)) {
            $data->maxWidth = $this->maxWidth;
        }
        if (isset($this->maxHeight)) {
            $data->maxHeight = $this->maxHeight;
        }
        if (isset($this->hideZeros)) {
            $data->hideZeros = $this->hideZeros;
        }
        return $data;
    }
}
