// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Constant variable specification
public class ConstantVariableSpec {
    @JsonProperty("name")
    public String name;
    @JsonProperty("query")
    public String query;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("current")
    public VariableOption current;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("label")
    public String label;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("hide")
    public VariableHide hide;
    @JsonProperty("skipUrlSync")
    public Boolean skipUrlSync;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("description")
    public String description;
    public ConstantVariableSpec() {
        this.name = "";
        this.query = "";
        this.current = new VariableOption(false, StringOrArrayOfString.createString(""), StringOrArrayOfString.createString(""));
        this.hide = VariableHide.DONT_HIDE;
        this.skipUrlSync = false;
    }
    public ConstantVariableSpec(String name,String query,VariableOption current,String label,VariableHide hide,Boolean skipUrlSync,String description) {
        this.name = name;
        this.query = query;
        this.current = current;
        this.label = label;
        this.hide = hide;
        this.skipUrlSync = skipUrlSync;
        this.description = description;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
