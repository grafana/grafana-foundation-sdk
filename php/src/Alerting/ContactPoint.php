<?php

namespace Grafana\Foundation\Alerting;

/**
 * EmbeddedContactPoint is the contact point type that is used
 * by grafanas embedded alertmanager implementation.
 */
class ContactPoint implements \JsonSerializable
{
    /**
     * EmbeddedContactPoint is the contact point type that is used
     * by grafanas embedded alertmanager implementation.
     */
    public ?bool $disableResolveMessage;

    /**
     * EmbeddedContactPoint is the contact point type that is used
     * by grafanas embedded alertmanager implementation.
     */
    public ?string $name;

    /**
     * EmbeddedContactPoint is the contact point type that is used
     * by grafanas embedded alertmanager implementation.
     */
    public ?string $provenance;

    /**
     * EmbeddedContactPoint is the contact point type that is used
     * by grafanas embedded alertmanager implementation.
     * @var mixed
     */
    public $settings;

    /**
     * EmbeddedContactPoint is the contact point type that is used
     * by grafanas embedded alertmanager implementation.
     */
    public \Grafana\Foundation\Alerting\ContactPointType $type;

    /**
     * EmbeddedContactPoint is the contact point type that is used
     * by grafanas embedded alertmanager implementation.
     */
    public ?string $uid;

    /**
     * @param bool|null $disableResolveMessage
     * @param string|null $name
     * @param string|null $provenance
     * @param mixed|null $settings
     * @param \Grafana\Foundation\Alerting\ContactPointType|null $type
     * @param string|null $uid
     */
    public function __construct(?bool $disableResolveMessage = null, ?string $name = null, ?string $provenance = null,  $settings = null, ?\Grafana\Foundation\Alerting\ContactPointType $type = null, ?string $uid = null)
    {
        $this->disableResolveMessage = $disableResolveMessage;
        $this->name = $name;
        $this->provenance = $provenance;
        $this->settings = $settings ?: null;
        $this->type = $type ?: \Grafana\Foundation\Alerting\ContactPointType::Alertmanager();
        $this->uid = $uid;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{disableResolveMessage?: bool, name?: string, provenance?: string, settings?: mixed, type?: string, uid?: string} $inputData */
        $data = $inputData;
        return new self(
            disableResolveMessage: $data["disableResolveMessage"] ?? null,
            name: $data["name"] ?? null,
            provenance: $data["provenance"] ?? null,
            settings: $data["settings"] ?? null,
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Alerting\ContactPointType::fromValue($input); })($data["type"]) : null,
            uid: $data["uid"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "settings" => $this->settings,
            "type" => $this->type,
        ];
        if (isset($this->disableResolveMessage)) {
            $data["disableResolveMessage"] = $this->disableResolveMessage;
        }
        if (isset($this->name)) {
            $data["name"] = $this->name;
        }
        if (isset($this->provenance)) {
            $data["provenance"] = $this->provenance;
        }
        if (isset($this->uid)) {
            $data["uid"] = $this->uid;
        }
        return $data;
    }
}
