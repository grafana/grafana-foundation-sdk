<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class VizTooltipOptions implements \JsonSerializable
{
    public \Grafana\Foundation\Common\TooltipDisplayMode $mode;

    public \Grafana\Foundation\Common\SortOrder $sort;

    /**
     * @param \Grafana\Foundation\Common\TooltipDisplayMode|null $mode
     * @param \Grafana\Foundation\Common\SortOrder|null $sort
     */
    public function __construct(?\Grafana\Foundation\Common\TooltipDisplayMode $mode = null, ?\Grafana\Foundation\Common\SortOrder $sort = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Common\TooltipDisplayMode::Single();
        $this->sort = $sort ?: \Grafana\Foundation\Common\SortOrder::Ascending();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, sort?: string} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Common\TooltipDisplayMode::fromValue($input); })($data["mode"]) : null,
            sort: isset($data["sort"]) ? (function($input) { return \Grafana\Foundation\Common\SortOrder::fromValue($input); })($data["sort"]) : null,
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
        return $data;
    }
}
