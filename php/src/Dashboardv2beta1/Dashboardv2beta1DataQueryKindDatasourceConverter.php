<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class Dashboardv2beta1DataQueryKindDatasourceConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasource $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasourceBuilder())',
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

