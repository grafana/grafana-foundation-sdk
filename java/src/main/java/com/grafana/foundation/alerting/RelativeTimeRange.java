// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// RelativeTimeRange is the per query start and end time
// for requests.
public class RelativeTimeRange {
    // RelativeTimeRange is the per query start and end time
    // for requests. 
    @JsonProperty("from")
    public Long from;
    // RelativeTimeRange is the per query start and end time
    // for requests. 
    @JsonProperty("to")
    public Long to;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
