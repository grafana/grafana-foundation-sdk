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

// Custom variable specification
public class CustomVariableSpec {
    @JsonProperty("name")
    public String name;
    @JsonProperty("query")
    public String query;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("current")
    public VariableOption current;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("options")
    public List<VariableOption> options;
    @JsonProperty("multi")
    public Boolean multi;
    @JsonProperty("includeAll")
    public Boolean includeAll;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("allValue")
    public String allValue;
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
    @JsonProperty("allowCustomValue")
    public Boolean allowCustomValue;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("valuesFormat")
    public CustomVariableSpecValuesFormat valuesFormat;
    public CustomVariableSpec() {
        this.name = "";
        this.query = "";
        this.current = new com.grafana.foundation.dashboardv2beta1.VariableOption();
        this.options = new LinkedList<>();
        this.multi = false;
        this.includeAll = false;
        this.hide = VariableHide.DONT_HIDE;
        this.skipUrlSync = false;
        this.allowCustomValue = true;
    }
    public CustomVariableSpec(String name,String query,VariableOption current,List<VariableOption> options,Boolean multi,Boolean includeAll,String allValue,String label,VariableHide hide,Boolean skipUrlSync,String description,Boolean allowCustomValue,CustomVariableSpecValuesFormat valuesFormat) {
        this.name = name;
        this.query = query;
        this.current = current;
        this.options = options;
        this.multi = multi;
        this.includeAll = includeAll;
        this.allValue = allValue;
        this.label = label;
        this.hide = hide;
        this.skipUrlSync = skipUrlSync;
        this.description = description;
        this.allowCustomValue = allowCustomValue;
        this.valuesFormat = valuesFormat;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
