<?php

namespace Grafana\Foundation\Heatmap;

/**
 * Controls exemplar options
 */
class ExemplarConfig implements \JsonSerializable
{
    /**
     * Sets the color of the exemplar markers
     */
    public string $color;

    /**
     * @param string|null $color
     */
    public function __construct(?string $color = null)
    {
        $this->color = $color ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{color?: string} $inputData */
        $data = $inputData;
        return new self(
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
        return $data;
    }
}
