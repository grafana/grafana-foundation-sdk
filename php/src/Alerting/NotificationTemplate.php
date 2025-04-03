<?php

namespace Grafana\Foundation\Alerting;

class NotificationTemplate implements \JsonSerializable
{
    public ?string $name;

    public string $provenance;

    public ?string $template;

    public ?string $version;

    /**
     * @param string|null $name
     * @param string|null $provenance
     * @param string|null $template
     * @param string|null $version
     */
    public function __construct(?string $name = null, ?string $provenance = null, ?string $template = null, ?string $version = null)
    {
        $this->name = $name;
        $this->provenance = $provenance ?: "";
        $this->template = $template;
        $this->version = $version;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, provenance?: string, template?: string, version?: string} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            provenance: $data["provenance"] ?? null,
            template: $data["template"] ?? null,
            version: $data["version"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "provenance" => $this->provenance,
        ];
        if (isset($this->name)) {
            $data["name"] = $this->name;
        }
        if (isset($this->template)) {
            $data["template"] = $this->template;
        }
        if (isset($this->version)) {
            $data["version"] = $this->version;
        }
        return $data;
    }
}
