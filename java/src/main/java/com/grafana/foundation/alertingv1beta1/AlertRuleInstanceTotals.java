// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

/**
 * Count of alert instances per state. error also counts instances whose evaluation
 * errored but were mapped to another state via execErrState, so it can overlap.
 */
public class AlertRuleInstanceTotals {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("healthy")
    public Long healthy;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("firing")
    public Long firing;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pending")
    public Long pending;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("recovering")
    public Long recovering;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("nodata")
    public Long nodata;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("error")
    public Long error;
    public AlertRuleInstanceTotals() {
    }
    public AlertRuleInstanceTotals(Long healthy,Long firing,Long pending,Long recovering,Long nodata,Long error) {
        this.healthy = healthy;
        this.firing = firing;
        this.pending = pending;
        this.recovering = recovering;
        this.nodata = nodata;
        this.error = error;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
