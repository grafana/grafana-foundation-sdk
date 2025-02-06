<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Counterpart for TypeScript's TimeOption type.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\TimeOption>
 */
class TimeOptionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\TimeOption $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\TimeOption();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\TimeOption
     */
    public function build()
    {
        return $this->internal;
    }

    public function display(string $display): static
    {
        $this->internal->display = $display;
    
        return $this;
    }
    public function from(string $from): static
    {
        $this->internal->from = $from;
    
        return $this;
    }
    public function to(string $to): static
    {
        $this->internal->to = $to;
    
        return $this;
    }

}
