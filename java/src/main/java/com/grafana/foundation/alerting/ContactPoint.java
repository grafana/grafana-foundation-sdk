// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
public class ContactPoint {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("disableResolveMessage")
    public Boolean disableResolveMessage;
    // Name is used as grouping key in the UI. Contact points with the
    // same name will be grouped in the UI.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("name")
    public String name;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("provenance")
    public String provenance;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("settings")
    public Object settings;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public ContactPointType type;
    // UID is the unique identifier of the contact point. The UID can be
    // set by the user.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("uid")
    public String uid;
    public ContactPoint() {
    }
    public ContactPoint(Boolean disableResolveMessage,String name,String provenance,Object settings,ContactPointType type,String uid) {
        this.disableResolveMessage = disableResolveMessage;
        this.name = name;
        this.provenance = provenance;
        this.settings = settings;
        this.type = type;
        this.uid = uid;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
