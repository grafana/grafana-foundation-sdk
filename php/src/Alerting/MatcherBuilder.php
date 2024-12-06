<?php

namespace Grafana\Foundation\Alerting;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alerting\Matcher>
 */
class MatcherBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alerting\Matcher $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alerting\Matcher();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alerting\Matcher
     */
    public function build()
    {
        return $this->internal;
    }

    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    public function type(\Grafana\Foundation\Alerting\MatchType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    public function value(string $value): static
    {
        $this->internal->value = $value;
    
        return $this;
    }

}
