<?php

namespace Grafana\Foundation\Dashboard;

final class AnnotationPermissionConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\AnnotationPermission $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\AnnotationPermissionBuilder())',
        ];
            if ($input->dashboard !== null) {
    
        
    $buffer = 'dashboard(';
        $arg0 = \Grafana\Foundation\Dashboard\AnnotationActionsConverter::convert($input->dashboard);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->organization !== null) {
    
        
    $buffer = 'organization(';
        $arg0 = \Grafana\Foundation\Dashboard\AnnotationActionsConverter::convert($input->organization);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

