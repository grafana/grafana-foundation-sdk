<?php

namespace Grafana\Foundation\Playlistv0alpha1;

class Playlist implements \JsonSerializable
{
    public string $title;

    public string $interval;

    /**
     * @var array<\Grafana\Foundation\Playlistv0alpha1\Item>
     */
    public array $items;

    /**
     * @param string|null $title
     * @param string|null $interval
     * @param array<\Grafana\Foundation\Playlistv0alpha1\Item>|null $items
     */
    public function __construct(?string $title = null, ?string $interval = null, ?array $items = null)
    {
        $this->title = $title ?: "";
        $this->interval = $interval ?: "";
        $this->items = $items ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{title?: string, interval?: string, items?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            title: $data["title"] ?? null,
            interval: $data["interval"] ?? null,
            items: array_filter(array_map((function($input) {
    	/** @var array{type?: string, value?: string} */
    $val = $input;
    	return \Grafana\Foundation\Playlistv0alpha1\Item::fromArray($val);
    }), $data["items"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->title = $this->title;
        $data->interval = $this->interval;
        $data->items = $this->items;
        return $data;
    }
}
