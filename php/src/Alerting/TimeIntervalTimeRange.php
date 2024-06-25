<?php

namespace Grafana\Foundation\Alerting;

class TimeIntervalTimeRange implements \JsonSerializable
{
    public ?string $endTime;

    public ?string $startTime;

    /**
     * @param string|null $endTime
     * @param string|null $startTime
     */
    public function __construct(?string $endTime = null, ?string $startTime = null)
    {
        $this->endTime = $endTime;
        $this->startTime = $startTime;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{end_time?: string, start_time?: string} $inputData */
        $data = $inputData;
        return new self(
            endTime: $data["end_time"] ?? null,
            startTime: $data["start_time"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->endTime)) {
            $data["end_time"] = $this->endTime;
        }
        if (isset($this->startTime)) {
            $data["start_time"] = $this->startTime;
        }
        return $data;
    }
}
