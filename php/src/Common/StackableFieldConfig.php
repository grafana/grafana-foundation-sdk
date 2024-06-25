<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class StackableFieldConfig implements \JsonSerializable
{
    public ?\Grafana\Foundation\Common\StackingConfig $stacking;

    /**
     * @param \Grafana\Foundation\Common\StackingConfig|null $stacking
     */
    public function __construct(?\Grafana\Foundation\Common\StackingConfig $stacking = null)
    {
        $this->stacking = $stacking;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{stacking?: mixed} $inputData */
        $data = $inputData;
        return new self(
            stacking: isset($data["stacking"]) ? (function($input) {
    	/** @var array{mode?: string, group?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\StackingConfig::fromArray($val);
    })($data["stacking"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->stacking)) {
            $data["stacking"] = $this->stacking;
        }
        return $data;
    }
}
