<?php

namespace Grafana\Foundation\Folderv1beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Folderv1beta1\Folder>
 */
class FolderBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Folderv1beta1\Folder $internal;

    public function __construct(string $title)
    {
    	$this->internal = new \Grafana\Foundation\Folderv1beta1\Folder();
    $this->internal->title = $title;
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Folderv1beta1\Folder
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

    public function description(string $description): static
    {
        $this->internal->description = $description;
    
        return $this;
    }

}
