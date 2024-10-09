<?php

namespace Grafana\Foundation\Dashboard;

final class SnapshotConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\Snapshot $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\SnapshotBuilder())',
        ];
            if ($input->expires !== "") {
    
        
    $buffer = 'expires(';
        $arg0 =\var_export($input->expires, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'external(';
        $arg0 =\var_export($input->external, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->externalUrl !== "") {
    
        
    $buffer = 'externalUrl(';
        $arg0 =\var_export($input->externalUrl, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->key !== "") {
    
        
    $buffer = 'key(';
        $arg0 =\var_export($input->key, true);
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
            
    
        {
    $buffer = 'orgId(';
        $arg0 =\var_export($input->orgId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->url !== null && $input->url !== "") {
    
        
    $buffer = 'url(';
        $arg0 =\var_export($input->url, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->dashboard !== null) {
    
        
    $buffer = 'dashboard(';
        $arg0 = \Grafana\Foundation\Dashboard\DashboardConverter::convert($input->dashboard);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

