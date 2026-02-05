// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class SwitchVariableSpec {
    @JsonProperty("name")
    public String name;
    @JsonProperty("current")
    public String current;
    @JsonProperty("enabledValue")
    public String enabledValue;
    @JsonProperty("disabledValue")
    public String disabledValue;
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
    public SwitchVariableSpec() {
        this.name = "";
        this.current = "false";
        this.enabledValue = "true";
        this.disabledValue = "false";
        this.hide = VariableHide.DONT_HIDE;
        this.skipUrlSync = false;
    }
    public SwitchVariableSpec(String name,String current,String enabledValue,String disabledValue,String label,VariableHide hide,Boolean skipUrlSync,String description) {
        this.name = name;
        this.current = current;
        this.enabledValue = enabledValue;
        this.disabledValue = disabledValue;
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
