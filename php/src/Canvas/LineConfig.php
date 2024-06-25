<?php

namespace Grafana\Foundation\Canvas;

class LineConfig implements \JsonSerializable
{
    public ?\Grafana\Foundation\Common\ColorDimensionConfig $color;

    public ?float $width;

    /**
     * @param \Grafana\Foundation\Common\ColorDimensionConfig|null $color
     * @param float|null $width
     */
    public function __construct(?\Grafana\Foundation\Common\ColorDimensionConfig $color = null, ?float $width = null)
    {
        $this->color = $color;
        $this->width = $width;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{color?: mixed, width?: float} $inputData */
        $data = $inputData;
        return new self(
            color: isset($data["color"]) ? (function($input) {
    	/** @var array{fixed?: string, field?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\ColorDimensionConfig::fromArray($val);
    })($data["color"]) : null,
            width: $data["width"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->color)) {
            $data["color"] = $this->color;
        }
        if (isset($this->width)) {
            $data["width"] = $this->width;
        }
        return $data;
    }
}
