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

    build(): alerting.ContactPoint {
        return this.internal;
    }

    // EmbeddedContactPoint is the contact point type that is used
    // by grafanas embedded alertmanager implementation.
    disableResolveMessage(disableResolveMessage: boolean): this {
        this.internal.disableResolveMessage = disableResolveMessage;
        return this;
    }

    // EmbeddedContactPoint is the contact point type that is used
    // by grafanas embedded alertmanager implementation.
    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    // EmbeddedContactPoint is the contact point type that is used
    // by grafanas embedded alertmanager implementation.
    provenance(provenance: string): this {
        this.internal.provenance = provenance;
        return this;
    }

    // EmbeddedContactPoint is the contact point type that is used
    // by grafanas embedded alertmanager implementation.
    settings(settings: alerting.Json): this {
        this.internal.settings = settings;
        return this;
    }

    // EmbeddedContactPoint is the contact point type that is used
    // by grafanas embedded alertmanager implementation.
    type(type: "alertmanager" | " dingding" | " discord" | " email" | " googlechat" | " kafka" | " line" | " opsgenie" | " pagerduty" | " pushover" | " sensugo" | " slack" | " teams" | " telegram" | " threema" | " victorops" | " webhook" | " wecom"): this {
        this.internal.type = type;
        return this;
    }

    // EmbeddedContactPoint is the contact point type that is used
    // by grafanas embedded alertmanager implementation.
    uid(uid: string): this {
        this.internal.uid = uid;
        return this;
    }
}
