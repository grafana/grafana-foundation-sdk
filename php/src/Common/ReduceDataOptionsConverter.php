<?php

namespace Grafana\Foundation\Common;

final class ReduceDataOptionsConverter
{
    public static function convert(\Grafana\Foundation\Common\ReduceDataOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\ReduceDataOptionsBuilder())',
        ];
            if ($input->values !== null) {
    
        
    $buffer = 'values(';
        $arg0 =\var_export($input->values, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->limit !== null) {
    
        
    $buffer = 'limit(';
        $arg0 =\var_export($input->limit, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->calcs) >= 1) {
    
        
    $buffer = 'calcs(';
        $tmparg0 = [];
        foreach ($input->calcs as $arg1) {
        $tmpcalcsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpcalcsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fields !== null && $input->fields !== "") {
    
        
    $buffer = 'fields(';
        $arg0 =\var_export($input->fields, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

