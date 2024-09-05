// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.List;

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
    @JsonProperty("filters")
    public List<String> filters;
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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<LegacyCloudMonitoringAnnotationQuery> {
        private final LegacyCloudMonitoringAnnotationQuery internal;
        
        public Builder() {
            this.internal = new LegacyCloudMonitoringAnnotationQuery();
        }
    public Builder projectName(String projectName) {
    this.internal.projectName = projectName;
        return this;
    }
    
    public Builder metricType(String metricType) {
    this.internal.metricType = metricType;
        return this;
    }
    
    public Builder refId(String refId) {
    this.internal.refId = refId;
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
    
    public Builder title(String title) {
    this.internal.title = title;
        return this;
    }
    
    public Builder text(String text) {
    this.internal.text = text;
        return this;
    }
    public LegacyCloudMonitoringAnnotationQuery build() {
            return this.internal;
        }
    }
}
