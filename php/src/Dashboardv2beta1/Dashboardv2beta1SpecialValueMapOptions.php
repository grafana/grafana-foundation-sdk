<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class Dashboardv2beta1SpecialValueMapOptions implements \JsonSerializable
{
    /**
     * Special value to match against
     */
    public \Grafana\Foundation\Dashboardv2beta1\SpecialValueMatch $match;

    /**
     * Config to apply when the value matches the special value
     */
    public \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult $result;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\SpecialValueMatch|null $match
     * @param \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult|null $result
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\SpecialValueMatch $match = null, ?\Grafana\Foundation\Dashboardv2beta1\ValueMappingResult $result = null)
    {
        $this->match = $match ?: \Grafana\Foundation\Dashboardv2beta1\SpecialValueMatch::True();
        $this->result = $result ?: new \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{match?: string, result?: mixed} $inputData */
        $data = $inputData;
        return new self(
            match: isset($data["match"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\SpecialValueMatch::fromValue($input); })($data["match"]) : null,
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
        $data->match = $this->match;
        $data->result = $this->result;
        return $data;
    }
}
