// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class LibraryPanelKindSpecBuilder implements cog.Builder<dashboardv2beta1.LibraryPanelKindSpec> {
    protected readonly internal: dashboardv2beta1.LibraryPanelKindSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultLibraryPanelKindSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.LibraryPanelKindSpec {
        return this.internal;
    }

    // Panel ID for the library panel in the dashboard
    id(id: number): this {
        this.internal.id = id;
        return this;
    }

    // Title for the library panel in the dashboard
    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    libraryPanel(libraryPanel: cog.Builder<dashboardv2beta1.LibraryPanelRef>): this {
        const libraryPanelResource = libraryPanel.build();
        this.internal.libraryPanel = libraryPanelResource;
        return this;
    }
}

