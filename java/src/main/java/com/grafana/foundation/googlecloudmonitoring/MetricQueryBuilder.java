// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import java.util.List;

public class MetricQueryBuilder implements com.grafana.foundation.cog.Builder<MetricQuery> {
    protected final MetricQuery internal;
    
    public MetricQueryBuilder() {
        this.internal = new MetricQuery();
    }
    public MetricQueryBuilder projectName(String projectName) {
    this.internal.projectName = projectName;
        return this;
    }
    
    public MetricQueryBuilder perSeriesAligner(String perSeriesAligner) {
    this.internal.perSeriesAligner = perSeriesAligner;
        return this;
    }
    
    public MetricQueryBuilder alignmentPeriod(String alignmentPeriod) {
    this.internal.alignmentPeriod = alignmentPeriod;
        return this;
    }
    
    public MetricQueryBuilder aliasBy(String aliasBy) {
    this.internal.aliasBy = aliasBy;
        return this;
    }
    
    public MetricQueryBuilder editorMode(String editorMode) {
    this.internal.editorMode = editorMode;
        return this;
    }
    
    public MetricQueryBuilder metricType(String metricType) {
    this.internal.metricType = metricType;
        return this;
    }
    
    public MetricQueryBuilder crossSeriesReducer(String crossSeriesReducer) {
    this.internal.crossSeriesReducer = crossSeriesReducer;
        return this;
    }
    
    public MetricQueryBuilder groupBys(List<String> groupBys) {
    this.internal.groupBys = groupBys;
        return this;
    }
    
    public MetricQueryBuilder filters(List<String> filters) {
    this.internal.filters = filters;
        return this;
    }
    
    public MetricQueryBuilder metricKind(MetricKind metricKind) {
    this.internal.metricKind = metricKind;
        return this;
    }
    
    public MetricQueryBuilder valueType(String valueType) {
    this.internal.valueType = valueType;
        return this;
    }
    
    public MetricQueryBuilder view(String view) {
    this.internal.view = view;
        return this;
    }
    
    public MetricQueryBuilder query(String query) {
    this.internal.query = query;
        return this;
    }
    
    public MetricQueryBuilder preprocessor(PreprocessorType preprocessor) {
    this.internal.preprocessor = preprocessor;
        return this;
    }
    
    public MetricQueryBuilder graphPeriod(String graphPeriod) {
    this.internal.graphPeriod = graphPeriod;
        return this;
    }
    public MetricQuery build() {
        return this.internal;
    }
}
