<?php

namespace Grafana\Foundation\Alerting;

class NotificationSettings implements \JsonSerializable
{
    /**
     * @var array<string>|null
     */
    public ?array $groupBy;

    public ?string $groupInterval;

    public ?string $groupWait;

    /**
     * @var array<string>|null
     */
    public ?array $muteTimeIntervals;

    public string $receiver;

    public ?string $repeatInterval;

    /**
     * @param array<string>|null $groupBy
     * @param string|null $groupInterval
     * @param string|null $groupWait
     * @param array<string>|null $muteTimeIntervals
     * @param string|null $receiver
     * @param string|null $repeatInterval
     */
    public function __construct(?array $groupBy = null, ?string $groupInterval = null, ?string $groupWait = null, ?array $muteTimeIntervals = null, ?string $receiver = null, ?string $repeatInterval = null)
    {
        $this->groupBy = $groupBy;
        $this->groupInterval = $groupInterval;
        $this->groupWait = $groupWait;
        $this->muteTimeIntervals = $muteTimeIntervals;
        $this->receiver = $receiver ?: "";
        $this->repeatInterval = $repeatInterval;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{group_by?: array<string>, group_interval?: string, group_wait?: string, mute_time_intervals?: array<string>, receiver?: string, repeat_interval?: string} $inputData */
        $data = $inputData;
        return new self(
            groupBy: $data["group_by"] ?? null,
            groupInterval: $data["group_interval"] ?? null,
            groupWait: $data["group_wait"] ?? null,
            muteTimeIntervals: $data["mute_time_intervals"] ?? null,
            receiver: $data["receiver"] ?? null,
            repeatInterval: $data["repeat_interval"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "receiver" => $this->receiver,
        ];
        if (isset($this->groupBy)) {
            $data["group_by"] = $this->groupBy;
        }
        if (isset($this->groupInterval)) {
            $data["group_interval"] = $this->groupInterval;
        }
        if (isset($this->groupWait)) {
            $data["group_wait"] = $this->groupWait;
        }
        if (isset($this->muteTimeIntervals)) {
            $data["mute_time_intervals"] = $this->muteTimeIntervals;
        }
        if (isset($this->repeatInterval)) {
            $data["repeat_interval"] = $this->repeatInterval;
        }
        return $data;
    }
}
