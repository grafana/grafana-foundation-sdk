<?php

namespace Grafana\Foundation\Common;

/**
 * Auto mode table cell options
 */
class TableAutoCellOptions implements \JsonSerializable
{
    public \Grafana\Foundation\Common\TableCellDisplayMode $type;

    public ?bool $wrapText;

    /**
     * @param bool|null $wrapText
     */
    public function __construct(?bool $wrapText = null)
    {
        $this->type = \Grafana\Foundation\Common\TableCellDisplayMode::auto();
        $this->wrapText = $wrapText;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: "auto", wrapText?: bool} $inputData */
        $data = $inputData;
        return new self(
            wrapText: $data["wrapText"] ?? null,
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
        if (isset($this->wrapText)) {
            $data["wrapText"] = $this->wrapText;
        }
        return $data;
    }
}
