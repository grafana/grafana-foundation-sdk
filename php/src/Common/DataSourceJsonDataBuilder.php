<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\DataSourceJsonData>
 */
class DataSourceJsonDataBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\DataSourceJsonData $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\DataSourceJsonData();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\DataSourceJsonData
     */
    public function build()
    {
        return $this->internal;
    }

    public function authType(string $authType): static
    {
        $this->internal->authType = $authType;
    
        return $this;
    }
    public function defaultRegion(string $defaultRegion): static
    {
        $this->internal->defaultRegion = $defaultRegion;
    
        return $this;
    }
    public function profile(string $profile): static
    {
        $this->internal->profile = $profile;
    
        return $this;
    }
    public function manageAlerts(bool $manageAlerts): static
    {
        $this->internal->manageAlerts = $manageAlerts;
    
        return $this;
    }
    public function alertmanagerUid(string $alertmanagerUid): static
    {
        $this->internal->alertmanagerUid = $alertmanagerUid;
    
        return $this;
    }

}
