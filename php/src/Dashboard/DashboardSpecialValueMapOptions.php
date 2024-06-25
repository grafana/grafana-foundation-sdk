<?php

namespace Grafana\Foundation\Dashboard;

class DashboardSpecialValueMapOptions implements \JsonSerializable
{
    /**
     * Special value to match against
     */
    public \Grafana\Foundation\Dashboard\SpecialValueMatch $match;

    /**
     * Config to apply when the value matches the special value
     */
    public \Grafana\Foundation\Dashboard\ValueMappingResult $result;

    /**
     * @param \Grafana\Foundation\Dashboard\SpecialValueMatch|null $match
     * @param \Grafana\Foundation\Dashboard\ValueMappingResult|null $result
     */
    public function __construct(?\Grafana\Foundation\Dashboard\SpecialValueMatch $match = null, ?\Grafana\Foundation\Dashboard\ValueMappingResult $result = null)
    {
        $this->match = $match ?: \Grafana\Foundation\Dashboard\SpecialValueMatch::True();
        $this->result = $result ?: new \Grafana\Foundation\Dashboard\ValueMappingResult();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{match?: string, result?: mixed} $inputData */
        $data = $inputData;
        return new self(
            match: isset($data["match"]) ? (function($input) { return \Grafana\Foundation\Dashboard\SpecialValueMatch::fromValue($input); })($data["match"]) : null,
            result: isset($data["result"]) ? (function($input) {
    	/** @var array{text?: string, color?: string, icon?: string, index?: int} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\ValueMappingResult::fromArray($val);
    })($data["result"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "match" => $this->match,
            "result" => $this->result,
        ];
        return $data;
    }
}
