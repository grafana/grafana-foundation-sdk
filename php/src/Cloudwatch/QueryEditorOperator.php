<?php

namespace Grafana\Foundation\Cloudwatch;

/**
 * TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer
 */
class QueryEditorOperator implements \JsonSerializable
{
    public ?string $name;

    /**
     * @var string|bool|int|array<string|bool|int>|null
     */
    public $value;

    /**
     * @param string|null $name
     * @param string|bool|int|array<string|bool|int>|null $value
     */
    public function __construct(?string $name = null,  $value = null)
    {
        $this->name = $name;
        $this->value = $value;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, value?: string|bool|int|array<string|bool|int>} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            value: isset($data["value"]) ? (function($input) {
        switch (true) {
        case is_string($input):
            return $input;
        case is_bool($input):
            return $input;
        case is_int($input):
            return $input;
        default:
            return $input;
    }
    })($data["value"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->name)) {
            $data->name = $this->name;
        }
        if (isset($this->value)) {
            $data->value = $this->value;
        }
        return $data;
    }
}
