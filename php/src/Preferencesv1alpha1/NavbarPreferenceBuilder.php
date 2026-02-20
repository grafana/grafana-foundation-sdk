<?php

namespace Grafana\Foundation\Preferencesv1alpha1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Preferencesv1alpha1\NavbarPreference>
 */
class NavbarPreferenceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Preferencesv1alpha1\NavbarPreference $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Preferencesv1alpha1\NavbarPreference();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Preferencesv1alpha1\NavbarPreference
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
