<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class V2QueryGroupSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\V2QueryGroupSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\V2QueryGroupSpecBuilder())',
        ];
            if (count($input->queries) >= 1) {
    
        
    $buffer = 'queries(';
        $tmparg0 = [];
        foreach ($input->queries as $arg1) {
        $tmpqueriesarg1 = \Grafana\Foundation\Dashboardv2beta1\TargetConverter::convert($arg1);
        $tmparg0[] = $tmpqueriesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->transformations) >= 1) {
    
        
    $buffer = 'transformations(';
        $tmparg0 = [];
        foreach ($input->transformations as $arg1) {
        $tmptransformationsarg1 = \Grafana\Foundation\Dashboardv2beta1\V2TransformationKindConverter::convert($arg1);
        $tmparg0[] = $tmptransformationsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'queryOptions(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\QueryOptionsSpecConverter::convert($input->queryOptions);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

