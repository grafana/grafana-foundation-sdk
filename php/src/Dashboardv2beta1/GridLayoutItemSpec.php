<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class GridLayoutItemSpec implements \JsonSerializable
{
    public int $x;

    public int $y;

    public int $width;

    public int $height;

    /**
     * reference to a PanelKind from dashboard.spec.elements Expressed as JSON Schema reference
     */
    public \Grafana\Foundation\Dashboardv2beta1\ElementReference $element;

    public ?\Grafana\Foundation\Dashboardv2beta1\RepeatOptions $repeat;

    /**
     * @param int|null $x
     * @param int|null $y
     * @param int|null $width
     * @param int|null $height
     * @param \Grafana\Foundation\Dashboardv2beta1\ElementReference|null $element
     * @param \Grafana\Foundation\Dashboardv2beta1\RepeatOptions|null $repeat
     */
    public function __construct(?int $x = null, ?int $y = null, ?int $width = null, ?int $height = null, ?\Grafana\Foundation\Dashboardv2beta1\ElementReference $element = null, ?\Grafana\Foundation\Dashboardv2beta1\RepeatOptions $repeat = null)
    {
        $this->x = $x ?: 0;
        $this->y = $y ?: 0;
        $this->width = $width ?: 0;
        $this->height = $height ?: 0;
        $this->element = $element ?: new \Grafana\Foundation\Dashboardv2beta1\ElementReference();
        $this->repeat = $repeat;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{x?: int, y?: int, width?: int, height?: int, element?: mixed, repeat?: mixed} $inputData */
        $data = $inputData;
        return new self(
            x: $data["x"] ?? null,
            y: $data["y"] ?? null,
            width: $data["width"] ?? null,
            height: $data["height"] ?? null,
            element: isset($data["element"]) ? (function($input) {
    	/** @var array{kind?: string, name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\ElementReference::fromArray($val);
    })($data["element"]) : null,
            repeat: isset($data["repeat"]) ? (function($input) {
    	/** @var array{mode?: "variable", value?: string, direction?: string, maxPerRow?: int} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\RepeatOptions::fromArray($val);
    })($data["repeat"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->x = $this->x;
        $data->y = $this->y;
        $data->width = $this->width;
        $data->height = $this->height;
        $data->element = $this->element;
        if (isset($this->repeat)) {
            $data->repeat = $this->repeat;
        }
        return $data;
    }
}
