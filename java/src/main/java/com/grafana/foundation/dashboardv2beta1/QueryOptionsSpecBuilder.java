// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class QueryOptionsSpecBuilder implements com.grafana.foundation.cog.Builder<QueryOptionsSpec> {
    protected final QueryOptionsSpec internal;
    
    public QueryOptionsSpecBuilder() {
        this.internal = new QueryOptionsSpec();
    }
    public QueryOptionsSpecBuilder timeFrom(String timeFrom) {
        this.internal.timeFrom = timeFrom;
        return this;
    }
    
    public QueryOptionsSpecBuilder maxDataPoints(Long maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public QueryOptionsSpecBuilder timeShift(String timeShift) {
        this.internal.timeShift = timeShift;
        return this;
    }
    
    public QueryOptionsSpecBuilder queryCachingTTL(Long queryCachingTTL) {
        this.internal.queryCachingTTL = queryCachingTTL;
        return this;
    }
    
    public QueryOptionsSpecBuilder interval(String interval) {
        this.internal.interval = interval;
        return this;
    }
    
    public QueryOptionsSpecBuilder cacheTimeout(String cacheTimeout) {
        this.internal.cacheTimeout = cacheTimeout;
        return this;
    }
    
    public QueryOptionsSpecBuilder hideTimeOverride(Boolean hideTimeOverride) {
        this.internal.hideTimeOverride = hideTimeOverride;
        return this;
    }
    
    public QueryOptionsSpecBuilder timeCompare(String timeCompare) {
        this.internal.timeCompare = timeCompare;
        return this;
    }
    public QueryOptionsSpec build() {
        return this.internal;
    }
}
