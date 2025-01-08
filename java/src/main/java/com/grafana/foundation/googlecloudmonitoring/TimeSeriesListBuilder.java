// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import java.util.List;

public class TimeSeriesListBuilder implements com.grafana.foundation.cog.Builder<TimeSeriesList> {
    protected final TimeSeriesList internal;
    
    public TimeSeriesListBuilder() {
        this.internal = new TimeSeriesList();
    }
    public TimeSeriesListBuilder projectName(String projectName) {
    this.internal.projectName = projectName;
        return this;
    }
    
    public TimeSeriesListBuilder crossSeriesReducer(String crossSeriesReducer) {
    this.internal.crossSeriesReducer = crossSeriesReducer;
        return this;
    }
    
    public TimeSeriesListBuilder alignmentPeriod(String alignmentPeriod) {
    this.internal.alignmentPeriod = alignmentPeriod;
        return this;
    }
    
    public TimeSeriesListBuilder perSeriesAligner(String perSeriesAligner) {
    this.internal.perSeriesAligner = perSeriesAligner;
        return this;
    }
    
    public TimeSeriesListBuilder groupBys(List<String> groupBys) {
    this.internal.groupBys = groupBys;
        return this;
    }
    
    public TimeSeriesListBuilder filters(List<String> filters) {
    this.internal.filters = filters;
        return this;
    }
    
    public TimeSeriesListBuilder view(String view) {
    this.internal.view = view;
        return this;
    }
    
    public TimeSeriesListBuilder title(String title) {
    this.internal.title = title;
        return this;
    }
    
    public TimeSeriesListBuilder text(String text) {
    this.internal.text = text;
        return this;
    }
    
    public TimeSeriesListBuilder secondaryCrossSeriesReducer(String secondaryCrossSeriesReducer) {
    this.internal.secondaryCrossSeriesReducer = secondaryCrossSeriesReducer;
        return this;
    }
    
    public TimeSeriesListBuilder secondaryAlignmentPeriod(String secondaryAlignmentPeriod) {
    this.internal.secondaryAlignmentPeriod = secondaryAlignmentPeriod;
        return this;
    }
    
    public TimeSeriesListBuilder secondaryPerSeriesAligner(String secondaryPerSeriesAligner) {
    this.internal.secondaryPerSeriesAligner = secondaryPerSeriesAligner;
        return this;
    }
    
    public TimeSeriesListBuilder secondaryGroupBys(List<String> secondaryGroupBys) {
    this.internal.secondaryGroupBys = secondaryGroupBys;
        return this;
    }
    
    public TimeSeriesListBuilder preprocessor(PreprocessorType preprocessor) {
    this.internal.preprocessor = preprocessor;
        return this;
    }
    public TimeSeriesList build() {
        return this.internal;
    }
}
