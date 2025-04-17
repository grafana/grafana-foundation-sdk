<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Maps numerical ranges to a display text and color.
 * For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
 */
class RangeMap implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboard\MappingType $type;

    /**
     * Range to match against and the result to apply when the value is within the range
     */
    public \Grafana\Foundation\Dashboard\DashboardRangeMapOptions $options;

    /**
     * @param \Grafana\Foundation\Dashboard\DashboardRangeMapOptions|null $options
     */
    public function __construct(?\Grafana\Foundation\Dashboard\DashboardRangeMapOptions $options = null)
    {
        $this->type = \Grafana\Foundation\Dashboard\MappingType::valueToText();
        $this->options = $options ?: new \Grafana\Foundation\Dashboard\DashboardRangeMapOptions();
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
    	return \Grafana\Foundation\Dashboard\DashboardRangeMapOptions::fromArray($val);
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
