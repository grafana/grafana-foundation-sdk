<?php

namespace Grafana\Foundation\Common;

class BaseDimensionConfig implements \JsonSerializable
{
    /**
     * fixed: T -- will be added by each element
     */
    public ?string $field;

    /**
     * @param string|null $field
     */
    public function __construct(?string $field = null)
    {
        $this->field = $field;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{field?: string} $inputData */
        $data = $inputData;
        return new self(
            field: $data["field"] ?? null,
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
        return $data;
    }
}
