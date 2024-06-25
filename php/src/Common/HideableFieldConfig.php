<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class HideableFieldConfig implements \JsonSerializable
{
    public ?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom;

    /**
     * @param \Grafana\Foundation\Common\HideSeriesConfig|null $hideFrom
     */
    public function __construct(?\Grafana\Foundation\Common\HideSeriesConfig $hideFrom = null)
    {
        $this->hideFrom = $hideFrom;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{hideFrom?: mixed} $inputData */
        $data = $inputData;
        return new self(
            hideFrom: isset($data["hideFrom"]) ? (function($input) {
    	/** @var array{tooltip?: bool, legend?: bool, viz?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Common\HideSeriesConfig::fromArray($val);
    })($data["hideFrom"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->hideFrom)) {
            $data["hideFrom"] = $this->hideFrom;
        }
        return $data;
    }
}
