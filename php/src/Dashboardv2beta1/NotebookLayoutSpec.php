<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class NotebookLayoutSpec implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemKind>
     */
    public array $cells;

    /**
     * @param array<\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemKind>|null $cells
     */
    public function __construct(?array $cells = null)
    {
        $this->cells = $cells ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{cells?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            cells: array_filter(array_map((function($input) {
    	/** @var array{kind?: string, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemKind::fromArray($val);
    }), $data["cells"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->cells = $this->cells;
        return $data;
    }
}
