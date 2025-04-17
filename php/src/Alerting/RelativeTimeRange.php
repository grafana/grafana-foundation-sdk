<?php

namespace Grafana\Foundation\Alerting;

/**
 * RelativeTimeRange is the per query start and end time
 * for requests.
 */
class RelativeTimeRange implements \JsonSerializable
{
    /**
     * A Duration represents the elapsed time between two instants
     * as an int64 nanosecond count. The representation limits the
     * largest representable duration to approximately 290 years.
     */
    public int $from;

    /**
     * A Duration represents the elapsed time between two instants
     * as an int64 nanosecond count. The representation limits the
     * largest representable duration to approximately 290 years.
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
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->from = $this->from;
        $data->to = $this->to;
        return $data;
    }
}
