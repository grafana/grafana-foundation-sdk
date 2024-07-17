// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class TermsSettings { 
    @JsonProperty("order")
    public TermsOrder order; 
    @JsonProperty("size")
    public String size; 
    @JsonProperty("min_doc_count")
    public String minDocCount; 
    @JsonProperty("orderBy")
    public String orderBy; 
    @JsonProperty("missing")
    public String missing;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TermsSettings> {
        private final TermsSettings internal;
        
        public Builder() {
            this.internal = new TermsSettings();
        }
    public Builder order(TermsOrder order) {
    this.internal.order = order;
        return this;
    }
    
    public Builder size(String size) {
    this.internal.size = size;
        return this;
    }
    
    public Builder minDocCount(String minDocCount) {
    this.internal.minDocCount = minDocCount;
        return this;
    }
    
    public Builder orderBy(String orderBy) {
    this.internal.orderBy = orderBy;
        return this;
    }
    
    public Builder missing(String missing) {
    this.internal.missing = missing;
        return this;
    }
    public TermsSettings build() {
            return this.internal;
        }
    }
}
