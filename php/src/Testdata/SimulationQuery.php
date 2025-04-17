<?php

namespace Grafana\Foundation\Testdata;

class SimulationQuery implements \JsonSerializable
{
    /**
     * @var mixed|null
     */
    public $config;

    public \Grafana\Foundation\Testdata\Key $key;

    public ?bool $last;

    public ?bool $stream;

    /**
     * @param mixed|null $config
     * @param \Grafana\Foundation\Testdata\Key|null $key
     * @param bool|null $last
     * @param bool|null $stream
     */
    public function __construct( $config = null, ?\Grafana\Foundation\Testdata\Key $key = null, ?bool $last = null, ?bool $stream = null)
    {
        $this->config = $config;
        $this->key = $key ?: new \Grafana\Foundation\Testdata\Key();
        $this->last = $last;
        $this->stream = $stream;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{config?: mixed, key?: mixed, last?: bool, stream?: bool} $inputData */
        $data = $inputData;
        return new self(
            config: $data["config"] ?? null,
            key: isset($data["key"]) ? (function($input) {
    	/** @var array{tick?: float, type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Testdata\Key::fromArray($val);
    })($data["key"]) : null,
            last: $data["last"] ?? null,
            stream: $data["stream"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->key = $this->key;
        if (isset($this->config)) {
            $data->config = $this->config;
        }
        if (isset($this->last)) {
            $data->last = $this->last;
        }
        if (isset($this->stream)) {
            $data->stream = $this->stream;
        }
        return $data;
    }
}
