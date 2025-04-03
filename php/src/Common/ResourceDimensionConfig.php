<?php

namespace Grafana\Foundation\Common;

/**
 * Links to a resource (image/svg path)
 */
class ResourceDimensionConfig implements \JsonSerializable
{
    public \Grafana\Foundation\Common\ResourceDimensionMode $mode;

    /**
     * fixed: T -- will be added by each element
     */
    public ?string $field;

    public ?string $fixed;

    /**
     * @param \Grafana\Foundation\Common\ResourceDimensionMode|null $mode
     * @param string|null $field
     * @param string|null $fixed
     */
    public function __construct(?\Grafana\Foundation\Common\ResourceDimensionMode $mode = null, ?string $field = null, ?string $fixed = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Common\ResourceDimensionMode::Fixed();
        $this->field = $field;
        $this->fixed = $fixed;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, field?: string, fixed?: string} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Common\ResourceDimensionMode::fromValue($input); })($data["mode"]) : null,
            field: $data["field"] ?? null,
            fixed: $data["fixed"] ?? null,
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
        if (isset($this->field)) {
            $data["field"] = $this->field;
        }
        if (isset($this->fixed)) {
            $data["fixed"] = $this->fixed;
        }
        return $data;
    }
}
