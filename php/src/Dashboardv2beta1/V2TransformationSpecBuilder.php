<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Dashboard v2 shape: no `id`, it is carried by the parent's `group`.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\V2TransformationSpec>
 */
class V2TransformationSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\V2TransformationSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\V2TransformationSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\V2TransformationSpec
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Disabled transformations are skipped
     */
    public function disabled(bool $disabled): static
    {
        $this->internal->disabled = $disabled;
    
        return $this;
    }

    /**
     * Optional frame matcher. When missing it will be applied to all results
     */
    public function filter(\Grafana\Foundation\Dashboardv2beta1\MatcherConfig $filter): static
    {
        $this->internal->filter = $filter;
    
        return $this;
    }

    /**
     * Where to pull DataFrames from as input to transformation
     */
    public function topic(\Grafana\Foundation\Dashboardv2beta1\DataTopic $topic): static
    {
        $this->internal->topic = $topic;
    
        return $this;
    }

    /**
     * Options to be passed to the transformer
     * Valid options depend on the transformer id
     * @param mixed $options
     */
    public function options( $options): static
    {
        $this->internal->options = $options;
    
        return $this;
    }

}
