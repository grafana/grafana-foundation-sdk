// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.librarypanel;


public class LibraryElementDTOMetaUserBuilder implements com.grafana.foundation.cog.Builder<LibraryElementDTOMetaUser> {
    protected final LibraryElementDTOMetaUser internal;
    
    public LibraryElementDTOMetaUserBuilder() {
        this.internal = new LibraryElementDTOMetaUser();
    }
    public LibraryElementDTOMetaUserBuilder id(Long id) {
    this.internal.id = id;
        return this;
    }
    
    public LibraryElementDTOMetaUserBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public LibraryElementDTOMetaUserBuilder avatarUrl(String avatarUrl) {
    this.internal.avatarUrl = avatarUrl;
        return this;
    }
    public LibraryElementDTOMetaUser build() {
        return this.internal;
    }
}
