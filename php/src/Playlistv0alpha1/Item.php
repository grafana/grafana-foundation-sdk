<?php

namespace Grafana\Foundation\Playlistv0alpha1;

class Item implements \JsonSerializable
{
    /**
     * type of the item.
     */
    public \Grafana\Foundation\Playlistv0alpha1\ItemType $type;

    /**
     * Value depends on type and describes the playlist item.
     *  - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
     *  is not portable as the numerical identifier is non-deterministic between different instances.
     *  Will be replaced by dashboard_by_uid in the future. (deprecated)
     *  - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
     *  dashboards behind the tag will be added to the playlist.
     *  - dashboard_by_uid: The value is the dashboard UID
     */
    public string $value;

    /**
     * @param \Grafana\Foundation\Playlistv0alpha1\ItemType|null $type
     * @param string|null $value
     */
    public function __construct(?\Grafana\Foundation\Playlistv0alpha1\ItemType $type = null, ?string $value = null)
    {
        $this->type = $type ?: \Grafana\Foundation\Playlistv0alpha1\ItemType::DashboardByTag();
        $this->value = $value ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, value?: string} $inputData */
        $data = $inputData;
        return new self(
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Playlistv0alpha1\ItemType::fromValue($input); })($data["type"]) : null,
            value: $data["value"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->type = $this->type;
        $data->value = $this->value;
        return $data;
    }
}
