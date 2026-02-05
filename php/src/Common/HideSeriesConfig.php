<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class HideSeriesConfig implements \JsonSerializable
{
    public bool $tooltip;

    public bool $legend;

    public bool $viz;

    /**
     * @param bool|null $tooltip
     * @param bool|null $legend
     * @param bool|null $viz
     */
    public function __construct(?bool $tooltip = null, ?bool $legend = null, ?bool $viz = null)
    {
        $this->tooltip = $tooltip ?: false;
        $this->legend = $legend ?: false;
        $this->viz = $viz ?: false;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{tooltip?: bool, legend?: bool, viz?: bool} $inputData */
        $data = $inputData;
        return new self(
            tooltip: $data["tooltip"] ?? null,
            legend: $data["legend"] ?? null,
            viz: $data["viz"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->tooltip = $this->tooltip;
        $data->legend = $this->legend;
        $data->viz = $this->viz;
        return $data;
    }
}
