<?php

namespace Grafana\Foundation\Azuremonitor;

class AzureTracesFilter implements \JsonSerializable
{
    /**
     * Property name, auto-populated based on available traces.
     */
    public string $property;

    /**
     * Comparison operator to use. Either equals or not equals.
     */
    public string $operation;

    /**
     * Values to filter by.
     * @var array<string>
     */
    public array $filters;

    /**
     * @param string|null $property
     * @param string|null $operation
     * @param array<string>|null $filters
     */
    public function __construct(?string $property = null, ?string $operation = null, ?array $filters = null)
    {
        $this->property = $property ?: "";
        $this->operation = $operation ?: "";
        $this->filters = $filters ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{property?: string, operation?: string, filters?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            property: $data["property"] ?? null,
            operation: $data["operation"] ?? null,
            filters: $data["filters"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "property" => $this->property,
            "operation" => $this->operation,
            "filters" => $this->filters,
        ];
        return $data;
    }
}
