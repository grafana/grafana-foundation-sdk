<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * Annotation event field mapping. Defines how to map a data frame field to an annotation event field.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AnnotationEventFieldMapping>
 */
class AnnotationEventFieldMappingBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\AnnotationEventFieldMapping $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\AnnotationEventFieldMapping();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\AnnotationEventFieldMapping
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Source type for the field value
     */
    public function source(string $source): static
    {
        $this->internal->source = $source;
    
        return $this;
    }

    /**
     * Constant value to use when source is "text"
     */
    public function value(string $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }

    /**
     * Regular expression to apply to the field value
     */
    public function regex(string $regex): static
    {
        $this->internal->regex = $regex;
    
        return $this;
    }

}
