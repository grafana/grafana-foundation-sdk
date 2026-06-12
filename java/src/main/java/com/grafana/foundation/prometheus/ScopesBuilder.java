// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.prometheus;

import java.util.List;
import java.util.LinkedList;

public class ScopesBuilder implements com.grafana.foundation.cog.Builder<Scopes> {
    protected final Scopes internal;
    
    public ScopesBuilder() {
        this.internal = new Scopes();
    }
    public ScopesBuilder defaultPath(List<String> defaultPath) {
        this.internal.defaultPath = defaultPath;
        return this;
    }
    
    public ScopesBuilder filters(List<com.grafana.foundation.cog.Builder<ScopesFilters>> filters) {
        List<ScopesFilters> filtersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<ScopesFilters> r1 : filters) {
                ScopesFilters filtersDepth1 = r1.build();
                filtersResources.add(filtersDepth1); 
        }
        this.internal.filters = filtersResources;
        return this;
    }
    
    public ScopesBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    public Scopes build() {
        return this.internal;
    }
}
