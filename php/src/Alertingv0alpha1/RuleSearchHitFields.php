<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * #RuleSearchHitFields is the per-kind field payload returned on each hit.
 * It carries the union of alert- and recording-rule search fields; only the
 * fields relevant to a hit's kind are populated. This maps to the kind's
 * declared searchFields.
 */
class RuleSearchHitFields implements \JsonSerializable
{
    public ?string $title;

    public ?string $folder;

    public ?string $type;

    public ?string $interval;

    public ?bool $paused;

    /**
     * @var array<string, string>|null
     */
    public ?array $labels;

    /**
     * @var array<string>|null
     */
    public ?array $datasourceUIDs;

    /**
     * Alert-rule fields.
     * @var array<string, string>|null
     */
    public ?array $annotations;

    public ?string $for;

    public ?string $keepFiringFor;

    public ?string $dashboardUID;

    public ?int $panelID;

    public ?string $receiver;

    public ?string $notificationType;

    public ?string $routingTree;

    /**
     * Recording-rule fields.
     */
    public ?string $metric;

    public ?string $targetDatasourceUID;

    /**
     * @param string|null $title
     * @param string|null $folder
     * @param string|null $type
     * @param string|null $interval
     * @param bool|null $paused
     * @param array<string, string>|null $labels
     * @param array<string>|null $datasourceUIDs
     * @param array<string, string>|null $annotations
     * @param string|null $for
     * @param string|null $keepFiringFor
     * @param string|null $dashboardUID
     * @param int|null $panelID
     * @param string|null $receiver
     * @param string|null $notificationType
     * @param string|null $routingTree
     * @param string|null $metric
     * @param string|null $targetDatasourceUID
     */
    public function __construct(?string $title = null, ?string $folder = null, ?string $type = null, ?string $interval = null, ?bool $paused = null, ?array $labels = null, ?array $datasourceUIDs = null, ?array $annotations = null, ?string $for = null, ?string $keepFiringFor = null, ?string $dashboardUID = null, ?int $panelID = null, ?string $receiver = null, ?string $notificationType = null, ?string $routingTree = null, ?string $metric = null, ?string $targetDatasourceUID = null)
    {
        $this->title = $title;
        $this->folder = $folder;
        $this->type = $type;
        $this->interval = $interval;
        $this->paused = $paused;
        $this->labels = $labels;
        $this->datasourceUIDs = $datasourceUIDs;
        $this->annotations = $annotations;
        $this->for = $for;
        $this->keepFiringFor = $keepFiringFor;
        $this->dashboardUID = $dashboardUID;
        $this->panelID = $panelID;
        $this->receiver = $receiver;
        $this->notificationType = $notificationType;
        $this->routingTree = $routingTree;
        $this->metric = $metric;
        $this->targetDatasourceUID = $targetDatasourceUID;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{title?: string, folder?: string, type?: string, interval?: string, paused?: bool, labels?: array<string, string>, datasourceUIDs?: array<string>, annotations?: array<string, string>, for?: string, keepFiringFor?: string, dashboardUID?: string, panelID?: int, receiver?: string, notificationType?: string, routingTree?: string, metric?: string, targetDatasourceUID?: string} $inputData */
        $data = $inputData;
        return new self(
            title: $data["title"] ?? null,
            folder: $data["folder"] ?? null,
            type: $data["type"] ?? null,
            interval: $data["interval"] ?? null,
            paused: $data["paused"] ?? null,
            labels: $data["labels"] ?? null,
            datasourceUIDs: $data["datasourceUIDs"] ?? null,
            annotations: $data["annotations"] ?? null,
            for: $data["for"] ?? null,
            keepFiringFor: $data["keepFiringFor"] ?? null,
            dashboardUID: $data["dashboardUID"] ?? null,
            panelID: $data["panelID"] ?? null,
            receiver: $data["receiver"] ?? null,
            notificationType: $data["notificationType"] ?? null,
            routingTree: $data["routingTree"] ?? null,
            metric: $data["metric"] ?? null,
            targetDatasourceUID: $data["targetDatasourceUID"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->title)) {
            $data->title = $this->title;
        }
        if (isset($this->folder)) {
            $data->folder = $this->folder;
        }
        if (isset($this->type)) {
            $data->type = $this->type;
        }
        if (isset($this->interval)) {
            $data->interval = $this->interval;
        }
        if (isset($this->paused)) {
            $data->paused = $this->paused;
        }
        if (isset($this->labels)) {
            $data->labels = $this->labels;
        }
        if (isset($this->datasourceUIDs)) {
            $data->datasourceUIDs = $this->datasourceUIDs;
        }
        if (isset($this->annotations)) {
            $data->annotations = $this->annotations;
        }
        if (isset($this->for)) {
            $data->for = $this->for;
        }
        if (isset($this->keepFiringFor)) {
            $data->keepFiringFor = $this->keepFiringFor;
        }
        if (isset($this->dashboardUID)) {
            $data->dashboardUID = $this->dashboardUID;
        }
        if (isset($this->panelID)) {
            $data->panelID = $this->panelID;
        }
        if (isset($this->receiver)) {
            $data->receiver = $this->receiver;
        }
        if (isset($this->notificationType)) {
            $data->notificationType = $this->notificationType;
        }
        if (isset($this->routingTree)) {
            $data->routingTree = $this->routingTree;
        }
        if (isset($this->metric)) {
            $data->metric = $this->metric;
        }
        if (isset($this->targetDatasourceUID)) {
            $data->targetDatasourceUID = $this->targetDatasourceUID;
        }
        return $data;
    }
}
