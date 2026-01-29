// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class AnnotationQueryBuilder implements cog.Builder<dashboardv2beta1.AnnotationQueryKind> {
    protected readonly internal: dashboardv2beta1.AnnotationQueryKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultAnnotationQueryKind();
        this.internal.kind = "AnnotationQuery";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AnnotationQueryKind {
        return this.internal;
    }

    spec(spec: cog.Builder<dashboardv2beta1.AnnotationQuerySpec>): this {
        const specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }

    query(query: cog.Builder<dashboardv2beta1.DataQueryKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        const queryResource = query.build();
        this.internal.spec.query = queryResource;
        return this;
    }

    enable(enable: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        this.internal.spec.enable = enable;
        return this;
    }

    hide(hide: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    iconColor(iconColor: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        this.internal.spec.iconColor = iconColor;
        return this;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        this.internal.spec.name = name;
        return this;
    }

    builtIn(builtIn: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        this.internal.spec.builtIn = builtIn;
        return this;
    }

    filter(filter: cog.Builder<dashboardv2beta1.AnnotationPanelFilter>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        const filterResource = filter.build();
        this.internal.spec.filter = filterResource;
        return this;
    }

    // Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
    placement(placement: "inControlsMenu"): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        this.internal.spec.placement = placement;
        return this;
    }

    // Mappings define how to convert data frame fields to annotation event fields.
    mappings(mappings: Record<string, cog.Builder<dashboardv2beta1.AnnotationEventFieldMapping>>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        const mappingsResource = (function() {
            let results1: Record<string, dashboardv2beta1.AnnotationEventFieldMapping> = {};
            for (const key1 in mappings) {
                const val1 = mappings[key1];
                results1[key1] = val1.build();
            }
            return results1;
        }());
        this.internal.spec.mappings = mappingsResource;
        return this;
    }

    // Catch-all field for datasource-specific properties. Should not be available in as code tooling.
    legacyOptions(legacyOptions: Record<string, any>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        this.internal.spec.legacyOptions = legacyOptions;
        return this;
    }
}

