<?php

namespace Grafana\Foundation\Nodegraph;

class EdgeOptions implements \JsonSerializable
{
    /**
     * Unit for the main stat to override what ever is set in the data frame.
     */
    public ?string $mainStatUnit;

    /**
     * Unit for the secondary stat to override what ever is set in the data frame.
     */
    public ?string $secondaryStatUnit;

    /**
     * @param string|null $mainStatUnit
     * @param string|null $secondaryStatUnit
     */
    public function __construct(?string $mainStatUnit = null, ?string $secondaryStatUnit = null)
    {
        $this->mainStatUnit = $mainStatUnit;
        $this->secondaryStatUnit = $secondaryStatUnit;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mainStatUnit?: string, secondaryStatUnit?: string} $inputData */
        $data = $inputData;
        return new self(
            mainStatUnit: $data["mainStatUnit"] ?? null,
            secondaryStatUnit: $data["secondaryStatUnit"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->mainStatUnit)) {
            $data["mainStatUnit"] = $this->mainStatUnit;
        }
        if (isset($this->secondaryStatUnit)) {
            $data["secondaryStatUnit"] = $this->secondaryStatUnit;
        }
        return $data;
    }
}
