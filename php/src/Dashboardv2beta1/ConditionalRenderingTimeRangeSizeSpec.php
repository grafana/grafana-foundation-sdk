<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class ConditionalRenderingTimeRangeSizeSpec implements \JsonSerializable
{
    public string $value;

    /**
     * @param string|null $value
     */
    public function __construct(?string $value = null)
    {
        $this->value = $value ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{value?: string} $inputData */
        $data = $inputData;
        return new self(
            value: $data["value"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->value = $this->value;
        return $data;
    }
}
