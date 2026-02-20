<?php

namespace Grafana\Foundation\Preferencesv1alpha1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Preferencesv1alpha1\QueryHistoryPreference>
 */
class QueryHistoryPreferenceBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Preferencesv1alpha1\QueryHistoryPreference $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Preferencesv1alpha1\QueryHistoryPreference();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Preferencesv1alpha1\QueryHistoryPreference
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
