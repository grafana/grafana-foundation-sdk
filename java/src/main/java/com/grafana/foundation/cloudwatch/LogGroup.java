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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<LogGroup> {
        private final LogGroup internal;
        
        public Builder() {
            this.internal = new LogGroup();
        }
    public Builder arn(String arn) {
    this.internal.arn = arn;
        return this;
    }
    
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder accountId(String accountId) {
    this.internal.accountId = accountId;
        return this;
    }
    
    public Builder accountLabel(String accountLabel) {
    this.internal.accountLabel = accountLabel;
        return this;
    }
    public LogGroup build() {
            return this.internal;
        }
    }
}
