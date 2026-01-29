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

// Datasource variable specification
public class DatasourceVariableSpec {
    @JsonProperty("name")
    public String name;
    @JsonProperty("pluginId")
    public String pluginId;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("refresh")
    public VariableRefresh refresh;
    @JsonProperty("regex")
    public String regex;
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
    public DatasourceVariableSpec() {
        this.name = "";
        this.pluginId = "";
        this.refresh = VariableRefresh.NEVER;
        this.regex = "";
        this.current = new VariableOption(false, StringOrArrayOfString.createString(""), StringOrArrayOfString.createString(""));
        this.options = new LinkedList<>();
        this.multi = false;
        this.includeAll = false;
        this.hide = VariableHide.DONT_HIDE;
        this.skipUrlSync = false;
        this.allowCustomValue = true;
    }
    public DatasourceVariableSpec(String name,String pluginId,VariableRefresh refresh,String regex,VariableOption current,List<VariableOption> options,Boolean multi,Boolean includeAll,String allValue,String label,VariableHide hide,Boolean skipUrlSync,String description,Boolean allowCustomValue) {
        this.name = name;
        this.pluginId = pluginId;
        this.refresh = refresh;
        this.regex = regex;
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
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
