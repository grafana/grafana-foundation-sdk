<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class StackingConfig implements \JsonSerializable
{
    public ?\Grafana\Foundation\Common\StackingMode $mode;

    public ?string $group;

    /**
     * @param \Grafana\Foundation\Common\StackingMode|null $mode
     * @param string|null $group
     */
    public function __construct(?\Grafana\Foundation\Common\StackingMode $mode = null, ?string $group = null)
    {
        $this->mode = $mode;
        $this->group = $group;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, group?: string} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Common\StackingMode::fromValue($input); })($data["mode"]) : null,
            group: $data["group"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->mode)) {
            $data->mode = $this->mode;
        }
        if (isset($this->group)) {
            $data->group = $this->group;
        }
        return $data;
    }
}
