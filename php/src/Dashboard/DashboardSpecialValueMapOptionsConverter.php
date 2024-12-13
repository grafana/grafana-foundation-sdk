<?php

namespace Grafana\Foundation\Dashboard;

final class DashboardSpecialValueMapOptionsConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptionsBuilder())',
        ];
            
    
        {
    $buffer = 'match(';
        $arg0 ='\Grafana\Foundation\Dashboard\SpecialValueMatch::fromValue("'.$input->match.'")';
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

