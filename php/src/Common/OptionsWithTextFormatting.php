<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class OptionsWithTextFormatting implements \JsonSerializable
{
    public ?\Grafana\Foundation\Common\VizTextDisplayOptions $text;

    /**
     * @param \Grafana\Foundation\Common\VizTextDisplayOptions|null $text
     */
    public function __construct(?\Grafana\Foundation\Common\VizTextDisplayOptions $text = null)
    {
        $this->text = $text;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{text?: mixed} $inputData */
        $data = $inputData;
        return new self(
            text: isset($data["text"]) ? (function($input) {
    	/** @var array{titleSize?: float, valueSize?: float, percentSize?: float} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizTextDisplayOptions::fromArray($val);
    })($data["text"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->text)) {
            $data["text"] = $this->text;
        }
        return $data;
    }
}
