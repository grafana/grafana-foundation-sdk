// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import java.util.List;

public class FiltersSettingsBuilder implements com.grafana.foundation.cog.Builder<FiltersSettings> {
    protected final FiltersSettings internal;
    
    public FiltersSettingsBuilder() {
        this.internal = new FiltersSettings();
    }
    public FiltersSettingsBuilder filters(com.grafana.foundation.cog.Builder<List<Filter>> filters) {
        this.internal.filters = filters.build();
        return this;
    }
    public FiltersSettings build() {
        return this.internal;
    }
}
