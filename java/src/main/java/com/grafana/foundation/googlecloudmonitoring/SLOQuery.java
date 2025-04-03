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
    public SLOQuery() {
    }
    public SLOQuery(String projectName,String perSeriesAligner,String alignmentPeriod,String selectorName,String serviceId,String serviceName,String sloId,String sloName,Double goal,String lookbackPeriod) {
        this.projectName = projectName;
        this.perSeriesAligner = perSeriesAligner;
        this.alignmentPeriod = alignmentPeriod;
        this.selectorName = selectorName;
        this.serviceId = serviceId;
        this.serviceName = serviceName;
        this.sloId = sloId;
        this.sloName = sloName;
        this.goal = goal;
        this.lookbackPeriod = lookbackPeriod;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
