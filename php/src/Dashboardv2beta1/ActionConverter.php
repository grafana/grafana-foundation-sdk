<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class ActionConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\Action $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\ActionBuilder())',
        ];
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\ActionType::fromValue("'.$input->type.'")';
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
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\FetchOptionsConverter::convert($input->fetch);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->infinity !== null) {
    
        
    $buffer = 'infinity(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\InfinityOptionsConverter::convert($input->infinity);
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
        $tmpvariablesarg1 = \Grafana\Foundation\Dashboardv2beta1\ActionVariableConverter::convert($arg1);
        $tmparg0[] = $tmpvariablesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->style !== null) {
    
        
    $buffer = 'style(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1ActionStyleConverter::convert($input->style);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

