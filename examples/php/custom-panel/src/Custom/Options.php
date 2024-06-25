<?php

namespace App\Custom;

class Options implements \JsonSerializable
{
    public bool $makeBeautiful;

    public function __construct(?bool $makeBeautiful = null)
    {
        $this->makeBeautiful = $makeBeautiful ?: false;
    }

    public function jsonSerialize(): array
    {
        return [
            "makeBeautiful" => $this->makeBeautiful,
        ];
    }

    /**
     * @param array{makeBeautiful?: bool} $data
     */
    public static function fromArray(array $data): self
    {
        return new self(
            makeBeautiful: $data["makeBeautiful"] ?? null,
        );
    }
}