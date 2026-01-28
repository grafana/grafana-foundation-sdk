<?php

namespace Grafana\Foundation\Logs;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Logs\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Logs\OptionsBuilder())',
        ];
            
    
        {
    $buffer = 'showLabels(';
        $arg0 =\var_export($input->showLabels, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'showCommonLabels(';
        $arg0 =\var_export($input->showCommonLabels, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'showTime(';
        $arg0 =\var_export($input->showTime, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'showLogContextToggle(';
        $arg0 =\var_export($input->showLogContextToggle, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'wrapLogMessage(';
        $arg0 =\var_export($input->wrapLogMessage, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'prettifyLogMessage(';
        $arg0 =\var_export($input->prettifyLogMessage, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'enableLogDetails(';
        $arg0 =\var_export($input->enableLogDetails, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'sortOrder(';
        $arg0 ='\Grafana\Foundation\Common\LogsSortOrder::fromValue("'.$input->sortOrder.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'dedupStrategy(';
        $arg0 ='\Grafana\Foundation\Common\LogsDedupStrategy::fromValue("'.$input->dedupStrategy.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

