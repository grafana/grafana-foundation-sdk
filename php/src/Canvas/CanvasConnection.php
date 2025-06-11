<?php

namespace Grafana\Foundation\Canvas;

class CanvasConnection implements \JsonSerializable
{
    public \Grafana\Foundation\Canvas\ConnectionCoordinates $source;

    public \Grafana\Foundation\Canvas\ConnectionCoordinates $target;

    public ?string $targetName;

    public \Grafana\Foundation\Canvas\ConnectionPath $path;

    public ?\Grafana\Foundation\Common\ColorDimensionConfig $color;

    public ?\Grafana\Foundation\Common\ScaleDimensionConfig $size;

    /**
     * @param \Grafana\Foundation\Canvas\ConnectionCoordinates|null $source
     * @param \Grafana\Foundation\Canvas\ConnectionCoordinates|null $target
     * @param string|null $targetName
     * @param \Grafana\Foundation\Common\ColorDimensionConfig|null $color
     * @param \Grafana\Foundation\Common\ScaleDimensionConfig|null $size
     */
    public function __construct(?\Grafana\Foundation\Canvas\ConnectionCoordinates $source = null, ?\Grafana\Foundation\Canvas\ConnectionCoordinates $target = null, ?string $targetName = null, ?\Grafana\Foundation\Common\ColorDimensionConfig $color = null, ?\Grafana\Foundation\Common\ScaleDimensionConfig $size = null)
    {
        $this->source = $source ?: new \Grafana\Foundation\Canvas\ConnectionCoordinates();
        $this->target = $target ?: new \Grafana\Foundation\Canvas\ConnectionCoordinates();
        $this->targetName = $targetName;
        $this->path = \Grafana\Foundation\Canvas\ConnectionPath::straight();
        $this->color = $color;
        $this->size = $size;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{source?: mixed, target?: mixed, targetName?: string, path?: "straight", color?: mixed, size?: mixed} $inputData */
        $data = $inputData;
        return new self(
            source: isset($data["source"]) ? (function($input) {
    	/** @var array{x?: float, y?: float} */
    $val = $input;
    	return \Grafana\Foundation\Canvas\ConnectionCoordinates::fromArray($val);
    })($data["source"]) : null,
            target: isset($data["target"]) ? (function($input) {
    	/** @var array{x?: float, y?: float} */
    $val = $input;
    	return \Grafana\Foundation\Canvas\ConnectionCoordinates::fromArray($val);
    })($data["target"]) : null,
            targetName: $data["targetName"] ?? null,
            color: isset($data["color"]) ? (function($input) {
    	/** @var array{fixed?: string, field?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\ColorDimensionConfig::fromArray($val);
    })($data["color"]) : null,
            size: isset($data["size"]) ? (function($input) {
    	/** @var array{min?: float, max?: float, fixed?: float, field?: string, mode?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\ScaleDimensionConfig::fromArray($val);
    })($data["size"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->source = $this->source;
        $data->target = $this->target;
        $data->path = $this->path;
        if (isset($this->targetName)) {
            $data->targetName = $this->targetName;
        }
        if (isset($this->color)) {
            $data->color = $this->color;
        }
        if (isset($this->size)) {
            $data->size = $this->size;
        }
        return $data;
    }
}
