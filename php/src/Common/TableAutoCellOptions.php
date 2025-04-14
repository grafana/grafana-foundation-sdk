<?php

namespace Grafana\Foundation\Common;

/**
 * Auto mode table cell options
 */
class TableAutoCellOptions implements \JsonSerializable
{
    public \Grafana\Foundation\Common\TableCellDisplayMode $type;

    public function __construct()
    {
        $this->type = \Grafana\Foundation\Common\TableCellDisplayMode::auto();
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
