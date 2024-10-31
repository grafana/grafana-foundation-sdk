<?php

namespace Grafana\Foundation\Common;

/**
 * Sort by field state
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableSortByFieldState>
 */
class TableSortByFieldStateBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\TableSortByFieldState $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\TableSortByFieldState();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\TableSortByFieldState
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Sets the display name of the field to sort by
     */
    public function displayName(string $displayName): static
    {
        $this->internal->displayName = $displayName;
    
        return $this;
    }
    /**
     * Flag used to indicate descending sort order
     */
    public function desc(bool $desc): static
    {
        $this->internal->desc = $desc;
    
        return $this;
    }

}
