// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';

/**
 * #SearchResults is the response envelope, mirroring
 * search.grafana.app SearchResults.
 */
export class SearchResultsBuilder implements cog.Builder<alertingv0alpha1.SearchResults> {
    protected readonly internal: alertingv0alpha1.SearchResults;

    constructor() {
        this.internal = alertingv0alpha1.defaultSearchResults();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.SearchResults {
        return this.internal;
    }

    metadata(metadata: cog.Builder<alertingv0alpha1.SearchResultsMetadata>): this {
        const metadataResource = metadata.build();
        this.internal.metadata = metadataResource;
        return this;
    }

    items(items: cog.Builder<alertingv0alpha1.SearchResultHit>[]): this {
        const itemsResources = items.map(builder1 => builder1.build());
        this.internal.items = itemsResources;
        return this;
    }

    facets(facets: Record<string, cog.Builder<alertingv0alpha1.FacetValue>[]>): this {
        const facetsResource = (function() {
            let results1: Record<string, alertingv0alpha1.FacetValue[]> = {};
            for (const key1 in facets) {
                const val1 = facets[key1];
                results1[key1] = val1.map(builder2 => builder2.build());
            }
            return results1;
        }());
        this.internal.facets = facetsResource;
        return this;
    }
}

