<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Maps numerical ranges to a display text and color.
 * For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
 */
class RangeMap implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\MappingType $type;

    /**
     * Range to match against and the result to apply when the value is within the range
     */
    public \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RangeMapOptions $options;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RangeMapOptions|null $options
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RangeMapOptions $options = null)
    {
        $this->type = \Grafana\Foundation\Dashboardv2beta1\MappingType::value();
        $this->options = $options ?: new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RangeMapOptions();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: "range", options?: mixed} $inputData */
        $data = $inputData;
        return new self(
            options: isset($data["options"]) ? (function($input) {
    	/** @var array{from?: float, to?: float, result?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RangeMapOptions::fromArray($val);
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
