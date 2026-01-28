<?php

namespace Grafana\Foundation\Dashboardlist;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardlist\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardlist\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardlist\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardlist\Options
     */
    public function build()
    {
        return $this->internal;
    }

    public function keepTime(bool $keepTime): static
    {
        $this->internal->keepTime = $keepTime;
    
        return $this;
    }

    public function includeVars(bool $includeVars): static
    {
        $this->internal->includeVars = $includeVars;
    
        return $this;
    }

    public function showStarred(bool $showStarred): static
    {
        $this->internal->showStarred = $showStarred;
    
        return $this;
    }

    public function showRecentlyViewed(bool $showRecentlyViewed): static
    {
        $this->internal->showRecentlyViewed = $showRecentlyViewed;
    
        return $this;
    }

    public function showSearch(bool $showSearch): static
    {
        $this->internal->showSearch = $showSearch;
    
        return $this;
    }

    public function showHeadings(bool $showHeadings): static
    {
        $this->internal->showHeadings = $showHeadings;
    
        return $this;
    }

    public function maxItems(int $maxItems): static
    {
        $this->internal->maxItems = $maxItems;
    
        return $this;
    }

    public function query(string $query): static
    {
        $this->internal->query = $query;
    
        return $this;
    }

    /**
     * @param array<string> $tags
     */
    public function tags(array $tags): static
    {
        $this->internal->tags = $tags;
    
        return $this;
    }

    /**
     * folderId is deprecated, and migrated to folderUid on panel init
     */
    public function folderId(int $folderId): static
    {
        $this->internal->folderId = $folderId;
    
        return $this;
    }

    public function folderUID(string $folderUID): static
    {
        $this->internal->folderUID = $folderUID;
    
        return $this;
    }

}
