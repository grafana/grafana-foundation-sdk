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

// The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.
// Each column within this structure is called a field. A field can represent a single time series or table column.
// Field options allow you to change how the data is displayed in your visualizations.
public class FieldConfigSource {
    // Defaults are the options applied to all fields.
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("defaults")
    public FieldConfig defaults;
    // Overrides are the options applied to specific fields overriding the defaults.
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("overrides")
    public List<Dashboardv2beta1FieldConfigSourceOverrides> overrides;
    public FieldConfigSource() {
        this.defaults = new com.grafana.foundation.dashboardv2beta1.FieldConfig();
        this.overrides = new LinkedList<>();
    }
    public FieldConfigSource(FieldConfig defaults,List<Dashboardv2beta1FieldConfigSourceOverrides> overrides) {
        this.defaults = defaults;
        this.overrides = overrides;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
