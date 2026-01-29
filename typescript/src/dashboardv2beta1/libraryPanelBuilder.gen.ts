// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class LibraryPanelBuilder implements cog.Builder<dashboardv2beta1.LibraryPanelKind> {
    protected readonly internal: dashboardv2beta1.LibraryPanelKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultLibraryPanelKind();
        this.internal.kind = "LibraryPanel";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.LibraryPanelKind {
        return this.internal;
    }

    spec(spec: cog.Builder<dashboardv2beta1.LibraryPanelKindSpec>): this {
        const specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }

    // Panel ID for the library panel in the dashboard
    id(id: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultLibraryPanelKindSpec();
        }
        this.internal.spec.id = id;
        return this;
    }

    // Title for the library panel in the dashboard
    title(title: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultLibraryPanelKindSpec();
        }
        this.internal.spec.title = title;
        return this;
    }

    libraryPanel(libraryPanel: cog.Builder<dashboardv2beta1.LibraryPanelRef>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultLibraryPanelKindSpec();
        }
        const libraryPanelResource = libraryPanel.build();
        this.internal.spec.libraryPanel = libraryPanelResource;
        return this;
    }
}

