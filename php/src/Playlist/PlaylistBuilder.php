<?php

namespace Grafana\Foundation\Playlist;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Playlist\Playlist>
 */
class PlaylistBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Playlist\Playlist $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Playlist\Playlist();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Playlist\Playlist
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Unique playlist identifier. Generated on creation, either by the
     * creator of the playlist of by the application.
     */
    public function uid(string $uid): static
    {
        $this->internal->uid = $uid;
    
        return $this;
    }
    /**
     * Name of the playlist.
     */
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    /**
     * Interval sets the time between switching views in a playlist.
     * FIXME: Is this based on a standardized format or what options are available? Can datemath be used?
     */
    public function interval(string $interval): static
    {
        $this->internal->interval = $interval;
    
        return $this;
    }
    /**
     * The ordered list of items that the playlist will iterate over.
     * FIXME! This should not be optional, but changing it makes the godegen awkward
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Playlist\PlaylistItem>> $items
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

}
