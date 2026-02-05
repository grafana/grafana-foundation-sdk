<?php

namespace Grafana\Foundation\Debug;

class Options implements \JsonSerializable
{
    public \Grafana\Foundation\Debug\DebugMode $mode;

    public ?\Grafana\Foundation\Debug\UpdateConfig $counters;

    /**
     * @param \Grafana\Foundation\Debug\DebugMode|null $mode
     * @param \Grafana\Foundation\Debug\UpdateConfig|null $counters
     */
    public function __construct(?\Grafana\Foundation\Debug\DebugMode $mode = null, ?\Grafana\Foundation\Debug\UpdateConfig $counters = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Debug\DebugMode::Render();
        $this->counters = $counters;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, counters?: mixed} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Debug\DebugMode::fromValue($input); })($data["mode"]) : null,
            counters: isset($data["counters"]) ? (function($input) {
    	/** @var array{render?: bool, dataChanged?: bool, schemaChanged?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Debug\UpdateConfig::fromArray($val);
    })($data["counters"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->mode = $this->mode;
        if (isset($this->counters)) {
            $data->counters = $this->counters;
        }
        return $data;
    }
}
