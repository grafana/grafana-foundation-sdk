// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as starsv1alpha1 from '../starsv1alpha1';
import * as resource from '../resource';

export class StarsBuilder implements cog.Builder<starsv1alpha1.Stars> {
    protected readonly internal: starsv1alpha1.Stars;

    constructor() {
        this.internal = starsv1alpha1.defaultStars();
    }

    /**
     * Builds the object.
     */
    build(): starsv1alpha1.Stars {
        return this.internal;
    }

    resources(resource: cog.Builder<starsv1alpha1.Resource>[]): this {
        const resourceResources = resource.map(builder1 => builder1.build());
        this.internal.resource = resourceResources;
        return this;
    }

    resource(resource: cog.Builder<starsv1alpha1.Resource>): this {
        if (!this.internal.resource) {
            this.internal.resource = [];
        }
        const resourceResource = resource.build();
        this.internal.resource.push(resourceResource);
        return this;
    }
}

/**
 * Creates a resource manifest from Stars.
 */
export function manifest(name: string,stars: cog.Builder<starsv1alpha1.Stars>): cog.Builder<resource.Manifest> {
	const builder = new resource.ManifestBuilder();
	builder.apiVersion("preferences.grafana.app/v1alpha1");
	builder.metadata(resource.named(name));
	builder.kind("Stars");
	builder.spec(stars.build());

	return builder;
}

