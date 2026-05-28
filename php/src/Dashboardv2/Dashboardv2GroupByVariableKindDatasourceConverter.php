<?php

namespace Grafana\Foundation\Dashboardv2;

final class Dashboardv2GroupByVariableKindDatasourceConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2\Dashboardv2GroupByVariableKindDatasource $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2\Dashboardv2GroupByVariableKindDatasourceBuilder())',
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

