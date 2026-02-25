<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class LibraryPanelConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\LibraryPanelBuilder())',
        ];
            
    
        {
    $buffer = 'id(';
        $arg0 =\var_export($input->spec->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->spec->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'libraryPanel(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\LibraryPanelRefConverter::convert($input->spec->libraryPanel);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

