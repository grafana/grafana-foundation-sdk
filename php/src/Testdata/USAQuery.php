<?php

namespace Grafana\Foundation\Testdata;

class USAQuery implements \JsonSerializable
{
    public ?string $mode;

    public ?string $period;

    /**
     * @var array<string>|null
     */
    public ?array $fields;

    /**
     * @var array<string>|null
     */
    public ?array $states;

    /**
     * @param string|null $mode
     * @param string|null $period
     * @param array<string>|null $fields
     * @param array<string>|null $states
     */
    public function __construct(?string $mode = null, ?string $period = null, ?array $fields = null, ?array $states = null)
    {
        $this->mode = $mode;
        $this->period = $period;
        $this->fields = $fields;
        $this->states = $states;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, period?: string, fields?: array<string>, states?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            mode: $data["mode"] ?? null,
            period: $data["period"] ?? null,
            fields: $data["fields"] ?? null,
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
        if (isset($this->mode)) {
            $data["mode"] = $this->mode;
        }
        if (isset($this->period)) {
            $data["period"] = $this->period;
        }
        if (isset($this->fields)) {
            $data["fields"] = $this->fields;
        }
        if (isset($this->states)) {
            $data["states"] = $this->states;
        }
        return $data;
    }
}
