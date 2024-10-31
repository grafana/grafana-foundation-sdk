<?php

namespace Grafana\Foundation\Publicdashboard;

final class PublicDashboardConverter
{
    public static function convert(\Grafana\Foundation\Publicdashboard\PublicDashboard $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Publicdashboard\PublicDashboardBuilder())',
        ];
            if ($input->uid !== "") {
    
        
    $buffer = 'uid(';
        $arg0 =\var_export($input->uid, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->dashboardUid !== "") {
    
        
    $buffer = 'dashboardUid(';
        $arg0 =\var_export($input->dashboardUid, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->accessToken !== null && $input->accessToken !== "") {
    
        
    $buffer = 'accessToken(';
        $arg0 =\var_export($input->accessToken, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'isEnabled(';
        $arg0 =\var_export($input->isEnabled, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'annotationsEnabled(';
        $arg0 =\var_export($input->annotationsEnabled, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'timeSelectionEnabled(';
        $arg0 =\var_export($input->timeSelectionEnabled, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

