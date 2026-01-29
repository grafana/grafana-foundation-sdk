<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class ThresholdsConfig implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\ThresholdsMode $mode;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\Threshold>
     */
    public array $steps;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\ThresholdsMode|null $mode
     * @param array<\Grafana\Foundation\Dashboardv2beta1\Threshold>|null $steps
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\ThresholdsMode $mode = null, ?array $steps = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Dashboardv2beta1\ThresholdsMode::Absolute();
        $this->steps = $steps ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, steps?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\ThresholdsMode::fromValue($input); })($data["mode"]) : null,
            steps: array_filter(array_map((function($input) {
    	/** @var array{value?: float, color?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\Threshold::fromArray($val);
    }), $data["steps"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->mode = $this->mode;
        $data->steps = $this->steps;
        return $data;
    }
}
