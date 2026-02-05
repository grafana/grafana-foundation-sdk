<?php

namespace Grafana\Foundation\Datagrid;

class Options implements \JsonSerializable
{
    public int $selectedSeries;

    /**
     * @param int|null $selectedSeries
     */
    public function __construct(?int $selectedSeries = null)
    {
        $this->selectedSeries = $selectedSeries ?: 0;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{selectedSeries?: int} $inputData */
        $data = $inputData;
        return new self(
            selectedSeries: $data["selectedSeries"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->selectedSeries = $this->selectedSeries;
        return $data;
    }
}
