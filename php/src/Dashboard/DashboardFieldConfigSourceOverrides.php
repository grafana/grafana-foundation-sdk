<?php

namespace Grafana\Foundation\Dashboard;

class DashboardFieldConfigSourceOverrides implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboard\MatcherConfig $matcher;

    /**
     * @var array<\Grafana\Foundation\Dashboard\DynamicConfigValue>
     */
    public array $properties;

    /**
     * @param \Grafana\Foundation\Dashboard\MatcherConfig|null $matcher
     * @param array<\Grafana\Foundation\Dashboard\DynamicConfigValue>|null $properties
     */
    public function __construct(?\Grafana\Foundation\Dashboard\MatcherConfig $matcher = null, ?array $properties = null)
    {
        $this->matcher = $matcher ?: new \Grafana\Foundation\Dashboard\MatcherConfig();
        $this->properties = $properties ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{matcher?: mixed, properties?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            matcher: isset($data["matcher"]) ? (function($input) {
    	/** @var array{id?: string, options?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\MatcherConfig::fromArray($val);
    })($data["matcher"]) : null,
            properties: array_filter(array_map((function($input) {
    	/** @var array{id?: string, value?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DynamicConfigValue::fromArray($val);
    }), $data["properties"] ?? [])),
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "matcher" => $this->matcher,
            "properties" => $this->properties,
        ];
        return $data;
    }
}
