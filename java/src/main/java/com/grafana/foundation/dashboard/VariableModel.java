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
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public VariableHide hide;
    // Whether the variable value should be managed by URL query params or not
    @JsonInclude(JsonInclude.Include.NON_NULL)
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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class QueryVariableBuilder implements com.grafana.foundation.cog.Builder<VariableModel> {
        private final VariableModel internal;
        
        public QueryVariableBuilder(String name) {
            this.internal = new VariableModel();
    this.internal.name = name;
    this.internal.type = VariableType.QUERY;
        }
    public QueryVariableBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public QueryVariableBuilder label(String label) {
    this.internal.label = label;
        return this;
    }
    
    public QueryVariableBuilder hide(VariableHide hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public QueryVariableBuilder description(String description) {
    this.internal.description = description;
        return this;
    }
    
    public QueryVariableBuilder query(StringOrMap query) {
    this.internal.query = query;
        return this;
    }
    
    public QueryVariableBuilder datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    
    public QueryVariableBuilder current(VariableOption current) {
    this.internal.current = current;
        return this;
    }
    
    public QueryVariableBuilder multi(Boolean multi) {
    this.internal.multi = multi;
        return this;
    }
    
    public QueryVariableBuilder options(List<VariableOption> options) {
    this.internal.options = options;
        return this;
    }
    
    public QueryVariableBuilder refresh(VariableRefresh refresh) {
    this.internal.refresh = refresh;
        return this;
    }
    
    public QueryVariableBuilder sort(VariableSort sort) {
    this.internal.sort = sort;
        return this;
    }
    
    public QueryVariableBuilder includeAll(Boolean includeAll) {
    this.internal.includeAll = includeAll;
        return this;
    }
    
    public QueryVariableBuilder allValue(String allValue) {
    this.internal.allValue = allValue;
        return this;
    }
    
    public QueryVariableBuilder regex(String regex) {
    this.internal.regex = regex;
        return this;
    }
    public VariableModel build() {
            return this.internal;
        }
    }
    
    public static class AdHocVariableBuilder implements com.grafana.foundation.cog.Builder<VariableModel> {
        private final VariableModel internal;
        
        public AdHocVariableBuilder(String name) {
            this.internal = new VariableModel();
    this.internal.name = name;
    this.internal.type = VariableType.ADHOC;
        }
    public AdHocVariableBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public AdHocVariableBuilder label(String label) {
    this.internal.label = label;
        return this;
    }
    
    public AdHocVariableBuilder hide(VariableHide hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public AdHocVariableBuilder description(String description) {
    this.internal.description = description;
        return this;
    }
    
    public AdHocVariableBuilder datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    public VariableModel build() {
            return this.internal;
        }
    }
    
    public static class ConstantVariableBuilder implements com.grafana.foundation.cog.Builder<VariableModel> {
        private final VariableModel internal;
        
        public ConstantVariableBuilder(String name) {
            this.internal = new VariableModel();
    this.internal.name = name;
    this.internal.type = VariableType.CONSTANT;
        }
    public ConstantVariableBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public ConstantVariableBuilder label(String label) {
    this.internal.label = label;
        return this;
    }
    
    public ConstantVariableBuilder description(String description) {
    this.internal.description = description;
        return this;
    }
    
    public ConstantVariableBuilder value(StringOrMap query) {
    this.internal.query = query;
        return this;
    }
    public VariableModel build() {
            return this.internal;
        }
    }
    
    public static class DatasourceVariableBuilder implements com.grafana.foundation.cog.Builder<VariableModel> {
        private final VariableModel internal;
        
        public DatasourceVariableBuilder(String name) {
            this.internal = new VariableModel();
    this.internal.name = name;
    this.internal.type = VariableType.DATASOURCE;
        }
    public DatasourceVariableBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public DatasourceVariableBuilder label(String label) {
    this.internal.label = label;
        return this;
    }
    
    public DatasourceVariableBuilder hide(VariableHide hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public DatasourceVariableBuilder description(String description) {
    this.internal.description = description;
        return this;
    }
    
    public DatasourceVariableBuilder type(String string) {
		if (this.internal.query == null) {
			this.internal.query = new com.grafana.foundation.dashboard.StringOrMap();
		}
    this.internal.query.string = string;
        return this;
    }
    
    public DatasourceVariableBuilder current(VariableOption current) {
    this.internal.current = current;
        return this;
    }
    
    public DatasourceVariableBuilder multi(Boolean multi) {
    this.internal.multi = multi;
        return this;
    }
    
    public DatasourceVariableBuilder includeAll(Boolean includeAll) {
    this.internal.includeAll = includeAll;
        return this;
    }
    
    public DatasourceVariableBuilder allValue(String allValue) {
    this.internal.allValue = allValue;
        return this;
    }
    
    public DatasourceVariableBuilder regex(String regex) {
    this.internal.regex = regex;
        return this;
    }
    public VariableModel build() {
            return this.internal;
        }
    }
    
    public static class IntervalVariableBuilder implements com.grafana.foundation.cog.Builder<VariableModel> {
        private final VariableModel internal;
        
        public IntervalVariableBuilder(String name) {
            this.internal = new VariableModel();
    this.internal.name = name;
    this.internal.type = VariableType.INTERVAL;
        }
    public IntervalVariableBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public IntervalVariableBuilder label(String label) {
    this.internal.label = label;
        return this;
    }
    
    public IntervalVariableBuilder hide(VariableHide hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public IntervalVariableBuilder description(String description) {
    this.internal.description = description;
        return this;
    }
    
    public IntervalVariableBuilder values(StringOrMap query) {
    this.internal.query = query;
        return this;
    }
    
    public IntervalVariableBuilder current(VariableOption current) {
    this.internal.current = current;
        return this;
    }
    
    public IntervalVariableBuilder options(List<VariableOption> options) {
    this.internal.options = options;
        return this;
    }
    public VariableModel build() {
            return this.internal;
        }
    }
    
    public static class TextBoxVariableBuilder implements com.grafana.foundation.cog.Builder<VariableModel> {
        private final VariableModel internal;
        
        public TextBoxVariableBuilder(String name) {
            this.internal = new VariableModel();
    this.internal.name = name;
    this.internal.type = VariableType.TEXTBOX;
        }
    public TextBoxVariableBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public TextBoxVariableBuilder label(String label) {
    this.internal.label = label;
        return this;
    }
    
    public TextBoxVariableBuilder hide(VariableHide hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public TextBoxVariableBuilder description(String description) {
    this.internal.description = description;
        return this;
    }
    
    public TextBoxVariableBuilder defaultValue(StringOrMap query) {
    this.internal.query = query;
        return this;
    }
    
    public TextBoxVariableBuilder current(VariableOption current) {
    this.internal.current = current;
        return this;
    }
    
    public TextBoxVariableBuilder options(List<VariableOption> options) {
    this.internal.options = options;
        return this;
    }
    public VariableModel build() {
            return this.internal;
        }
    }
    
    public static class CustomVariableBuilder implements com.grafana.foundation.cog.Builder<VariableModel> {
        private final VariableModel internal;
        
        public CustomVariableBuilder(String name) {
            this.internal = new VariableModel();
    this.internal.name = name;
    this.internal.type = VariableType.CUSTOM;
        }
    public CustomVariableBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public CustomVariableBuilder label(String label) {
    this.internal.label = label;
        return this;
    }
    
    public CustomVariableBuilder hide(VariableHide hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public CustomVariableBuilder description(String description) {
    this.internal.description = description;
        return this;
    }
    
    public CustomVariableBuilder values(StringOrMap query) {
    this.internal.query = query;
        return this;
    }
    
    public CustomVariableBuilder current(VariableOption current) {
    this.internal.current = current;
        return this;
    }
    
    public CustomVariableBuilder multi(Boolean multi) {
    this.internal.multi = multi;
        return this;
    }
    
    public CustomVariableBuilder options(List<VariableOption> options) {
    this.internal.options = options;
        return this;
    }
    
    public CustomVariableBuilder includeAll(Boolean includeAll) {
    this.internal.includeAll = includeAll;
        return this;
    }
    
    public CustomVariableBuilder allValue(String allValue) {
    this.internal.allValue = allValue;
        return this;
    }
    public VariableModel build() {
            return this.internal;
        }
    }
}
