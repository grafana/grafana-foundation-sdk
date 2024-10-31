<?php

namespace Grafana\Foundation\Librarypanel;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Librarypanel\LibraryElementDTOMeta>
 */
class LibraryElementDTOMetaBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Librarypanel\LibraryElementDTOMeta $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Librarypanel\LibraryElementDTOMeta();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Librarypanel\LibraryElementDTOMeta
     */
    public function build()
    {
        return $this->internal;
    }

    public function folderName(string $folderName): static
    {
        $this->internal->folderName = $folderName;
    
        return $this;
    }
    public function folderUid(string $folderUid): static
    {
        $this->internal->folderUid = $folderUid;
    
        return $this;
    }
    public function connectedDashboards(int $connectedDashboards): static
    {
        $this->internal->connectedDashboards = $connectedDashboards;
    
        return $this;
    }
    public function created(string $created): static
    {
        $this->internal->created = $created;
    
        return $this;
    }
    public function updated(string $updated): static
    {
        $this->internal->updated = $updated;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser> $createdBy
     */
    public function createdBy(\Grafana\Foundation\Cog\Builder $createdBy): static
    {
        $createdByResource = $createdBy->build();
        $this->internal->createdBy = $createdByResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser> $updatedBy
     */
    public function updatedBy(\Grafana\Foundation\Cog\Builder $updatedBy): static
    {
        $updatedByResource = $updatedBy->build();
        $this->internal->updatedBy = $updatedByResource;
    
        return $this;
    }

}
