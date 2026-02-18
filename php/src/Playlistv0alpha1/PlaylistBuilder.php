<?php

namespace Grafana\Foundation\Playlistv0alpha1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Playlistv0alpha1\Playlist>
 */
class PlaylistBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Playlistv0alpha1\Playlist $internal;

    public function __construct(string $title)
    {
    	$this->internal = new \Grafana\Foundation\Playlistv0alpha1\Playlist();
    $this->internal->title = $title;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Playlistv0alpha1\Playlist
     */
    public function build()
    {
        return $this->internal;
    }

    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }

    public function interval(string $interval): static
    {
        $this->internal->interval = $interval;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Playlistv0alpha1\Item>> $items
     */
    public function items(array $items): static
    {
            $itemsResources = [];
            foreach ($items as $r1) {
                    $itemsResources[] = $r1->build();
            }
        $this->internal->items = $itemsResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Playlistv0alpha1\Item> $item
     */
    public function item(\Grafana\Foundation\Cog\Builder $item): static
    {
        $itemResource = $item->build();
        $this->internal->items[] = $itemResource;
    
        return $this;
    }

}
