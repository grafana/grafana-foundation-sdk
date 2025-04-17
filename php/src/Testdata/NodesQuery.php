<?php

namespace Grafana\Foundation\Testdata;

class NodesQuery implements \JsonSerializable
{
    public ?\Grafana\Foundation\Testdata\NodesQueryType $type;

    public ?int $count;

    /**
     * @param \Grafana\Foundation\Testdata\NodesQueryType|null $type
     * @param int|null $count
     */
    public function __construct(?\Grafana\Foundation\Testdata\NodesQueryType $type = null, ?int $count = null)
    {
        $this->type = $type;
        $this->count = $count;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, count?: int} $inputData */
        $data = $inputData;
        return new self(
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Testdata\NodesQueryType::fromValue($input); })($data["type"]) : null,
            count: $data["count"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->type)) {
            $data->type = $this->type;
        }
        if (isset($this->count)) {
            $data->count = $this->count;
        }
        return $data;
    }
}
