// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// Transformations allow to manipulate data returned by a query before the system applies a visualization.
// Using transformations you can: rename fields, join time series data, perform mathematical operations across queries,
// use the output of one transformation as the input to another transformation, etc.
public class DataTransformerConfig {
    // Unique identifier of transformer
    @JsonProperty("id")
    public String id;
    // Disabled transformations are skipped
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("disabled")
    public Boolean disabled;
    // Optional frame matcher. When missing it will be applied to all results
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("filter")
    public MatcherConfig filter;
    // Options to be passed to the transformer
    // Valid options depend on the transformer id
    @JsonProperty("options")
    public Object options;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
