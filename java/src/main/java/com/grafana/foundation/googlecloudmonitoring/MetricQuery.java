// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// @deprecated This type is for migration purposes only. Replaced by TimeSeriesList Metric sub-query properties.
public class MetricQuery {
    // GCP project to execute the query against.
    @JsonProperty("projectName")
    public String projectName;
    // Alignment function to be used. Defaults to ALIGN_MEAN.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("perSeriesAligner")
    public String perSeriesAligner;
    // Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("alignmentPeriod")
    public String alignmentPeriod;
    // Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
    @JsonInclude(JsonInclude.Include.NON_NULL)
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
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("groupBys")
    public List<String> groupBys;
    // Array of filters to query data by. Labels that can be filtered on are defined by the metric.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("filters")
    public List<String> filters;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("metricKind")
    public MetricKind metricKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("valueType")
    public String valueType;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("view")
    public String view;
    // MQL query to be executed.
    @JsonProperty("query")
    public String query;
    // Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("preprocessor")
    public PreprocessorType preprocessor;
    // To disable the graphPeriod, it should explictly be set to 'disabled'.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("graphPeriod")
    public String graphPeriod;
    public MetricQuery() {
        this.projectName = "";
        this.editorMode = "";
        this.metricType = "";
        this.crossSeriesReducer = "";
        this.query = "";
        this.graphPeriod = "disabled";
    }
    public MetricQuery(String projectName,String perSeriesAligner,String alignmentPeriod,String aliasBy,String editorMode,String metricType,String crossSeriesReducer,List<String> groupBys,List<String> filters,MetricKind metricKind,String valueType,String view,String query,PreprocessorType preprocessor,String graphPeriod) {
        this.projectName = projectName;
        this.perSeriesAligner = perSeriesAligner;
        this.alignmentPeriod = alignmentPeriod;
        this.aliasBy = aliasBy;
        this.editorMode = editorMode;
        this.metricType = metricType;
        this.crossSeriesReducer = crossSeriesReducer;
        this.groupBys = groupBys;
        this.filters = filters;
        this.metricKind = metricKind;
        this.valueType = valueType;
        this.view = view;
        this.query = query;
        this.preprocessor = preprocessor;
        this.graphPeriod = graphPeriod;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
