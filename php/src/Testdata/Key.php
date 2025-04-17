<?php

namespace Grafana\Foundation\Testdata;

class Key implements \JsonSerializable
{
    public float $tick;

    public string $type;

    public ?string $uid;

    /**
     * @param float|null $tick
     * @param string|null $type
     * @param string|null $uid
     */
    public function __construct(?float $tick = null, ?string $type = null, ?string $uid = null)
    {
        $this->tick = $tick ?: 0;
        $this->type = $type ?: "";
        $this->uid = $uid;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{tick?: float, type?: string, uid?: string} $inputData */
        $data = $inputData;
        return new self(
            tick: $data["tick"] ?? null,
            type: $data["type"] ?? null,
            uid: $data["uid"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->tick = $this->tick;
        $data->type = $this->type;
        if (isset($this->uid)) {
            $data->uid = $this->uid;
        }
        return $data;
    }
}
