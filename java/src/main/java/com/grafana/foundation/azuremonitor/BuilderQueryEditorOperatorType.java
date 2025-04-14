// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class BuilderQueryEditorOperatorType {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("String")
    public String string;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("Bool")
    public Boolean bool;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("Float64")
    public Double float64;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("SelectableValue")
    public SelectableValue selectableValue;
    public BuilderQueryEditorOperatorType() {
    }
    public BuilderQueryEditorOperatorType(String string,Boolean bool,Double float64,SelectableValue selectableValue) {
        this.string = string;
        this.bool = bool;
        this.float64 = float64;
        this.selectableValue = selectableValue;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
