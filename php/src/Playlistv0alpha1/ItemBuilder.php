<?php

namespace Grafana\Foundation\Playlistv0alpha1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Playlistv0alpha1\Item>
 */
class ItemBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Playlistv0alpha1\Item $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Playlistv0alpha1\Item();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Playlistv0alpha1\Item
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * type of the item.
     */
    public function type(\Grafana\Foundation\Playlistv0alpha1\ItemType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }

    /**
     * Value depends on type and describes the playlist item.
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

}
