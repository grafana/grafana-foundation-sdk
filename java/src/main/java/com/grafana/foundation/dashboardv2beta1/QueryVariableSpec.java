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

// Query variable specification
public class QueryVariableSpec {
    @JsonProperty("name")
    public String name;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("current")
    public VariableOption current;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("label")
    public String label;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("hide")
    public VariableHide hide;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("refresh")
    public VariableRefresh refresh;
    @JsonProperty("skipUrlSync")
    public Boolean skipUrlSync;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("description")
    public String description;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("query")
    public DataQueryKind query;
    @JsonProperty("regex")
    public String regex;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("regexApplyTo")
    public VariableRegexApplyTo regexApplyTo;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("sort")
    public VariableSort sort;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("definition")
    public String definition;
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
    @JsonProperty("placeholder")
    public String placeholder;
    @JsonProperty("allowCustomValue")
    public Boolean allowCustomValue;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("staticOptions")
    public List<VariableOption> staticOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("staticOptionsOrder")
    public QueryVariableSpecStaticOptionsOrder staticOptionsOrder;
    public QueryVariableSpec() {
        this.name = "";
        this.current = new VariableOption(false, StringOrArrayOfString.createString(""), StringOrArrayOfString.createString(""));
        this.hide = VariableHide.DONT_HIDE;
        this.refresh = VariableRefresh.NEVER;
        this.skipUrlSync = false;
        this.query = new com.grafana.foundation.dashboardv2beta1.DataQueryKind();
        this.regex = "";
        this.regexApplyTo = VariableRegexApplyTo.VALUE;
        this.sort = VariableSort.DISABLED;
        this.options = new LinkedList<>();
        this.multi = false;
        this.includeAll = false;
        this.allowCustomValue = true;
    }
    public QueryVariableSpec(String name,VariableOption current,String label,VariableHide hide,VariableRefresh refresh,Boolean skipUrlSync,String description,DataQueryKind query,String regex,VariableRegexApplyTo regexApplyTo,VariableSort sort,String definition,List<VariableOption> options,Boolean multi,Boolean includeAll,String allValue,String placeholder,Boolean allowCustomValue,List<VariableOption> staticOptions,QueryVariableSpecStaticOptionsOrder staticOptionsOrder) {
        this.name = name;
        this.current = current;
        this.label = label;
        this.hide = hide;
        this.refresh = refresh;
        this.skipUrlSync = skipUrlSync;
        this.description = description;
        this.query = query;
        this.regex = regex;
        this.regexApplyTo = regexApplyTo;
        this.sort = sort;
        this.definition = definition;
        this.options = options;
        this.multi = multi;
        this.includeAll = includeAll;
        this.allValue = allValue;
        this.placeholder = placeholder;
        this.allowCustomValue = allowCustomValue;
        this.staticOptions = staticOptions;
        this.staticOptionsOrder = staticOptionsOrder;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
