import * as cog from '../cog';
import * as librarypanel from '../librarypanel';
export declare class LibraryElementDTOMetaUserBuilder implements cog.Builder<librarypanel.LibraryElementDTOMetaUser> {
    protected readonly internal: librarypanel.LibraryElementDTOMetaUser;
    constructor();
    /**
     * Builds the object.
     */
    build(): librarypanel.LibraryElementDTOMetaUser;
    id(id: number): this;
    name(name: string): this;
    avatarUrl(avatarUrl: string): this;
}
