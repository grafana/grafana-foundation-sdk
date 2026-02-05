// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ResourceGroupsQuery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("rawQuery")
    public String rawQuery;
    @JsonProperty("kind")
    public String kind;
    @JsonProperty("subscription")
    public String subscription;
    public ResourceGroupsQuery() {
        this.kind = "";
        this.subscription = "";
    }
    public ResourceGroupsQuery(String rawQuery,String kind,String subscription) {
        this.rawQuery = rawQuery;
        this.kind = kind;
        this.subscription = subscription;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
