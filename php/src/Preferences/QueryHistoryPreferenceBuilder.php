<?php

namespace Grafana\Foundation\Preferences;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Preferences\QueryHistoryPreference>
 */
class QueryHistoryPreferenceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Preferences\QueryHistoryPreference $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Preferences\QueryHistoryPreference();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Preferences\QueryHistoryPreference
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * one of: '' | 'query' | 'starred';
     */
    public function homeTab(string $homeTab): static
    {
        $this->internal->homeTab = $homeTab;
    
        return $this;
    }

}
