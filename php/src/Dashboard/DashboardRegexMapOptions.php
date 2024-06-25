<?php

namespace Grafana\Foundation\Dashboard;

class DashboardRegexMapOptions implements \JsonSerializable
{
    /**
     * Regular expression to match against
     */
    public string $pattern;

    /**
     * Config to apply when the value matches the regex
     */
    public \Grafana\Foundation\Dashboard\ValueMappingResult $result;

    /**
     * @param string|null $pattern
     * @param \Grafana\Foundation\Dashboard\ValueMappingResult|null $result
     */
    public function __construct(?string $pattern = null, ?\Grafana\Foundation\Dashboard\ValueMappingResult $result = null)
    {
        $this->pattern = $pattern ?: "";
        $this->result = $result ?: new \Grafana\Foundation\Dashboard\ValueMappingResult();
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
            "pattern" => $this->pattern,
            "result" => $this->result,
        ];
        return $data;
    }
}
