<?php

namespace Grafana\Foundation\Alerting;

class MuteTiming implements \JsonSerializable
{
    public ?string $name;

    /**
     * @var array<\Grafana\Foundation\Alerting\TimeInterval>|null
     */
    public ?array $timeIntervals;

    /**
     * @param string|null $name
     * @param array<\Grafana\Foundation\Alerting\TimeInterval>|null $timeIntervals
     */
    public function __construct(?string $name = null, ?array $timeIntervals = null)
    {
        $this->name = $name;
        $this->timeIntervals = $timeIntervals;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, time_intervals?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            timeIntervals: array_filter(array_map((function($input) {
    	/** @var array{times?: array<mixed>, weekdays?: array<mixed>, days_of_month?: array<mixed>, months?: array<mixed>, years?: array<mixed>, location?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alerting\TimeInterval::fromArray($val);
    }), $data["time_intervals"] ?? [])),
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->name)) {
            $data["name"] = $this->name;
        }
        if (isset($this->timeIntervals)) {
            $data["time_intervals"] = $this->timeIntervals;
        }
        return $data;
    }
}
