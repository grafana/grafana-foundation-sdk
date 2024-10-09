<?php

namespace Grafana\Foundation\Dashboard;

final class AnnotationActionsConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\AnnotationActions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\AnnotationActionsBuilder())',
        ];
            if ($input->canAdd !== null) {
    
        
    $buffer = 'canAdd(';
        $arg0 =\var_export($input->canAdd, true);
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

        return \implode("\n\t->", $calls);
    }
}

