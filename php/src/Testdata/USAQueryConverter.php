<?php

namespace Grafana\Foundation\Testdata;

final class USAQueryConverter
{
    public static function convert(\Grafana\Foundation\Testdata\USAQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Testdata\USAQueryBuilder())',
        ];
            if ($input->mode !== null && $input->mode !== "") {
    
        
    $buffer = 'mode(';
        $arg0 =\var_export($input->mode, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->period !== null && $input->period !== "") {
    
        
    $buffer = 'period(';
        $arg0 =\var_export($input->period, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fields !== null && count($input->fields) >= 1) {
    
        
    $buffer = 'fields(';
        $tmparg0 = [];
        foreach ($input->fields as $arg1) {
        $tmpfieldsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpfieldsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->states !== null && count($input->states) >= 1) {
    
        
    $buffer = 'states(';
        $tmparg0 = [];
        foreach ($input->states as $arg1) {
        $tmpstatesarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpstatesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

