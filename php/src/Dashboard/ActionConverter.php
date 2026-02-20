<?php

namespace Grafana\Foundation\Dashboard;

final class ActionConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\Action $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\ActionBuilder())',
        ];
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Dashboard\ActionType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fetch !== null) {
    
        
    $buffer = 'fetch(';
        $arg0 = \Grafana\Foundation\Dashboard\FetchOptionsConverter::convert($input->fetch);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->infinity !== null) {
    
        
    $buffer = 'infinity(';
        $arg0 = \Grafana\Foundation\Dashboard\InfinityOptionsConverter::convert($input->infinity);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->confirmation !== null && $input->confirmation !== "") {
    
        
    $buffer = 'confirmation(';
        $arg0 =\var_export($input->confirmation, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->oneClick !== null) {
    
        
    $buffer = 'oneClick(';
        $arg0 =\var_export($input->oneClick, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->variables !== null && count($input->variables) >= 1) {
    
        
    $buffer = 'variables(';
        $tmparg0 = [];
        foreach ($input->variables as $arg1) {
        $tmpvariablesarg1 = \Grafana\Foundation\Dashboard\ActionVariableConverter::convert($arg1);
        $tmparg0[] = $tmpvariablesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->style !== null) {
    
        
    $buffer = 'style(';
        $arg0 = \Grafana\Foundation\Dashboard\DashboardActionStyleConverter::convert($input->style);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

