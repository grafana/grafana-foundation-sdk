<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class ScaleDistributionConfig implements \JsonSerializable
{
    public \Grafana\Foundation\Common\ScaleDistribution $type;

    public ?float $log;

    public ?float $linearThreshold;

    /**
     * @param \Grafana\Foundation\Common\ScaleDistribution|null $type
     * @param float|null $log
     * @param float|null $linearThreshold
     */
    public function __construct(?\Grafana\Foundation\Common\ScaleDistribution $type = null, ?float $log = null, ?float $linearThreshold = null)
    {
        $this->type = $type ?: \Grafana\Foundation\Common\ScaleDistribution::Linear();
        $this->log = $log;
        $this->linearThreshold = $linearThreshold;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, log?: float, linearThreshold?: float} $inputData */
        $data = $inputData;
        return new self(
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Common\ScaleDistribution::fromValue($input); })($data["type"]) : null,
            log: $data["log"] ?? null,
            linearThreshold: $data["linearThreshold"] ?? null,
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
        if (isset($this->log)) {
            $data["log"] = $this->log;
        }
        if (isset($this->linearThreshold)) {
            $data["linearThreshold"] = $this->linearThreshold;
        }
        return $data;
    }
}
