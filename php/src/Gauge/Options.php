<?php

namespace Grafana\Foundation\Gauge;

class Options implements \JsonSerializable
{
    public bool $showThresholdLabels;

    public \Grafana\Foundation\Common\ReduceDataOptions $reduceOptions;

    public ?\Grafana\Foundation\Common\VizTextDisplayOptions $text;

    public bool $showThresholdMarkers;

    public \Grafana\Foundation\Common\VizOrientation $orientation;

    /**
     * @param bool|null $showThresholdLabels
     * @param \Grafana\Foundation\Common\ReduceDataOptions|null $reduceOptions
     * @param \Grafana\Foundation\Common\VizTextDisplayOptions|null $text
     * @param bool|null $showThresholdMarkers
     * @param \Grafana\Foundation\Common\VizOrientation|null $orientation
     */
    public function __construct(?bool $showThresholdLabels = null, ?\Grafana\Foundation\Common\ReduceDataOptions $reduceOptions = null, ?\Grafana\Foundation\Common\VizTextDisplayOptions $text = null, ?bool $showThresholdMarkers = null, ?\Grafana\Foundation\Common\VizOrientation $orientation = null)
    {
        $this->showThresholdLabels = $showThresholdLabels ?: false;
        $this->reduceOptions = $reduceOptions ?: new \Grafana\Foundation\Common\ReduceDataOptions();
        $this->text = $text;
        $this->showThresholdMarkers = $showThresholdMarkers ?: true;
        $this->orientation = $orientation ?: \Grafana\Foundation\Common\VizOrientation::Auto();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{showThresholdLabels?: bool, reduceOptions?: mixed, text?: mixed, showThresholdMarkers?: bool, orientation?: string} $inputData */
        $data = $inputData;
        return new self(
            showThresholdLabels: $data["showThresholdLabels"] ?? null,
            reduceOptions: isset($data["reduceOptions"]) ? (function($input) {
    	/** @var array{values?: bool, limit?: float, calcs?: array<string>, fields?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\ReduceDataOptions::fromArray($val);
    })($data["reduceOptions"]) : null,
            text: isset($data["text"]) ? (function($input) {
    	/** @var array{titleSize?: float, valueSize?: float} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizTextDisplayOptions::fromArray($val);
    })($data["text"]) : null,
            showThresholdMarkers: $data["showThresholdMarkers"] ?? null,
            orientation: isset($data["orientation"]) ? (function($input) { return \Grafana\Foundation\Common\VizOrientation::fromValue($input); })($data["orientation"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->showThresholdLabels = $this->showThresholdLabels;
        $data->reduceOptions = $this->reduceOptions;
        $data->showThresholdMarkers = $this->showThresholdMarkers;
        $data->orientation = $this->orientation;
        if (isset($this->text)) {
            $data->text = $this->text;
        }
        return $data;
    }
}
