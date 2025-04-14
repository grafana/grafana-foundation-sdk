// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.news;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Options {
    // empty/missing will default to grafana blog
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("feedUrl")
    public String feedUrl;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showImage")
    public Boolean showImage;
    public Options() {
        this.showImage = true;
    }
    public Options(String feedUrl,Boolean showImage) {
        this.feedUrl = feedUrl;
        this.showImage = showImage;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
