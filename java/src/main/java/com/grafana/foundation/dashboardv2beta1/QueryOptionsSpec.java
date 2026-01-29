// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class QueryOptionsSpec {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeFrom")
    public String timeFrom;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxDataPoints")
    public Long maxDataPoints;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeShift")
    public String timeShift;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryCachingTTL")
    public Long queryCachingTTL;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("interval")
    public String interval;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("cacheTimeout")
    public String cacheTimeout;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hideTimeOverride")
    public Boolean hideTimeOverride;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeCompare")
    public String timeCompare;
    public QueryOptionsSpec() {
    }
    public QueryOptionsSpec(String timeFrom,Long maxDataPoints,String timeShift,Long queryCachingTTL,String interval,String cacheTimeout,Boolean hideTimeOverride,String timeCompare) {
        this.timeFrom = timeFrom;
        this.maxDataPoints = maxDataPoints;
        this.timeShift = timeShift;
        this.queryCachingTTL = queryCachingTTL;
        this.interval = interval;
        this.cacheTimeout = cacheTimeout;
        this.hideTimeOverride = hideTimeOverride;
        this.timeCompare = timeCompare;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
