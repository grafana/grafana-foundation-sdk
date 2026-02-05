// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Text variable specification
public class TextVariableSpec {
    @JsonProperty("name")
    public String name;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("current")
    public VariableOption current;
    @JsonProperty("query")
    public String query;
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
    public TextVariableSpec() {
        this.name = "";
        this.current = new VariableOption(false, StringOrArrayOfString.createString(""), StringOrArrayOfString.createString(""));
        this.query = "";
        this.hide = VariableHide.DONT_HIDE;
        this.skipUrlSync = false;
    }
    public TextVariableSpec(String name,VariableOption current,String query,String label,VariableHide hide,Boolean skipUrlSync,String description) {
        this.name = name;
        this.current = current;
        this.query = query;
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
