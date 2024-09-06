// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.tempo;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class TraceqlFilter {
    // Uniquely identify the filter, will not be used in the query generation
    @JsonProperty("id")
    public String id;
    // The tag for the search filter, for example: .http.status_code, .service.name, status
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("tag")
    public String tag;
    // The operator that connects the tag to the value, for example: =, >, !=, =~
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("operator")
    public String operator;
    // The value for the search filter
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("value")
    public StringOrArrayOfString value;
    // The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("valueType")
    public String valueType;
    // The scope of the filter, can either be unscoped/all scopes, resource or span
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("scope")
    public TraceqlSearchScope scope;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TraceqlFilter> {
        private final TraceqlFilter internal;
        
        public Builder() {
            this.internal = new TraceqlFilter();
        }
    public Builder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder tag(String tag) {
    this.internal.tag = tag;
        return this;
    }
    
    public Builder operator(String operator) {
    this.internal.operator = operator;
        return this;
    }
    
    public Builder value(StringOrArrayOfString value) {
    this.internal.value = value;
        return this;
    }
    
    public Builder valueType(String valueType) {
    this.internal.valueType = valueType;
        return this;
    }
    
    public Builder scope(TraceqlSearchScope scope) {
    this.internal.scope = scope;
        return this;
    }
    public TraceqlFilter build() {
            return this.internal;
        }
    }
}
