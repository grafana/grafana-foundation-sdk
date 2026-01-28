<?php

namespace Grafana\Foundation\Table;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Table\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Table\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Table\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Table\Options
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Represents the index of the selected frame
     */
    public function frameIndex(float $frameIndex): static
    {
        $this->internal->frameIndex = $frameIndex;
    
        return $this;
    }

    /**
     * Controls whether the panel should show the header
     */
    public function showHeader(bool $showHeader): static
    {
        $this->internal->showHeader = $showHeader;
    
        return $this;
    }

    /**
     * Controls whether the header should show icons for the column types
     */
    public function showTypeIcons(bool $showTypeIcons): static
    {
        $this->internal->showTypeIcons = $showTypeIcons;
    
        return $this;
    }

    /**
     * Used to control row sorting
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableSortByFieldState>> $sortBy
     */
    public function sortBy(array $sortBy): static
    {
            $sortByResources = [];
            foreach ($sortBy as $r1) {
                    $sortByResources[] = $r1->build();
            }
        $this->internal->sortBy = $sortByResources;
    
        return $this;
    }

    /**
     * Controls footer options
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\TableFooterOptions> $footer
     */
    public function footer(\Grafana\Foundation\Cog\Builder $footer): static
    {
        $footerResource = $footer->build();
        $this->internal->footer = $footerResource;
    
        return $this;
    }

    /**
     * Controls the height of the rows
     */
    public function cellHeight(\Grafana\Foundation\Common\TableCellHeight $cellHeight): static
    {
        $this->internal->cellHeight = $cellHeight;
    
        return $this;
    }

}
