// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class AnnotationQuerySpecBuilder implements cog.Builder<dashboardv2beta1.AnnotationQuerySpec> {
    protected readonly internal: dashboardv2beta1.AnnotationQuerySpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultAnnotationQuerySpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AnnotationQuerySpec {
        return this.internal;
    }

    query(query: cog.Builder<dashboardv2beta1.DataQueryKind>): this {
        const queryResource = query.build();
        this.internal.query = queryResource;
        return this;
    }

    enable(enable: boolean): this {
        this.internal.enable = enable;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }

    iconColor(iconColor: string): this {
        this.internal.iconColor = iconColor;
        return this;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    builtIn(builtIn: boolean): this {
        this.internal.builtIn = builtIn;
        return this;
    }

    filter(filter: cog.Builder<dashboardv2beta1.AnnotationPanelFilter>): this {
        const filterResource = filter.build();
        this.internal.filter = filterResource;
        return this;
    }

    // Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
    placement(placement: "inControlsMenu"): this {
        this.internal.placement = placement;
        return this;
    }

    // Mappings define how to convert data frame fields to annotation event fields.
    mappings(mappings: Record<string, cog.Builder<dashboardv2beta1.AnnotationEventFieldMapping>>): this {
        const mappingsResource = (function() {
            let results1: Record<string, dashboardv2beta1.AnnotationEventFieldMapping> = {};
            for (const key1 in mappings) {
                const val1 = mappings[key1];
                results1[key1] = val1.build();
            }
            return results1;
        }());
        this.internal.mappings = mappingsResource;
        return this;
    }

    // Catch-all field for datasource-specific properties. Should not be available in as code tooling.
    legacyOptions(legacyOptions: Record<string, any>): this {
        this.internal.legacyOptions = legacyOptions;
        return this;
    }
}

