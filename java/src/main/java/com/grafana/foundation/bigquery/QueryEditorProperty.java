// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bigquery;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class QueryEditorProperty {
    @JsonProperty("type")
    public QueryEditorPropertyType type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("name")
    public String name;
    public QueryEditorProperty() {
        this.type = QueryEditorPropertyType.STRING;
    }
    public QueryEditorProperty(String name) {
        this.type = QueryEditorPropertyType.STRING;
        this.name = name;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
