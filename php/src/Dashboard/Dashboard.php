<?php

namespace Grafana\Foundation\Dashboard;

class Dashboard implements \JsonSerializable
{
    /**
     * Unique numeric identifier for the dashboard.
     * `id` is internal to a specific Grafana instance. `uid` should be used to identify a dashboard across Grafana instances.
     */
    public ?int $id;

    /**
     * Unique dashboard identifier that can be generated by anyone. string (8-40)
     */
    public ?string $uid;

    /**
     * Title of dashboard.
     */
    public ?string $title;

    /**
     * Description of dashboard.
     */
    public ?string $description;

    /**
     * This property should only be used in dashboards defined by plugins.  It is a quick check
     * to see if the version has changed since the last time.
     */
    public ?int $revision;

    /**
     * ID of a dashboard imported from the https://grafana.com/grafana/dashboards/ portal
     */
    public ?string $gnetId;

    /**
     * Tags associated with dashboard.
     * @var array<string>|null
     */
    public ?array $tags;

    /**
     * Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
     */
    public ?string $timezone;

    /**
     * Whether a dashboard is editable or not.
     */
    public ?bool $editable;

    /**
     * Configuration of dashboard cursor sync behavior.
     * Accepted values are 0 (sync turned off), 1 (shared crosshair), 2 (shared crosshair and tooltip).
     */
    public ?\Grafana\Foundation\Dashboard\DashboardCursorSync $graphTooltip;

    /**
     * Time range for dashboard.
     * Accepted values are relative time strings like {from: 'now-6h', to: 'now'} or absolute time strings like {from: '2020-07-10T08:00:00.000Z', to: '2020-07-10T14:00:00.000Z'}.
     */
    public ?\Grafana\Foundation\Dashboard\DashboardDashboardTime $time;

    /**
     * Configuration of the time picker shown at the top of a dashboard.
     */
    public ?\Grafana\Foundation\Dashboard\TimePickerConfig $timepicker;

    /**
     * The month that the fiscal year starts on.  0 = January, 11 = December
     */
    public ?int $fiscalYearStartMonth;

    /**
     * When set to true, the dashboard will redraw panels at an interval matching the pixel width.
     * This will keep data "moving left" regardless of the query refresh rate. This setting helps
     * avoid dashboards presenting stale live data
     */
    public ?bool $liveNow;

    /**
     * Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
     */
    public ?string $weekStart;

    /**
     * Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
     */
    public ?string $refresh;

    /**
     * Version of the JSON schema, incremented each time a Grafana update brings
     * changes to said schema.
     */
    public int $schemaVersion;

    /**
     * Version of the dashboard, incremented each time the dashboard is updated.
     */
    public ?int $version;

    /**
     * List of dashboard panels
     * @var array<\Grafana\Foundation\Dashboard\Panel|\Grafana\Foundation\Dashboard\RowPanel>|null
     */
    public ?array $panels;

    /**
     * Configured template variables
     */
    public \Grafana\Foundation\Dashboard\DashboardDashboardTemplating $templating;

    /**
     * Contains the list of annotations that are associated with the dashboard.
     * Annotations are used to overlay event markers and overlay event tags on graphs.
     * Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
     * See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
     */
    public \Grafana\Foundation\Dashboard\AnnotationContainer $annotations;

    /**
     * Links with references to other dashboards or external websites.
     * @var array<\Grafana\Foundation\Dashboard\DashboardLink>|null
     */
    public ?array $links;

    /**
     * Snapshot options. They are present only if the dashboard is a snapshot.
     */
    public ?\Grafana\Foundation\Dashboard\Snapshot $snapshot;

    /**
     * @param int|null $id
     * @param string|null $uid
     * @param string|null $title
     * @param string|null $description
     * @param int|null $revision
     * @param string|null $gnetId
     * @param array<string>|null $tags
     * @param string|null $timezone
     * @param bool|null $editable
     * @param \Grafana\Foundation\Dashboard\DashboardCursorSync|null $graphTooltip
     * @param \Grafana\Foundation\Dashboard\DashboardDashboardTime|null $time
     * @param \Grafana\Foundation\Dashboard\TimePickerConfig|null $timepicker
     * @param int|null $fiscalYearStartMonth
     * @param bool|null $liveNow
     * @param string|null $weekStart
     * @param string|null $refresh
     * @param int|null $schemaVersion
     * @param int|null $version
     * @param array<\Grafana\Foundation\Dashboard\Panel|\Grafana\Foundation\Dashboard\RowPanel>|null $panels
     * @param \Grafana\Foundation\Dashboard\DashboardDashboardTemplating|null $templating
     * @param \Grafana\Foundation\Dashboard\AnnotationContainer|null $annotations
     * @param array<\Grafana\Foundation\Dashboard\DashboardLink>|null $links
     * @param \Grafana\Foundation\Dashboard\Snapshot|null $snapshot
     */
    public function __construct(?int $id = null, ?string $uid = null, ?string $title = null, ?string $description = null, ?int $revision = null, ?string $gnetId = null, ?array $tags = null, ?string $timezone = null, ?bool $editable = null, ?\Grafana\Foundation\Dashboard\DashboardCursorSync $graphTooltip = null, ?\Grafana\Foundation\Dashboard\DashboardDashboardTime $time = null, ?\Grafana\Foundation\Dashboard\TimePickerConfig $timepicker = null, ?int $fiscalYearStartMonth = null, ?bool $liveNow = null, ?string $weekStart = null, ?string $refresh = null, ?int $schemaVersion = null, ?int $version = null, ?array $panels = null, ?\Grafana\Foundation\Dashboard\DashboardDashboardTemplating $templating = null, ?\Grafana\Foundation\Dashboard\AnnotationContainer $annotations = null, ?array $links = null, ?\Grafana\Foundation\Dashboard\Snapshot $snapshot = null)
    {
        $this->id = $id;
        $this->uid = $uid;
        $this->title = $title;
        $this->description = $description;
        $this->revision = $revision;
        $this->gnetId = $gnetId;
        $this->tags = $tags;
        $this->timezone = $timezone;
        $this->editable = $editable;
        $this->graphTooltip = $graphTooltip;
        $this->time = $time;
        $this->timepicker = $timepicker;
        $this->fiscalYearStartMonth = $fiscalYearStartMonth;
        $this->liveNow = $liveNow;
        $this->weekStart = $weekStart;
        $this->refresh = $refresh;
        $this->schemaVersion = $schemaVersion ?: 36;
        $this->version = $version;
        $this->panels = $panels;
        $this->templating = $templating ?: new \Grafana\Foundation\Dashboard\DashboardDashboardTemplating();
        $this->annotations = $annotations ?: new \Grafana\Foundation\Dashboard\AnnotationContainer();
        $this->links = $links;
        $this->snapshot = $snapshot;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{id?: int, uid?: string, title?: string, description?: string, revision?: int, gnetId?: string, tags?: array<string>, timezone?: string, editable?: bool, graphTooltip?: int, time?: mixed, timepicker?: mixed, fiscalYearStartMonth?: int, liveNow?: bool, weekStart?: string, refresh?: string, schemaVersion?: int, version?: int, panels?: array<mixed|mixed>, templating?: mixed, annotations?: mixed, links?: array<mixed>, snapshot?: mixed} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
            uid: $data["uid"] ?? null,
            title: $data["title"] ?? null,
            description: $data["description"] ?? null,
            revision: $data["revision"] ?? null,
            gnetId: $data["gnetId"] ?? null,
            tags: $data["tags"] ?? null,
            timezone: $data["timezone"] ?? null,
            editable: $data["editable"] ?? null,
            graphTooltip: isset($data["graphTooltip"]) ? (function($input) { return \Grafana\Foundation\Dashboard\DashboardCursorSync::fromValue($input); })($data["graphTooltip"]) : null,
            time: isset($data["time"]) ? (function($input) {
    	/** @var array{from?: string, to?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DashboardDashboardTime::fromArray($val);
    })($data["time"]) : null,
            timepicker: isset($data["timepicker"]) ? (function($input) {
    	/** @var array{hidden?: bool, refresh_intervals?: array<string>, time_options?: array<string>, nowDelay?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\TimePickerConfig::fromArray($val);
    })($data["timepicker"]) : null,
            fiscalYearStartMonth: $data["fiscalYearStartMonth"] ?? null,
            liveNow: $data["liveNow"] ?? null,
            weekStart: $data["weekStart"] ?? null,
            refresh: $data["refresh"] ?? null,
            schemaVersion: $data["schemaVersion"] ?? null,
            version: $data["version"] ?? null,
            panels: !empty($data["panels"]) ? array_map((function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
        /** @var array<string, mixed> $input */
        switch ($input["type"]) {
        case "row":
            return RowPanel::fromArray($input);
        default:
            return Panel::fromArray($input);
    }
    }), $data["panels"]) : null,
            templating: isset($data["templating"]) ? (function($input) {
    	/** @var array{list?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DashboardDashboardTemplating::fromArray($val);
    })($data["templating"]) : null,
            annotations: isset($data["annotations"]) ? (function($input) {
    	/** @var array{list?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\AnnotationContainer::fromArray($val);
    })($data["annotations"]) : null,
            links: array_filter(array_map((function($input) {
    	/** @var array{title?: string, type?: string, icon?: string, tooltip?: string, url?: string, tags?: array<string>, asDropdown?: bool, targetBlank?: bool, includeVars?: bool, keepTime?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DashboardLink::fromArray($val);
    }), $data["links"] ?? [])),
            snapshot: isset($data["snapshot"]) ? (function($input) {
    	/** @var array{created?: string, expires?: string, external?: bool, externalUrl?: string, originalUrl?: string, id?: int, key?: string, name?: string, orgId?: int, updated?: string, url?: string, userId?: int, dashboard?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\Snapshot::fromArray($val);
    })($data["snapshot"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "schemaVersion" => $this->schemaVersion,
            "templating" => $this->templating,
            "annotations" => $this->annotations,
        ];
        if (isset($this->id)) {
            $data["id"] = $this->id;
        }
        if (isset($this->uid)) {
            $data["uid"] = $this->uid;
        }
        if (isset($this->title)) {
            $data["title"] = $this->title;
        }
        if (isset($this->description)) {
            $data["description"] = $this->description;
        }
        if (isset($this->revision)) {
            $data["revision"] = $this->revision;
        }
        if (isset($this->gnetId)) {
            $data["gnetId"] = $this->gnetId;
        }
        if (isset($this->tags)) {
            $data["tags"] = $this->tags;
        }
        if (isset($this->timezone)) {
            $data["timezone"] = $this->timezone;
        }
        if (isset($this->editable)) {
            $data["editable"] = $this->editable;
        }
        if (isset($this->graphTooltip)) {
            $data["graphTooltip"] = $this->graphTooltip;
        }
        if (isset($this->time)) {
            $data["time"] = $this->time;
        }
        if (isset($this->timepicker)) {
            $data["timepicker"] = $this->timepicker;
        }
        if (isset($this->fiscalYearStartMonth)) {
            $data["fiscalYearStartMonth"] = $this->fiscalYearStartMonth;
        }
        if (isset($this->liveNow)) {
            $data["liveNow"] = $this->liveNow;
        }
        if (isset($this->weekStart)) {
            $data["weekStart"] = $this->weekStart;
        }
        if (isset($this->refresh)) {
            $data["refresh"] = $this->refresh;
        }
        if (isset($this->version)) {
            $data["version"] = $this->version;
        }
        if (isset($this->panels)) {
            $data["panels"] = $this->panels;
        }
        if (isset($this->links)) {
            $data["links"] = $this->links;
        }
        if (isset($this->snapshot)) {
            $data["snapshot"] = $this->snapshot;
        }
        return $data;
    }
}
