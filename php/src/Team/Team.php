<?php

namespace Grafana\Foundation\Team;

class Team implements \JsonSerializable
{
    public ?string $email;

    public string $name;

    /**
     * @param string|null $email
     * @param string|null $name
     */
    public function __construct(?string $email = null, ?string $name = null)
    {
        $this->email = $email;
        $this->name = $name ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{email?: string, name?: string} $inputData */
        $data = $inputData;
        return new self(
            email: $data["email"] ?? null,
            name: $data["name"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->name = $this->name;
        if (isset($this->email)) {
            $data->email = $this->email;
        }
        return $data;
    }
}
