<?php

namespace Grafana\Foundation\Publicdashboard;

class PublicDashboard implements \JsonSerializable
{
    /**
     * Unique public dashboard identifier
     */
    public string $uid;

    /**
     * Dashboard unique identifier referenced by this public dashboard
     */
    public string $dashboardUid;

    /**
     * Unique public access token
     */
    public ?string $accessToken;

    /**
     * Flag that indicates if the public dashboard is enabled
     */
    public bool $isEnabled;

    /**
     * Flag that indicates if annotations are enabled
     */
    public bool $annotationsEnabled;

    /**
     * Flag that indicates if the time range picker is enabled
     */
    public bool $timeSelectionEnabled;

    /**
     * @param string|null $uid
     * @param string|null $dashboardUid
     * @param string|null $accessToken
     * @param bool|null $isEnabled
     * @param bool|null $annotationsEnabled
     * @param bool|null $timeSelectionEnabled
     */
    public function __construct(?string $uid = null, ?string $dashboardUid = null, ?string $accessToken = null, ?bool $isEnabled = null, ?bool $annotationsEnabled = null, ?bool $timeSelectionEnabled = null)
    {
        $this->uid = $uid ?: "";
        $this->dashboardUid = $dashboardUid ?: "";
        $this->accessToken = $accessToken;
        $this->isEnabled = $isEnabled ?: false;
        $this->annotationsEnabled = $annotationsEnabled ?: false;
        $this->timeSelectionEnabled = $timeSelectionEnabled ?: false;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{uid?: string, dashboardUid?: string, accessToken?: string, isEnabled?: bool, annotationsEnabled?: bool, timeSelectionEnabled?: bool} $inputData */
        $data = $inputData;
        return new self(
            uid: $data["uid"] ?? null,
            dashboardUid: $data["dashboardUid"] ?? null,
            accessToken: $data["accessToken"] ?? null,
            isEnabled: $data["isEnabled"] ?? null,
            annotationsEnabled: $data["annotationsEnabled"] ?? null,
            timeSelectionEnabled: $data["timeSelectionEnabled"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "uid" => $this->uid,
            "dashboardUid" => $this->dashboardUid,
            "isEnabled" => $this->isEnabled,
            "annotationsEnabled" => $this->annotationsEnabled,
            "timeSelectionEnabled" => $this->timeSelectionEnabled,
        ];
        if (isset($this->accessToken)) {
            $data["accessToken"] = $this->accessToken;
        }
        return $data;
    }
}
