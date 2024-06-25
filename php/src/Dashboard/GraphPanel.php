<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Support for legacy graph panel.
 * @deprecated this a deprecated panel type
 */
class GraphPanel implements \JsonSerializable
{
    public string $type;

    /**
     * @deprecated this is part of deprecated graph panel
     */
    public ?\Grafana\Foundation\Dashboard\DashboardGraphPanelLegend $legend;

    /**
     * @param \Grafana\Foundation\Dashboard\DashboardGraphPanelLegend|null $legend
     */
    public function __construct(?\Grafana\Foundation\Dashboard\DashboardGraphPanelLegend $legend = null)
    {
        $this->type = "graph";
    
        $this->legend = $legend;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, legend?: mixed} $inputData */
        $data = $inputData;
        return new self(
            legend: isset($data["legend"]) ? (function($input) {
    	/** @var array{show?: bool, sort?: string, sortDesc?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DashboardGraphPanelLegend::fromArray($val);
    })($data["legend"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
        ];
        if (isset($this->legend)) {
            $data["legend"] = $this->legend;
        }
        return $data;
    }
}
