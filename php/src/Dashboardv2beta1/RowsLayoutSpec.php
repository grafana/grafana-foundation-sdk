<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class RowsLayoutSpec implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowKind>
     */
    public array $rows;

    /**
     * @param array<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowKind>|null $rows
     */
    public function __construct(?array $rows = null)
    {
        $this->rows = $rows ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{rows?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            rows: array_filter(array_map((function($input) {
    	/** @var array{kind?: string, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowKind::fromArray($val);
    }), $data["rows"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->rows = $this->rows;
        return $data;
    }
}
