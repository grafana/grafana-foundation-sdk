// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';

/**
 * #SearchResultResource is the full identity of a hit.
 */
export class SearchResultResourceBuilder implements cog.Builder<alertingv0alpha1.SearchResultResource> {
    protected readonly internal: alertingv0alpha1.SearchResultResource;

    constructor() {
        this.internal = alertingv0alpha1.defaultSearchResultResource();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.SearchResultResource {
        return this.internal;
    }

    group(group: string): this {
        this.internal.group = group;
        return this;
    }

    resource(resource: string): this {
        this.internal.resource = resource;
        return this;
    }

    kind(kind: string): this {
        this.internal.kind = kind;
        return this;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }
}

