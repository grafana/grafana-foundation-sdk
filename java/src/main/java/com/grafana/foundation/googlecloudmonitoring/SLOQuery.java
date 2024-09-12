// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// SLO sub-query properties.
public class SLOQuery {
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
    // SLO selector.
    @JsonProperty("selectorName")
    public String selectorName;
    // ID for the service the SLO is in.
    @JsonProperty("serviceId")
    public String serviceId;
    // Name for the service the SLO is in.
    @JsonProperty("serviceName")
    public String serviceName;
    // ID for the SLO.
    @JsonProperty("sloId")
    public String sloId;
    // Name of the SLO.
    @JsonProperty("sloName")
    public String sloName;
    // SLO goal value.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("goal")
    public Double goal;
    // Specific lookback period for the SLO.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lookbackPeriod")
    public String lookbackPeriod;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<SLOQuery> {
        private final SLOQuery internal;
        
        public Builder() {
            this.internal = new SLOQuery();
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
    
    public Builder selectorName(String selectorName) {
    this.internal.selectorName = selectorName;
        return this;
    }
    
    public Builder serviceId(String serviceId) {
    this.internal.serviceId = serviceId;
        return this;
    }
    
    public Builder serviceName(String serviceName) {
    this.internal.serviceName = serviceName;
        return this;
    }
    
    public Builder sloId(String sloId) {
    this.internal.sloId = sloId;
        return this;
    }
    
    public Builder sloName(String sloName) {
    this.internal.sloName = sloName;
        return this;
    }
    
    public Builder goal(Double goal) {
    this.internal.goal = goal;
        return this;
    }
    
    public Builder lookbackPeriod(String lookbackPeriod) {
    this.internal.lookbackPeriod = lookbackPeriod;
        return this;
    }
    public SLOQuery build() {
            return this.internal;
        }
    }
}
