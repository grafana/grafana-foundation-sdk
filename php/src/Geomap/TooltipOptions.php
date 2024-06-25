<?php

namespace Grafana\Foundation\Geomap;

class TooltipOptions implements \JsonSerializable
{
    public \Grafana\Foundation\Geomap\TooltipMode $mode;

    /**
     * @param \Grafana\Foundation\Geomap\TooltipMode|null $mode
     */
    public function __construct(?\Grafana\Foundation\Geomap\TooltipMode $mode = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Geomap\TooltipMode::None();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Geomap\TooltipMode::fromValue($input); })($data["mode"]) : null,
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
