<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class OptionsWithTooltip implements \JsonSerializable
{
    public \Grafana\Foundation\Common\VizTooltipOptions $tooltip;

    /**
     * @param \Grafana\Foundation\Common\VizTooltipOptions|null $tooltip
     */
    public function __construct(?\Grafana\Foundation\Common\VizTooltipOptions $tooltip = null)
    {
        $this->tooltip = $tooltip ?: new \Grafana\Foundation\Common\VizTooltipOptions();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{tooltip?: mixed} $inputData */
        $data = $inputData;
        return new self(
            tooltip: isset($data["tooltip"]) ? (function($input) {
    	/** @var array{mode?: string, sort?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizTooltipOptions::fromArray($val);
    })($data["tooltip"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "tooltip" => $this->tooltip,
        ];
        return $data;
    }
}
