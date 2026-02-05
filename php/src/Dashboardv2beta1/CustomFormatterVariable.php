<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Custom formatter variable
 */
class CustomFormatterVariable implements \JsonSerializable
{
    public string $name;

    public \Grafana\Foundation\Dashboardv2beta1\VariableType $type;

    public bool $multi;

    public bool $includeAll;

    /**
     * @param string|null $name
     * @param \Grafana\Foundation\Dashboardv2beta1\VariableType|null $type
     * @param bool|null $multi
     * @param bool|null $includeAll
     */
    public function __construct(?string $name = null, ?\Grafana\Foundation\Dashboardv2beta1\VariableType $type = null, ?bool $multi = null, ?bool $includeAll = null)
    {
        $this->name = $name ?: "";
        $this->type = $type ?: \Grafana\Foundation\Dashboardv2beta1\VariableType::Query();
        $this->multi = $multi ?: false;
        $this->includeAll = $includeAll ?: false;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, type?: string, multi?: bool, includeAll?: bool} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\VariableType::fromValue($input); })($data["type"]) : null,
            multi: $data["multi"] ?? null,
            includeAll: $data["includeAll"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->name = $this->name;
        $data->type = $this->type;
        $data->multi = $this->multi;
        $data->includeAll = $this->includeAll;
        return $data;
    }
}
