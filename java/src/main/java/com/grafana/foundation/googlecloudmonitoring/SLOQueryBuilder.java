// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;


public class SLOQueryBuilder implements com.grafana.foundation.cog.Builder<SLOQuery> {
    protected final SLOQuery internal;
    
    public SLOQueryBuilder() {
        this.internal = new SLOQuery();
    }
    public SLOQueryBuilder projectName(String projectName) {
    this.internal.projectName = projectName;
        return this;
    }
    
    public SLOQueryBuilder perSeriesAligner(String perSeriesAligner) {
    this.internal.perSeriesAligner = perSeriesAligner;
        return this;
    }
    
    public SLOQueryBuilder alignmentPeriod(String alignmentPeriod) {
    this.internal.alignmentPeriod = alignmentPeriod;
        return this;
    }
    
    public SLOQueryBuilder selectorName(String selectorName) {
    this.internal.selectorName = selectorName;
        return this;
    }
    
    public SLOQueryBuilder serviceId(String serviceId) {
    this.internal.serviceId = serviceId;
        return this;
    }
    
    public SLOQueryBuilder serviceName(String serviceName) {
    this.internal.serviceName = serviceName;
        return this;
    }
    
    public SLOQueryBuilder sloId(String sloId) {
    this.internal.sloId = sloId;
        return this;
    }
    
    public SLOQueryBuilder sloName(String sloName) {
    this.internal.sloName = sloName;
        return this;
    }
    
    public SLOQueryBuilder goal(Double goal) {
    this.internal.goal = goal;
        return this;
    }
    
    public SLOQueryBuilder lookbackPeriod(String lookbackPeriod) {
    this.internal.lookbackPeriod = lookbackPeriod;
        return this;
    }
    public SLOQuery build() {
        return this.internal;
    }
}
