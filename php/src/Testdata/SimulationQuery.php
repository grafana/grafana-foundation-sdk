<?php

namespace Grafana\Foundation\Testdata;

class SimulationQuery implements \JsonSerializable
{
    public \Grafana\Foundation\Testdata\Key $key;

    /**
     * @var array<string, mixed>|null
     */
    public ?array $config;

    public ?bool $stream;

    public ?bool $last;

    /**
     * @param \Grafana\Foundation\Testdata\Key|null $key
     * @param array<string, mixed>|null $config
     * @param bool|null $stream
     * @param bool|null $last
     */
    public function __construct(?\Grafana\Foundation\Testdata\Key $key = null, ?array $config = null, ?bool $stream = null, ?bool $last = null)
    {
        $this->key = $key ?: new \Grafana\Foundation\Testdata\Key();
        $this->config = $config;
        $this->stream = $stream;
        $this->last = $last;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{key?: mixed, config?: array<string, mixed>, stream?: bool, last?: bool} $inputData */
        $data = $inputData;
        return new self(
            key: isset($data["key"]) ? (function($input) {
    	/** @var array{type?: string, tick?: float, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Testdata\Key::fromArray($val);
    })($data["key"]) : null,
            config: $data["config"] ?? null,
            stream: $data["stream"] ?? null,
            last: $data["last"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "key" => $this->key,
        ];
        if (isset($this->config)) {
            $data["config"] = $this->config;
        }
        if (isset($this->stream)) {
            $data["stream"] = $this->stream;
        }
        if (isset($this->last)) {
            $data["last"] = $this->last;
        }
        return $data;
    }
}
