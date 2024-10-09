<?php

namespace Grafana\Foundation\Dashboard;

final class AnnotationPanelFilterConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\AnnotationPanelFilter $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\AnnotationPanelFilterBuilder())',
        ];
            if ($input->exclude !== null && $input->exclude !== false) {
    
        
    $buffer = 'exclude(';
        $arg0 =\var_export($input->exclude, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->ids) >= 1) {
    
        
    $buffer = 'ids(';
        $tmparg0 = [];
        foreach ($input->ids as $arg1) {
        $tmpidsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpidsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

