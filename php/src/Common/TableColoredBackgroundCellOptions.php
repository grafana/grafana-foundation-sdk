<?php

namespace Grafana\Foundation\Common;

/**
 * Colored background cell options
 */
class TableColoredBackgroundCellOptions implements \JsonSerializable
{
    public string $type;

    public ?\Grafana\Foundation\Common\TableCellBackgroundDisplayMode $mode;

    /**
     * @param \Grafana\Foundation\Common\TableCellBackgroundDisplayMode|null $mode
     */
    public function __construct(?\Grafana\Foundation\Common\TableCellBackgroundDisplayMode $mode = null)
    {
        $this->type = "color-background";
    
        $this->mode = $mode;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, mode?: string} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Common\TableCellBackgroundDisplayMode::fromValue($input); })($data["mode"]) : null,
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
        return $data;
    }
}
