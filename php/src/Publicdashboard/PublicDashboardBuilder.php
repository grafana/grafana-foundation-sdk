<?php

namespace Grafana\Foundation\Publicdashboard;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Publicdashboard\PublicDashboard>
 */
class PublicDashboardBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Publicdashboard\PublicDashboard $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Publicdashboard\PublicDashboard();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Publicdashboard\PublicDashboard
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Unique public dashboard identifier
     */
    public function uid(string $uid): static
    {
        $this->internal->uid = $uid;
    
        return $this;
    }
    /**
     * Dashboard unique identifier referenced by this public dashboard
     */
    public function dashboardUid(string $dashboardUid): static
    {
        $this->internal->dashboardUid = $dashboardUid;
    
        return $this;
    }
    /**
     * Unique public access token
     */
    public function accessToken(string $accessToken): static
    {
        $this->internal->accessToken = $accessToken;
    
        return $this;
    }
    /**
     * Flag that indicates if the public dashboard is enabled
     */
    public function isEnabled(bool $isEnabled): static
    {
        $this->internal->isEnabled = $isEnabled;
    
        return $this;
    }
    /**
     * Flag that indicates if annotations are enabled
     */
    public function annotationsEnabled(bool $annotationsEnabled): static
    {
        $this->internal->annotationsEnabled = $annotationsEnabled;
    
        return $this;
    }
    /**
     * Flag that indicates if the time range picker is enabled
     */
    public function timeSelectionEnabled(bool $timeSelectionEnabled): static
    {
        $this->internal->timeSelectionEnabled = $timeSelectionEnabled;
    
        return $this;
    }

}
