<?php

namespace Grafana\Foundation\Playlist;

class PlaylistItem implements \JsonSerializable
{
    /**
     * Type of the item.
     */
    public \Grafana\Foundation\Playlist\PlaylistItemType $type;

    /**
     * Value depends on type and describes the playlist item.
     * 
     *  - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
     *  is not portable as the numerical identifier is non-deterministic between different instances.
     *  Will be replaced by dashboard_by_uid in the future. (deprecated)
     *  - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
     *  dashboards behind the tag will be added to the playlist.
     *  - dashboard_by_uid: The value is the dashboard UID
     */
    public string $value;

    /**
     * Title is an unused property -- it will be removed in the future
     */
    public ?string $title;

    /**
     * @param \Grafana\Foundation\Playlist\PlaylistItemType|null $type
     * @param string|null $value
     * @param string|null $title
     */
    public function __construct(?\Grafana\Foundation\Playlist\PlaylistItemType $type = null, ?string $value = null, ?string $title = null)
    {
        $this->type = $type ?: \Grafana\Foundation\Playlist\PlaylistItemType::DashboardByUid();
        $this->value = $value ?: "";
        $this->title = $title;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, value?: string, title?: string} $inputData */
        $data = $inputData;
        return new self(
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Playlist\PlaylistItemType::fromValue($input); })($data["type"]) : null,
            value: $data["value"] ?? null,
            title: $data["title"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "type" => $this->type,
            "value" => $this->value,
        ];
        if (isset($this->title)) {
            $data["title"] = $this->title;
        }
        return $data;
    }
}
