// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonInclude;

// @deprecated Use TimeSeriesList instead. Legacy annotation query properties for migration purposes.
public class LegacyCloudMonitoringAnnotationQuery {
    // GCP project to execute the query against.
    @JsonProperty("projectName")
    public String projectName;
    @JsonProperty("metricType")
    public String metricType;
    // Query refId.
    @JsonProperty("refId")
    public String refId;
    // Array of filters to query data by. Labels that can be filtered on are defined by the metric.
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("filters")
    public List<String> filters;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("metricKind")
    public MetricKind metricKind;
    @JsonProperty("valueType")
    public String valueType;
    // Annotation title.
    @JsonProperty("title")
    public String title;
    // Annotation text.
    @JsonProperty("text")
    public String text;
    public LegacyCloudMonitoringAnnotationQuery() {
    }
    
    public LegacyCloudMonitoringAnnotationQuery(String projectName,String metricType,String refId,List<String> filters,MetricKind metricKind,String valueType,String title,String text) {
        this.projectName = projectName;
        this.metricType = metricType;
        this.refId = refId;
        this.filters = filters;
        this.metricKind = metricKind;
        this.valueType = valueType;
        this.title = title;
        this.text = text;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
