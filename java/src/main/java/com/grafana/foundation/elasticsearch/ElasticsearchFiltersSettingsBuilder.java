// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import java.util.List;
import java.util.LinkedList;

public class ElasticsearchFiltersSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchFiltersSettings> {
    protected final ElasticsearchFiltersSettings internal;
    
    public ElasticsearchFiltersSettingsBuilder() {
        this.internal = new ElasticsearchFiltersSettings();
    }
    public ElasticsearchFiltersSettingsBuilder filters(List<com.grafana.foundation.cog.Builder<Filter>> filters) {
        List<Filter> filtersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Filter> r1 : filters) {
                Filter filtersDepth1 = r1.build();
                filtersResources.add(filtersDepth1); 
        }
        this.internal.filters = filtersResources;
        return this;
    }
    public ElasticsearchFiltersSettings build() {
        return this.internal;
    }
}
