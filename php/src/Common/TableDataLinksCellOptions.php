<?php

namespace Grafana\Foundation\Common;

/**
 * Show data links in the cell
 */
class TableDataLinksCellOptions implements \JsonSerializable
{
    public string $type;

    public function __construct()
    {
        $this->type = "data-links";
    
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
