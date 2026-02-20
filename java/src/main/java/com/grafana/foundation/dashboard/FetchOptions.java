// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// Fetch options
public class FetchOptions {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("method")
    public HttpRequestMethod method;
    @JsonProperty("url")
    public String url;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("body")
    public String body;
    // These are 2D arrays of strings, each representing a key-value pair
    // We are defining this way because we can't generate a go struct that
    // that would have exactly two strings in each sub-array
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryParams")
    public List<List<String>> queryParams;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("headers")
    public List<List<String>> headers;
    public FetchOptions() {
        this.method = HttpRequestMethod.GET;
        this.url = "";
    }
    public FetchOptions(HttpRequestMethod method,String url,String body,List<List<String>> queryParams,List<List<String>> headers) {
        this.method = method;
        this.url = url;
        this.body = body;
        this.queryParams = queryParams;
        this.headers = headers;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
