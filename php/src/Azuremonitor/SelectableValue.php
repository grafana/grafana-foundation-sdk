<?php

namespace Grafana\Foundation\Azuremonitor;

class SelectableValue implements \JsonSerializable
{
    public string $label;

    public string $value;

    /**
     * @param string|null $label
     * @param string|null $value
     */
    public function __construct(?string $label = null, ?string $value = null)
    {
        $this->label = $label ?: "";
        $this->value = $value ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{label?: string, value?: string} $inputData */
        $data = $inputData;
        return new self(
            label: $data["label"] ?? null,
            value: $data["value"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "label" => $this->label,
            "value" => $this->value,
        ];
        return $data;
    }
}
