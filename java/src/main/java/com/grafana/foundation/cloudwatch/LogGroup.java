// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class LogGroup {
    // ARN of the log group
    @JsonProperty("arn")
    public String arn;
    // Name of the log group
    @JsonProperty("name")
    public String name;
    // AccountId of the log group
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("accountId")
    public String accountId;
    // Label of the log group
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("accountLabel")
    public String accountLabel;
    public LogGroup() {
    }
    public LogGroup(String arn,String name,String accountId,String accountLabel) {
        this.arn = arn;
        this.name = name;
        this.accountId = accountId;
        this.accountLabel = accountLabel;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
