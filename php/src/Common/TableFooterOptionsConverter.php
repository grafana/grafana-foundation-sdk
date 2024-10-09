<?php

namespace Grafana\Foundation\Common;

final class TableFooterOptionsConverter
{
    public static function convert(\Grafana\Foundation\Common\TableFooterOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\TableFooterOptionsBuilder())',
        ];
            
    
        {
    $buffer = 'show(';
        $arg0 =\var_export($input->show, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->reducer) >= 1) {
    
        
    $buffer = 'reducer(';
        $tmparg0 = [];
        foreach ($input->reducer as $arg1) {
        $tmpreducerarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpreducerarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
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
            if ($input->enablePagination !== null) {
    
        
    $buffer = 'enablePagination(';
        $arg0 =\var_export($input->enablePagination, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->countRows !== null) {
    
        
    $buffer = 'countRows(';
        $arg0 =\var_export($input->countRows, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

