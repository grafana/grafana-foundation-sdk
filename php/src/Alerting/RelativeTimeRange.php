<?php

namespace Grafana\Foundation\Alerting;

/**
 * RelativeTimeRange is the per query start and end time
 * for requests.
 */
class RelativeTimeRange implements \JsonSerializable
{
    /**
     * RelativeTimeRange is the per query start and end time
     * for requests.
     */
    public int $from;

    /**
     * RelativeTimeRange is the per query start and end time
     * for requests.
     */
    public int $to;

    /**
     * @param int|null $from
     * @param int|null $to
     */
    public function __construct(?int $from = null, ?int $to = null)
    {
        $this->from = $from ?: 0;
        $this->to = $to ?: 0;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{from?: int, to?: int} $inputData */
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
            "from" => $this->from,
            "to" => $this->to,
        ];
        return $data;
    }
}
