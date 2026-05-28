<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\TransformationKind>
 */
class TransformationBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2\TransformationKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2\TransformationKind();
    $this->internal->kind = "Transformation";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2\TransformationKind
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The group is the transformation ID
     */
    public function group(string $group): static
    {
        $this->internal->group = $group;
    
        return $this;
    }

    /**
     * Disabled transformations are skipped
     */
    public function disabled(bool $disabled): static
    {
        $this->internal->spec->disabled = $disabled;
    
        return $this;
    }

    /**
     * Optional frame matcher. When missing it will be applied to all results
     */
    public function filter(\Grafana\Foundation\Dashboardv2\MatcherConfig $filter): static
    {
        $this->internal->spec->filter = $filter;
    
        return $this;
    }

    /**
     * Where to pull DataFrames from as input to transformation
     */
    public function topic(\Grafana\Foundation\Dashboardv2\DataTopic $topic): static
    {
        $this->internal->spec->topic = $topic;
    
        return $this;
    }

    /**
     * Options to be passed to the transformer
     * Valid options depend on the transformer id
     * @param mixed $options
     */
    public function options( $options): static
    {
        $this->internal->spec->options = $options;
    
        return $this;
    }

}
