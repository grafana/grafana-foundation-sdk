<?php

namespace Grafana\Foundation\Athena;

class ConnectionArgs implements \JsonSerializable
{
    public ?string $region;

    public ?string $catalog;

    public ?string $database;

    public ?bool $resultReuseEnabled;

    public ?float $resultReuseMaxAgeInMinutes;

    /**
     * @param string|null $region
     * @param string|null $catalog
     * @param string|null $database
     * @param bool|null $resultReuseEnabled
     * @param float|null $resultReuseMaxAgeInMinutes
     */
    public function __construct(?string $region = null, ?string $catalog = null, ?string $database = null, ?bool $resultReuseEnabled = null, ?float $resultReuseMaxAgeInMinutes = null)
    {
        $this->region = $region;
        $this->catalog = $catalog;
        $this->database = $database;
        $this->resultReuseEnabled = $resultReuseEnabled;
        $this->resultReuseMaxAgeInMinutes = $resultReuseMaxAgeInMinutes;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{region?: string, catalog?: string, database?: string, resultReuseEnabled?: bool, resultReuseMaxAgeInMinutes?: float} $inputData */
        $data = $inputData;
        return new self(
            region: $data["region"] ?? null,
            catalog: $data["catalog"] ?? null,
            database: $data["database"] ?? null,
            resultReuseEnabled: $data["resultReuseEnabled"] ?? null,
            resultReuseMaxAgeInMinutes: $data["resultReuseMaxAgeInMinutes"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->region)) {
            $data["region"] = $this->region;
        }
        if (isset($this->catalog)) {
            $data["catalog"] = $this->catalog;
        }
        if (isset($this->database)) {
            $data["database"] = $this->database;
        }
        if (isset($this->resultReuseEnabled)) {
            $data["resultReuseEnabled"] = $this->resultReuseEnabled;
        }
        if (isset($this->resultReuseMaxAgeInMinutes)) {
            $data["resultReuseMaxAgeInMinutes"] = $this->resultReuseMaxAgeInMinutes;
        }
        return $data;
    }
}
