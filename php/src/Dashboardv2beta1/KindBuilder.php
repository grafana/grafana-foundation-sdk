<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * --- Common types ---
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Kind>
 */
class KindBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\Kind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\Kind();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\Kind
     */
    public function build()
    {
        return $this->internal;
    }

    public function kind(string $kind): static
    {
        $this->internal->kind = $kind;
    
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

    /**
     * @param mixed $metadata
     */
    public function metadata( $metadata): static
    {
        $this->internal->metadata = $metadata;
    
        return $this;
    }

}
