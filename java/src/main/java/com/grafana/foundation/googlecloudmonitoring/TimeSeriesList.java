// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// Time Series List sub-query properties.
public class TimeSeriesList {
    // GCP project to execute the query against.
    @JsonProperty("projectName")
    public String projectName;
    // Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
    @JsonProperty("crossSeriesReducer")
    public String crossSeriesReducer;
    // Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("alignmentPeriod")
    public String alignmentPeriod;
    // Alignment function to be used. Defaults to ALIGN_MEAN.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("perSeriesAligner")
    public String perSeriesAligner;
    // Array of labels to group data by.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("groupBys")
    public List<String> groupBys;
    // Array of filters to query data by. Labels that can be filtered on are defined by the metric.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("filters")
    public List<String> filters;
    // Data view, defaults to FULL.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("view")
    public String view;
    // Annotation title.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("title")
    public String title;
    // Annotation text.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("text")
    public String text;
    // Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("secondaryCrossSeriesReducer")
    public String secondaryCrossSeriesReducer;
    // Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("secondaryAlignmentPeriod")
    public String secondaryAlignmentPeriod;
    // Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("secondaryPerSeriesAligner")
    public String secondaryPerSeriesAligner;
    // Only present if a preprocessor is selected. Array of labels to group data by.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("secondaryGroupBys")
    public List<String> secondaryGroupBys;
    // Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("preprocessor")
    public PreprocessorType preprocessor;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TimeSeriesList> {
        private final TimeSeriesList internal;
        
        public Builder() {
            this.internal = new TimeSeriesList();
        }
    public Builder projectName(String projectName) {
    this.internal.projectName = projectName;
        return this;
    }
    
    public Builder crossSeriesReducer(String crossSeriesReducer) {
    this.internal.crossSeriesReducer = crossSeriesReducer;
        return this;
    }
    
    public Builder alignmentPeriod(String alignmentPeriod) {
    this.internal.alignmentPeriod = alignmentPeriod;
        return this;
    }
    
    public Builder perSeriesAligner(String perSeriesAligner) {
    this.internal.perSeriesAligner = perSeriesAligner;
        return this;
    }
    
    public Builder groupBys(List<String> groupBys) {
    this.internal.groupBys = groupBys;
        return this;
    }
    
    public Builder filters(List<String> filters) {
    this.internal.filters = filters;
        return this;
    }
    
    public Builder view(String view) {
    this.internal.view = view;
        return this;
    }
    
    public Builder title(String title) {
    this.internal.title = title;
        return this;
    }
    
    public Builder text(String text) {
    this.internal.text = text;
        return this;
    }
    
    public Builder secondaryCrossSeriesReducer(String secondaryCrossSeriesReducer) {
    this.internal.secondaryCrossSeriesReducer = secondaryCrossSeriesReducer;
        return this;
    }
    
    public Builder secondaryAlignmentPeriod(String secondaryAlignmentPeriod) {
    this.internal.secondaryAlignmentPeriod = secondaryAlignmentPeriod;
        return this;
    }
    
    public Builder secondaryPerSeriesAligner(String secondaryPerSeriesAligner) {
    this.internal.secondaryPerSeriesAligner = secondaryPerSeriesAligner;
        return this;
    }
    
    public Builder secondaryGroupBys(List<String> secondaryGroupBys) {
    this.internal.secondaryGroupBys = secondaryGroupBys;
        return this;
    }
    
    public Builder preprocessor(PreprocessorType preprocessor) {
    this.internal.preprocessor = preprocessor;
        return this;
    }
    public TimeSeriesList build() {
            return this.internal;
        }
    }
}
