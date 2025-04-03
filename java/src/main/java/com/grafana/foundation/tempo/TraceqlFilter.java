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
    public TraceqlFilter() {
    }
    public TraceqlFilter(String id,String tag,String operator,StringOrArrayOfString value,String valueType,TraceqlSearchScope scope) {
        this.id = id;
        this.tag = tag;
        this.operator = operator;
        this.value = value;
        this.valueType = valueType;
        this.scope = scope;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
