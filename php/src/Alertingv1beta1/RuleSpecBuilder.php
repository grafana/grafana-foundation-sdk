<?php

namespace Grafana\Foundation\Alertingv1beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\RuleSpec>
 */
class RuleSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv1beta1\RuleSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv1beta1\RuleSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv1beta1\RuleSpec
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

    public function paused(bool $paused): static
    {
        $this->internal->paused = $paused;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\IntervalTrigger> $trigger
     */
    public function trigger(\Grafana\Foundation\Cog\Builder $trigger): static
    {
        $triggerResource = $trigger->build();
        $this->internal->trigger = $triggerResource;
    
        return $this;
    }

    /**
     * @param array<string, string> $labels
     */
    public function labels(array $labels): static
    {
        $this->internal->labels = $labels;
    
        return $this;
    }

    /**
     * @param array<string, \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\Expression>> $expressions
     */
    public function expressions(array $expressions): static
    {
            $expressionsResources = [];
            foreach ($expressions as $key1 => $val1) {
                    $expressionsResources[$key1] = $val1->build();
            }
        $this->internal->expressions = $expressionsResources;
    
        return $this;
    }

}
