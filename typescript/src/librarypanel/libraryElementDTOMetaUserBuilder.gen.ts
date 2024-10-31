// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as librarypanel from '../librarypanel';

export class LibraryElementDTOMetaUserBuilder implements cog.Builder<librarypanel.LibraryElementDTOMetaUser> {
    protected readonly internal: librarypanel.LibraryElementDTOMetaUser;

    constructor() {
        this.internal = librarypanel.defaultLibraryElementDTOMetaUser();
    }

    /**
     * Builds the object.
     */
    build(): librarypanel.LibraryElementDTOMetaUser {
        return this.internal;
    }

    id(id: number): this {
        this.internal.id = id;
        return this;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    avatarUrl(avatarUrl: string): this {
        this.internal.avatarUrl = avatarUrl;
        return this;
    }
}
