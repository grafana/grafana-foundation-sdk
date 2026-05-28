<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class Threshold implements \JsonSerializable
{
    /**
     * Value null means -Infinity
     */
    public ?float $value;

    public string $color;

    /**
     * @param float|null $value
     * @param string|null $color
     */
    public function __construct(?float $value = null, ?string $color = null)
    {
        $this->value = $value;
        $this->color = $color ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{value?: float, color?: string} $inputData */
        $data = $inputData;
        return new self(
            value: $data["value"] ?? null,
            color: $data["color"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->color = $this->color;
        if (isset($this->value)) {
            $data->value = $this->value;
        }
        return $data;
    }
}
