// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class PanelBuilder implements cog.Builder<dashboardv2.PanelKind> {
    protected readonly internal: dashboardv2.PanelKind;

    constructor() {
        this.internal = dashboardv2.defaultPanelKind();
        this.internal.kind = "Panel";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.PanelKind {
        return this.internal;
    }

    id(id: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultPanelSpec();
        }
        if (!(id >= 0)) {
            throw new Error("id must be >= 0");
        }
        this.internal.spec.id = id;
        return this;
    }

    title(title: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultPanelSpec();
        }
        this.internal.spec.title = title;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultPanelSpec();
        }
        this.internal.spec.description = description;
        return this;
    }

    links(links: cog.Builder<dashboardv2.DataLink>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultPanelSpec();
        }
        const linksResources = links.map(builder1 => builder1.build());
        this.internal.spec.links = linksResources;
        return this;
    }

    data(data: cog.Builder<dashboardv2.QueryGroupKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultPanelSpec();
        }
        const dataResource = data.build();
        this.internal.spec.data = dataResource;
        return this;
    }

    visualization(vizConfig: cog.Builder<dashboardv2.VizConfigKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultPanelSpec();
        }
        const vizConfigResource = vizConfig.build();
        this.internal.spec.vizConfig = vizConfigResource;
        return this;
    }

    transparent(transparent: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultPanelSpec();
        }
        this.internal.spec.transparent = transparent;
        return this;
    }
}

