<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Maps text values to a color or different display text and color.
 * For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
 */
class ValueMap implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\MappingType $type;

    /**
     * Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }
     * @var array<string, \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult>
     */
    public array $options;

    /**
     * @param array<string, \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult>|null $options
     */
    public function __construct(?array $options = null)
    {
        $this->type = \Grafana\Foundation\Dashboardv2beta1\MappingType::value();
        $this->options = $options ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: "value", options?: array<string, mixed>} $inputData */
        $data = $inputData;
        return new self(
            options: isset($data["options"]) ? (function($input) {
        /** @var array<string, \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult> $results */
        $results = [];
        foreach ($input as $key => $val) {
            $results[$key] = isset($val) ? (function($input) {
    	/** @var array{text?: string, color?: string, icon?: string, index?: int} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult::fromArray($val);
    })($val) : null;
        }
        return array_filter($results);
    })($data["options"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->type = $this->type;
        $data->options = $this->options;
        return $data;
    }
}
