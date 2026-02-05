// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import java.util.List;

public class LegacyCloudMonitoringAnnotationQueryBuilder implements com.grafana.foundation.cog.Builder<LegacyCloudMonitoringAnnotationQuery> {
    protected final LegacyCloudMonitoringAnnotationQuery internal;
    
    public LegacyCloudMonitoringAnnotationQueryBuilder() {
        this.internal = new LegacyCloudMonitoringAnnotationQuery();
    }
    public LegacyCloudMonitoringAnnotationQueryBuilder projectName(String projectName) {
        this.internal.projectName = projectName;
        return this;
    }
    
    public LegacyCloudMonitoringAnnotationQueryBuilder metricType(String metricType) {
        this.internal.metricType = metricType;
        return this;
    }
    
    public LegacyCloudMonitoringAnnotationQueryBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public LegacyCloudMonitoringAnnotationQueryBuilder filters(List<String> filters) {
        this.internal.filters = filters;
        return this;
    }
    
    public LegacyCloudMonitoringAnnotationQueryBuilder metricKind(MetricKind metricKind) {
        this.internal.metricKind = metricKind;
        return this;
    }
    
    public LegacyCloudMonitoringAnnotationQueryBuilder valueType(String valueType) {
        this.internal.valueType = valueType;
        return this;
    }
    
    public LegacyCloudMonitoringAnnotationQueryBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public LegacyCloudMonitoringAnnotationQueryBuilder text(String text) {
        this.internal.text = text;
        return this;
    }
    public LegacyCloudMonitoringAnnotationQuery build() {
        return this.internal;
    }
}
