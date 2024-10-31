<?php

namespace Grafana\Foundation\Folder;

/**
 * TODO:
 * common metadata will soon support setting the parent folder in the metadata
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Folder\Folder>
 */
class FolderBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Folder\Folder $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Folder\Folder();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Folder\Folder
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Unique folder id. (will be k8s name)
     */
    public function uid(string $uid): static
    {
        $this->internal->uid = $uid;
    
        return $this;
    }
    /**
     * Folder title
     */
    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }
    /**
     * Description of the folder.
     */
    public function description(string $description): static
    {
        $this->internal->description = $description;
    
        return $this;
    }

}
