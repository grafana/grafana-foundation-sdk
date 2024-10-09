<?php

namespace Grafana\Foundation\Librarypanel;

final class LibraryElementDTOMetaUserConverter
{
    public static function convert(\Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUser $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Librarypanel\LibraryElementDTOMetaUserBuilder())',
        ];
            
    
        {
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
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
            if ($input->avatarUrl !== "") {
    
        
    $buffer = 'avatarUrl(';
        $arg0 =\var_export($input->avatarUrl, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

