<?php

namespace Grafana\Foundation\Common;

/**
 * Json view cell options
 */
class TableImageCellOptions implements \JsonSerializable
{
    public string $type;

    public ?string $alt;

    public ?string $title;

    /**
     * @param string|null $alt
     * @param string|null $title
     */
    public function __construct(?string $alt = null, ?string $title = null)
    {
        $this->type = "image";
    
        $this->alt = $alt;
        $this->title = $title;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, alt?: string, title?: string} $inputData */
        $data = $inputData;
        return new self(
            alt: $data["alt"] ?? null,
            title: $data["title"] ?? null,
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
        if (isset($this->alt)) {
            $data["alt"] = $this->alt;
        }
        if (isset($this->title)) {
            $data["title"] = $this->title;
        }
        return $data;
    }
}
