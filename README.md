# postcardgolf

Utilities to help me create some code golf'd ASCII art with a simple RLE compression.

The output of this can be fed to the following program:

```
echo ... | \
  perl -ne '@_=map{ord($_)-33}/\S/g;while(@_)
   {$a=pop@_;$a<32&&$a!=10&&($l=$a,$a=pop@_);
   $l=print chr($a)x$l}'
```
