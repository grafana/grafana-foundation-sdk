// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class RangeMapBuilder implements com.grafana.foundation.cog.Builder<RangeMap> {
    protected final RangeMap internal;
    
    public RangeMapBuilder() {
        this.internal = new RangeMap();
    this.internal.type = "range";
    }
    public RangeMapBuilder options(com.grafana.foundation.cog.Builder<DashboardRangeMapOptions> options) {
    this.internal.options = options.build();
        return this;
    }
    public RangeMap build() {
        return this.internal;
    }
}
