// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class BuilderQueryEditorOperator {
    @JsonProperty("name")
    public String name;
    @JsonProperty("value")
    public String value;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("labelValue")
    public String labelValue;
    public BuilderQueryEditorOperator() {
    }
    public BuilderQueryEditorOperator(String name,String value,String labelValue) {
        this.name = name;
        this.value = value;
        this.labelValue = labelValue;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
