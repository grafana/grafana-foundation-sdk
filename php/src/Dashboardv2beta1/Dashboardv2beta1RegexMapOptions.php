<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class Dashboardv2beta1RegexMapOptions implements \JsonSerializable
{
    /**
     * Regular expression to match against
     */
    public string $pattern;

    /**
     * Config to apply when the value matches the regex
     */
    public \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult $result;

    /**
     * @param string|null $pattern
     * @param \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult|null $result
     */
    public function __construct(?string $pattern = null, ?\Grafana\Foundation\Dashboardv2beta1\ValueMappingResult $result = null)
    {
        $this->pattern = $pattern ?: "";
        $this->result = $result ?: new \Grafana\Foundation\Dashboardv2beta1\ValueMappingResult();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{pattern?: string, result?: mixed} $inputData */
        $data = $inputData;
        return new self(
            pattern: $data["pattern"] ?? null,
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
        $data->pattern = $this->pattern;
        $data->result = $this->result;
        return $data;
    }
}
