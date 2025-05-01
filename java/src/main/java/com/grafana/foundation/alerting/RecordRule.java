// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class RecordRule {
    // Which expression node should be used as the input for the recorded metric.
    @JsonProperty("from")
    public String from;
    // Name of the recorded metric.
    @JsonProperty("metric")
    public String metric;
    // Which data source should be used to write the output of the recording rule, specified by UID.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("target_datasource_uid")
    public String targetDatasourceUid;
    public RecordRule() {
        this.from = "";
        this.metric = "";
    }
    public RecordRule(String from,String metric,String targetDatasourceUid) {
        this.from = from;
        this.metric = metric;
        this.targetDatasourceUid = targetDatasourceUid;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
