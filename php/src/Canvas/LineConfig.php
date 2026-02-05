<?php

namespace Grafana\Foundation\Canvas;

class LineConfig implements \JsonSerializable
{
    public ?\Grafana\Foundation\Common\ColorDimensionConfig $color;

    public ?float $width;

    public ?float $radius;

    /**
     * @param \Grafana\Foundation\Common\ColorDimensionConfig|null $color
     * @param float|null $width
     * @param float|null $radius
     */
    public function __construct(?\Grafana\Foundation\Common\ColorDimensionConfig $color = null, ?float $width = null, ?float $radius = null)
    {
        $this->color = $color;
        $this->width = $width;
        $this->radius = $radius;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{color?: mixed, width?: float, radius?: float} $inputData */
        $data = $inputData;
        return new self(
            color: isset($data["color"]) ? (function($input) {
    	/** @var array{fixed?: string, field?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\ColorDimensionConfig::fromArray($val);
    })($data["color"]) : null,
            width: $data["width"] ?? null,
            radius: $data["radius"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->color)) {
            $data->color = $this->color;
        }
        if (isset($this->width)) {
            $data->width = $this->width;
        }
        if (isset($this->radius)) {
            $data->radius = $this->radius;
        }
        return $data;
    }
}
