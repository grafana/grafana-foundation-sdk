<?php

namespace Grafana\Foundation\Dashboard;

/**
 * A dashboard snapshot shares an interactive dashboard publicly.
 * It is a read-only version of a dashboard, and is not editable.
 * It is possible to create a snapshot of a snapshot.
 * Grafana strips away all sensitive information from the dashboard.
 * Sensitive information stripped: queries (metric, template,annotation) and panel links.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\Snapshot>
 */
class SnapshotBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\Snapshot $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\Snapshot();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\Snapshot
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Time when the snapshot expires, default is never to expire
     */
    public function expires(string $expires): static
    {
        $this->internal->expires = $expires;
    
        return $this;
    }
    /**
     * Is the snapshot saved in an external grafana instance
     */
    public function external(bool $external): static
    {
        $this->internal->external = $external;
    
        return $this;
    }
    /**
     * external url, if snapshot was shared in external grafana instance
     */
    public function externalUrl(string $externalUrl): static
    {
        $this->internal->externalUrl = $externalUrl;
    
        return $this;
    }
    /**
     * Unique identifier of the snapshot
     */
    public function id(int $id): static
    {
        $this->internal->id = $id;
    
        return $this;
    }
    /**
     * Optional, defined the unique key of the snapshot, required if external is true
     */
    public function key(string $key): static
    {
        $this->internal->key = $key;
    
        return $this;
    }
    /**
     * Optional, name of the snapshot
     */
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    /**
     * org id of the snapshot
     */
    public function orgId(int $orgId): static
    {
        $this->internal->orgId = $orgId;
    
        return $this;
    }
    /**
     * url of the snapshot, if snapshot was shared internally
     */
    public function url(string $url): static
    {
        $this->internal->url = $url;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\Dashboard> $dashboard
     */
    public function dashboard(\Grafana\Foundation\Cog\Builder $dashboard): static
    {
        $dashboardResource = $dashboard->build();
        $this->internal->dashboard = $dashboardResource;
    
        return $this;
    }

}
