<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class SwitchVariableKindConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\SwitchVariableKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\SwitchVariableKindBuilder())',
        ];
            
    
        {
    $buffer = 'spec(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\SwitchVariableSpecConverter::convert($input->spec);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

