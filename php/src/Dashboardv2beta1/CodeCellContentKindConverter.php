<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class CodeCellContentKindConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\CodeCellContentKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\CodeCellContentKindBuilder())',
        ];
            
    
        {
    $buffer = 'spec(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\CodeCellContentSpecConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

