<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\OptionsWithTimezones>
 */
class OptionsWithTimezonesBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\OptionsWithTimezones $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\OptionsWithTimezones();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\OptionsWithTimezones
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<string> $timezone
     */
    public function timezone(array $timezone): static
    {
        $this->internal->timezone = $timezone;
    
        return $this;
    }

}
