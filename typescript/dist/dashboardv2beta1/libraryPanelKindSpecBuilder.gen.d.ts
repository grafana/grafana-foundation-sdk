import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class LibraryPanelKindSpecBuilder implements cog.Builder<dashboardv2beta1.LibraryPanelKindSpec> {
    protected readonly internal: dashboardv2beta1.LibraryPanelKindSpec;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.LibraryPanelKindSpec;
    id(id: number): this;
    title(title: string): this;
    libraryPanel(libraryPanel: cog.Builder<dashboardv2beta1.LibraryPanelRef>): this;
}
