import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class LibraryPanelBuilder implements cog.Builder<dashboardv2beta1.LibraryPanelKind> {
    protected readonly internal: dashboardv2beta1.LibraryPanelKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.LibraryPanelKind;
    spec(spec: cog.Builder<dashboardv2beta1.LibraryPanelKindSpec>): this;
    id(id: number): this;
    title(title: string): this;
    libraryPanel(libraryPanel: cog.Builder<dashboardv2beta1.LibraryPanelRef>): this;
}
