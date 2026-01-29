// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class PanelBuilder implements cog.Builder<dashboardv2beta1.PanelKind> {
    protected readonly internal: dashboardv2beta1.PanelKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultPanelKind();
        this.internal.kind = "Panel";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.PanelKind {
        return this.internal;
    }

    id(id: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelSpec();
        }
        this.internal.spec.id = id;
        return this;
    }

    title(title: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelSpec();
        }
        this.internal.spec.title = title;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelSpec();
        }
        this.internal.spec.description = description;
        return this;
    }

    links(links: cog.Builder<dashboardv2beta1.DataLink>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelSpec();
        }
        const linksResources = links.map(builder1 => builder1.build());
        this.internal.spec.links = linksResources;
        return this;
    }

    data(data: cog.Builder<dashboardv2beta1.QueryGroupKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelSpec();
        }
        const dataResource = data.build();
        this.internal.spec.data = dataResource;
        return this;
    }

    visualization(vizConfig: cog.Builder<dashboardv2beta1.VizConfigKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelSpec();
        }
        const vizConfigResource = vizConfig.build();
        this.internal.spec.vizConfig = vizConfigResource;
        return this;
    }

    transparent(transparent: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelSpec();
        }
        this.internal.spec.transparent = transparent;
        return this;
    }
}

