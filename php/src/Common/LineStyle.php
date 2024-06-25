<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class LineStyle implements \JsonSerializable
{
    public ?\Grafana\Foundation\Common\LineStyleFill $fill;

    /**
     * @var array<float>|null
     */
    public ?array $dash;

    /**
     * @param \Grafana\Foundation\Common\LineStyleFill|null $fill
     * @param array<float>|null $dash
     */
    public function __construct(?\Grafana\Foundation\Common\LineStyleFill $fill = null, ?array $dash = null)
    {
        $this->fill = $fill;
        $this->dash = $dash;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{fill?: string, dash?: array<float>} $inputData */
        $data = $inputData;
        return new self(
            fill: isset($data["fill"]) ? (function($input) { return \Grafana\Foundation\Common\LineStyleFill::fromValue($input); })($data["fill"]) : null,
            dash: $data["dash"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->fill)) {
            $data["fill"] = $this->fill;
        }
        if (isset($this->dash)) {
            $data["dash"] = $this->dash;
        }
        return $data;
    }
}
