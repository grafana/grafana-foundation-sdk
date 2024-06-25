<?php

namespace Grafana\Foundation\Canvas;

class BackgroundConfig implements \JsonSerializable
{
    public ?\Grafana\Foundation\Common\ColorDimensionConfig $color;

    public ?\Grafana\Foundation\Common\ResourceDimensionConfig $image;

    public ?\Grafana\Foundation\Canvas\BackgroundImageSize $size;

    /**
     * @param \Grafana\Foundation\Common\ColorDimensionConfig|null $color
     * @param \Grafana\Foundation\Common\ResourceDimensionConfig|null $image
     * @param \Grafana\Foundation\Canvas\BackgroundImageSize|null $size
     */
    public function __construct(?\Grafana\Foundation\Common\ColorDimensionConfig $color = null, ?\Grafana\Foundation\Common\ResourceDimensionConfig $image = null, ?\Grafana\Foundation\Canvas\BackgroundImageSize $size = null)
    {
        $this->color = $color;
        $this->image = $image;
        $this->size = $size;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{color?: mixed, image?: mixed, size?: string} $inputData */
        $data = $inputData;
        return new self(
            color: isset($data["color"]) ? (function($input) {
    	/** @var array{fixed?: string, field?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\ColorDimensionConfig::fromArray($val);
    })($data["color"]) : null,
            image: isset($data["image"]) ? (function($input) {
    	/** @var array{mode?: string, field?: string, fixed?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\ResourceDimensionConfig::fromArray($val);
    })($data["image"]) : null,
            size: isset($data["size"]) ? (function($input) { return \Grafana\Foundation\Canvas\BackgroundImageSize::fromValue($input); })($data["size"]) : null,
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
        if (isset($this->image)) {
            $data["image"] = $this->image;
        }
        if (isset($this->size)) {
            $data["size"] = $this->size;
        }
        return $data;
    }
}
