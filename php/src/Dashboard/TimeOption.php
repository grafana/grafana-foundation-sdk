<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Counterpart for TypeScript's TimeOption type.
 */
class TimeOption implements \JsonSerializable
{
    public string $display;

    public string $from;

    public string $to;

    /**
     * @param string|null $display
     * @param string|null $from
     * @param string|null $to
     */
    public function __construct(?string $display = null, ?string $from = null, ?string $to = null)
    {
        $this->display = $display ?: "";
        $this->from = $from ?: "";
        $this->to = $to ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{display?: string, from?: string, to?: string} $inputData */
        $data = $inputData;
        return new self(
            display: $data["display"] ?? null,
            from: $data["from"] ?? null,
            to: $data["to"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->display = $this->display;
        $data->from = $this->from;
        $data->to = $this->to;
        return $data;
    }
}
