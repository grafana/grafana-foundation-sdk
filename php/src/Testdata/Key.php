<?php

namespace Grafana\Foundation\Testdata;

class Key implements \JsonSerializable
{
    public string $type;

    public float $tick;

    public ?string $uid;

    /**
     * @param string|null $type
     * @param float|null $tick
     * @param string|null $uid
     */
    public function __construct(?string $type = null, ?float $tick = null, ?string $uid = null)
    {
        $this->type = $type ?: "";
        $this->tick = $tick ?: 0;
        $this->uid = $uid;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, tick?: float, uid?: string} $inputData */
        $data = $inputData;
        return new self(
            type: $data["type"] ?? null,
            tick: $data["tick"] ?? null,
            uid: $data["uid"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
            "tick" => $this->tick,
        ];
        if (isset($this->uid)) {
            $data["uid"] = $this->uid;
        }
        return $data;
    }
}
