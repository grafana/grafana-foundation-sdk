<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class QueryGroupConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\QueryGroupKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\QueryGroupBuilder())',
        ];
            if (count($input->spec->queries) >= 1) {
    
        
    $buffer = 'targets(';
        $tmparg0 = [];
        foreach ($input->spec->queries as $arg1) {
        $tmpqueriesarg1 = \Grafana\Foundation\Dashboardv2beta1\TargetConverter::convert($arg1);
        $tmparg0[] = $tmpqueriesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->spec->transformations) >= 1) {
    
        
    $buffer = 'transformations(';
        $tmparg0 = [];
        foreach ($input->spec->transformations as $arg1) {
        $tmptransformationsarg1 = \Grafana\Foundation\Dashboardv2beta1\TransformationConverter::convert($arg1);
        $tmparg0[] = $tmptransformationsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'queryOptions(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\QueryOptionsSpecConverter::convert($input->spec->queryOptions);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

