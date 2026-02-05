import * as cog from '../cog';
import * as librarypanel from '../librarypanel';
export declare class LibraryPanelBuilder implements cog.Builder<librarypanel.LibraryPanel> {
    protected readonly internal: librarypanel.LibraryPanel;
    constructor();
    /**
     * Builds the object.
     */
    build(): librarypanel.LibraryPanel;
    folderUid(folderUid: string): this;
    uid(uid: string): this;
    name(name: string): this;
    description(description: string): this;
    type(type: string): this;
    schemaVersion(schemaVersion: number): this;
    version(version: number): this;
    model(model: cog.Builder<librarypanel.PanelModel>): this;
    meta(meta: cog.Builder<librarypanel.LibraryElementDTOMeta>): this;
}
