<?php

namespace Grafana\Foundation\Librarypanel;

final class LibraryPanelConverter
{
    public static function convert(\Grafana\Foundation\Librarypanel\LibraryPanel $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Librarypanel\LibraryPanelBuilder())',
        ];
            if ($input->folderUid !== null && $input->folderUid !== "") {
    
        
    $buffer = 'folderUid(';
        $arg0 =\var_export($input->folderUid, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->uid !== "") {
    
        
    $buffer = 'uid(';
        $arg0 =\var_export($input->uid, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->description !== null && $input->description !== "") {
    
        
    $buffer = 'description(';
        $arg0 =\var_export($input->description, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->type !== "") {
    
        
    $buffer = 'type(';
        $arg0 =\var_export($input->type, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->schemaVersion !== null) {
    
        
    $buffer = 'schemaVersion(';
        $arg0 =\var_export($input->schemaVersion, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'version(';
        $arg0 =\var_export($input->version, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'model(';
        $arg0 = \Grafana\Foundation\Librarypanel\LibrarypanelLibraryPanelModelConverter::convert($input->model);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->meta !== null) {
    
        
    $buffer = 'meta(';
        $arg0 = \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaConverter::convert($input->meta);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

