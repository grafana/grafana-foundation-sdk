<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class GraphThresholdsStyleConfig implements \JsonSerializable
{
    public \Grafana\Foundation\Common\GraphTresholdsStyleMode $mode;

    /**
     * @param \Grafana\Foundation\Common\GraphTresholdsStyleMode|null $mode
     */
    public function __construct(?\Grafana\Foundation\Common\GraphTresholdsStyleMode $mode = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Common\GraphTresholdsStyleMode::Off();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Common\GraphTresholdsStyleMode::fromValue($input); })($data["mode"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "mode" => $this->mode,
        ];
        return $data;
    }
}
