<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class GridLayoutSpec implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\GridLayoutItemKind>
     */
    public array $items;

    /**
     * @param array<\Grafana\Foundation\Dashboardv2beta1\GridLayoutItemKind>|null $items
     */
    public function __construct(?array $items = null)
    {
        $this->items = $items ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{items?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            items: array_filter(array_map((function($input) {
    	/** @var array{kind?: string, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\GridLayoutItemKind::fromArray($val);
    }), $data["items"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->items = $this->items;
        return $data;
    }
}
