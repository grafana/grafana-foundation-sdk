// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import java.util.List;
import java.util.LinkedList;

public class FiltersSettingsBuilder implements com.grafana.foundation.cog.Builder<FiltersSettings> {
    protected final FiltersSettings internal;
    
    public FiltersSettingsBuilder() {
        this.internal = new FiltersSettings();
    }
    public FiltersSettingsBuilder filters(List<com.grafana.foundation.cog.Builder<Filter>> filters) {
        List<Filter> filtersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Filter> r1 : filters) {
                Filter filtersDepth1 = r1.build();
                filtersResources.add(filtersDepth1); 
        }
        this.internal.filters = filtersResources;
        return this;
    }
    public FiltersSettings build() {
        return this.internal;
    }
}
