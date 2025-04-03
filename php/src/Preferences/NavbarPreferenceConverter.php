<?php

namespace Grafana\Foundation\Preferences;

final class NavbarPreferenceConverter
{
    public static function convert(\Grafana\Foundation\Preferences\NavbarPreference $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Preferences\NavbarPreferenceBuilder())',
        ];
            if (count($input->bookmarkUrls) >= 1) {
    
        
    $buffer = 'bookmarkUrls(';
        $tmparg0 = [];
        foreach ($input->bookmarkUrls as $arg1) {
        $tmpbookmarkUrlsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpbookmarkUrlsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

