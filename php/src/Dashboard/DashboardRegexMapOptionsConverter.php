<?php

namespace Grafana\Foundation\Dashboard;

final class DashboardRegexMapOptionsConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\DashboardRegexMapOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\DashboardRegexMapOptionsBuilder())',
        ];
            if ($input->pattern !== "") {
    
        
    $buffer = 'pattern(';
        $arg0 =\var_export($input->pattern, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'result(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\ValueMappingResult('.(($input->result->text !== null) ? 'text: '.\var_export($input->result->text, true).', ' : '').''.(($input->result->color !== null) ? 'color: '.\var_export($input->result->color, true).', ' : '').''.(($input->result->icon !== null) ? 'icon: '.\var_export($input->result->icon, true).', ' : '').''.(($input->result->index !== null) ? 'index: '.\var_export($input->result->index, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

