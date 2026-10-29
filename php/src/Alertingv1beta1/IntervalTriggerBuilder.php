<?php

namespace Grafana\Foundation\Alertingv1beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv1beta1\IntervalTrigger>
 */
class IntervalTriggerBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv1beta1\IntervalTrigger $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv1beta1\IntervalTrigger();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv1beta1\IntervalTrigger
     */
    public function build()
    {
        return $this->internal;
    }

    public function interval(string $interval): static
    {
        if (!preg_match("/^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?|0)$/", $interval)) {
            throw new \ValueError('$interval must match regex ^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?|0)$');
        }
        if (preg_match("/hmuµn/", $interval)) {
            throw new \ValueError('$interval must not match regex hmuµn');
        }
        $this->internal->interval = $interval;
    
        return $this;
    }

}
