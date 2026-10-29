<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * #FacetValue is a single value/count pair in a facet breakdown.
 */
class FacetValue implements \JsonSerializable
{
    public string $value;

    public int $count;

    /**
     * @param string|null $value
     * @param int|null $count
     */
    public function __construct(?string $value = null, ?int $count = null)
    {
        $this->value = $value ?: "";
        $this->count = $count ?: 0;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{value?: string, count?: int} $inputData */
        $data = $inputData;
        return new self(
            value: $data["value"] ?? null,
            count: $data["count"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->value = $this->value;
        $data->count = $this->count;
        return $data;
    }
}
