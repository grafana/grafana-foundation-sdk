<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Variable option specification
 */
class VariableOption implements \JsonSerializable
{
    /**
     * Whether the option is selected or not
     */
    public ?bool $selected;

    /**
     * Text to be displayed for the option
     * @var string|array<string>
     */
    public $text;

    /**
     * Value of the option
     * @var string|array<string>
     */
    public $value;

    /**
     * Additional properties for multi-props variables
     * @var array<string, string>|null
     */
    public ?array $properties;

    /**
     * @param bool|null $selected
     * @param string|array<string>|null $text
     * @param string|array<string>|null $value
     * @param array<string, string>|null $properties
     */
    public function __construct(?bool $selected = null,  $text = null,  $value = null, ?array $properties = null)
    {
        $this->selected = $selected;
        $this->text = $text ?: "";
        $this->value = $value ?: "";
        $this->properties = $properties;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{selected?: bool, text?: string|array<string>, value?: string|array<string>, properties?: array<string, string>} $inputData */
        $data = $inputData;
        return new self(
            selected: $data["selected"] ?? null,
            text: isset($data["text"]) ? (function($input) {
        switch (true) {
        case is_string($input):
            return $input;
        default:
            return $input;
    }
    })($data["text"]) : null,
            value: isset($data["value"]) ? (function($input) {
        switch (true) {
        case is_string($input):
            return $input;
        default:
            return $input;
    }
    })($data["value"]) : null,
            properties: $data["properties"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->text = $this->text;
        $data->value = $this->value;
        if (isset($this->selected)) {
            $data->selected = $this->selected;
        }
        if (isset($this->properties)) {
            $data->properties = $this->properties;
        }
        return $data;
    }
}
