// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import java.util.List;

public class AnnotationQueryBuilder implements com.grafana.foundation.cog.Builder<AnnotationQuery> {
    protected final AnnotationQuery internal;
    
    public AnnotationQueryBuilder() {
        this.internal = new AnnotationQuery();
    }
    public AnnotationQueryBuilder projectName(String projectName) {
    this.internal.projectName = projectName;
        return this;
    }
    
    public AnnotationQueryBuilder crossSeriesReducer(String crossSeriesReducer) {
    this.internal.crossSeriesReducer = crossSeriesReducer;
        return this;
    }
    
    public AnnotationQueryBuilder alignmentPeriod(String alignmentPeriod) {
    this.internal.alignmentPeriod = alignmentPeriod;
        return this;
    }
    
    public AnnotationQueryBuilder perSeriesAligner(String perSeriesAligner) {
    this.internal.perSeriesAligner = perSeriesAligner;
        return this;
    }
    
    public AnnotationQueryBuilder groupBys(List<String> groupBys) {
    this.internal.groupBys = groupBys;
        return this;
    }
    
    public AnnotationQueryBuilder filters(List<String> filters) {
    this.internal.filters = filters;
        return this;
    }
    
    public AnnotationQueryBuilder view(String view) {
    this.internal.view = view;
        return this;
    }
    
    public AnnotationQueryBuilder secondaryCrossSeriesReducer(String secondaryCrossSeriesReducer) {
    this.internal.secondaryCrossSeriesReducer = secondaryCrossSeriesReducer;
        return this;
    }
    
    public AnnotationQueryBuilder secondaryAlignmentPeriod(String secondaryAlignmentPeriod) {
    this.internal.secondaryAlignmentPeriod = secondaryAlignmentPeriod;
        return this;
    }
    
    public AnnotationQueryBuilder secondaryPerSeriesAligner(String secondaryPerSeriesAligner) {
    this.internal.secondaryPerSeriesAligner = secondaryPerSeriesAligner;
        return this;
    }
    
    public AnnotationQueryBuilder secondaryGroupBys(List<String> secondaryGroupBys) {
    this.internal.secondaryGroupBys = secondaryGroupBys;
        return this;
    }
    
    public AnnotationQueryBuilder title(String title) {
    this.internal.title = title;
        return this;
    }
    
    public AnnotationQueryBuilder preprocessor(PreprocessorType preprocessor) {
    this.internal.preprocessor = preprocessor;
        return this;
    }
    
    public AnnotationQueryBuilder text(String text) {
    this.internal.text = text;
        return this;
    }
    public AnnotationQuery build() {
        return this.internal;
    }
}
