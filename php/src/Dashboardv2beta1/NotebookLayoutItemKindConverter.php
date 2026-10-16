<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class NotebookLayoutItemKindConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemKindBuilder())',
        ];
            
    
        {
    $buffer = 'spec(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemSpecConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

