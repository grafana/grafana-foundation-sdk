// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.librarypanel;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class LibraryElementDTOMetaUser {
    @JsonProperty("id")
    public Long id;
    @JsonProperty("name")
    public String name;
    @JsonProperty("avatarUrl")
    public String avatarUrl;
    public LibraryElementDTOMetaUser() {
        this.id = 0L;
        this.name = "";
        this.avatarUrl = "";
    }
    public LibraryElementDTOMetaUser(Long id,String name,String avatarUrl) {
        this.id = id;
        this.name = name;
        this.avatarUrl = avatarUrl;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
