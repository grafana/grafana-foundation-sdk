// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.prometheus;


public class PrometheusDataqueryScopeBuilder implements com.grafana.foundation.cog.Builder<PrometheusDataqueryScope> {
    protected final PrometheusDataqueryScope internal;
    
    public PrometheusDataqueryScopeBuilder() {
        this.internal = new PrometheusDataqueryScope();
    }
    public PrometheusDataqueryScopeBuilder matchers(String matchers) {
        this.internal.matchers = matchers;
        return this;
    }
    public PrometheusDataqueryScope build() {
        return this.internal;
    }
}
