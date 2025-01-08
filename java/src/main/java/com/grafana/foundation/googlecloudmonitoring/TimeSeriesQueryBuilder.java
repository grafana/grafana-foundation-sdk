// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;


public class TimeSeriesQueryBuilder implements com.grafana.foundation.cog.Builder<TimeSeriesQuery> {
    protected final TimeSeriesQuery internal;
    
    public TimeSeriesQueryBuilder() {
        this.internal = new TimeSeriesQuery();
    }
    public TimeSeriesQueryBuilder projectName(String projectName) {
    this.internal.projectName = projectName;
        return this;
    }
    
    public TimeSeriesQueryBuilder query(String query) {
    this.internal.query = query;
        return this;
    }
    
    public TimeSeriesQueryBuilder graphPeriod(String graphPeriod) {
    this.internal.graphPeriod = graphPeriod;
        return this;
    }
    public TimeSeriesQuery build() {
        return this.internal;
    }
}
