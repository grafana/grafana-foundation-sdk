// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class AppInsightsMetricNameQueryBuilder implements com.grafana.foundation.cog.Builder<AppInsightsMetricNameQuery> {
    protected final AppInsightsMetricNameQuery internal;
    
    public AppInsightsMetricNameQueryBuilder() {
        this.internal = new AppInsightsMetricNameQuery();
        this.internal.kind = "AppInsightsMetricNameQuery";
    }
    public AppInsightsMetricNameQueryBuilder rawQuery(String rawQuery) {
        this.internal.rawQuery = rawQuery;
        return this;
    }
    public AppInsightsMetricNameQuery build() {
        return this.internal;
    }
}
