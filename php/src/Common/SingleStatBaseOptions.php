<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class SingleStatBaseOptions implements \JsonSerializable
{
    public \Grafana\Foundation\Common\ReduceDataOptions $reduceOptions;

    public ?\Grafana\Foundation\Common\VizTextDisplayOptions $text;

    public \Grafana\Foundation\Common\VizOrientation $orientation;

    /**
     * @param \Grafana\Foundation\Common\ReduceDataOptions|null $reduceOptions
     * @param \Grafana\Foundation\Common\VizTextDisplayOptions|null $text
     * @param \Grafana\Foundation\Common\VizOrientation|null $orientation
     */
    public function __construct(?\Grafana\Foundation\Common\ReduceDataOptions $reduceOptions = null, ?\Grafana\Foundation\Common\VizTextDisplayOptions $text = null, ?\Grafana\Foundation\Common\VizOrientation $orientation = null)
    {
        $this->reduceOptions = $reduceOptions ?: new \Grafana\Foundation\Common\ReduceDataOptions();
        $this->text = $text;
        $this->orientation = $orientation ?: \Grafana\Foundation\Common\VizOrientation::Auto();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{reduceOptions?: mixed, text?: mixed, orientation?: string} $inputData */
        $data = $inputData;
        return new self(
            reduceOptions: isset($data["reduceOptions"]) ? (function($input) {
    	/** @var array{values?: bool, limit?: float, calcs?: array<string>, fields?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\ReduceDataOptions::fromArray($val);
    })($data["reduceOptions"]) : null,
            text: isset($data["text"]) ? (function($input) {
    	/** @var array{titleSize?: float, valueSize?: float, percentSize?: float} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizTextDisplayOptions::fromArray($val);
    })($data["text"]) : null,
            orientation: isset($data["orientation"]) ? (function($input) { return \Grafana\Foundation\Common\VizOrientation::fromValue($input); })($data["orientation"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->reduceOptions = $this->reduceOptions;
        $data->orientation = $this->orientation;
        if (isset($this->text)) {
            $data->text = $this->text;
        }
        return $data;
    }
}
