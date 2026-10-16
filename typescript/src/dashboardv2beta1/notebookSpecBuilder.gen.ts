// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

/**
 * A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time
 * settings, and a map of elements (panels and cells) referenced by the layout.
 */
export class NotebookSpecBuilder implements cog.Builder<dashboardv2beta1.NotebookSpec> {
    protected readonly internal: dashboardv2beta1.NotebookSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultNotebookSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.NotebookSpec {
        return this.internal;
    }

    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    description(description: string): this {
        this.internal.description = description;
        return this;
    }

    tags(tags: string[]): this {
        this.internal.tags = tags;
        return this;
    }

    timeSettings(timeSettings: cog.Builder<dashboardv2beta1.TimeSettingsSpec>): this {
        const timeSettingsResource = timeSettings.build();
        this.internal.timeSettings = timeSettingsResource;
        return this;
    }

    elements(elements: Record<string, cog.Builder<dashboardv2beta1.NotebookElement>>): this {
        const elementsResource = (function() {
            let results1: Record<string, dashboardv2beta1.NotebookElement> = {};
            for (const key1 in elements) {
                const val1 = elements[key1];
                results1[key1] = val1.build();
            }
            return results1;
        }());
        this.internal.elements = elementsResource;
        return this;
    }

    layout(layout: cog.Builder<dashboardv2beta1.NotebookLayoutKind>): this {
        const layoutResource = layout.build();
        this.internal.layout = layoutResource;
        return this;
    }
}

