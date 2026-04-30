// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class LibraryPanelBuilder implements cog.Builder<dashboardv2.LibraryPanelKind> {
    protected readonly internal: dashboardv2.LibraryPanelKind;

    constructor() {
        this.internal = dashboardv2.defaultLibraryPanelKind();
        this.internal.kind = "LibraryPanel";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.LibraryPanelKind {
        return this.internal;
    }

    // Panel ID for the library panel in the dashboard
    id(id: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultLibraryPanelKindSpec();
        }
        if (!(id >= 0)) {
            throw new Error("id must be >= 0");
        }
        this.internal.spec.id = id;
        return this;
    }

    // Title for the library panel in the dashboard
    title(title: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultLibraryPanelKindSpec();
        }
        this.internal.spec.title = title;
        return this;
    }

    libraryPanel(libraryPanel: cog.Builder<dashboardv2.LibraryPanelRef>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultLibraryPanelKindSpec();
        }
        const libraryPanelResource = libraryPanel.build();
        this.internal.spec.libraryPanel = libraryPanelResource;
        return this;
    }
}

