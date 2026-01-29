<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class Dashboardv2beta1RangeMapOptions implements \JsonSerializable
{
    /**
     * Min value of the range. It can be null which means -Infinity
     */
    public ?float $from;

    /**
     * Max value of the range. It can be null which means +Infinity
     */
    public ?float $to;

    /**
     * Config to apply when the value is within the range
     */
    public \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult $result;

    /**
     * @param float|null $from
     * @param float|null $to
     * @param \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult|null $result
     */
    public function __construct(?float $from = null, ?float $to = null, ?\Grafana\Foundation\Dashboardv2beta1\ValueMappingResult $result = null)
    {
        $this->from = $from;
        $this->to = $to;
        $this->result = $result ?: new \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{from?: float, to?: float, result?: mixed} $inputData */
        $data = $inputData;
        return new self(
            from: $data["from"] ?? null,
            to: $data["to"] ?? null,
            result: isset($data["result"]) ? (function($input) {
    	/** @var array{text?: string, color?: string, icon?: string, index?: int} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult::fromArray($val);
    })($data["result"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->result = $this->result;
        if (isset($this->from)) {
            $data->from = $this->from;
        }
        if (isset($this->to)) {
            $data->to = $this->to;
        }
        return $data;
    }
}
