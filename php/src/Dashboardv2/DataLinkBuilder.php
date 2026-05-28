<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\DataLink>
 */
class DataLinkBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\DataLink $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\DataLink();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\DataLink
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

    public function url(string $url): static
    {
        $this->internal->url = $url;
    
        return $this;
    }

    public function targetBlank(bool $targetBlank): static
    {
        $this->internal->targetBlank = $targetBlank;
    
        return $this;
    }

}
