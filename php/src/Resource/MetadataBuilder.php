<?php

namespace Grafana\Foundation\Resource;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Resource\Metadata>
 */
class MetadataBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Resource\Metadata $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Resource\Metadata();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Resource\Metadata
     */
    public function build()
    {
        return $this->internal;
    }

    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }

    public function namespace(string $namespace): static
    {
        $this->internal->namespace = $namespace;
    
        return $this;
    }

    /**
     * @param array<string, string> $labels
     */
    public function labels(array $labels): static
    {
        $this->internal->labels = $labels;
    
        return $this;
    }

    /**
     * @param array<string, string> $annotations
     */
    public function annotations(array $annotations): static
    {
        $this->internal->annotations = $annotations;
    
        return $this;
    }

    public function uid(string $uid): static
    {
        $this->internal->uid = $uid;
    
        return $this;
    }

    public function resourceVersion(string $resourceVersion): static
    {
        $this->internal->resourceVersion = $resourceVersion;
    
        return $this;
    }

    public function generation(int $generation): static
    {
        $this->internal->generation = $generation;
    
        return $this;
    }

    public function creationTimestamp(string $creationTimestamp): static
    {
        $this->internal->creationTimestamp = $creationTimestamp;
    
        return $this;
    }

    public function updateTimestamp(string $updateTimestamp): static
    {
        $this->internal->updateTimestamp = $updateTimestamp;
    
        return $this;
    }

    public function deletionTimestamp(string $deletionTimestamp): static
    {
        $this->internal->deletionTimestamp = $deletionTimestamp;
    
        return $this;
    }

}
