<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Maps text values to a color or different display text and color.
 * For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
 */
class ValueMap implements \JsonSerializable
{
    public string $type;

    /**
     * Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }
     * @var array<string, \Grafana\Foundation\Dashboard\ValueMappingResult>
     */
    public array $options;

    /**
     * @param array<string, \Grafana\Foundation\Dashboard\ValueMappingResult>|null $options
     */
    public function __construct(?array $options = null)
    {
        $this->type = "value";
    
        $this->options = $options ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, options?: array<string, mixed>} $inputData */
        $data = $inputData;
        return new self(
            options: isset($data["options"]) ? array_map((function($input) {
    	/** @var array{text?: string, color?: string, icon?: string, index?: int} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\ValueMappingResult::fromArray($val);
    }), $data["options"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
            "options" => $this->options,
        ];
        return $data;
    }
}
