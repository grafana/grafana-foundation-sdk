// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Time Series sub-query properties.
public class TimeSeriesQuery {
    // GCP project to execute the query against.
    @JsonProperty("projectName")
    public String projectName;
    // MQL query to be executed.
    @JsonProperty("query")
    public String query;
    // To disable the graphPeriod, it should explictly be set to 'disabled'.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("graphPeriod")
    public String graphPeriod;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TimeSeriesQuery> {
        protected final TimeSeriesQuery internal;
        
        public Builder() {
            this.internal = new TimeSeriesQuery();
        this.graphPeriod("disabled");
        }
    public Builder projectName(String projectName) {
    this.internal.projectName = projectName;
        return this;
    }
    
    public Builder query(String query) {
    this.internal.query = query;
        return this;
    }
    
    public Builder graphPeriod(String graphPeriod) {
    this.internal.graphPeriod = graphPeriod;
        return this;
    }
    public TimeSeriesQuery build() {
            return this.internal;
        }
    }
}
