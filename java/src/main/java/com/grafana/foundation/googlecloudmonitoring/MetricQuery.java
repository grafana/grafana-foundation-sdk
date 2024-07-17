// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// @deprecated This type is for migration purposes only. Replaced by TimeSeriesList Metric sub-query properties.
public class MetricQuery {
    // GCP project to execute the query against. 
    @JsonProperty("projectName")
    public String projectName;
    // Alignment function to be used. Defaults to ALIGN_MEAN. 
    @JsonProperty("perSeriesAligner")
    public String perSeriesAligner;
    // Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto. 
    @JsonProperty("alignmentPeriod")
    public String alignmentPeriod;
    // Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail. 
    @JsonProperty("aliasBy")
    public String aliasBy; 
    @JsonProperty("editorMode")
    public String editorMode; 
    @JsonProperty("metricType")
    public String metricType;
    // Reducer applied across a set of time-series values. Defaults to REDUCE_NONE. 
    @JsonProperty("crossSeriesReducer")
    public String crossSeriesReducer;
    // Array of labels to group data by. 
    @JsonProperty("groupBys")
    public List<String> groupBys;
    // Array of filters to query data by. Labels that can be filtered on are defined by the metric. 
    @JsonProperty("filters")
    public List<String> filters; 
    @JsonProperty("metricKind")
    public MetricKind metricKind; 
    @JsonProperty("valueType")
    public String valueType; 
    @JsonProperty("view")
    public String view;
    // MQL query to be executed. 
    @JsonProperty("query")
    public String query;
    // Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters 
    @JsonProperty("preprocessor")
    public PreprocessorType preprocessor;
    // To disable the graphPeriod, it should explictly be set to 'disabled'. 
    @JsonProperty("graphPeriod")
    public String graphPeriod;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<MetricQuery> {
        private final MetricQuery internal;
        
        public Builder() {
            this.internal = new MetricQuery();
        this.graphPeriod("disabled");
        }
    public Builder projectName(String projectName) {
    this.internal.projectName = projectName;
        return this;
    }
    
    public Builder perSeriesAligner(String perSeriesAligner) {
    this.internal.perSeriesAligner = perSeriesAligner;
        return this;
    }
    
    public Builder alignmentPeriod(String alignmentPeriod) {
    this.internal.alignmentPeriod = alignmentPeriod;
        return this;
    }
    
    public Builder aliasBy(String aliasBy) {
    this.internal.aliasBy = aliasBy;
        return this;
    }
    
    public Builder editorMode(String editorMode) {
    this.internal.editorMode = editorMode;
        return this;
    }
    
    public Builder metricType(String metricType) {
    this.internal.metricType = metricType;
        return this;
    }
    
    public Builder crossSeriesReducer(String crossSeriesReducer) {
    this.internal.crossSeriesReducer = crossSeriesReducer;
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
    
    public Builder metricKind(MetricKind metricKind) {
    this.internal.metricKind = metricKind;
        return this;
    }
    
    public Builder valueType(String valueType) {
    this.internal.valueType = valueType;
        return this;
    }
    
    public Builder view(String view) {
    this.internal.view = view;
        return this;
    }
    
    public Builder query(String query) {
    this.internal.query = query;
        return this;
    }
    
    public Builder preprocessor(PreprocessorType preprocessor) {
    this.internal.preprocessor = preprocessor;
        return this;
    }
    
    public Builder graphPeriod(String graphPeriod) {
    this.internal.graphPeriod = graphPeriod;
        return this;
    }
    public MetricQuery build() {
            return this.internal;
        }
    }
}
