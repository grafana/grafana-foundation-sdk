<?php

namespace Grafana\Foundation\Dashboard;

final class AdHocVariableConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\VariableModel $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\AdHocVariableBuilder('.\var_export($input->name, true).'))',
        ];
            if ($input->id !== "" && $input->id !== "00000000-0000-0000-0000-000000000000") {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->label !== null && $input->label !== "") {
    
        
    $buffer = 'label(';
        $arg0 =\var_export($input->label, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'hide(';
        $arg0 ='\Grafana\Foundation\Dashboard\VariableHide::fromValue("'.$input->hide.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->description !== null && $input->description !== "") {
    
        
    $buffer = 'description(';
        $arg0 =\var_export($input->description, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\DataSourceRef('.(($input->datasource->type !== null) ? 'type: '.\var_export($input->datasource->type, true).', ' : '').''.(($input->datasource->uid !== null) ? 'uid: '.\var_export($input->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->allFormat !== null && $input->allFormat !== "") {
    
        
    $buffer = 'allFormat(';
        $arg0 =\var_export($input->allFormat, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

