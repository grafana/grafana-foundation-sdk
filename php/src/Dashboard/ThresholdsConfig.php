<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Thresholds configuration for the panel
 */
class ThresholdsConfig implements \JsonSerializable
{
    /**
     * Thresholds mode.
     */
    public \Grafana\Foundation\Dashboard\ThresholdsMode $mode;

    /**
     * Must be sorted by 'value', first value is always -Infinity
     * @var array<\Grafana\Foundation\Dashboard\Threshold>
     */
    public array $steps;

    /**
     * @param \Grafana\Foundation\Dashboard\ThresholdsMode|null $mode
     * @param array<\Grafana\Foundation\Dashboard\Threshold>|null $steps
     */
    public function __construct(?\Grafana\Foundation\Dashboard\ThresholdsMode $mode = null, ?array $steps = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Dashboard\ThresholdsMode::Absolute();
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
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Dashboard\ThresholdsMode::fromValue($input); })($data["mode"]) : null,
            steps: array_filter(array_map((function($input) {
    	/** @var array{value?: float, color?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\Threshold::fromArray($val);
    }), $data["steps"] ?? [])),
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "mode" => $this->mode,
            "steps" => $this->steps,
        ];
        return $data;
    }
}
