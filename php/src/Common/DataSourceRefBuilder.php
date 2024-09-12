<?php

namespace Grafana\Foundation\Common;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\DataSourceRef>
 */
class DataSourceRefBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\DataSourceRef $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\DataSourceRef();
    }

    /**
     * @return \Grafana\Foundation\Common\DataSourceRef
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The plugin type-id
     */
    public function type(string $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    /**
     * Specific datasource instance
     */
    public function uid(string $uid): static
    {
        $this->internal->uid = $uid;
    
        return $this;
    }
    /**
     * Datasource API version
     */
    public function apiVersion(string $apiVersion): static
    {
        $this->internal->apiVersion = $apiVersion;
    
        return $this;
    }

}
