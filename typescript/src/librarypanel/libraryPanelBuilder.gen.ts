// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as librarypanel from '../librarypanel';

export class LibraryPanelBuilder implements cog.Builder<librarypanel.LibraryPanel> {
    protected readonly internal: librarypanel.LibraryPanel;

    constructor() {
        this.internal = librarypanel.defaultLibraryPanel();
    }

    /**
     * Builds the object.
     */
    build(): librarypanel.LibraryPanel {
        return this.internal;
    }

    // Folder UID
    folderUid(folderUid: string): this {
        this.internal.folderUid = folderUid;
        return this;
    }

    // Library element UID
    uid(uid: string): this {
        this.internal.uid = uid;
        return this;
    }

    // Panel name (also saved in the model)
    name(name: string): this {
        if (!(name.length >= 1)) {
            throw new Error("name.length must be >= 1");
        }
        this.internal.name = name;
        return this;
    }

    // Panel description
    description(description: string): this {
        this.internal.description = description;
        return this;
    }

    // The panel type (from inside the model)
    type(type: string): this {
        if (!(type.length >= 1)) {
            throw new Error("type.length must be >= 1");
        }
        this.internal.type = type;
        return this;
    }

    // Dashboard version when this was saved (zero if unknown)
    schemaVersion(schemaVersion: number): this {
        this.internal.schemaVersion = schemaVersion;
        return this;
    }

    // panel version, incremented each time the dashboard is updated.
    version(version: number): this {
        this.internal.version = version;
        return this;
    }

    // TODO: should be the same panel schema defined in dashboard
    // Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;
    model(model: cog.Builder<librarypanel.PanelModel>): this {
        const modelResource = model.build();
        this.internal.model = modelResource;
        return this;
    }

    // Object storage metadata
    meta(meta: cog.Builder<librarypanel.LibraryElementDTOMeta>): this {
        const metaResource = meta.build();
        this.internal.meta = metaResource;
        return this;
    }
}

