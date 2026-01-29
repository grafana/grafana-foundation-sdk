<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class QueryOptionsSpec implements \JsonSerializable
{
    public ?string $timeFrom;

    public ?int $maxDataPoints;

    public ?string $timeShift;

    public ?int $queryCachingTTL;

    public ?string $interval;

    public ?string $cacheTimeout;

    public ?bool $hideTimeOverride;

    public ?string $timeCompare;

    /**
     * @param string|null $timeFrom
     * @param int|null $maxDataPoints
     * @param string|null $timeShift
     * @param int|null $queryCachingTTL
     * @param string|null $interval
     * @param string|null $cacheTimeout
     * @param bool|null $hideTimeOverride
     * @param string|null $timeCompare
     */
    public function __construct(?string $timeFrom = null, ?int $maxDataPoints = null, ?string $timeShift = null, ?int $queryCachingTTL = null, ?string $interval = null, ?string $cacheTimeout = null, ?bool $hideTimeOverride = null, ?string $timeCompare = null)
    {
        $this->timeFrom = $timeFrom;
        $this->maxDataPoints = $maxDataPoints;
        $this->timeShift = $timeShift;
        $this->queryCachingTTL = $queryCachingTTL;
        $this->interval = $interval;
        $this->cacheTimeout = $cacheTimeout;
        $this->hideTimeOverride = $hideTimeOverride;
        $this->timeCompare = $timeCompare;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{timeFrom?: string, maxDataPoints?: int, timeShift?: string, queryCachingTTL?: int, interval?: string, cacheTimeout?: string, hideTimeOverride?: bool, timeCompare?: string} $inputData */
        $data = $inputData;
        return new self(
            timeFrom: $data["timeFrom"] ?? null,
            maxDataPoints: $data["maxDataPoints"] ?? null,
            timeShift: $data["timeShift"] ?? null,
            queryCachingTTL: $data["queryCachingTTL"] ?? null,
            interval: $data["interval"] ?? null,
            cacheTimeout: $data["cacheTimeout"] ?? null,
            hideTimeOverride: $data["hideTimeOverride"] ?? null,
            timeCompare: $data["timeCompare"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->timeFrom)) {
            $data->timeFrom = $this->timeFrom;
        }
        if (isset($this->maxDataPoints)) {
            $data->maxDataPoints = $this->maxDataPoints;
        }
        if (isset($this->timeShift)) {
            $data->timeShift = $this->timeShift;
        }
        if (isset($this->queryCachingTTL)) {
            $data->queryCachingTTL = $this->queryCachingTTL;
        }
        if (isset($this->interval)) {
            $data->interval = $this->interval;
        }
        if (isset($this->cacheTimeout)) {
            $data->cacheTimeout = $this->cacheTimeout;
        }
        if (isset($this->hideTimeOverride)) {
            $data->hideTimeOverride = $this->hideTimeOverride;
        }
        if (isset($this->timeCompare)) {
            $data->timeCompare = $this->timeCompare;
        }
        return $data;
    }
}
