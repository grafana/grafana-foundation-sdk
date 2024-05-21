// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

export class NotificationTemplateBuilder implements cog.Builder<alerting.NotificationTemplate> {
    protected readonly internal: alerting.NotificationTemplate;

    constructor() {
        this.internal = alerting.defaultNotificationTemplate();
    }

    build(): alerting.NotificationTemplate {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    provenance(provenance: alerting.Provenance): this {
        this.internal.provenance = provenance;
        return this;
    }

    template(template: string): this {
        this.internal.template = template;
        return this;
    }
}
