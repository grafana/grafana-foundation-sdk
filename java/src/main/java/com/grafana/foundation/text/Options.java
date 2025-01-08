// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.text;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Options {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("mode")
    public TextMode mode;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("code")
    public CodeOptions code;
    @JsonProperty("content")
    public String content;
    public Options() {
        this.mode = TextMode.MARKDOWN;
        this.content = "# Title\n\nFor markdown syntax help: [commonmark.org/help](https://commonmark.org/help/)";
    }
    
    public Options(TextMode mode,CodeOptions code,String content) {
        this.mode = mode;
        this.code = code;
        this.content = content;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
