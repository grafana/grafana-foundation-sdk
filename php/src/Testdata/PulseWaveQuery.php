<?php

namespace Grafana\Foundation\Testdata;

class PulseWaveQuery implements \JsonSerializable
{
    public ?int $offCount;

    public ?float $offValue;

    public ?int $onCount;

    public ?float $onValue;

    public ?int $timeStep;

    /**
     * @param int|null $offCount
     * @param float|null $offValue
     * @param int|null $onCount
     * @param float|null $onValue
     * @param int|null $timeStep
     */
    public function __construct(?int $offCount = null, ?float $offValue = null, ?int $onCount = null, ?float $onValue = null, ?int $timeStep = null)
    {
        $this->offCount = $offCount;
        $this->offValue = $offValue;
        $this->onCount = $onCount;
        $this->onValue = $onValue;
        $this->timeStep = $timeStep;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{offCount?: int, offValue?: float, onCount?: int, onValue?: float, timeStep?: int} $inputData */
        $data = $inputData;
        return new self(
            offCount: $data["offCount"] ?? null,
            offValue: $data["offValue"] ?? null,
            onCount: $data["onCount"] ?? null,
            onValue: $data["onValue"] ?? null,
            timeStep: $data["timeStep"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->offCount)) {
            $data["offCount"] = $this->offCount;
        }
        if (isset($this->offValue)) {
            $data["offValue"] = $this->offValue;
        }
        if (isset($this->onCount)) {
            $data["onCount"] = $this->onCount;
        }
        if (isset($this->onValue)) {
            $data["onValue"] = $this->onValue;
        }
        if (isset($this->timeStep)) {
            $data["timeStep"] = $this->timeStep;
        }
        return $data;
    }
}
