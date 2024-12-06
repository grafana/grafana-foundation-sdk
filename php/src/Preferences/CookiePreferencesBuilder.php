<?php

namespace Grafana\Foundation\Preferences;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Preferences\CookiePreferences>
 */
class CookiePreferencesBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Preferences\CookiePreferences $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Preferences\CookiePreferences();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Preferences\CookiePreferences
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param mixed $analytics
     */
    public function analytics( $analytics): static
    {
        $this->internal->analytics = $analytics;
    
        return $this;
    }
    /**
     * @param mixed $performance
     */
    public function performance( $performance): static
    {
        $this->internal->performance = $performance;
    
        return $this;
    }
    /**
     * @param mixed $functional
     */
    public function functional( $functional): static
    {
        $this->internal->functional = $functional;
    
        return $this;
    }

}
