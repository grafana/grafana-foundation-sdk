// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;
import java.util.Map;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;

// A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
public class NotificationPolicy {
    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("continue")
    public Boolean continueArg;
    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("group_by")
    public List<String> groupBy;
    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("group_interval")
    public String groupInterval;
    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("group_wait")
    public String groupWait;
    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("match")
    public Map<String, String> match;
    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("match_re")
    public Map<String, Object> matchRe;
    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("matchers")
    public List<Matcher> matchers;
    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mute_time_intervals")
    public List<String> muteTimeIntervals;
    // Matchers is a slice of Matchers that is sortable, implements Stringer, and
    // provides a Matches method to match a LabelSet against all Matchers in the
    // slice. Note that some users of Matchers might require it to be sorted.
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("object_matchers")
    public List<Matcher> objectMatchers;
    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("provenance")
    public String provenance;
    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("receiver")
    public String receiver;
    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("repeat_interval")
    public String repeatInterval;
    // A Route is a node that contains definitions of how to handle alerts. This is modified
    // from the upstream alertmanager in that it adds the ObjectMatchers property.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("routes")
    public List<NotificationPolicy> routes;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<NotificationPolicy> {
        protected final NotificationPolicy internal;
        
        public Builder() {
            this.internal = new NotificationPolicy();
        }
    public Builder continueArg(Boolean continueArg) {
    this.internal.continueArg = continueArg;
        return this;
    }
    
    public Builder groupBy(List<String> groupBy) {
    this.internal.groupBy = groupBy;
        return this;
    }
    
    public Builder groupInterval(String groupInterval) {
    this.internal.groupInterval = groupInterval;
        return this;
    }
    
    public Builder groupWait(String groupWait) {
    this.internal.groupWait = groupWait;
        return this;
    }
    
    public Builder match(Map<String, String> match) {
    this.internal.match = match;
        return this;
    }
    
    public Builder matchRe(Map<String, Object> matchRe) {
    this.internal.matchRe = matchRe;
        return this;
    }
    
    public Builder matchers(List<Matcher> matchers) {
    this.internal.matchers = matchers;
        return this;
    }
    
    public Builder muteTimeIntervals(List<String> muteTimeIntervals) {
    this.internal.muteTimeIntervals = muteTimeIntervals;
        return this;
    }
    
    public Builder objectMatchers(com.grafana.foundation.cog.Builder<List<Matcher>> objectMatchers) {
    this.internal.objectMatchers = objectMatchers.build();
        return this;
    }
    
    public Builder provenance(String provenance) {
    this.internal.provenance = provenance;
        return this;
    }
    
    public Builder receiver(String receiver) {
    this.internal.receiver = receiver;
        return this;
    }
    
    public Builder repeatInterval(String repeatInterval) {
    this.internal.repeatInterval = repeatInterval;
        return this;
    }
    
    public Builder routes(com.grafana.foundation.cog.Builder<List<NotificationPolicy>> routes) {
    this.internal.routes = routes.build();
        return this;
    }
    public NotificationPolicy build() {
            return this.internal;
        }
    }
}
