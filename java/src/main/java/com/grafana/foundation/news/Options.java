// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.news;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class Options {
    // empty/missing will default to grafana blog 
    @JsonProperty("feedUrl")
    public String feedUrl; 
    @JsonProperty("showImage")
    public Boolean showImage;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
