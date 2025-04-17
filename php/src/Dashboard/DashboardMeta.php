<?php

namespace Grafana\Foundation\Dashboard;

class DashboardMeta implements \JsonSerializable
{
    /**
     * +k8s:deepcopy-gen=true
     */
    public ?\Grafana\Foundation\Dashboard\AnnotationPermission $annotationsPermissions;

    public ?string $apiVersion;

    public ?bool $canAdmin;

    public ?bool $canDelete;

    public ?bool $canEdit;

    public ?bool $canSave;

    public ?bool $canStar;

    public ?string $created;

    public ?string $createdBy;

    public ?string $expires;

    /**
     * Deprecated: use FolderUID instead
     */
    public ?int $folderId;

    public ?string $folderTitle;

    public ?string $folderUid;

    public ?string $folderUrl;

    public ?bool $hasAcl;

    public ?bool $isFolder;

    public ?bool $isSnapshot;

    public ?bool $isStarred;

    public ?bool $provisioned;

    public ?string $provisionedExternalId;

    public ?bool $publicDashboardEnabled;

    public ?string $slug;

    public ?string $type;

    public ?string $updated;

    public ?string $updatedBy;

    public ?string $url;

    public ?int $version;

    /**
     * @param \Grafana\Foundation\Dashboard\AnnotationPermission|null $annotationsPermissions
     * @param string|null $apiVersion
     * @param bool|null $canAdmin
     * @param bool|null $canDelete
     * @param bool|null $canEdit
     * @param bool|null $canSave
     * @param bool|null $canStar
     * @param string|null $created
     * @param string|null $createdBy
     * @param string|null $expires
     * @param int|null $folderId
     * @param string|null $folderTitle
     * @param string|null $folderUid
     * @param string|null $folderUrl
     * @param bool|null $hasAcl
     * @param bool|null $isFolder
     * @param bool|null $isSnapshot
     * @param bool|null $isStarred
     * @param bool|null $provisioned
     * @param string|null $provisionedExternalId
     * @param bool|null $publicDashboardEnabled
     * @param string|null $slug
     * @param string|null $type
     * @param string|null $updated
     * @param string|null $updatedBy
     * @param string|null $url
     * @param int|null $version
     */
    public function __construct(?\Grafana\Foundation\Dashboard\AnnotationPermission $annotationsPermissions = null, ?string $apiVersion = null, ?bool $canAdmin = null, ?bool $canDelete = null, ?bool $canEdit = null, ?bool $canSave = null, ?bool $canStar = null, ?string $created = null, ?string $createdBy = null, ?string $expires = null, ?int $folderId = null, ?string $folderTitle = null, ?string $folderUid = null, ?string $folderUrl = null, ?bool $hasAcl = null, ?bool $isFolder = null, ?bool $isSnapshot = null, ?bool $isStarred = null, ?bool $provisioned = null, ?string $provisionedExternalId = null, ?bool $publicDashboardEnabled = null, ?string $slug = null, ?string $type = null, ?string $updated = null, ?string $updatedBy = null, ?string $url = null, ?int $version = null)
    {
        $this->annotationsPermissions = $annotationsPermissions;
        $this->apiVersion = $apiVersion;
        $this->canAdmin = $canAdmin;
        $this->canDelete = $canDelete;
        $this->canEdit = $canEdit;
        $this->canSave = $canSave;
        $this->canStar = $canStar;
        $this->created = $created;
        $this->createdBy = $createdBy;
        $this->expires = $expires;
        $this->folderId = $folderId;
        $this->folderTitle = $folderTitle;
        $this->folderUid = $folderUid;
        $this->folderUrl = $folderUrl;
        $this->hasAcl = $hasAcl;
        $this->isFolder = $isFolder;
        $this->isSnapshot = $isSnapshot;
        $this->isStarred = $isStarred;
        $this->provisioned = $provisioned;
        $this->provisionedExternalId = $provisionedExternalId;
        $this->publicDashboardEnabled = $publicDashboardEnabled;
        $this->slug = $slug;
        $this->type = $type;
        $this->updated = $updated;
        $this->updatedBy = $updatedBy;
        $this->url = $url;
        $this->version = $version;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{annotationsPermissions?: mixed, apiVersion?: string, canAdmin?: bool, canDelete?: bool, canEdit?: bool, canSave?: bool, canStar?: bool, created?: string, createdBy?: string, expires?: string, folderId?: int, folderTitle?: string, folderUid?: string, folderUrl?: string, hasAcl?: bool, isFolder?: bool, isSnapshot?: bool, isStarred?: bool, provisioned?: bool, provisionedExternalId?: string, publicDashboardEnabled?: bool, slug?: string, type?: string, updated?: string, updatedBy?: string, url?: string, version?: int} $inputData */
        $data = $inputData;
        return new self(
            annotationsPermissions: isset($data["annotationsPermissions"]) ? (function($input) {
    	/** @var array{dashboard?: mixed, organization?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\AnnotationPermission::fromArray($val);
    })($data["annotationsPermissions"]) : null,
            apiVersion: $data["apiVersion"] ?? null,
            canAdmin: $data["canAdmin"] ?? null,
            canDelete: $data["canDelete"] ?? null,
            canEdit: $data["canEdit"] ?? null,
            canSave: $data["canSave"] ?? null,
            canStar: $data["canStar"] ?? null,
            created: $data["created"] ?? null,
            createdBy: $data["createdBy"] ?? null,
            expires: $data["expires"] ?? null,
            folderId: $data["folderId"] ?? null,
            folderTitle: $data["folderTitle"] ?? null,
            folderUid: $data["folderUid"] ?? null,
            folderUrl: $data["folderUrl"] ?? null,
            hasAcl: $data["hasAcl"] ?? null,
            isFolder: $data["isFolder"] ?? null,
            isSnapshot: $data["isSnapshot"] ?? null,
            isStarred: $data["isStarred"] ?? null,
            provisioned: $data["provisioned"] ?? null,
            provisionedExternalId: $data["provisionedExternalId"] ?? null,
            publicDashboardEnabled: $data["publicDashboardEnabled"] ?? null,
            slug: $data["slug"] ?? null,
            type: $data["type"] ?? null,
            updated: $data["updated"] ?? null,
            updatedBy: $data["updatedBy"] ?? null,
            url: $data["url"] ?? null,
            version: $data["version"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->annotationsPermissions)) {
            $data->annotationsPermissions = $this->annotationsPermissions;
        }
        if (isset($this->apiVersion)) {
            $data->apiVersion = $this->apiVersion;
        }
        if (isset($this->canAdmin)) {
            $data->canAdmin = $this->canAdmin;
        }
        if (isset($this->canDelete)) {
            $data->canDelete = $this->canDelete;
        }
        if (isset($this->canEdit)) {
            $data->canEdit = $this->canEdit;
        }
        if (isset($this->canSave)) {
            $data->canSave = $this->canSave;
        }
        if (isset($this->canStar)) {
            $data->canStar = $this->canStar;
        }
        if (isset($this->created)) {
            $data->created = $this->created;
        }
        if (isset($this->createdBy)) {
            $data->createdBy = $this->createdBy;
        }
        if (isset($this->expires)) {
            $data->expires = $this->expires;
        }
        if (isset($this->folderId)) {
            $data->folderId = $this->folderId;
        }
        if (isset($this->folderTitle)) {
            $data->folderTitle = $this->folderTitle;
        }
        if (isset($this->folderUid)) {
            $data->folderUid = $this->folderUid;
        }
        if (isset($this->folderUrl)) {
            $data->folderUrl = $this->folderUrl;
        }
        if (isset($this->hasAcl)) {
            $data->hasAcl = $this->hasAcl;
        }
        if (isset($this->isFolder)) {
            $data->isFolder = $this->isFolder;
        }
        if (isset($this->isSnapshot)) {
            $data->isSnapshot = $this->isSnapshot;
        }
        if (isset($this->isStarred)) {
            $data->isStarred = $this->isStarred;
        }
        if (isset($this->provisioned)) {
            $data->provisioned = $this->provisioned;
        }
        if (isset($this->provisionedExternalId)) {
            $data->provisionedExternalId = $this->provisionedExternalId;
        }
        if (isset($this->publicDashboardEnabled)) {
            $data->publicDashboardEnabled = $this->publicDashboardEnabled;
        }
        if (isset($this->slug)) {
            $data->slug = $this->slug;
        }
        if (isset($this->type)) {
            $data->type = $this->type;
        }
        if (isset($this->updated)) {
            $data->updated = $this->updated;
        }
        if (isset($this->updatedBy)) {
            $data->updatedBy = $this->updatedBy;
        }
        if (isset($this->url)) {
            $data->url = $this->url;
        }
        if (isset($this->version)) {
            $data->version = $this->version;
        }
        return $data;
    }
}
