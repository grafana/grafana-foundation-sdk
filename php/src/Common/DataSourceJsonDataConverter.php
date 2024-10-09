<?php

namespace Grafana\Foundation\Common;

final class DataSourceJsonDataConverter
{
    public static function convert(\Grafana\Foundation\Common\DataSourceJsonData $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\DataSourceJsonDataBuilder())',
        ];
            if ($input->authType !== null && $input->authType !== "") {
    
        
    $buffer = 'authType(';
        $arg0 =\var_export($input->authType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->defaultRegion !== null && $input->defaultRegion !== "") {
    
        
    $buffer = 'defaultRegion(';
        $arg0 =\var_export($input->defaultRegion, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->profile !== null && $input->profile !== "") {
    
        
    $buffer = 'profile(';
        $arg0 =\var_export($input->profile, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->manageAlerts !== null) {
    
        
    $buffer = 'manageAlerts(';
        $arg0 =\var_export($input->manageAlerts, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->alertmanagerUid !== null && $input->alertmanagerUid !== "") {
    
        
    $buffer = 'alertmanagerUid(';
        $arg0 =\var_export($input->alertmanagerUid, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

