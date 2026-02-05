<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class Dashboardv2beta1FieldConfigSourceOverrides implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\MatcherConfig $matcher;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\DynamicConfigValue>
     */
    public array $properties;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\MatcherConfig|null $matcher
     * @param array<\Grafana\Foundation\Dashboardv2beta1\DynamicConfigValue>|null $properties
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\MatcherConfig $matcher = null, ?array $properties = null)
    {
        $this->matcher = $matcher ?: new \Grafana\Foundation\Dashboardv2beta1\MatcherConfig();
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
    	return \Grafana\Foundation\Dashboardv2beta1\MatcherConfig::fromArray($val);
    })($data["matcher"]) : null,
            properties: array_filter(array_map((function($input) {
    	/** @var array{id?: string, value?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\DynamicConfigValue::fromArray($val);
    }), $data["properties"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->matcher = $this->matcher;
        $data->properties = $this->properties;
        return $data;
    }
}
