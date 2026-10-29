<?php

namespace Grafana\Foundation\Alertingv0alpha1;

class SimplifiedRouting implements \JsonSerializable
{
    public \Grafana\Foundation\Alertingv0alpha1\NotificationSettingsType $type;

    public string $receiver;

    /**
     * @var array<string>|null
     */
    public ?array $groupBy;

    public string $groupWait;

    public string $groupInterval;

    public string $repeatInterval;

    /**
     * @var array<string>|null
     */
    public ?array $muteTimeIntervals;

    /**
     * @var array<string>|null
     */
    public ?array $activeTimeIntervals;

    /**
     * @param string|null $receiver
     * @param array<string>|null $groupBy
     * @param string|null $groupWait
     * @param string|null $groupInterval
     * @param string|null $repeatInterval
     * @param array<string>|null $muteTimeIntervals
     * @param array<string>|null $activeTimeIntervals
     */
    public function __construct(?string $receiver = null, ?array $groupBy = null, ?string $groupWait = null, ?string $groupInterval = null, ?string $repeatInterval = null, ?array $muteTimeIntervals = null, ?array $activeTimeIntervals = null)
    {
        $this->type = \Grafana\Foundation\Alertingv0alpha1\NotificationSettingsType::simplifiedRouting();
        $this->receiver = $receiver ?: "";
        $this->groupBy = $groupBy;
        $this->groupWait = $groupWait ?: "";
        $this->groupInterval = $groupInterval ?: "";
        $this->repeatInterval = $repeatInterval ?: "";
        $this->muteTimeIntervals = $muteTimeIntervals;
        $this->activeTimeIntervals = $activeTimeIntervals;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: "SimplifiedRouting", receiver?: string, groupBy?: array<string>, groupWait?: string, groupInterval?: string, repeatInterval?: string, muteTimeIntervals?: array<string>, activeTimeIntervals?: array<string>} $inputData */
        $data = $inputData;
        return new self(
            receiver: $data["receiver"] ?? null,
            groupBy: $data["groupBy"] ?? null,
            groupWait: $data["groupWait"] ?? null,
            groupInterval: $data["groupInterval"] ?? null,
            repeatInterval: $data["repeatInterval"] ?? null,
            muteTimeIntervals: $data["muteTimeIntervals"] ?? null,
            activeTimeIntervals: $data["activeTimeIntervals"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->type = $this->type;
        $data->receiver = $this->receiver;
        $data->groupWait = $this->groupWait;
        $data->groupInterval = $this->groupInterval;
        $data->repeatInterval = $this->repeatInterval;
        if (isset($this->groupBy)) {
            $data->groupBy = $this->groupBy;
        }
        if (isset($this->muteTimeIntervals)) {
            $data->muteTimeIntervals = $this->muteTimeIntervals;
        }
        if (isset($this->activeTimeIntervals)) {
            $data->activeTimeIntervals = $this->activeTimeIntervals;
        }
        return $data;
    }
}
