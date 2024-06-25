<?php

namespace Grafana\Foundation\Common;

/**
 * Colored text cell options
 */
class TableColorTextCellOptions implements \JsonSerializable
{
    public string $type;

    public function __construct()
    {
        $this->type = "color-text";
    
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        return new self(
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
        ];
        return $data;
    }
}
