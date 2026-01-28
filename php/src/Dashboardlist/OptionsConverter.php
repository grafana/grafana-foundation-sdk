<?php

namespace Grafana\Foundation\Dashboardlist;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Dashboardlist\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardlist\OptionsBuilder())',
        ];
            if ($input->keepTime !== false) {
    
        
    $buffer = 'keepTime(';
        $arg0 =\var_export($input->keepTime, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->includeVars !== false) {
    
        
    $buffer = 'includeVars(';
        $arg0 =\var_export($input->includeVars, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showStarred !== true) {
    
        
    $buffer = 'showStarred(';
        $arg0 =\var_export($input->showStarred, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showRecentlyViewed !== false) {
    
        
    $buffer = 'showRecentlyViewed(';
        $arg0 =\var_export($input->showRecentlyViewed, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showSearch !== false) {
    
        
    $buffer = 'showSearch(';
        $arg0 =\var_export($input->showSearch, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showHeadings !== true) {
    
        
    $buffer = 'showHeadings(';
        $arg0 =\var_export($input->showHeadings, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showFolderNames !== true) {
    
        
    $buffer = 'showFolderNames(';
        $arg0 =\var_export($input->showFolderNames, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->maxItems !== 10) {
    
        
    $buffer = 'maxItems(';
        $arg0 =\var_export($input->maxItems, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->query !== "") {
    
        
    $buffer = 'query(';
        $arg0 =\var_export($input->query, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->tags) >= 1) {
    
        
    $buffer = 'tags(';
        $tmparg0 = [];
        foreach ($input->tags as $arg1) {
        $tmptagsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmptagsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
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
            if ($input->folderUID !== null && $input->folderUID !== "") {
    
        
    $buffer = 'folderUID(';
        $arg0 =\var_export($input->folderUID, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

