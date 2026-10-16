// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class V2PanelSpecBuilder implements cog.Builder<dashboardv2beta1.V2PanelSpec> {
    protected readonly internal: dashboardv2beta1.V2PanelSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultV2PanelSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.V2PanelSpec {
        return this.internal;
    }

    id(id: number): this {
        this.internal.id = id;
        return this;
    }

    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    // Shown in a info icon tooltip next to panel title
    description(description: string): this {
        this.internal.description = description;
        return this;
    }

    // Shown in a sub header below the title.
    subtitle(subtitle: string): this {
        this.internal.subtitle = subtitle;
        return this;
    }

    links(links: cog.Builder<dashboardv2beta1.DataLink>[]): this {
        const linksResources = links.map(builder1 => builder1.build());
        this.internal.links = linksResources;
        return this;
    }

    data(data: cog.Builder<dashboardv2beta1.V2QueryGroupKind>): this {
        const dataResource = data.build();
        this.internal.data = dataResource;
        return this;
    }

    vizConfig(vizConfig: cog.Builder<dashboardv2beta1.VizConfigKind>): this {
        const vizConfigResource = vizConfig.build();
        this.internal.vizConfig = vizConfigResource;
        return this;
    }

    transparent(transparent: boolean): this {
        this.internal.transparent = transparent;
        return this;
    }
}

