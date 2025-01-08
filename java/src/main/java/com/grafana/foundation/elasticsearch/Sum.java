// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Sum {
    @JsonProperty("type")
    public String type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("field")
    public String field;
    @JsonProperty("id")
    public String id;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("settings")
    public ElasticsearchSumSettings settings;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    public Sum() {
    }
    
    public Sum(String type,String field,String id,ElasticsearchSumSettings settings,Boolean hide) {
        this.type = type;
        this.field = field;
        this.id = id;
        this.settings = settings;
        this.hide = hide;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
