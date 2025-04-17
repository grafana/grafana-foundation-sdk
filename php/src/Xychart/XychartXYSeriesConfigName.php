<?php

namespace Grafana\Foundation\Xychart;

class XychartXYSeriesConfigName implements \JsonSerializable
{
    public ?string $fixed;

    /**
     * @param string|null $fixed
     */
    public function __construct(?string $fixed = null)
    {
        $this->fixed = $fixed;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{fixed?: string} $inputData */
        $data = $inputData;
        return new self(
            fixed: $data["fixed"] ?? null,
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
        return $data;
    }
}
