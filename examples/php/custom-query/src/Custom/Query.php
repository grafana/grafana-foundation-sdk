<?php

namespace App\Custom;

use Grafana\Foundation\Cog;

class Query implements \JsonSerializable, Cog\Dataquery
{
    // RefId and Hide are expected on all queries
    public string $refId;
    public ?bool $hide;

    // Expr is specific to the Query type
    public string $expr;

    /**
     * @param string|null $expr
     * @param string|null $refId
     * @param bool|null $hide
     */
    public function __construct(?string $expr = null, ?string $refId = null, ?bool $hide = null)
    {
        $this->expr = $expr ?: "";
        $this->refId = $refId ?: "";
        $this->hide = $hide;
    }

    /**
     * @param array{expr?: string, refId?: string, hide?: bool} $data
     */
    public static function fromArray(array $data): self
    {
        return new self(
            expr: $data["expr"] ?? null,
            refId: $data["refId"] ?? null,
            hide: $data["hide"] ?? null,
        );
    }

    public function jsonSerialize(): array
    {
        $data = [
            "expr" => $this->expr,
            "refId" => $this->refId,
        ];
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        return $data;
    }

    public function dataqueryType(): string
    {
        return "custom";
    }
}