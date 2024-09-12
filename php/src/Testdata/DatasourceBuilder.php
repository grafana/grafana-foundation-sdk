<?php

namespace Grafana\Foundation\Testdata;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Testdata\Datasource>
 */
class DatasourceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Testdata\Datasource $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Testdata\Datasource();
    }

    /**
     * @return \Grafana\Foundation\Testdata\Datasource
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The apiserver version
     */
    public function apiVersion(string $apiVersion): static
    {
        $this->internal->apiVersion = $apiVersion;
    
        return $this;
    }
    /**
     * The datasource plugin type
     */
    public function type(string $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    /**
     * Datasource UID (NOTE: name in k8s)
     */
    public function uid(string $uid): static
    {
        $this->internal->uid = $uid;
    
        return $this;
    }

}
