<?php

namespace Grafana\Foundation\Dashboard;

class DashboardFieldConfigSourceOverrides implements \JsonSerializable
{
    public ?string $systemRef;

    public \Grafana\Foundation\Dashboard\MatcherConfig $matcher;

    /**
     * @var array<\Grafana\Foundation\Dashboard\DynamicConfigValue>
     */
    public array $properties;

    /**
     * @param string|null $systemRef
     * @param \Grafana\Foundation\Dashboard\MatcherConfig|null $matcher
     * @param array<\Grafana\Foundation\Dashboard\DynamicConfigValue>|null $properties
     */
    public function __construct(?string $systemRef = null, ?\Grafana\Foundation\Dashboard\MatcherConfig $matcher = null, ?array $properties = null)
    {
        $this->systemRef = $systemRef;
        $this->matcher = $matcher ?: new \Grafana\Foundation\Dashboard\MatcherConfig();
        $this->properties = $properties ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{__systemRef?: string, matcher?: mixed, properties?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            systemRef: $data["__systemRef"] ?? null,
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
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->matcher = $this->matcher;
        $data->properties = $this->properties;
        if (isset($this->systemRef)) {
            $data->__systemRef = $this->systemRef;
        }
        return $data;
    }
}
