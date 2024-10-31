<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchDerivativeSettings implements \JsonSerializable
{
    public ?string $unit;

    /**
     * @param string|null $unit
     */
    public function __construct(?string $unit = null)
    {
        $this->unit = $unit;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{unit?: string} $inputData */
        $data = $inputData;
        return new self(
            unit: $data["unit"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->unit)) {
            $data["unit"] = $this->unit;
        }
        return $data;
    }
}
