<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\SimplifiedRouting>
 */
class SimplifiedRoutingBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv0alpha1\SimplifiedRouting $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv0alpha1\SimplifiedRouting();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv0alpha1\SimplifiedRouting
     */
    public function build()
    {
        return $this->internal;
    }

    public function receiver(string $receiver): static
    {
        $this->internal->receiver = $receiver;
    
        return $this;
    }

    /**
     * @param array<string> $groupBy
     */
    public function groupBy(array $groupBy): static
    {
        $this->internal->groupBy = $groupBy;
    
        return $this;
    }

    public function groupWait(string $groupWait): static
    {
        if (!preg_match("/^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?|0)$/", $groupWait)) {
            throw new \ValueError('$groupWait must match regex ^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?|0)$');
        }
        if (preg_match("/hmuµn/", $groupWait)) {
            throw new \ValueError('$groupWait must not match regex hmuµn');
        }
        $this->internal->groupWait = $groupWait;
    
        return $this;
    }

    public function groupInterval(string $groupInterval): static
    {
        if (!preg_match("/^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?|0)$/", $groupInterval)) {
            throw new \ValueError('$groupInterval must match regex ^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?|0)$');
        }
        if (preg_match("/hmuµn/", $groupInterval)) {
            throw new \ValueError('$groupInterval must not match regex hmuµn');
        }
        $this->internal->groupInterval = $groupInterval;
    
        return $this;
    }

    public function repeatInterval(string $repeatInterval): static
    {
        if (!preg_match("/^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?|0)$/", $repeatInterval)) {
            throw new \ValueError('$repeatInterval must match regex ^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?|0)$');
        }
        if (preg_match("/hmuµn/", $repeatInterval)) {
            throw new \ValueError('$repeatInterval must not match regex hmuµn');
        }
        $this->internal->repeatInterval = $repeatInterval;
    
        return $this;
    }

    /**
     * @param array<string> $muteTimeIntervals
     */
    public function muteTimeIntervals(array $muteTimeIntervals): static
    {
        $this->internal->muteTimeIntervals = $muteTimeIntervals;
    
        return $this;
    }

    /**
     * @param array<string> $activeTimeIntervals
     */
    public function activeTimeIntervals(array $activeTimeIntervals): static
    {
        $this->internal->activeTimeIntervals = $activeTimeIntervals;
    
        return $this;
    }

}
