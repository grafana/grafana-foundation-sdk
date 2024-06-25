<?php

namespace Grafana\Foundation\Alertgroups;

class Options implements \JsonSerializable
{
    /**
     * Comma-separated list of values used to filter alert results
     */
    public string $labels;

    /**
     * Name of the alertmanager used as a source for alerts
     */
    public string $alertmanager;

    /**
     * Expand all alert groups by default
     */
    public bool $expandAll;

    /**
     * @param string|null $labels
     * @param string|null $alertmanager
     * @param bool|null $expandAll
     */
    public function __construct(?string $labels = null, ?string $alertmanager = null, ?bool $expandAll = null)
    {
        $this->labels = $labels ?: "";
        $this->alertmanager = $alertmanager ?: "";
        $this->expandAll = $expandAll ?: false;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{labels?: string, alertmanager?: string, expandAll?: bool} $inputData */
        $data = $inputData;
        return new self(
            labels: $data["labels"] ?? null,
            alertmanager: $data["alertmanager"] ?? null,
            expandAll: $data["expandAll"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "labels" => $this->labels,
            "alertmanager" => $this->alertmanager,
            "expandAll" => $this->expandAll,
        ];
        return $data;
    }
}
