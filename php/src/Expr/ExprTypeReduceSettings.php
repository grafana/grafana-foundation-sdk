<?php

namespace Grafana\Foundation\Expr;

class ExprTypeReduceSettings implements \JsonSerializable
{
    /**
     * Non-number reduce behavior
     * Possible enum values:
     *  - `"dropNN"` Drop non-numbers
     *  - `"replaceNN"` Replace non-numbers
     */
    public \Grafana\Foundation\Expr\ExprTypeReduceSettingsMode $mode;

    /**
     * Only valid when mode is replace
     */
    public ?float $replaceWithValue;

    /**
     * @param \Grafana\Foundation\Expr\ExprTypeReduceSettingsMode|null $mode
     * @param float|null $replaceWithValue
     */
    public function __construct(?\Grafana\Foundation\Expr\ExprTypeReduceSettingsMode $mode = null, ?float $replaceWithValue = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Expr\ExprTypeReduceSettingsMode::DropNN();
        $this->replaceWithValue = $replaceWithValue;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, replaceWithValue?: float} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Expr\ExprTypeReduceSettingsMode::fromValue($input); })($data["mode"]) : null,
            replaceWithValue: $data["replaceWithValue"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->mode = $this->mode;
        if (isset($this->replaceWithValue)) {
            $data->replaceWithValue = $this->replaceWithValue;
        }
        return $data;
    }
}
