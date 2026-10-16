<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class CellKindConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\CellKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\CellKindBuilder())',
        ];
            
    
        {
    $buffer = 'spec(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\CellSpecConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

