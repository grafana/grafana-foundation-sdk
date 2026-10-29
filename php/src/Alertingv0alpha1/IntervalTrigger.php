<?php

namespace Grafana\Foundation\Alertingv0alpha1;

class IntervalTrigger implements \JsonSerializable
{
    public string $interval;

    /**
     * @param string|null $interval
     */
    public function __construct(?string $interval = null)
    {
        $this->interval = $interval ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{interval?: string} $inputData */
        $data = $inputData;
        return new self(
            interval: $data["interval"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->interval = $this->interval;
        return $data;
    }
}
