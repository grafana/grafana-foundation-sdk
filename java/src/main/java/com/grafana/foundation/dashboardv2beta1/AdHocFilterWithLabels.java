// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// Define the AdHocFilterWithLabels type
public class AdHocFilterWithLabels {
    @JsonProperty("key")
    public String key;
    @JsonProperty("operator")
    public String operator;
    @JsonProperty("value")
    public String value;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("values")
    public List<String> values;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("keyLabel")
    public String keyLabel;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("valueLabels")
    public List<String> valueLabels;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("forceEdit")
    public Boolean forceEdit;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("origin")
    public String origin;
    // @deprecated
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("condition")
    public String condition;
    public AdHocFilterWithLabels() {
        this.key = "";
        this.operator = "";
        this.value = "";
        this.origin = Constants.FilterOrigin;
    }
    public AdHocFilterWithLabels(String key,String operator,String value,List<String> values,String keyLabel,List<String> valueLabels,Boolean forceEdit,String condition) {
        this.key = key;
        this.operator = operator;
        this.value = value;
        this.values = values;
        this.keyLabel = keyLabel;
        this.valueLabels = valueLabels;
        this.forceEdit = forceEdit;
        this.origin = Constants.FilterOrigin;
        this.condition = condition;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
