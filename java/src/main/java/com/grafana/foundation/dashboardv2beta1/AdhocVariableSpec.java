// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonInclude;

// Adhoc variable specification
public class AdhocVariableSpec {
    @JsonProperty("name")
    public String name;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("baseFilters")
    public List<AdHocFilterWithLabels> baseFilters;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("filters")
    public List<AdHocFilterWithLabels> filters;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("defaultKeys")
    public List<MetricFindValue> defaultKeys;
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
    public AdhocVariableSpec() {
        this.name = "";
        this.baseFilters = new LinkedList<>();
        this.filters = new LinkedList<>();
        this.defaultKeys = new LinkedList<>();
        this.hide = VariableHide.DONT_HIDE;
        this.skipUrlSync = false;
        this.allowCustomValue = true;
    }
    public AdhocVariableSpec(String name,List<AdHocFilterWithLabels> baseFilters,List<AdHocFilterWithLabels> filters,List<MetricFindValue> defaultKeys,String label,VariableHide hide,Boolean skipUrlSync,String description,Boolean allowCustomValue) {
        this.name = name;
        this.baseFilters = baseFilters;
        this.filters = filters;
        this.defaultKeys = defaultKeys;
        this.label = label;
        this.hide = hide;
        this.skipUrlSync = skipUrlSync;
        this.description = description;
        this.allowCustomValue = allowCustomValue;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
