<?php

namespace Grafana\Foundation\Canvas;

class CanvasOptionsRoot implements \JsonSerializable
{
    /**
     * Name of the root element
     */
    public string $name;

    /**
     * Type of root element (frame)
     */
    public string $type;

    /**
     * The list of canvas elements attached to the root element
     * @var array<\Grafana\Foundation\Canvas\CanvasElementOptions>
     */
    public array $elements;

    /**
     * @param string|null $name
     * @param array<\Grafana\Foundation\Canvas\CanvasElementOptions>|null $elements
     */
    public function __construct(?string $name = null, ?array $elements = null)
    {
        $this->name = $name ?: "";
        $this->type = "frame";
    
        $this->elements = $elements ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, type?: string, elements?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            elements: array_filter(array_map((function($input) {
    	/** @var array{name?: string, type?: string, config?: mixed, constraint?: mixed, placement?: mixed, background?: mixed, border?: mixed, connections?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Canvas\CanvasElementOptions::fromArray($val);
    }), $data["elements"] ?? [])),
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "name" => $this->name,
            "type" => $this->type,
            "elements" => $this->elements,
        ];
        return $data;
    }
}
