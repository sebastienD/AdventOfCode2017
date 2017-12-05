use strict;
use warnings;

my $file = 'input.txt';
open my $info, $file or die "Could not open $file: $!";

my @lines = ();
while( my $line = <$info>)  {   
    print $line;
    push(@lines, $line)
    #last if $. == 2;
}

close $info;

my $currenti = 0;
my $steps = 0;

while ($currenti < @lines) {
    my $v = $lines[$currenti];
    my $lasti = $currenti;
    $currenti = $v + $currenti;
    $lines[$lasti]++;
    $steps++;
}

print "Steps $steps\n";

