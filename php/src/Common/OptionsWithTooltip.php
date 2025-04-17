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
    	/** @var array{mode?: string, sort?: string, maxWidth?: float, maxHeight?: float} */
    $val = $input;
    	return \Grafana\Foundation\Common\VizTooltipOptions::fromArray($val);
    })($data["tooltip"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->tooltip = $this->tooltip;
        return $data;
    }
}
