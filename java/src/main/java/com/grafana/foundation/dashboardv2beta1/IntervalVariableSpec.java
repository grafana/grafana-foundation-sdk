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

// Interval variable specification
public class IntervalVariableSpec {
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
    @JsonProperty("auto")
    public Boolean auto;
    @JsonProperty("auto_min")
    public String autoMin;
    @JsonProperty("auto_count")
    public Long autoCount;
    @JsonProperty("refresh")
    public String refresh;
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
    public IntervalVariableSpec() {
        this.name = "";
        this.query = "";
        this.current = new VariableOption(false, StringOrArrayOfString.createString(""), StringOrArrayOfString.createString(""));
        this.options = new LinkedList<>();
        this.auto = false;
        this.autoMin = "";
        this.autoCount = 0L;
        this.refresh = "";
        this.hide = VariableHide.DONT_HIDE;
        this.skipUrlSync = false;
    }
    public IntervalVariableSpec(String name,String query,VariableOption current,List<VariableOption> options,Boolean auto,String autoMin,Long autoCount,String refresh,String label,VariableHide hide,Boolean skipUrlSync,String description) {
        this.name = name;
        this.query = query;
        this.current = current;
        this.options = options;
        this.auto = auto;
        this.autoMin = autoMin;
        this.autoCount = autoCount;
        this.refresh = refresh;
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
