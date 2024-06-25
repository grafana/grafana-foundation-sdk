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
    public \Grafana\Foundation\Expr\TypeReduceMode $mode;

    /**
     * Only valid when mode is replace
     */
    public ?float $replaceWithValue;

    /**
     * @param \Grafana\Foundation\Expr\TypeReduceMode|null $mode
     * @param float|null $replaceWithValue
     */
    public function __construct(?\Grafana\Foundation\Expr\TypeReduceMode $mode = null, ?float $replaceWithValue = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Expr\TypeReduceMode::DropNN();
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
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Expr\TypeReduceMode::fromValue($input); })($data["mode"]) : null,
            replaceWithValue: $data["replaceWithValue"] ?? null,
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
        if (isset($this->replaceWithValue)) {
            $data["replaceWithValue"] = $this->replaceWithValue;
        }
        return $data;
    }
}
