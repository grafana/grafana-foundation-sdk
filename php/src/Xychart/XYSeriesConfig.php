<?php

namespace Grafana\Foundation\Xychart;

class XYSeriesConfig implements \JsonSerializable
{
    public ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigName $name;

    public ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigFrame $frame;

    public ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigX $x;

    public ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigY $y;

    public ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigColor $color;

    public ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigSize $size;

    /**
     * @param \Grafana\Foundation\Xychart\XychartXYSeriesConfigName|null $name
     * @param \Grafana\Foundation\Xychart\XychartXYSeriesConfigFrame|null $frame
     * @param \Grafana\Foundation\Xychart\XychartXYSeriesConfigX|null $x
     * @param \Grafana\Foundation\Xychart\XychartXYSeriesConfigY|null $y
     * @param \Grafana\Foundation\Xychart\XychartXYSeriesConfigColor|null $color
     * @param \Grafana\Foundation\Xychart\XychartXYSeriesConfigSize|null $size
     */
    public function __construct(?\Grafana\Foundation\Xychart\XychartXYSeriesConfigName $name = null, ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigFrame $frame = null, ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigX $x = null, ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigY $y = null, ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigColor $color = null, ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigSize $size = null)
    {
        $this->name = $name;
        $this->frame = $frame;
        $this->x = $x;
        $this->y = $y;
        $this->color = $color;
        $this->size = $size;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: mixed, frame?: mixed, x?: mixed, y?: mixed, color?: mixed, size?: mixed} $inputData */
        $data = $inputData;
        return new self(
            name: isset($data["name"]) ? (function($input) {
    	/** @var array{fixed?: string} */
    $val = $input;
    	return \Grafana\Foundation\Xychart\XychartXYSeriesConfigName::fromArray($val);
    })($data["name"]) : null,
            frame: isset($data["frame"]) ? (function($input) {
    	/** @var array{matcher?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Xychart\XychartXYSeriesConfigFrame::fromArray($val);
    })($data["frame"]) : null,
            x: isset($data["x"]) ? (function($input) {
    	/** @var array{matcher?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Xychart\XychartXYSeriesConfigX::fromArray($val);
    })($data["x"]) : null,
            y: isset($data["y"]) ? (function($input) {
    	/** @var array{matcher?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Xychart\XychartXYSeriesConfigY::fromArray($val);
    })($data["y"]) : null,
            color: isset($data["color"]) ? (function($input) {
    	/** @var array{matcher?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Xychart\XychartXYSeriesConfigColor::fromArray($val);
    })($data["color"]) : null,
            size: isset($data["size"]) ? (function($input) {
    	/** @var array{matcher?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Xychart\XychartXYSeriesConfigSize::fromArray($val);
    })($data["size"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->name)) {
            $data->name = $this->name;
        }
        if (isset($this->frame)) {
            $data->frame = $this->frame;
        }
        if (isset($this->x)) {
            $data->x = $this->x;
        }
        if (isset($this->y)) {
            $data->y = $this->y;
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
