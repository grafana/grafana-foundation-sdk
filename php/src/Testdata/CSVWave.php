<?php

namespace Grafana\Foundation\Testdata;

class CSVWave implements \JsonSerializable
{
    public ?int $timeStep;

    public ?string $name;

    public ?string $valuesCSV;

    public ?string $labels;

    /**
     * @param int|null $timeStep
     * @param string|null $name
     * @param string|null $valuesCSV
     * @param string|null $labels
     */
    public function __construct(?int $timeStep = null, ?string $name = null, ?string $valuesCSV = null, ?string $labels = null)
    {
        $this->timeStep = $timeStep;
        $this->name = $name;
        $this->valuesCSV = $valuesCSV;
        $this->labels = $labels;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{timeStep?: int, name?: string, valuesCSV?: string, labels?: string} $inputData */
        $data = $inputData;
        return new self(
            timeStep: $data["timeStep"] ?? null,
            name: $data["name"] ?? null,
            valuesCSV: $data["valuesCSV"] ?? null,
            labels: $data["labels"] ?? null,
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
        if (isset($this->name)) {
            $data->name = $this->name;
        }
        if (isset($this->valuesCSV)) {
            $data->valuesCSV = $this->valuesCSV;
        }
        if (isset($this->labels)) {
            $data->labels = $this->labels;
        }
        return $data;
    }
}
