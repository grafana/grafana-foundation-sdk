// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// RelativeTimeRange is the per query start and end time
// for requests.
public class RelativeTimeRange {
    // A Duration represents the elapsed time between two instants
    // as an int64 nanosecond count. The representation limits the
    // largest representable duration to approximately 290 years.
    @JsonProperty("from")
    public Long from;
    // A Duration represents the elapsed time between two instants
    // as an int64 nanosecond count. The representation limits the
    // largest representable duration to approximately 290 years.
    @JsonProperty("to")
    public Long to;
    public RelativeTimeRange() {
    }
    public RelativeTimeRange(Long from,Long to) {
        this.from = from;
        this.to = to;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
