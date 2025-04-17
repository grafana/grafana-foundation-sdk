<?php

namespace Grafana\Foundation\Nodegraph;

class Options implements \JsonSerializable
{
    public ?\Grafana\Foundation\Nodegraph\NodeOptions $nodes;

    public ?\Grafana\Foundation\Nodegraph\EdgeOptions $edges;

    /**
     * How to handle zoom/scroll events in the node graph
     */
    public ?\Grafana\Foundation\Nodegraph\ZoomMode $zoomMode;

    /**
     * How to layout the nodes in the node graph
     */
    public ?\Grafana\Foundation\Nodegraph\LayoutAlgorithm $layoutAlgorithm;

    /**
     * @param \Grafana\Foundation\Nodegraph\NodeOptions|null $nodes
     * @param \Grafana\Foundation\Nodegraph\EdgeOptions|null $edges
     * @param \Grafana\Foundation\Nodegraph\ZoomMode|null $zoomMode
     * @param \Grafana\Foundation\Nodegraph\LayoutAlgorithm|null $layoutAlgorithm
     */
    public function __construct(?\Grafana\Foundation\Nodegraph\NodeOptions $nodes = null, ?\Grafana\Foundation\Nodegraph\EdgeOptions $edges = null, ?\Grafana\Foundation\Nodegraph\ZoomMode $zoomMode = null, ?\Grafana\Foundation\Nodegraph\LayoutAlgorithm $layoutAlgorithm = null)
    {
        $this->nodes = $nodes;
        $this->edges = $edges;
        $this->zoomMode = $zoomMode;
        $this->layoutAlgorithm = $layoutAlgorithm;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{nodes?: mixed, edges?: mixed, zoomMode?: string, layoutAlgorithm?: string} $inputData */
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
            zoomMode: isset($data["zoomMode"]) ? (function($input) { return \Grafana\Foundation\Nodegraph\ZoomMode::fromValue($input); })($data["zoomMode"]) : null,
            layoutAlgorithm: isset($data["layoutAlgorithm"]) ? (function($input) { return \Grafana\Foundation\Nodegraph\LayoutAlgorithm::fromValue($input); })($data["layoutAlgorithm"]) : null,
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
        if (isset($this->zoomMode)) {
            $data->zoomMode = $this->zoomMode;
        }
        if (isset($this->layoutAlgorithm)) {
            $data->layoutAlgorithm = $this->layoutAlgorithm;
        }
        return $data;
    }
}
