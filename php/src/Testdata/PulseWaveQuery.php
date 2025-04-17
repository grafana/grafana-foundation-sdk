<?php

namespace Grafana\Foundation\Testdata;

class PulseWaveQuery implements \JsonSerializable
{
    public ?int $timeStep;

    public ?int $onCount;

    public ?int $offCount;

    public ?float $onValue;

    public ?float $offValue;

    /**
     * @param int|null $timeStep
     * @param int|null $onCount
     * @param int|null $offCount
     * @param float|null $onValue
     * @param float|null $offValue
     */
    public function __construct(?int $timeStep = null, ?int $onCount = null, ?int $offCount = null, ?float $onValue = null, ?float $offValue = null)
    {
        $this->timeStep = $timeStep;
        $this->onCount = $onCount;
        $this->offCount = $offCount;
        $this->onValue = $onValue;
        $this->offValue = $offValue;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{timeStep?: int, onCount?: int, offCount?: int, onValue?: float, offValue?: float} $inputData */
        $data = $inputData;
        return new self(
            timeStep: $data["timeStep"] ?? null,
            onCount: $data["onCount"] ?? null,
            offCount: $data["offCount"] ?? null,
            onValue: $data["onValue"] ?? null,
            offValue: $data["offValue"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->timeStep)) {
            $data->timeStep = $this->timeStep;
        }
        if (isset($this->onCount)) {
            $data->onCount = $this->onCount;
        }
        if (isset($this->offCount)) {
            $data->offCount = $this->offCount;
        }
        if (isset($this->onValue)) {
            $data->onValue = $this->onValue;
        }
        if (isset($this->offValue)) {
            $data->offValue = $this->offValue;
        }
        return $data;
    }
}
