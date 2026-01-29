<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class LibraryPanelKindSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKindSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\LibraryPanelKindSpecBuilder())',
        ];
            
    
        {
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'libraryPanel(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\LibraryPanelRefConverter::convert($input->libraryPanel);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

