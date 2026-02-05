// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class ExprTypeThresholdResultAssertions {
    // Maximum frame count
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxFrames")
    public Long maxFrames;
    // Type asserts that the frame matches a known type structure.
    // Possible enum values:
    //  - `""` 
    //  - `"timeseries-wide"` 
    //  - `"timeseries-long"` 
    //  - `"timeseries-many"` 
    //  - `"timeseries-multi"` 
    //  - `"directory-listing"` 
    //  - `"table"` 
    //  - `"numeric-wide"` 
    //  - `"numeric-multi"` 
    //  - `"numeric-long"` 
    //  - `"log-lines"` 
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("type")
    public ExprTypeThresholdResultAssertionsType type;
    // TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
    // contract documentation https://grafana.github.io/dataplane/contract/.
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("typeVersion")
    public List<Long> typeVersion;
    public ExprTypeThresholdResultAssertions() {
        this.typeVersion = new LinkedList<>();
    }
    public ExprTypeThresholdResultAssertions(Long maxFrames,ExprTypeThresholdResultAssertionsType type,List<Long> typeVersion) {
        this.maxFrames = maxFrames;
        this.type = type;
        this.typeVersion = typeVersion;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
