<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchRateSettings implements \JsonSerializable
{
    public ?string $unit;

    public ?string $mode;

    /**
     * @param string|null $unit
     * @param string|null $mode
     */
    public function __construct(?string $unit = null, ?string $mode = null)
    {
        $this->unit = $unit;
        $this->mode = $mode;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{unit?: string, mode?: string} $inputData */
        $data = $inputData;
        return new self(
            unit: $data["unit"] ?? null,
            mode: $data["mode"] ?? null,
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
        if (isset($this->mode)) {
            $data["mode"] = $this->mode;
        }
        return $data;
    }
}
