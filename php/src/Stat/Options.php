<?php

namespace Grafana\Foundation\Stat;

class Options implements \JsonSerializable
{
    public \Grafana\Foundation\Common\BigValueGraphMode $graphMode;

    public \Grafana\Foundation\Common\BigValueColorMode $colorMode;

    public \Grafana\Foundation\Common\BigValueJustifyMode $justifyMode;

    public \Grafana\Foundation\Common\ReduceDataOptions $reduceOptions;

    public ?\Grafana\Foundation\Common\VizTextDisplayOptions $text;

    public \Grafana\Foundation\Common\BigValueTextMode $textMode;

    public \Grafana\Foundation\Common\VizOrientation $orientation;

    /**
     * @param \Grafana\Foundation\Common\BigValueGraphMode|null $graphMode
     * @param \Grafana\Foundation\Common\BigValueColorMode|null $colorMode
     * @param \Grafana\Foundation\Common\BigValueJustifyMode|null $justifyMode
     * @param \Grafana\Foundation\Common\ReduceDataOptions|null $reduceOptions
     * @param \Grafana\Foundation\Common\VizTextDisplayOptions|null $text
     * @param \Grafana\Foundation\Common\BigValueTextMode|null $textMode
     * @param \Grafana\Foundation\Common\VizOrientation|null $orientation
     */
    public function __construct(?\Grafana\Foundation\Common\BigValueGraphMode $graphMode = null, ?\Grafana\Foundation\Common\BigValueColorMode $colorMode = null, ?\Grafana\Foundation\Common\BigValueJustifyMode $justifyMode = null, ?\Grafana\Foundation\Common\ReduceDataOptions $reduceOptions = null, ?\Grafana\Foundation\Common\VizTextDisplayOptions $text = null, ?\Grafana\Foundation\Common\BigValueTextMode $textMode = null, ?\Grafana\Foundation\Common\VizOrientation $orientation = null)
    {
        $this->graphMode = $graphMode ?: \Grafana\Foundation\Common\BigValueGraphMode::area();
        $this->colorMode = $colorMode ?: \Grafana\Foundation\Common\BigValueColorMode::value();
        $this->justifyMode = $justifyMode ?: \Grafana\Foundation\Common\BigValueJustifyMode::auto();
        $this->reduceOptions = $reduceOptions ?: new \Grafana\Foundation\Common\ReduceDataOptions();
        $this->text = $text;
        $this->textMode = $textMode ?: \Grafana\Foundation\Common\BigValueTextMode::auto();
        $this->orientation = $orientation ?: \Grafana\Foundation\Common\VizOrientation::Auto();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{graphMode?: string, colorMode?: string, justifyMode?: string, reduceOptions?: mixed, text?: mixed, textMode?: string, orientation?: string} $inputData */
        $data = $inputData;
        return new self(
            graphMode: isset($data["graphMode"]) ? (function($input) { return \Grafana\Foundation\Common\BigValueGraphMode::fromValue($input); })($data["graphMode"]) : null,
            colorMode: isset($data["colorMode"]) ? (function($input) { return \Grafana\Foundation\Common\BigValueColorMode::fromValue($input); })($data["colorMode"]) : null,
            justifyMode: isset($data["justifyMode"]) ? (function($input) { return \Grafana\Foundation\Common\BigValueJustifyMode::fromValue($input); })($data["justifyMode"]) : null,
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
            textMode: isset($data["textMode"]) ? (function($input) { return \Grafana\Foundation\Common\BigValueTextMode::fromValue($input); })($data["textMode"]) : null,
            orientation: isset($data["orientation"]) ? (function($input) { return \Grafana\Foundation\Common\VizOrientation::fromValue($input); })($data["orientation"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->graphMode = $this->graphMode;
        $data->colorMode = $this->colorMode;
        $data->justifyMode = $this->justifyMode;
        $data->reduceOptions = $this->reduceOptions;
        $data->textMode = $this->textMode;
        $data->orientation = $this->orientation;
        if (isset($this->text)) {
            $data->text = $this->text;
        }
        return $data;
    }
}
