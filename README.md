# postcardgolf

Utilities to help me create some code golf'd ASCII art with a simple RLE compression.

## Usage

Edit `input.txt`, and adjust the hex values as necessary.

Type `make`.

The encoded string at the end of the `make` output can be fed to the following program:

```
echo '...' | \
  perl -ne '@_=map{ord($_)-32}/\S/g;while(@_)
   {$a=pop@_;$a<32&&$a!=10&&($l=$a,$a=pop@_);
   $l=print chr($a)x$l}'
```
