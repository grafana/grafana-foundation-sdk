<?php

namespace Grafana\Foundation\Preferencesv1alpha1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Preferencesv1alpha1\Preferences>
 */
class PreferencesBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Preferencesv1alpha1\Preferences $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Preferencesv1alpha1\Preferences();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Preferencesv1alpha1\Preferences
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * UID for the home dashboard
     */
    public function homeDashboardUID(string $homeDashboardUID): static
    {
        $this->internal->homeDashboardUID = $homeDashboardUID;
    
        return $this;
    }

    /**
     * The timezone selection
     * TODO: this should use the timezone defined in common
     */
    public function timezone(string $timezone): static
    {
        $this->internal->timezone = $timezone;
    
        return $this;
    }

    /**
     * day of the week (sunday, monday, etc)
     */
    public function weekStart(string $weekStart): static
    {
        $this->internal->weekStart = $weekStart;
    
        return $this;
    }

    /**
     * light, dark, empty is default
     */
    public function theme(string $theme): static
    {
        $this->internal->theme = $theme;
    
        return $this;
    }

    /**
     * Selected language (beta)
     */
    public function language(string $language): static
    {
        $this->internal->language = $language;
    
        return $this;
    }

    /**
     * Selected locale (beta)
     */
    public function regionalFormat(string $regionalFormat): static
    {
        $this->internal->regionalFormat = $regionalFormat;
    
        return $this;
    }

    /**
     * Explore query history preferences
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Preferencesv1alpha1\QueryHistoryPreference> $queryHistory
     */
    public function queryHistory(\Grafana\Foundation\Cog\Builder $queryHistory): static
    {
        $queryHistoryResource = $queryHistory->build();
        $this->internal->queryHistory = $queryHistoryResource;
    
        return $this;
    }

    /**
     * Cookie preferences
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Preferencesv1alpha1\CookiePreferences> $cookiePreferences
     */
    public function cookiePreferences(\Grafana\Foundation\Cog\Builder $cookiePreferences): static
    {
        $cookiePreferencesResource = $cookiePreferences->build();
        $this->internal->cookiePreferences = $cookiePreferencesResource;
    
        return $this;
    }

    /**
     * Navigation preferences
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Preferencesv1alpha1\NavbarPreference> $navbar
     */
    public function navbar(\Grafana\Foundation\Cog\Builder $navbar): static
    {
        $navbarResource = $navbar->build();
        $this->internal->navbar = $navbarResource;
    
        return $this;
    }

}
