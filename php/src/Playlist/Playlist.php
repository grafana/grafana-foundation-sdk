<?php

namespace Grafana\Foundation\Playlist;

class Playlist implements \JsonSerializable
{
    /**
     * Unique playlist identifier. Generated on creation, either by the
     * creator of the playlist of by the application.
     */
    public string $uid;

    /**
     * Name of the playlist.
     */
    public string $name;

    /**
     * Interval sets the time between switching views in a playlist.
     * FIXME: Is this based on a standardized format or what options are available? Can datemath be used?
     */
    public string $interval;

    /**
     * The ordered list of items that the playlist will iterate over.
     * FIXME! This should not be optional, but changing it makes the godegen awkward
     * @var array<\Grafana\Foundation\Playlist\PlaylistItem>|null
     */
    public ?array $items;

    /**
     * @param string|null $uid
     * @param string|null $name
     * @param string|null $interval
     * @param array<\Grafana\Foundation\Playlist\PlaylistItem>|null $items
     */
    public function __construct(?string $uid = null, ?string $name = null, ?string $interval = null, ?array $items = null)
    {
        $this->uid = $uid ?: "";
        $this->name = $name ?: "";
        $this->interval = $interval ?: "5m";
        $this->items = $items;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{uid?: string, name?: string, interval?: string, items?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            uid: $data["uid"] ?? null,
            name: $data["name"] ?? null,
            interval: $data["interval"] ?? null,
            items: array_filter(array_map((function($input) {
    	/** @var array{type?: string, value?: string, title?: string} */
    $val = $input;
    	return \Grafana\Foundation\Playlist\PlaylistItem::fromArray($val);
    }), $data["items"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->uid = $this->uid;
        $data->name = $this->name;
        $data->interval = $this->interval;
        if (isset($this->items)) {
            $data->items = $this->items;
        }
        return $data;
    }
}
