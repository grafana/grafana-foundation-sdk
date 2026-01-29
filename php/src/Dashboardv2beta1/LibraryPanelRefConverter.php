<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class LibraryPanelRefConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\LibraryPanelRef $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\LibraryPanelRefBuilder())',
        ];
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
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

        return \implode("\n\t->", $calls);
    }
}

