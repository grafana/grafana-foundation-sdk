// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class AppInsightsGroupByQueryBuilder implements com.grafana.foundation.cog.Builder<AppInsightsGroupByQuery> {
    protected final AppInsightsGroupByQuery internal;
    
    public AppInsightsGroupByQueryBuilder() {
        this.internal = new AppInsightsGroupByQuery();
    this.internal.kind = "AppInsightsGroupByQuery";
    }
    public AppInsightsGroupByQueryBuilder rawQuery(String rawQuery) {
    this.internal.rawQuery = rawQuery;
        return this;
    }
    
    public AppInsightsGroupByQueryBuilder metricName(String metricName) {
    this.internal.metricName = metricName;
        return this;
    }
    public AppInsightsGroupByQuery build() {
        return this.internal;
    }
}
