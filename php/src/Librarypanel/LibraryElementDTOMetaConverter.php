<?php

namespace Grafana\Foundation\Librarypanel;

final class LibraryElementDTOMetaConverter
{
    public static function convert(\Grafana\Foundation\Librarypanel\LibraryElementDTOMeta $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaBuilder())',
        ];
            if ($input->folderName !== "") {
    
        
    $buffer = 'folderName(';
        $arg0 =\var_export($input->folderName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->folderUid !== "") {
    
        
    $buffer = 'folderUid(';
        $arg0 =\var_export($input->folderUid, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'connectedDashboards(';
        $arg0 =\var_export($input->connectedDashboards, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'created(';
        $arg0 =\var_export($input->created, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'updated(';
        $arg0 =\var_export($input->updated, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'createdBy(';
        $arg0 = \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUserConverter::convert($input->createdBy);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'updatedBy(';
        $arg0 = \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUserConverter::convert($input->updatedBy);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

