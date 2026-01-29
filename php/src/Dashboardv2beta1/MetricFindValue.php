<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Define the MetricFindValue type
 */
class MetricFindValue implements \JsonSerializable
{
    public string $text;

    /**
     * @var string|float|null
     */
    public $value;

    public ?string $group;

    public ?bool $expandable;

    /**
     * @param string|null $text
     * @param string|float|null $value
     * @param string|null $group
     * @param bool|null $expandable
     */
    public function __construct(?string $text = null,  $value = null, ?string $group = null, ?bool $expandable = null)
    {
        $this->text = $text ?: "";
        $this->value = $value;
        $this->group = $group;
        $this->expandable = $expandable;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{text?: string, value?: string|float, group?: string, expandable?: bool} $inputData */
        $data = $inputData;
        return new self(
            text: $data["text"] ?? null,
            value: isset($data["value"]) ? (function($input) {
        switch (true) {
        case is_string($input):
            return $input;
        case is_float($input):
            return $input;
        default:
            throw new \ValueError('incorrect value for disjunction');
    }
    })($data["value"]) : null,
            group: $data["group"] ?? null,
            expandable: $data["expandable"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->text = $this->text;
        if (isset($this->value)) {
            $data->value = $this->value;
        }
        if (isset($this->group)) {
            $data->group = $this->group;
        }
        if (isset($this->expandable)) {
            $data->expandable = $this->expandable;
        }
        return $data;
    }
}
