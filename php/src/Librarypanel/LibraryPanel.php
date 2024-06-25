<?php

namespace Grafana\Foundation\Librarypanel;

class LibraryPanel implements \JsonSerializable
{
    /**
     * Folder UID
     */
    public ?string $folderUid;

    /**
     * Library element UID
     */
    public string $uid;

    /**
     * Panel name (also saved in the model)
     */
    public string $name;

    /**
     * Panel description
     */
    public ?string $description;

    /**
     * The panel type (from inside the model)
     */
    public string $type;

    /**
     * Dashboard version when this was saved (zero if unknown)
     */
    public ?int $schemaVersion;

    /**
     * panel version, incremented each time the dashboard is updated.
     */
    public int $version;

    /**
     * TODO: should be the same panel schema defined in dashboard
     * Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;
     */
    public \Grafana\Foundation\Librarypanel\LibrarypanelLibraryPanelModel $model;

    /**
     * Object storage metadata
     */
    public ?\Grafana\Foundation\Librarypanel\LibraryElementDTOMeta $meta;

    /**
     * @param string|null $folderUid
     * @param string|null $uid
     * @param string|null $name
     * @param string|null $description
     * @param string|null $type
     * @param int|null $schemaVersion
     * @param int|null $version
     * @param \Grafana\Foundation\Librarypanel\LibrarypanelLibraryPanelModel|null $model
     * @param \Grafana\Foundation\Librarypanel\LibraryElementDTOMeta|null $meta
     */
    public function __construct(?string $folderUid = null, ?string $uid = null, ?string $name = null, ?string $description = null, ?string $type = null, ?int $schemaVersion = null, ?int $version = null, ?\Grafana\Foundation\Librarypanel\LibrarypanelLibraryPanelModel $model = null, ?\Grafana\Foundation\Librarypanel\LibraryElementDTOMeta $meta = null)
    {
        $this->folderUid = $folderUid;
        $this->uid = $uid ?: "";
        $this->name = $name ?: "";
        $this->description = $description;
        $this->type = $type ?: "";
        $this->schemaVersion = $schemaVersion;
        $this->version = $version ?: 0;
        $this->model = $model ?: new \Grafana\Foundation\Librarypanel\LibrarypanelLibraryPanelModel();
        $this->meta = $meta;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{folderUid?: string, uid?: string, name?: string, description?: string, type?: string, schemaVersion?: int, version?: int, model?: mixed, meta?: mixed} $inputData */
        $data = $inputData;
        return new self(
            folderUid: $data["folderUid"] ?? null,
            uid: $data["uid"] ?? null,
            name: $data["name"] ?? null,
            description: $data["description"] ?? null,
            type: $data["type"] ?? null,
            schemaVersion: $data["schemaVersion"] ?? null,
            version: $data["version"] ?? null,
            model: isset($data["model"]) ? (function($input) {
    	/** @var array{type?: string, pluginVersion?: string, tags?: array<string>, targets?: array<mixed>, title?: string, description?: string, transparent?: bool, datasource?: mixed, links?: array<mixed>, repeat?: string, repeatDirection?: string, repeatPanelId?: int, maxDataPoints?: float, transformations?: array<mixed>, interval?: string, timeFrom?: string, timeShift?: string, options?: mixed, fieldConfig?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Librarypanel\LibrarypanelLibraryPanelModel::fromArray($val);
    })($data["model"]) : null,
            meta: isset($data["meta"]) ? (function($input) {
    	/** @var array{folderName?: string, folderUid?: string, connectedDashboards?: int, created?: string, updated?: string, createdBy?: mixed, updatedBy?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Librarypanel\LibraryElementDTOMeta::fromArray($val);
    })($data["meta"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "uid" => $this->uid,
            "name" => $this->name,
            "type" => $this->type,
            "version" => $this->version,
            "model" => $this->model,
        ];
        if (isset($this->folderUid)) {
            $data["folderUid"] = $this->folderUid;
        }
        if (isset($this->description)) {
            $data["description"] = $this->description;
        }
        if (isset($this->schemaVersion)) {
            $data["schemaVersion"] = $this->schemaVersion;
        }
        if (isset($this->meta)) {
            $data["meta"] = $this->meta;
        }
        return $data;
    }
}
