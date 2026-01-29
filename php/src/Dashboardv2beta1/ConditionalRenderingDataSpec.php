<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class ConditionalRenderingDataSpec implements \JsonSerializable
{
    public bool $value;

    /**
     * @param bool|null $value
     */
    public function __construct(?bool $value = null)
    {
        $this->value = $value ?: false;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{value?: bool} $inputData */
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
