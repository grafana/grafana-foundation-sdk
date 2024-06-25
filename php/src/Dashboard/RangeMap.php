<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Maps numerical ranges to a display text and color.
 * For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
 */
class RangeMap implements \JsonSerializable
{
    public string $type;

    /**
     * Range to match against and the result to apply when the value is within the range
     */
    public \Grafana\Foundation\Dashboard\DashboardRangeMapOptions $options;

    /**
     * @param \Grafana\Foundation\Dashboard\DashboardRangeMapOptions|null $options
     */
    public function __construct(?\Grafana\Foundation\Dashboard\DashboardRangeMapOptions $options = null)
    {
        $this->type = "range";
    
        $this->options = $options ?: new \Grafana\Foundation\Dashboard\DashboardRangeMapOptions();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, options?: mixed} $inputData */
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
