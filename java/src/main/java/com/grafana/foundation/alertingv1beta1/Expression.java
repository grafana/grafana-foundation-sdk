// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Expression {
    // The type of query if this is a query expression
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryType")
    public String queryType;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("relativeTimeRange")
    public RelativeTimeRange relativeTimeRange;
    // The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource
    @JsonProperty("datasourceUID")
    public String datasourceUID;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("model")
    public Object model;
    // Used to mark the expression to be used as the final source for the rule evaluation
    // Only one expression in a rule can be marked as the source
    // For AlertRules, this is the expression that will be evaluated against the alerting condition
    // For RecordingRules, this is the expression that will be recorded
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("source")
    public Boolean source;
    public Expression() {
        this.model = new Object();
    }
    public Expression(String queryType,RelativeTimeRange relativeTimeRange,String datasourceUID,Object model,Boolean source) {
        this.queryType = queryType;
        this.relativeTimeRange = relativeTimeRange;
        this.datasourceUID = datasourceUID;
        this.model = model;
        this.source = source;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
