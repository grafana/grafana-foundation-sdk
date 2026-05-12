// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Annotation event field mapping. Defines how to map a data frame field to an annotation event field.
public class AnnotationEventFieldMapping {
    // Source type for the field value
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("source")
    public String source;
    // Constant value to use when source is "text"
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("value")
    public String value;
    // Regular expression to apply to the field value
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("regex")
    public String regex;
    public AnnotationEventFieldMapping() {
        this.source = "field";
    }
    public AnnotationEventFieldMapping(String source,String value,String regex) {
        this.source = source;
        this.value = value;
        this.regex = regex;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
