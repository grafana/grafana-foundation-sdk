<?php

namespace Grafana\Foundation\Testdata;

class USAQuery implements \JsonSerializable
{
    /**
     * @var array<string>|null
     */
    public ?array $fields;

    public ?string $mode;

    public ?string $period;

    /**
     * @var array<string>|null
     */
    public ?array $states;

    /**
     * @param array<string>|null $fields
     * @param string|null $mode
     * @param string|null $period
     * @param array<string>|null $states
     */
    public function __construct(?array $fields = null, ?string $mode = null, ?string $period = null, ?array $states = null)
    {
        $this->fields = $fields;
        $this->mode = $mode;
        $this->period = $period;
        $this->states = $states;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{fields?: array<string>, mode?: string, period?: string, states?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            fields: $data["fields"] ?? null,
            mode: $data["mode"] ?? null,
            period: $data["period"] ?? null,
            states: $data["states"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->fields)) {
            $data["fields"] = $this->fields;
        }
        if (isset($this->mode)) {
            $data["mode"] = $this->mode;
        }
        if (isset($this->period)) {
            $data["period"] = $this->period;
        }
        if (isset($this->states)) {
            $data["states"] = $this->states;
        }
        return $data;
    }
}
