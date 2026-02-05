// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

// GroupBy variable specification
public class GroupByVariableSpec {
    @JsonProperty("name")
    public String name;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("defaultValue")
    public VariableOption defaultValue;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("current")
    public VariableOption current;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("options")
    public List<VariableOption> options;
    @JsonProperty("multi")
    public Boolean multi;
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
    public GroupByVariableSpec() {
        this.name = "";
        this.current = new VariableOption(false, StringOrArrayOfString.createString(""), StringOrArrayOfString.createString(""));
        this.options = new LinkedList<>();
        this.multi = false;
        this.hide = VariableHide.DONT_HIDE;
        this.skipUrlSync = false;
    }
    public GroupByVariableSpec(String name,VariableOption defaultValue,VariableOption current,List<VariableOption> options,Boolean multi,String label,VariableHide hide,Boolean skipUrlSync,String description) {
        this.name = name;
        this.defaultValue = defaultValue;
        this.current = current;
        this.options = options;
        this.multi = multi;
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
