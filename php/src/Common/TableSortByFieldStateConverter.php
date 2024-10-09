<?php

namespace Grafana\Foundation\Common;

final class TableSortByFieldStateConverter
{
    public static function convert(\Grafana\Foundation\Common\TableSortByFieldState $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\TableSortByFieldStateBuilder())',
        ];
            if ($input->displayName !== "") {
    
        
    $buffer = 'displayName(';
        $arg0 =\var_export($input->displayName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->desc !== null) {
    
        
    $buffer = 'desc(';
        $arg0 =\var_export($input->desc, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

