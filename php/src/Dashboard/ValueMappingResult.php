<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Result used as replacement with text and color when the value matches
 */
class ValueMappingResult implements \JsonSerializable
{
    /**
     * Text to display when the value matches
     */
    public ?string $text;

    /**
     * Text to use when the value matches
     */
    public ?string $color;

    /**
     * Icon to display when the value matches. Only specific visualizations.
     */
    public ?string $icon;

    /**
     * Position in the mapping array. Only used internally.
     */
    public ?int $index;

    /**
     * @param string|null $text
     * @param string|null $color
     * @param string|null $icon
     * @param int|null $index
     */
    public function __construct(?string $text = null, ?string $color = null, ?string $icon = null, ?int $index = null)
    {
        $this->text = $text;
        $this->color = $color;
        $this->icon = $icon;
        $this->index = $index;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{text?: string, color?: string, icon?: string, index?: int} $inputData */
        $data = $inputData;
        return new self(
            text: $data["text"] ?? null,
            color: $data["color"] ?? null,
            icon: $data["icon"] ?? null,
            index: $data["index"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->text)) {
            $data->text = $this->text;
        }
        if (isset($this->color)) {
            $data->color = $this->color;
        }
        if (isset($this->icon)) {
            $data->icon = $this->icon;
        }
        if (isset($this->index)) {
            $data->index = $this->index;
        }
        return $data;
    }
}
