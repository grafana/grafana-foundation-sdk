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
     * Explicit percent text size
     */
    public ?float $percentSize;

    /**
     * @param float|null $titleSize
     * @param float|null $valueSize
     * @param float|null $percentSize
     */
    public function __construct(?float $titleSize = null, ?float $valueSize = null, ?float $percentSize = null)
    {
        $this->titleSize = $titleSize;
        $this->valueSize = $valueSize;
        $this->percentSize = $percentSize;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{titleSize?: float, valueSize?: float, percentSize?: float} $inputData */
        $data = $inputData;
        return new self(
            titleSize: $data["titleSize"] ?? null,
            valueSize: $data["valueSize"] ?? null,
            percentSize: $data["percentSize"] ?? null,
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
        if (isset($this->percentSize)) {
            $data->percentSize = $this->percentSize;
        }
        return $data;
    }
}
