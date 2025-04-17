<?php

namespace Grafana\Foundation\Dashboard;

/**
 * A dashboard snapshot shares an interactive dashboard publicly.
 * It is a read-only version of a dashboard, and is not editable.
 * It is possible to create a snapshot of a snapshot.
 * Grafana strips away all sensitive information from the dashboard.
 * Sensitive information stripped: queries (metric, template,annotation) and panel links.
 */
class Snapshot implements \JsonSerializable
{
    /**
     * Time when the snapshot was created
     */
    public string $created;

    /**
     * Time when the snapshot expires, default is never to expire
     */
    public string $expires;

    /**
     * Is the snapshot saved in an external grafana instance
     */
    public bool $external;

    /**
     * external url, if snapshot was shared in external grafana instance
     */
    public string $externalUrl;

    /**
     * original url, url of the dashboard that was snapshotted
     */
    public string $originalUrl;

    /**
     * Unique identifier of the snapshot
     */
    public int $id;

    /**
     * Optional, defined the unique key of the snapshot, required if external is true
     */
    public string $key;

    /**
     * Optional, name of the snapshot
     */
    public string $name;

    /**
     * org id of the snapshot
     */
    public int $orgId;

    /**
     * last time when the snapshot was updated
     */
    public string $updated;

    /**
     * url of the snapshot, if snapshot was shared internally
     */
    public ?string $url;

    /**
     * user id of the snapshot creator
     */
    public int $userId;

    public ?\Grafana\Foundation\Dashboard\Dashboard $dashboard;

    /**
     * @param string|null $created
     * @param string|null $expires
     * @param bool|null $external
     * @param string|null $externalUrl
     * @param string|null $originalUrl
     * @param int|null $id
     * @param string|null $key
     * @param string|null $name
     * @param int|null $orgId
     * @param string|null $updated
     * @param string|null $url
     * @param int|null $userId
     * @param \Grafana\Foundation\Dashboard\Dashboard|null $dashboard
     */
    public function __construct(?string $created = null, ?string $expires = null, ?bool $external = null, ?string $externalUrl = null, ?string $originalUrl = null, ?int $id = null, ?string $key = null, ?string $name = null, ?int $orgId = null, ?string $updated = null, ?string $url = null, ?int $userId = null, ?\Grafana\Foundation\Dashboard\Dashboard $dashboard = null)
    {
        $this->created = $created ?: "";
        $this->expires = $expires ?: "";
        $this->external = $external ?: false;
        $this->externalUrl = $externalUrl ?: "";
        $this->originalUrl = $originalUrl ?: "";
        $this->id = $id ?: 0;
        $this->key = $key ?: "";
        $this->name = $name ?: "";
        $this->orgId = $orgId ?: 0;
        $this->updated = $updated ?: "";
        $this->url = $url;
        $this->userId = $userId ?: 0;
        $this->dashboard = $dashboard;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{created?: string, expires?: string, external?: bool, externalUrl?: string, originalUrl?: string, id?: int, key?: string, name?: string, orgId?: int, updated?: string, url?: string, userId?: int, dashboard?: mixed} $inputData */
        $data = $inputData;
        return new self(
            created: $data["created"] ?? null,
            expires: $data["expires"] ?? null,
            external: $data["external"] ?? null,
            externalUrl: $data["externalUrl"] ?? null,
            originalUrl: $data["originalUrl"] ?? null,
            id: $data["id"] ?? null,
            key: $data["key"] ?? null,
            name: $data["name"] ?? null,
            orgId: $data["orgId"] ?? null,
            updated: $data["updated"] ?? null,
            url: $data["url"] ?? null,
            userId: $data["userId"] ?? null,
            dashboard: isset($data["dashboard"]) ? (function($input) {
    	/** @var array{id?: int, uid?: string, title?: string, description?: string, revision?: int, gnetId?: string, tags?: array<string>, timezone?: string, editable?: bool, graphTooltip?: int, time?: mixed, timepicker?: mixed, fiscalYearStartMonth?: int, liveNow?: bool, weekStart?: string, refresh?: string, schemaVersion?: int, version?: int, panels?: array<mixed|mixed>, templating?: mixed, annotations?: mixed, links?: array<mixed>, snapshot?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\Dashboard::fromArray($val);
    })($data["dashboard"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->created = $this->created;
        $data->expires = $this->expires;
        $data->external = $this->external;
        $data->externalUrl = $this->externalUrl;
        $data->originalUrl = $this->originalUrl;
        $data->id = $this->id;
        $data->key = $this->key;
        $data->name = $this->name;
        $data->orgId = $this->orgId;
        $data->updated = $this->updated;
        $data->userId = $this->userId;
        if (isset($this->url)) {
            $data->url = $this->url;
        }
        if (isset($this->dashboard)) {
            $data->dashboard = $this->dashboard;
        }
        return $data;
    }
}
