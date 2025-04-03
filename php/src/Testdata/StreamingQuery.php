<?php

namespace Grafana\Foundation\Testdata;

class StreamingQuery implements \JsonSerializable
{
    public ?int $bands;

    public float $noise;

    public float $speed;

    public float $spread;

    /**
     * Possible enum values:
     *  - `"fetch"` 
     *  - `"logs"` 
     *  - `"signal"` 
     *  - `"traces"` 
     */
    public \Grafana\Foundation\Testdata\StreamingQueryType $type;

    public ?string $url;

    /**
     * @param int|null $bands
     * @param float|null $noise
     * @param float|null $speed
     * @param float|null $spread
     * @param \Grafana\Foundation\Testdata\StreamingQueryType|null $type
     * @param string|null $url
     */
    public function __construct(?int $bands = null, ?float $noise = null, ?float $speed = null, ?float $spread = null, ?\Grafana\Foundation\Testdata\StreamingQueryType $type = null, ?string $url = null)
    {
        $this->bands = $bands;
        $this->noise = $noise ?: 0;
        $this->speed = $speed ?: 0;
        $this->spread = $spread ?: 0;
        $this->type = $type ?: \Grafana\Foundation\Testdata\StreamingQueryType::Fetch();
        $this->url = $url;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{bands?: int, noise?: float, speed?: float, spread?: float, type?: string, url?: string} $inputData */
        $data = $inputData;
        return new self(
            bands: $data["bands"] ?? null,
            noise: $data["noise"] ?? null,
            speed: $data["speed"] ?? null,
            spread: $data["spread"] ?? null,
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Testdata\StreamingQueryType::fromValue($input); })($data["type"]) : null,
            url: $data["url"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "noise" => $this->noise,
            "speed" => $this->speed,
            "spread" => $this->spread,
            "type" => $this->type,
        ];
        if (isset($this->bands)) {
            $data["bands"] = $this->bands;
        }
        if (isset($this->url)) {
            $data["url"] = $this->url;
        }
        return $data;
    }
}
