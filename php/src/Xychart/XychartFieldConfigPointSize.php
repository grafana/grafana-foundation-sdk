<?php

namespace Grafana\Foundation\Xychart;

class XychartFieldConfigPointSize implements \JsonSerializable
{
    public ?int $fixed;

    public ?int $min;

    public ?int $max;

    /**
     * @param int|null $fixed
     * @param int|null $min
     * @param int|null $max
     */
    public function __construct(?int $fixed = null, ?int $min = null, ?int $max = null)
    {
        $this->fixed = $fixed;
        $this->min = $min;
        $this->max = $max;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{fixed?: int, min?: int, max?: int} $inputData */
        $data = $inputData;
        return new self(
            fixed: $data["fixed"] ?? null,
            min: $data["min"] ?? null,
            max: $data["max"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->fixed)) {
            $data["fixed"] = $this->fixed;
        }
        if (isset($this->min)) {
            $data["min"] = $this->min;
        }
        if (isset($this->max)) {
            $data["max"] = $this->max;
        }
        return $data;
    }
}
