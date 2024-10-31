<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 */
class DataSourceJsonData implements \JsonSerializable
{
    public ?string $authType;

    public ?string $defaultRegion;

    public ?string $profile;

    public ?bool $manageAlerts;

    public ?string $alertmanagerUid;

    /**
     * @param string|null $authType
     * @param string|null $defaultRegion
     * @param string|null $profile
     * @param bool|null $manageAlerts
     * @param string|null $alertmanagerUid
     */
    public function __construct(?string $authType = null, ?string $defaultRegion = null, ?string $profile = null, ?bool $manageAlerts = null, ?string $alertmanagerUid = null)
    {
        $this->authType = $authType;
        $this->defaultRegion = $defaultRegion;
        $this->profile = $profile;
        $this->manageAlerts = $manageAlerts;
        $this->alertmanagerUid = $alertmanagerUid;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{authType?: string, defaultRegion?: string, profile?: string, manageAlerts?: bool, alertmanagerUid?: string} $inputData */
        $data = $inputData;
        return new self(
            authType: $data["authType"] ?? null,
            defaultRegion: $data["defaultRegion"] ?? null,
            profile: $data["profile"] ?? null,
            manageAlerts: $data["manageAlerts"] ?? null,
            alertmanagerUid: $data["alertmanagerUid"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->authType)) {
            $data["authType"] = $this->authType;
        }
        if (isset($this->defaultRegion)) {
            $data["defaultRegion"] = $this->defaultRegion;
        }
        if (isset($this->profile)) {
            $data["profile"] = $this->profile;
        }
        if (isset($this->manageAlerts)) {
            $data["manageAlerts"] = $this->manageAlerts;
        }
        if (isset($this->alertmanagerUid)) {
            $data["alertmanagerUid"] = $this->alertmanagerUid;
        }
        return $data;
    }
}
