<?php

namespace Grafana\Foundation\Playlist;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Playlist\PlaylistItem>
 */
class PlaylistItemBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Playlist\PlaylistItem $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Playlist\PlaylistItem();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Playlist\PlaylistItem
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Type of the item.
     */
    public function type(\Grafana\Foundation\Playlist\PlaylistItemType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
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
    public function value(string $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }
    /**
     * Title is an unused property -- it will be removed in the future
     */
    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }

}
