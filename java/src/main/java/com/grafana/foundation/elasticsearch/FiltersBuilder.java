// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class FiltersBuilder implements com.grafana.foundation.cog.Builder<Filters> {
    protected final Filters internal;
    
    public FiltersBuilder() {
        this.internal = new Filters();
    this.internal.type = "filters";
    }
    public FiltersBuilder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public FiltersBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchFiltersSettings> settings) {
    this.internal.settings = settings.build();
        return this;
    }
    public Filters build() {
        return this.internal;
    }
}
