<?php

namespace Grafana\Foundation\Dashboard;

class DashboardGraphPanelLegend implements \JsonSerializable
{
    public bool $show;

    public ?string $sort;

    public ?bool $sortDesc;

    /**
     * @param bool|null $show
     * @param string|null $sort
     * @param bool|null $sortDesc
     */
    public function __construct(?bool $show = null, ?string $sort = null, ?bool $sortDesc = null)
    {
        $this->show = $show ?: true;
        $this->sort = $sort;
        $this->sortDesc = $sortDesc;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{show?: bool, sort?: string, sortDesc?: bool} $inputData */
        $data = $inputData;
        return new self(
            show: $data["show"] ?? null,
            sort: $data["sort"] ?? null,
            sortDesc: $data["sortDesc"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->show = $this->show;
        if (isset($this->sort)) {
            $data->sort = $this->sort;
        }
        if (isset($this->sortDesc)) {
            $data->sortDesc = $this->sortDesc;
        }
        return $data;
    }
}
