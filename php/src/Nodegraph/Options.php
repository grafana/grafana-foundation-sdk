<?php

namespace Grafana\Foundation\Nodegraph;

class Options implements \JsonSerializable
{
    public ?\Grafana\Foundation\Nodegraph\NodeOptions $nodes;

    public ?\Grafana\Foundation\Nodegraph\EdgeOptions $edges;

    /**
     * @param \Grafana\Foundation\Nodegraph\NodeOptions|null $nodes
     * @param \Grafana\Foundation\Nodegraph\EdgeOptions|null $edges
     */
    public function __construct(?\Grafana\Foundation\Nodegraph\NodeOptions $nodes = null, ?\Grafana\Foundation\Nodegraph\EdgeOptions $edges = null)
    {
        $this->nodes = $nodes;
        $this->edges = $edges;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{nodes?: mixed, edges?: mixed} $inputData */
        $data = $inputData;
        return new self(
            nodes: isset($data["nodes"]) ? (function($input) {
    	/** @var array{mainStatUnit?: string, secondaryStatUnit?: string, arcs?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Nodegraph\NodeOptions::fromArray($val);
    })($data["nodes"]) : null,
            edges: isset($data["edges"]) ? (function($input) {
    	/** @var array{mainStatUnit?: string, secondaryStatUnit?: string} */
    $val = $input;
    	return \Grafana\Foundation\Nodegraph\EdgeOptions::fromArray($val);
    })($data["edges"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->nodes)) {
            $data->nodes = $this->nodes;
        }
        if (isset($this->edges)) {
            $data->edges = $this->edges;
        }
        return $data;
    }
}
