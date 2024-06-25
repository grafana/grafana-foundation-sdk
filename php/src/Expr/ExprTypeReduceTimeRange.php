<?php

namespace Grafana\Foundation\Expr;

class ExprTypeReduceTimeRange implements \JsonSerializable
{
    /**
     * From is the start time of the query.
     */
    public string $from;

    /**
     * To is the end time of the query.
     */
    public string $to;

    /**
     * @param string|null $from
     * @param string|null $to
     */
    public function __construct(?string $from = null, ?string $to = null)
    {
        $this->from = $from ?: "now-6h";
        $this->to = $to ?: "now";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{from?: string, to?: string} $inputData */
        $data = $inputData;
        return new self(
            from: $data["from"] ?? null,
            to: $data["to"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "from" => $this->from,
            "to" => $this->to,
        ];
        return $data;
    }
}
