<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DataLink>
 */
class DataLinkBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\DataLink $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\DataLink();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\DataLink
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
