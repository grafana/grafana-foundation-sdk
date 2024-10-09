<?php

namespace Grafana\Foundation\Dashboard;

final class DashboardMetaConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\DashboardMeta $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\DashboardMetaBuilder())',
        ];
            if ($input->annotationsPermissions !== null) {
    
        
    $buffer = 'annotationsPermissions(';
        $arg0 = \Grafana\Foundation\Dashboard\AnnotationPermissionConverter::convert($input->annotationsPermissions);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->canAdmin !== null) {
    
        
    $buffer = 'canAdmin(';
        $arg0 =\var_export($input->canAdmin, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->canDelete !== null) {
    
        
    $buffer = 'canDelete(';
        $arg0 =\var_export($input->canDelete, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->canEdit !== null) {
    
        
    $buffer = 'canEdit(';
        $arg0 =\var_export($input->canEdit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->canSave !== null) {
    
        
    $buffer = 'canSave(';
        $arg0 =\var_export($input->canSave, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->canStar !== null) {
    
        
    $buffer = 'canStar(';
        $arg0 =\var_export($input->canStar, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->created !== null) {
    
        
    $buffer = 'created(';
        $arg0 =\var_export($input->created, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->createdBy !== null && $input->createdBy !== "") {
    
        
    $buffer = 'createdBy(';
        $arg0 =\var_export($input->createdBy, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->expires !== null) {
    
        
    $buffer = 'expires(';
        $arg0 =\var_export($input->expires, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->folderId !== null) {
    
        
    $buffer = 'folderId(';
        $arg0 =\var_export($input->folderId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->folderTitle !== null && $input->folderTitle !== "") {
    
        
    $buffer = 'folderTitle(';
        $arg0 =\var_export($input->folderTitle, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->folderUid !== null && $input->folderUid !== "") {
    
        
    $buffer = 'folderUid(';
        $arg0 =\var_export($input->folderUid, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->folderUrl !== null && $input->folderUrl !== "") {
    
        
    $buffer = 'folderUrl(';
        $arg0 =\var_export($input->folderUrl, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hasAcl !== null) {
    
        
    $buffer = 'hasAcl(';
        $arg0 =\var_export($input->hasAcl, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->isFolder !== null) {
    
        
    $buffer = 'isFolder(';
        $arg0 =\var_export($input->isFolder, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->isSnapshot !== null) {
    
        
    $buffer = 'isSnapshot(';
        $arg0 =\var_export($input->isSnapshot, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->isStarred !== null) {
    
        
    $buffer = 'isStarred(';
        $arg0 =\var_export($input->isStarred, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->provisioned !== null) {
    
        
    $buffer = 'provisioned(';
        $arg0 =\var_export($input->provisioned, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->provisionedExternalId !== null && $input->provisionedExternalId !== "") {
    
        
    $buffer = 'provisionedExternalId(';
        $arg0 =\var_export($input->provisionedExternalId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->publicDashboardAccessToken !== null && $input->publicDashboardAccessToken !== "") {
    
        
    $buffer = 'publicDashboardAccessToken(';
        $arg0 =\var_export($input->publicDashboardAccessToken, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->publicDashboardEnabled !== null) {
    
        
    $buffer = 'publicDashboardEnabled(';
        $arg0 =\var_export($input->publicDashboardEnabled, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->publicDashboardUid !== null && $input->publicDashboardUid !== "") {
    
        
    $buffer = 'publicDashboardUid(';
        $arg0 =\var_export($input->publicDashboardUid, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->slug !== null && $input->slug !== "") {
    
        
    $buffer = 'slug(';
        $arg0 =\var_export($input->slug, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->type !== null && $input->type !== "") {
    
        
    $buffer = 'type(';
        $arg0 =\var_export($input->type, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->updated !== null) {
    
        
    $buffer = 'updated(';
        $arg0 =\var_export($input->updated, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->updatedBy !== null && $input->updatedBy !== "") {
    
        
    $buffer = 'updatedBy(';
        $arg0 =\var_export($input->updatedBy, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->url !== null && $input->url !== "") {
    
        
    $buffer = 'url(';
        $arg0 =\var_export($input->url, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->version !== null) {
    
        
    $buffer = 'version(';
        $arg0 =\var_export($input->version, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

