<?php

namespace Grafana\Foundation\Azuremonitor;

class AzureMetricDimension implements \JsonSerializable
{
    /**
     * Name of Dimension to be filtered on.
     */
    public ?string $dimension;

    /**
     * String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
     */
    public ?string $operator;

    /**
     * Values to match with the filter.
     * @var array<string>|null
     */
    public ?array $filters;

    /**
     * @deprecated filter is deprecated in favour of filters to support multiselect.
     */
    public ?string $filter;

    /**
     * @param string|null $dimension
     * @param string|null $operator
     * @param array<string>|null $filters
     * @param string|null $filter
     */
    public function __construct(?string $dimension = null, ?string $operator = null, ?array $filters = null, ?string $filter = null)
    {
        $this->dimension = $dimension;
        $this->operator = $operator;
        $this->filters = $filters;
        $this->filter = $filter;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{dimension?: string, operator?: string, filters?: array<string>, filter?: string} $inputData */
        $data = $inputData;
        return new self(
            dimension: $data["dimension"] ?? null,
            operator: $data["operator"] ?? null,
            filters: $data["filters"] ?? null,
            filter: $data["filter"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->dimension)) {
            $data->dimension = $this->dimension;
        }
        if (isset($this->operator)) {
            $data->operator = $this->operator;
        }
        if (isset($this->filters)) {
            $data->filters = $this->filters;
        }
        if (isset($this->filter)) {
            $data->filter = $this->filter;
        }
        return $data;
    }
}
