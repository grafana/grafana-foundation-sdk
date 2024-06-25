<?php

namespace Grafana\Foundation\Team;

class Team implements \JsonSerializable
{
    /**
     * Name of the team.
     */
    public string $name;

    /**
     * Email of the team.
     */
    public ?string $email;

    /**
     * @param string|null $name
     * @param string|null $email
     */
    public function __construct(?string $name = null, ?string $email = null)
    {
        $this->name = $name ?: "";
        $this->email = $email;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, email?: string} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            email: $data["email"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "name" => $this->name,
        ];
        if (isset($this->email)) {
            $data["email"] = $this->email;
        }
        return $data;
    }
}
