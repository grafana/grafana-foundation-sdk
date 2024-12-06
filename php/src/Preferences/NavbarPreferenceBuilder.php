<?php

namespace Grafana\Foundation\Preferences;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Preferences\NavbarPreference>
 */
class NavbarPreferenceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Preferences\NavbarPreference $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Preferences\NavbarPreference();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Preferences\NavbarPreference
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<string> $bookmarkUrls
     */
    public function bookmarkUrls(array $bookmarkUrls): static
    {
        $this->internal->bookmarkUrls = $bookmarkUrls;
    
        return $this;
    }

}
