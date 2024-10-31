<?php

namespace Grafana\Foundation\Librarypanel;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Librarypanel\LibraryPanel>
 */
class LibraryPanelBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Librarypanel\LibraryPanel $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Librarypanel\LibraryPanel();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Librarypanel\LibraryPanel
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Folder UID
     */
    public function folderUid(string $folderUid): static
    {
        $this->internal->folderUid = $folderUid;
    
        return $this;
    }
    /**
     * Library element UID
     */
    public function uid(string $uid): static
    {
        $this->internal->uid = $uid;
    
        return $this;
    }
    /**
     * Panel name (also saved in the model)
     */
    public function name(string $name): static
    {
        if (!(strlen($name) >= 1)) {
            throw new \ValueError('strlen($name) must be >= 1');
        }
        $this->internal->name = $name;
    
        return $this;
    }
    /**
     * Panel description
     */
    public function description(string $description): static
    {
        $this->internal->description = $description;
    
        return $this;
    }
    /**
     * The panel type (from inside the model)
     */
    public function type(string $type): static
    {
        if (!(strlen($type) >= 1)) {
            throw new \ValueError('strlen($type) must be >= 1');
        }
        $this->internal->type = $type;
    
        return $this;
    }
    /**
     * Dashboard version when this was saved (zero if unknown)
     */
    public function schemaVersion(int $schemaVersion): static
    {
        $this->internal->schemaVersion = $schemaVersion;
    
        return $this;
    }
    /**
     * panel version, incremented each time the dashboard is updated.
     */
    public function version(int $version): static
    {
        $this->internal->version = $version;
    
        return $this;
    }
    /**
     * TODO: should be the same panel schema defined in dashboard
     * Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Librarypanel\LibrarypanelLibraryPanelModel> $model
     */
    public function model(\Grafana\Foundation\Cog\Builder $model): static
    {
        $modelResource = $model->build();
        $this->internal->model = $modelResource;
    
        return $this;
    }
    /**
     * Object storage metadata
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Librarypanel\LibraryElementDTOMeta> $meta
     */
    public function meta(\Grafana\Foundation\Cog\Builder $meta): static
    {
        $metaResource = $meta->build();
        $this->internal->meta = $metaResource;
    
        return $this;
    }

}
