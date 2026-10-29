// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

/**
 * #FacetValue is a single value/count pair in a facet breakdown.
 */
public class FacetValue {
    @JsonProperty("value")
    public String value;
    @JsonProperty("count")
    public Long count;
    public FacetValue() {
        this.value = "";
        this.count = 0L;
    }
    public FacetValue(String value,Long count) {
        this.value = value;
        this.count = count;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
