<?php

namespace Grafana\Foundation\Common;

class ColorDimensionConfig implements \JsonSerializable
{
    /**
     * color value
     */
    public ?string $fixed;

    /**
     * fixed: T -- will be added by each element
     */
    public ?string $field;

    /**
     * @param string|null $fixed
     * @param string|null $field
     */
    public function __construct(?string $fixed = null, ?string $field = null)
    {
        $this->fixed = $fixed;
        $this->field = $field;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{fixed?: string, field?: string} $inputData */
        $data = $inputData;
        return new self(
            fixed: $data["fixed"] ?? null,
            field: $data["field"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->fixed)) {
            $data->fixed = $this->fixed;
        }
        if (isset($this->field)) {
            $data->field = $this->field;
        }
        return $data;
    }
}
