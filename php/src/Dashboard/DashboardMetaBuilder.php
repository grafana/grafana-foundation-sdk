<?php

namespace Grafana\Foundation\Dashboard;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardMeta>
 */
class DashboardMetaBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\DashboardMeta $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\DashboardMeta();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\DashboardMeta
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\AnnotationPermission> $annotationsPermissions
     */
    public function annotationsPermissions(\Grafana\Foundation\Cog\Builder $annotationsPermissions): static
    {
        $annotationsPermissionsResource = $annotationsPermissions->build();
        $this->internal->annotationsPermissions = $annotationsPermissionsResource;
    
        return $this;
    }
    public function canAdmin(bool $canAdmin): static
    {
        $this->internal->canAdmin = $canAdmin;
    
        return $this;
    }
    public function canDelete(bool $canDelete): static
    {
        $this->internal->canDelete = $canDelete;
    
        return $this;
    }
    public function canEdit(bool $canEdit): static
    {
        $this->internal->canEdit = $canEdit;
    
        return $this;
    }
    public function canSave(bool $canSave): static
    {
        $this->internal->canSave = $canSave;
    
        return $this;
    }
    public function canStar(bool $canStar): static
    {
        $this->internal->canStar = $canStar;
    
        return $this;
    }
    public function created(string $created): static
    {
        $this->internal->created = $created;
    
        return $this;
    }
    public function createdBy(string $createdBy): static
    {
        $this->internal->createdBy = $createdBy;
    
        return $this;
    }
    public function expires(string $expires): static
    {
        $this->internal->expires = $expires;
    
        return $this;
    }
    public function folderId(int $folderId): static
    {
        $this->internal->folderId = $folderId;
    
        return $this;
    }
    public function folderTitle(string $folderTitle): static
    {
        $this->internal->folderTitle = $folderTitle;
    
        return $this;
    }
    public function folderUid(string $folderUid): static
    {
        $this->internal->folderUid = $folderUid;
    
        return $this;
    }
    public function folderUrl(string $folderUrl): static
    {
        $this->internal->folderUrl = $folderUrl;
    
        return $this;
    }
    public function hasAcl(bool $hasAcl): static
    {
        $this->internal->hasAcl = $hasAcl;
    
        return $this;
    }
    public function isFolder(bool $isFolder): static
    {
        $this->internal->isFolder = $isFolder;
    
        return $this;
    }
    public function isSnapshot(bool $isSnapshot): static
    {
        $this->internal->isSnapshot = $isSnapshot;
    
        return $this;
    }
    public function isStarred(bool $isStarred): static
    {
        $this->internal->isStarred = $isStarred;
    
        return $this;
    }
    public function provisioned(bool $provisioned): static
    {
        $this->internal->provisioned = $provisioned;
    
        return $this;
    }
    public function provisionedExternalId(string $provisionedExternalId): static
    {
        $this->internal->provisionedExternalId = $provisionedExternalId;
    
        return $this;
    }
    public function publicDashboardAccessToken(string $publicDashboardAccessToken): static
    {
        $this->internal->publicDashboardAccessToken = $publicDashboardAccessToken;
    
        return $this;
    }
    public function publicDashboardEnabled(bool $publicDashboardEnabled): static
    {
        $this->internal->publicDashboardEnabled = $publicDashboardEnabled;
    
        return $this;
    }
    public function publicDashboardUid(string $publicDashboardUid): static
    {
        $this->internal->publicDashboardUid = $publicDashboardUid;
    
        return $this;
    }
    public function slug(string $slug): static
    {
        $this->internal->slug = $slug;
    
        return $this;
    }
    public function type(string $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    public function updated(string $updated): static
    {
        $this->internal->updated = $updated;
    
        return $this;
    }
    public function updatedBy(string $updatedBy): static
    {
        $this->internal->updatedBy = $updatedBy;
    
        return $this;
    }
    public function url(string $url): static
    {
        $this->internal->url = $url;
    
        return $this;
    }
    public function version(int $version): static
    {
        $this->internal->version = $version;
    
        return $this;
    }

}
