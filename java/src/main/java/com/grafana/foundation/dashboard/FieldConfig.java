// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.
// Each column within this structure is called a field. A field can represent a single time series or table column.
// Field options allow you to change how the data is displayed in your visualizations.
public class FieldConfig {
    // The display value for this field.  This supports template variables blank is auto 
    @JsonProperty("displayName")
    public String displayName;
    // This can be used by data sources that return and explicit naming structure for values and labels
    // When this property is configured, this value is used rather than the default naming strategy. 
    @JsonProperty("displayNameFromDS")
    public String displayNameFromDS;
    // Human readable field metadata 
    @JsonProperty("description")
    public String description;
    // An explicit path to the field in the datasource.  When the frame meta includes a path,
    // This will default to `${frame.meta.path}/${field.name}
    // 
    // When defined, this value can be used as an identifier within the datasource scope, and
    // may be used to update the results 
    @JsonProperty("path")
    public String path;
    // True if data source can write a value to the path. Auth/authz are supported separately 
    @JsonProperty("writeable")
    public Boolean writeable;
    // True if data source field supports ad-hoc filters 
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
    @JsonProperty("unit")
    public String unit;
    // Specify the number of decimals Grafana includes in the rendered value.
    // If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
    // For example 1.1234 will display as 1.12 and 100.456 will display as 100.
    // To display all decimals, set the unit to `String`. 
    @JsonProperty("decimals")
    public Double decimals;
    // The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields. 
    @JsonProperty("min")
    public Double min;
    // The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields. 
    @JsonProperty("max")
    public Double max;
    // Convert input values into a display string 
    @JsonProperty("mappings")
    public List<ValueMapping> mappings;
    // Map numeric values to states 
    @JsonProperty("thresholds")
    public ThresholdsConfig thresholds;
    // Panel color configuration 
    @JsonProperty("color")
    public FieldColor color;
    // The behavior when clicking on a result 
    @JsonProperty("links")
    public List<Object> links;
    // Alternative to empty string 
    @JsonProperty("noValue")
    public String noValue;
    // custom is specified by the FieldConfig field
    // in panel plugin schemas. 
    @JsonProperty("custom")
    public Object custom;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
