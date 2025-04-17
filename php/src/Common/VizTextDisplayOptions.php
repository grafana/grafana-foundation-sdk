<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class VizTextDisplayOptions implements \JsonSerializable
{
    /**
     * Explicit title text size
     */
    public ?float $titleSize;

    /**
     * Explicit value text size
     */
    public ?float $valueSize;

    /**
     * @param float|null $titleSize
     * @param float|null $valueSize
     */
    public function __construct(?float $titleSize = null, ?float $valueSize = null)
    {
        $this->titleSize = $titleSize;
        $this->valueSize = $valueSize;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{titleSize?: float, valueSize?: float} $inputData */
        $data = $inputData;
        return new self(
            titleSize: $data["titleSize"] ?? null,
            valueSize: $data["valueSize"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->titleSize)) {
            $data->titleSize = $this->titleSize;
        }
        if (isset($this->valueSize)) {
            $data->valueSize = $this->valueSize;
        }
        return $data;
    }
}
