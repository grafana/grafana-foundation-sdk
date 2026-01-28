<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class Dashboardv2beta1GroupByVariableKindDatasourceConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1GroupByVariableKindDatasource $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1GroupByVariableKindDatasourceBuilder())',
        ];
            if ($input->name !== null && $input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

