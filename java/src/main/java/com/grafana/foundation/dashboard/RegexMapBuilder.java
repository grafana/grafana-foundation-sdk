// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class RegexMapBuilder implements com.grafana.foundation.cog.Builder<RegexMap> {
    protected final RegexMap internal;
    
    public RegexMapBuilder() {
        this.internal = new RegexMap();
    this.internal.type = "regex";
    }
    public RegexMapBuilder options(com.grafana.foundation.cog.Builder<DashboardRegexMapOptions> options) {
    this.internal.options = options.build();
        return this;
    }
    public RegexMap build() {
        return this.internal;
    }
}
