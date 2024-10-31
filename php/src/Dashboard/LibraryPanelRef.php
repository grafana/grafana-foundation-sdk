<?php

namespace Grafana\Foundation\Dashboard;

/**
 * A library panel is a reusable panel that you can use in any dashboard.
 * When you make a change to a library panel, that change propagates to all instances of where the panel is used.
 * Library panels streamline reuse of panels across multiple dashboards.
 */
class LibraryPanelRef implements \JsonSerializable
{
    /**
     * Library panel name
     */
    public string $name;

    /**
     * Library panel uid
     */
    public string $uid;

    /**
     * @param string|null $name
     * @param string|null $uid
     */
    public function __construct(?string $name = null, ?string $uid = null)
    {
        $this->name = $name ?: "";
        $this->uid = $uid ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, uid?: string} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            uid: $data["uid"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "name" => $this->name,
            "uid" => $this->uid,
        ];
        return $data;
    }
}
