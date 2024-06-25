<?php

namespace Grafana\Foundation\Testdata;

class CSVWave implements \JsonSerializable
{
    public ?string $labels;

    public ?string $name;

    public ?int $timeStep;

    public ?string $valuesCSV;

    /**
     * @param string|null $labels
     * @param string|null $name
     * @param int|null $timeStep
     * @param string|null $valuesCSV
     */
    public function __construct(?string $labels = null, ?string $name = null, ?int $timeStep = null, ?string $valuesCSV = null)
    {
        $this->labels = $labels;
        $this->name = $name;
        $this->timeStep = $timeStep;
        $this->valuesCSV = $valuesCSV;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{labels?: string, name?: string, timeStep?: int, valuesCSV?: string} $inputData */
        $data = $inputData;
        return new self(
            labels: $data["labels"] ?? null,
            name: $data["name"] ?? null,
            timeStep: $data["timeStep"] ?? null,
            valuesCSV: $data["valuesCSV"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->labels)) {
            $data["labels"] = $this->labels;
        }
        if (isset($this->name)) {
            $data["name"] = $this->name;
        }
        if (isset($this->timeStep)) {
            $data["timeStep"] = $this->timeStep;
        }
        if (isset($this->valuesCSV)) {
            $data["valuesCSV"] = $this->valuesCSV;
        }
        return $data;
    }
}
