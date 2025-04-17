<?php

namespace Grafana\Foundation\Nodegraph;

class NodeOptions implements \JsonSerializable
{
    /**
     * Unit for the main stat to override what ever is set in the data frame.
     */
    public ?string $mainStatUnit;

    /**
     * Unit for the secondary stat to override what ever is set in the data frame.
     */
    public ?string $secondaryStatUnit;

    /**
     * Define which fields are shown as part of the node arc (colored circle around the node).
     * @var array<\Grafana\Foundation\Nodegraph\ArcOption>|null
     */
    public ?array $arcs;

    /**
     * @param string|null $mainStatUnit
     * @param string|null $secondaryStatUnit
     * @param array<\Grafana\Foundation\Nodegraph\ArcOption>|null $arcs
     */
    public function __construct(?string $mainStatUnit = null, ?string $secondaryStatUnit = null, ?array $arcs = null)
    {
        $this->mainStatUnit = $mainStatUnit;
        $this->secondaryStatUnit = $secondaryStatUnit;
        $this->arcs = $arcs;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mainStatUnit?: string, secondaryStatUnit?: string, arcs?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            mainStatUnit: $data["mainStatUnit"] ?? null,
            secondaryStatUnit: $data["secondaryStatUnit"] ?? null,
            arcs: array_filter(array_map((function($input) {
    	/** @var array{field?: string, color?: string} */
    $val = $input;
    	return \Grafana\Foundation\Nodegraph\ArcOption::fromArray($val);
    }), $data["arcs"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->mainStatUnit)) {
            $data->mainStatUnit = $this->mainStatUnit;
        }
        if (isset($this->secondaryStatUnit)) {
            $data->secondaryStatUnit = $this->secondaryStatUnit;
        }
        if (isset($this->arcs)) {
            $data->arcs = $this->arcs;
        }
        return $data;
    }
}
