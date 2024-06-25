<?php

namespace Grafana\Foundation\Librarypanel;

class LibraryElementDTOMeta implements \JsonSerializable
{
    public string $folderName;

    public string $folderUid;

    public int $connectedDashboards;

    public string $created;

    public string $updated;

    public \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser $createdBy;

    public \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser $updatedBy;

    /**
     * @param string|null $folderName
     * @param string|null $folderUid
     * @param int|null $connectedDashboards
     * @param string|null $created
     * @param string|null $updated
     * @param \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser|null $createdBy
     * @param \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser|null $updatedBy
     */
    public function __construct(?string $folderName = null, ?string $folderUid = null, ?int $connectedDashboards = null, ?string $created = null, ?string $updated = null, ?\Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser $createdBy = null, ?\Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser $updatedBy = null)
    {
        $this->folderName = $folderName ?: "";
        $this->folderUid = $folderUid ?: "";
        $this->connectedDashboards = $connectedDashboards ?: 0;
        $this->created = $created ?: "";
        $this->updated = $updated ?: "";
        $this->createdBy = $createdBy ?: new \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser();
        $this->updatedBy = $updatedBy ?: new \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{folderName?: string, folderUid?: string, connectedDashboards?: int, created?: string, updated?: string, createdBy?: mixed, updatedBy?: mixed} $inputData */
        $data = $inputData;
        return new self(
            folderName: $data["folderName"] ?? null,
            folderUid: $data["folderUid"] ?? null,
            connectedDashboards: $data["connectedDashboards"] ?? null,
            created: $data["created"] ?? null,
            updated: $data["updated"] ?? null,
            createdBy: isset($data["createdBy"]) ? (function($input) {
    	/** @var array{id?: int, name?: string, avatarUrl?: string} */
    $val = $input;
    	return \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser::fromArray($val);
    })($data["createdBy"]) : null,
            updatedBy: isset($data["updatedBy"]) ? (function($input) {
    	/** @var array{id?: int, name?: string, avatarUrl?: string} */
    $val = $input;
    	return \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser::fromArray($val);
    })($data["updatedBy"]) : null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "folderName" => $this->folderName,
            "folderUid" => $this->folderUid,
            "connectedDashboards" => $this->connectedDashboards,
            "created" => $this->created,
            "updated" => $this->updated,
            "createdBy" => $this->createdBy,
            "updatedBy" => $this->updatedBy,
        ];
        return $data;
    }
}
