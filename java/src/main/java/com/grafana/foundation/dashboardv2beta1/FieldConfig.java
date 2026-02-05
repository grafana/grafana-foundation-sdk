// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.
// Each column within this structure is called a field. A field can represent a single time series or table column.
// Field options allow you to change how the data is displayed in your visualizations.
public class FieldConfig {
    // The display value for this field.  This supports template variables blank is auto
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("displayName")
    public String displayName;
    // This can be used by data sources that return and explicit naming structure for values and labels
    // When this property is configured, this value is used rather than the default naming strategy.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("displayNameFromDS")
    public String displayNameFromDS;
    // Human readable field metadata
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("description")
    public String description;
    // An explicit path to the field in the datasource.  When the frame meta includes a path,
    // This will default to `${frame.meta.path}/${field.name}
    // 
    // When defined, this value can be used as an identifier within the datasource scope, and
    // may be used to update the results
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("path")
    public String path;
    // True if data source can write a value to the path. Auth/authz are supported separately
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("writeable")
    public Boolean writeable;
    // True if data source field supports ad-hoc filters
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("filterable")
    public Boolean filterable;
    // Unit a field should use. The unit you select is applied to all fields except time.
    // You can use the units ID availables in Grafana or a custom unit.
    // Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts
    // As custom unit, you can use the following formats:
    // `suffix:<suffix>` for custom unit that should go after value.
    // `prefix:<prefix>` for custom unit that should go before value.
    // `time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.
    // `si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.
    // `count:<unit>` for a custom count unit.
    // `currency:<unit>` for custom a currency unit.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("unit")
    public String unit;
    // Specify the number of decimals Grafana includes in the rendered value.
    // If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
    // For example 1.1234 will display as 1.12 and 100.456 will display as 100.
    // To display all decimals, set the unit to `String`.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("decimals")
    public Double decimals;
    // The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("min")
    public Double min;
    // The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("max")
    public Double max;
    // Convert input values into a display string
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mappings")
    public List<ValueMapping> mappings;
    // Map numeric values to states
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("thresholds")
    public ThresholdsConfig thresholds;
    // Panel color configuration
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("color")
    public FieldColor color;
    // The behavior when clicking on a result
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("links")
    public List<Object> links;
    // Define interactive HTTP requests that can be triggered from data visualizations.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("actions")
    public List<Action> actions;
    // Alternative to empty string
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("noValue")
    public String noValue;
    // custom is specified by the FieldConfig field
    // in panel plugin schemas.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("custom")
    public Object custom;
    public FieldConfig() {
    }
    public FieldConfig(String displayName,String displayNameFromDS,String description,String path,Boolean writeable,Boolean filterable,String unit,Double decimals,Double min,Double max,List<ValueMapping> mappings,ThresholdsConfig thresholds,FieldColor color,List<Object> links,List<Action> actions,String noValue,Object custom) {
        this.displayName = displayName;
        this.displayNameFromDS = displayNameFromDS;
        this.description = description;
        this.path = path;
        this.writeable = writeable;
        this.filterable = filterable;
        this.unit = unit;
        this.decimals = decimals;
        this.min = min;
        this.max = max;
        this.mappings = mappings;
        this.thresholds = thresholds;
        this.color = color;
        this.links = links;
        this.actions = actions;
        this.noValue = noValue;
        this.custom = custom;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
