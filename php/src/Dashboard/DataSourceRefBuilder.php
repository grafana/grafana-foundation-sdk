<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Ref to a DataSource instance
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DataSourceRef>
 */
class DataSourceRefBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\DataSourceRef $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\DataSourceRef();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\DataSourceRef
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

}
