<?php

namespace Grafana\Foundation\Resource;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Resource\Manifest>
 */
class ManifestBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Resource\Manifest $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Resource\Manifest();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Resource\Manifest
     */
    public function build()
    {
        return $this->internal;
    }

    public function apiVersion(string $apiVersion): static
    {
        $this->internal->apiVersion = $apiVersion;
    
        return $this;
    }

    public function kind(string $kind): static
    {
        $this->internal->kind = $kind;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Resource\Metadata> $metadata
     */
    public function metadata(\Grafana\Foundation\Cog\Builder $metadata): static
    {
        $metadataResource = $metadata->build();
        $this->internal->metadata = $metadataResource;
    
        return $this;
    }

    /**
     * @param mixed $spec
     */
    public function spec( $spec): static
    {
        $this->internal->spec = $spec;
    
        return $this;
    }

}
