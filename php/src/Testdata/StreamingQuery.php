<?php

namespace Grafana\Foundation\Testdata;

class StreamingQuery implements \JsonSerializable
{
    public \Grafana\Foundation\Testdata\StreamingQueryType $type;

    public int $speed;

    public int $spread;

    public int $noise;

    public ?int $bands;

    public ?string $url;

    /**
     * @param \Grafana\Foundation\Testdata\StreamingQueryType|null $type
     * @param int|null $speed
     * @param int|null $spread
     * @param int|null $noise
     * @param int|null $bands
     * @param string|null $url
     */
    public function __construct(?\Grafana\Foundation\Testdata\StreamingQueryType $type = null, ?int $speed = null, ?int $spread = null, ?int $noise = null, ?int $bands = null, ?string $url = null)
    {
        $this->type = $type ?: \Grafana\Foundation\Testdata\StreamingQueryType::Signal();
        $this->speed = $speed ?: 0;
        $this->spread = $spread ?: 0;
        $this->noise = $noise ?: 0;
        $this->bands = $bands;
        $this->url = $url;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, speed?: int, spread?: int, noise?: int, bands?: int, url?: string} $inputData */
        $data = $inputData;
        return new self(
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Testdata\StreamingQueryType::fromValue($input); })($data["type"]) : null,
            speed: $data["speed"] ?? null,
            spread: $data["spread"] ?? null,
            noise: $data["noise"] ?? null,
            bands: $data["bands"] ?? null,
            url: $data["url"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->type = $this->type;
        $data->speed = $this->speed;
        $data->spread = $this->spread;
        $data->noise = $this->noise;
        if (isset($this->bands)) {
            $data->bands = $this->bands;
        }
        if (isset($this->url)) {
            $data->url = $this->url;
        }
        return $data;
    }
}
