// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// Annotation sub-query properties.
public class AnnotationQuery {
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
    // Annotation title.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("title")
    public String title;
    // Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("preprocessor")
    public PreprocessorType preprocessor;
    // Annotation text.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("text")
    public String text;
    public AnnotationQuery() {
    }
    
    public AnnotationQuery(String projectName,String crossSeriesReducer,String alignmentPeriod,String perSeriesAligner,List<String> groupBys,List<String> filters,String view,String secondaryCrossSeriesReducer,String secondaryAlignmentPeriod,String secondaryPerSeriesAligner,List<String> secondaryGroupBys,String title,PreprocessorType preprocessor,String text) {
        this.projectName = projectName;
        this.crossSeriesReducer = crossSeriesReducer;
        this.alignmentPeriod = alignmentPeriod;
        this.perSeriesAligner = perSeriesAligner;
        this.groupBys = groupBys;
        this.filters = filters;
        this.view = view;
        this.secondaryCrossSeriesReducer = secondaryCrossSeriesReducer;
        this.secondaryAlignmentPeriod = secondaryAlignmentPeriod;
        this.secondaryPerSeriesAligner = secondaryPerSeriesAligner;
        this.secondaryGroupBys = secondaryGroupBys;
        this.title = title;
        this.preprocessor = preprocessor;
        this.text = text;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
