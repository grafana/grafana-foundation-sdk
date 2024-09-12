<?php

namespace Grafana\Foundation\Canvas;

class CanvasElementOptions implements \JsonSerializable
{
    public string $name;

    public string $type;

    /**
     * TODO: figure out how to define this (element config(s))
     * @var mixed|null
     */
    public $config;

    public ?\Grafana\Foundation\Canvas\Constraint $constraint;

    public ?\Grafana\Foundation\Canvas\Placement $placement;

    public ?\Grafana\Foundation\Canvas\BackgroundConfig $background;

    public ?\Grafana\Foundation\Canvas\LineConfig $border;

    /**
     * @var array<\Grafana\Foundation\Canvas\CanvasConnection>|null
     */
    public ?array $connections;

    /**
     * @param string|null $name
     * @param string|null $type
     * @param mixed|null $config
     * @param \Grafana\Foundation\Canvas\Constraint|null $constraint
     * @param \Grafana\Foundation\Canvas\Placement|null $placement
     * @param \Grafana\Foundation\Canvas\BackgroundConfig|null $background
     * @param \Grafana\Foundation\Canvas\LineConfig|null $border
     * @param array<\Grafana\Foundation\Canvas\CanvasConnection>|null $connections
     */
    public function __construct(?string $name = null, ?string $type = null,  $config = null, ?\Grafana\Foundation\Canvas\Constraint $constraint = null, ?\Grafana\Foundation\Canvas\Placement $placement = null, ?\Grafana\Foundation\Canvas\BackgroundConfig $background = null, ?\Grafana\Foundation\Canvas\LineConfig $border = null, ?array $connections = null)
    {
        $this->name = $name ?: "";
        $this->type = $type ?: "";
        $this->config = $config;
        $this->constraint = $constraint;
        $this->placement = $placement;
        $this->background = $background;
        $this->border = $border;
        $this->connections = $connections;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, type?: string, config?: mixed, constraint?: mixed, placement?: mixed, background?: mixed, border?: mixed, connections?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            type: $data["type"] ?? null,
            config: $data["config"] ?? null,
            constraint: isset($data["constraint"]) ? (function($input) {
    	/** @var array{horizontal?: string, vertical?: string} */
    $val = $input;
    	return \Grafana\Foundation\Canvas\Constraint::fromArray($val);
    })($data["constraint"]) : null,
            placement: isset($data["placement"]) ? (function($input) {
    	/** @var array{top?: float, left?: float, right?: float, bottom?: float, width?: float, height?: float, rotation?: float} */
    $val = $input;
    	return \Grafana\Foundation\Canvas\Placement::fromArray($val);
    })($data["placement"]) : null,
            background: isset($data["background"]) ? (function($input) {
    	/** @var array{color?: mixed, image?: mixed, size?: string} */
    $val = $input;
    	return \Grafana\Foundation\Canvas\BackgroundConfig::fromArray($val);
    })($data["background"]) : null,
            border: isset($data["border"]) ? (function($input) {
    	/** @var array{color?: mixed, width?: float, radius?: float} */
    $val = $input;
    	return \Grafana\Foundation\Canvas\LineConfig::fromArray($val);
    })($data["border"]) : null,
            connections: array_filter(array_map((function($input) {
    	/** @var array{source?: mixed, target?: mixed, targetName?: string, path?: string, color?: mixed, size?: mixed, vertices?: array<mixed>, sourceOriginal?: mixed, targetOriginal?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Canvas\CanvasConnection::fromArray($val);
    }), $data["connections"] ?? [])),
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "name" => $this->name,
            "type" => $this->type,
        ];
        if (isset($this->config)) {
            $data["config"] = $this->config;
        }
        if (isset($this->constraint)) {
            $data["constraint"] = $this->constraint;
        }
        if (isset($this->placement)) {
            $data["placement"] = $this->placement;
        }
        if (isset($this->background)) {
            $data["background"] = $this->background;
        }
        if (isset($this->border)) {
            $data["border"] = $this->border;
        }
        if (isset($this->connections)) {
            $data["connections"] = $this->connections;
        }
        return $data;
    }
}
