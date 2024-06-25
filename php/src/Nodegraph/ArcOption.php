<?php

namespace Grafana\Foundation\Nodegraph;

class ArcOption implements \JsonSerializable
{
    /**
     * Field from which to get the value. Values should be less than 1, representing fraction of a circle.
     */
    public ?string $field;

    /**
     * The color of the arc.
     */
    public ?string $color;

    /**
     * @param string|null $field
     * @param string|null $color
     */
    public function __construct(?string $field = null, ?string $color = null)
    {
        $this->field = $field;
        $this->color = $color;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{field?: string, color?: string} $inputData */
        $data = $inputData;
        return new self(
            field: $data["field"] ?? null,
            color: $data["color"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->field)) {
            $data["field"] = $this->field;
        }
        if (isset($this->color)) {
            $data["color"] = $this->color;
        }
        return $data;
    }
}
