<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\RuleRef>
 */
class RuleRefBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv0alpha1\RuleRef $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv0alpha1\RuleRef();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv0alpha1\RuleRef
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * name is the metadata.name of an AlertRule or RecordingRule resource.
     */
    public function name(string $name): static
    {
        if (!preg_match("/^[a-zA-Z0-9_-]+$/", $name)) {
            throw new \ValueError('$name must match regex ^[a-zA-Z0-9_-]+$');
        }
        $this->internal->name = $name;
    
        return $this;
    }

}
