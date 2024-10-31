// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

// EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
export class ContactPointBuilder implements cog.Builder<alerting.ContactPoint> {
    protected readonly internal: alerting.ContactPoint;

    constructor() {
        this.internal = alerting.defaultContactPoint();
    }

    /**
     * Builds the object.
     */
    build(): alerting.ContactPoint {
        return this.internal;
    }

    disableResolveMessage(disableResolveMessage: boolean): this {
        this.internal.disableResolveMessage = disableResolveMessage;
        return this;
    }

    // Name is used as grouping key in the UI. Contact points with the
    // same name will be grouped in the UI.
    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    provenance(provenance: string): this {
        this.internal.provenance = provenance;
        return this;
    }

    settings(settings: alerting.Json): this {
        this.internal.settings = settings;
        return this;
    }

    type(type: "alertmanager" | " dingding" | " discord" | " email" | " googlechat" | " kafka" | " line" | " opsgenie" | " pagerduty" | " pushover" | " sensugo" | " slack" | " teams" | " telegram" | " threema" | " victorops" | " webhook" | " wecom"): this {
        this.internal.type = type;
        return this;
    }

    // UID is the unique identifier of the contact point. The UID can be
    // set by the user.
    uid(uid: string): this {
        this.internal.uid = uid;
        return this;
    }
}
