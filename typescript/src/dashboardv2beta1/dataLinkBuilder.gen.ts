// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class DataLinkBuilder implements cog.Builder<dashboardv2beta1.DataLink> {
    protected readonly internal: dashboardv2beta1.DataLink;

    constructor() {
        this.internal = dashboardv2beta1.defaultDataLink();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataLink {
        return this.internal;
    }

    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    url(url: string): this {
        this.internal.url = url;
        return this;
    }

    targetBlank(targetBlank: boolean): this {
        this.internal.targetBlank = targetBlank;
        return this;
    }
}

