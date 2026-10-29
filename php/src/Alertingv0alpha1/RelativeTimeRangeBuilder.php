<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\RelativeTimeRange>
 */
class RelativeTimeRangeBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertingv0alpha1\RelativeTimeRange $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertingv0alpha1\RelativeTimeRange();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertingv0alpha1\RelativeTimeRange
     */
    public function build()
    {
        return $this->internal;
    }

    public function from(string $from): static
    {
        if (!preg_match("/^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)$/", $from)) {
            throw new \ValueError('$from must match regex ^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)$');
        }
        $this->internal->from = $from;
    
        return $this;
    }

    public function to(string $to): static
    {
        if (!preg_match("/^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)$/", $to)) {
            throw new \ValueError('$to must match regex ^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)$');
        }
        $this->internal->to = $to;
    
        return $this;
    }

}
