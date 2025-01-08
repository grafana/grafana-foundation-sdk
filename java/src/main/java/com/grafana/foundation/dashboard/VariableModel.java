// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
public class VariableModel {
    // Unique numeric identifier for the variable.
    @JsonProperty("id")
    public String id;
    // Type of variable
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public VariableType type;
    // Name of variable
    @JsonProperty("name")
    public String name;
    // Optional display name
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("label")
    public String label;
    // Visibility configuration for the variable
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("hide")
    public VariableHide hide;
    // Whether the variable value should be managed by URL query params or not
    @JsonProperty("skipUrlSync")
    public Boolean skipUrlSync;
    // Description of variable. It can be defined but `null`.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("description")
    public String description;
    // Query used to fetch values for a variable
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("query")
    public StringOrMap query;
    // Data source used to fetch values for a variable. It can be defined but `null`.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    // Format to use while fetching all values from data source, eg: wildcard, glob, regex, pipe, etc.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("allFormat")
    public String allFormat;
    // Shows current selected variable text/value on the dashboard
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("current")
    public VariableOption current;
    // Whether multiple values can be selected or not from variable value list
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("multi")
    public Boolean multi;
    // Options that can be selected for a variable.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("options")
    public List<VariableOption> options;
    // Options to config when to refresh a variable
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("refresh")
    public VariableRefresh refresh;
    // Options sort order
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("sort")
    public VariableSort sort;
    // Whether all value option is available or not
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("includeAll")
    public Boolean includeAll;
    // Custom all value
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("allValue")
    public String allValue;
    // Optional field, if you want to extract part of a series name or metric node segment.
    // Named capture groups can be used to separate the display text and value.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("regex")
    public String regex;
    // Dynamically calculates interval by dividing time range by the count specified.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("auto")
    public Boolean auto;
    // The minimum threshold below which the step count intervals will not divide the time.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("auto_min")
    public String autoMin;
    // How many times the current time range should be divided to calculate the value, similar to the Max data points query option.
    // For example, if the current visible time range is 30 minutes, then the auto interval groups the data into 30 one-minute increments.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("auto_count")
    public Integer autoCount;
    public VariableModel() {
        this.id = "00000000-0000-0000-0000-000000000000";
        this.skipUrlSync = false;
        this.multi = false;
        this.includeAll = false;
        this.auto = false;
        this.autoMin = "10s";
        this.autoCount = 30;
    }
    
    public VariableModel(String id,VariableType type,String name,String label,VariableHide hide,Boolean skipUrlSync,String description,StringOrMap query,DataSourceRef datasource,String allFormat,VariableOption current,Boolean multi,List<VariableOption> options,VariableRefresh refresh,VariableSort sort,Boolean includeAll,String allValue,String regex,Boolean auto,String autoMin,Integer autoCount) {
        this.id = id;
        this.type = type;
        this.name = name;
        this.label = label;
        this.hide = hide;
        this.skipUrlSync = skipUrlSync;
        this.description = description;
        this.query = query;
        this.datasource = datasource;
        this.allFormat = allFormat;
        this.current = current;
        this.multi = multi;
        this.options = options;
        this.refresh = refresh;
        this.sort = sort;
        this.includeAll = includeAll;
        this.allValue = allValue;
        this.regex = regex;
        this.auto = auto;
        this.autoMin = autoMin;
        this.autoCount = autoCount;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
