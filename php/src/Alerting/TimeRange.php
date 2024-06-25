<?php

namespace Grafana\Foundation\Alerting;

/**
 * Redefining this to avoid an import cycle
 */
class TimeRange implements \JsonSerializable
{
    /**
     * Redefining this to avoid an import cycle
     */
    public ?string $from;

    /**
     * Redefining this to avoid an import cycle
     */
    public ?string $to;

    /**
     * @param string|null $from
     * @param string|null $to
     */
    public function __construct(?string $from = null, ?string $to = null)
    {
        $this->from = $from;
        $this->to = $to;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{from?: string, to?: string} $inputData */
        $data = $inputData;
        return new self(
            from: $data["from"] ?? null,
            to: $data["to"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->from)) {
            $data["from"] = $this->from;
        }
        if (isset($this->to)) {
            $data["to"] = $this->to;
        }
        return $data;
    }
}
