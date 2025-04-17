<?php

namespace Grafana\Foundation\Alerting;

class Matcher implements \JsonSerializable
{
    public ?string $name;

    public ?\Grafana\Foundation\Alerting\MatchType $type;

    public ?string $value;

    /**
     * @param string|null $name
     * @param \Grafana\Foundation\Alerting\MatchType|null $type
     * @param string|null $value
     */
    public function __construct(?string $name = null, ?\Grafana\Foundation\Alerting\MatchType $type = null, ?string $value = null)
    {
        $this->name = $name;
        $this->type = $type;
        $this->value = $value;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{Name?: string, Type?: string, Value?: string} $inputData */
        $data = $inputData;
        return new self(
            name: $data["Name"] ?? null,
            type: isset($data["Type"]) ? (function($input) { return \Grafana\Foundation\Alerting\MatchType::fromValue($input); })($data["Type"]) : null,
            value: $data["Value"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->name)) {
            $data->Name = $this->name;
        }
        if (isset($this->type)) {
            $data->Type = $this->type;
        }
        if (isset($this->value)) {
            $data->Value = $this->value;
        }
        return $data;
    }
}
