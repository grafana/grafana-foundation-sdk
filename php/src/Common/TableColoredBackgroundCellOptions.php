<?php

namespace Grafana\Foundation\Common;

/**
 * Colored background cell options
 */
class TableColoredBackgroundCellOptions implements \JsonSerializable
{
    public string $type;

    public ?\Grafana\Foundation\Common\TableCellBackgroundDisplayMode $mode;

    public ?bool $applyToRow;

    /**
     * @param \Grafana\Foundation\Common\TableCellBackgroundDisplayMode|null $mode
     * @param bool|null $applyToRow
     */
    public function __construct(?\Grafana\Foundation\Common\TableCellBackgroundDisplayMode $mode = null, ?bool $applyToRow = null)
    {
        $this->type = "color-background";
    
        $this->mode = $mode;
        $this->applyToRow = $applyToRow;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, mode?: string, applyToRow?: bool} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Common\TableCellBackgroundDisplayMode::fromValue($input); })($data["mode"]) : null,
            applyToRow: $data["applyToRow"] ?? null,
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
        if (isset($this->mode)) {
            $data["mode"] = $this->mode;
        }
        if (isset($this->applyToRow)) {
            $data["applyToRow"] = $this->applyToRow;
        }
        return $data;
    }
}
