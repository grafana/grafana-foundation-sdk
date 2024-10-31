// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as librarypanel from '../librarypanel';

export class LibraryElementDTOMetaBuilder implements cog.Builder<librarypanel.LibraryElementDTOMeta> {
    protected readonly internal: librarypanel.LibraryElementDTOMeta;

    constructor() {
        this.internal = librarypanel.defaultLibraryElementDTOMeta();
    }

    /**
     * Builds the object.
     */
    build(): librarypanel.LibraryElementDTOMeta {
        return this.internal;
    }

    folderName(folderName: string): this {
        this.internal.folderName = folderName;
        return this;
    }

    folderUid(folderUid: string): this {
        this.internal.folderUid = folderUid;
        return this;
    }

    connectedDashboards(connectedDashboards: number): this {
        this.internal.connectedDashboards = connectedDashboards;
        return this;
    }

    created(created: string): this {
        this.internal.created = created;
        return this;
    }

    updated(updated: string): this {
        this.internal.updated = updated;
        return this;
    }

    createdBy(createdBy: cog.Builder<librarypanel.LibraryElementDTOMetaUser>): this {
        const createdByResource = createdBy.build();
        this.internal.createdBy = createdByResource;
        return this;
    }

    updatedBy(updatedBy: cog.Builder<librarypanel.LibraryElementDTOMetaUser>): this {
        const updatedByResource = updatedBy.build();
        this.internal.updatedBy = updatedByResource;
        return this;
    }
}
