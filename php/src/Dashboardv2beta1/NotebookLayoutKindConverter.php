<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class NotebookLayoutKindConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutKindBuilder())',
        ];
            
    
        {
    $buffer = 'spec(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutSpecConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

