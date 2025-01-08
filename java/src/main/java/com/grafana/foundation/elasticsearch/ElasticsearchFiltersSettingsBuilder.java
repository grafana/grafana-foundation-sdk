// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import java.util.List;

public class ElasticsearchFiltersSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchFiltersSettings> {
    protected final ElasticsearchFiltersSettings internal;
    
    public ElasticsearchFiltersSettingsBuilder() {
        this.internal = new ElasticsearchFiltersSettings();
    }
    public ElasticsearchFiltersSettingsBuilder filters(com.grafana.foundation.cog.Builder<List<Filter>> filters) {
    this.internal.filters = filters.build();
        return this;
    }
    public ElasticsearchFiltersSettings build() {
        return this.internal;
    }
}
