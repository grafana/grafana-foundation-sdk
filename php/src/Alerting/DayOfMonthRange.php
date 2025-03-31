<?php

namespace Grafana\Foundation\Alerting;

class DayOfMonthRange implements \JsonSerializable
{
    public ?int $begin;

    public ?int $end;

    /**
     * @param int|null $begin
     * @param int|null $end
     */
    public function __construct(?int $begin = null, ?int $end = null)
    {
        $this->begin = $begin;
        $this->end = $end;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{begin?: int, end?: int} $inputData */
        $data = $inputData;
        return new self(
            begin: $data["begin"] ?? null,
            end: $data["end"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->begin)) {
            $data["begin"] = $this->begin;
        }
        if (isset($this->end)) {
            $data["end"] = $this->end;
        }
        return $data;
    }
}
