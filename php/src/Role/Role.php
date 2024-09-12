<?php

namespace Grafana\Foundation\Role;

class Role implements \JsonSerializable
{
    /**
     * The role identifier `managed:builtins:editor:permissions`
     */
    public string $name;

    /**
     * Optional display
     */
    public ?string $displayName;

    /**
     * Name of the team.
     */
    public ?string $groupName;

    /**
     * Role description
     */
    public ?string $description;

    /**
     * Do not show this role
     */
    public bool $hidden;

    /**
     * @param string|null $name
     * @param string|null $displayName
     * @param string|null $groupName
     * @param string|null $description
     * @param bool|null $hidden
     */
    public function __construct(?string $name = null, ?string $displayName = null, ?string $groupName = null, ?string $description = null, ?bool $hidden = null)
    {
        $this->name = $name ?: "";
        $this->displayName = $displayName;
        $this->groupName = $groupName;
        $this->description = $description;
        $this->hidden = $hidden ?: false;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, displayName?: string, groupName?: string, description?: string, hidden?: bool} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            displayName: $data["displayName"] ?? null,
            groupName: $data["groupName"] ?? null,
            description: $data["description"] ?? null,
            hidden: $data["hidden"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "name" => $this->name,
            "hidden" => $this->hidden,
        ];
        if (isset($this->displayName)) {
            $data["displayName"] = $this->displayName;
        }
        if (isset($this->groupName)) {
            $data["groupName"] = $this->groupName;
        }
        if (isset($this->description)) {
            $data["description"] = $this->description;
        }
        return $data;
    }
}
