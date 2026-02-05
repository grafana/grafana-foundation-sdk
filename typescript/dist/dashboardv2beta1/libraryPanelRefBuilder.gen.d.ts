import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class LibraryPanelRefBuilder implements cog.Builder<dashboardv2beta1.LibraryPanelRef> {
    protected readonly internal: dashboardv2beta1.LibraryPanelRef;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.LibraryPanelRef;
    name(name: string): this;
    uid(uid: string): this;
}
