// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ElasticsearchTermsSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("order")
    public TermsOrder order;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("size")
    public String size;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("min_doc_count")
    public String minDocCount;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("orderBy")
    public String orderBy;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("missing")
    public String missing;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchTermsSettings> {
        private final ElasticsearchTermsSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchTermsSettings();
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
    public ElasticsearchTermsSettings build() {
            return this.internal;
        }
    }
}
