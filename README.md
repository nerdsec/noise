# noise

`noise` is a tool to help you visual the bytes of a file. It creates a black and
white image from the input file by mapping each byte to a pixel. The image below
on the left was created from a regular JPG; the image on the right was created
from the same JPG after the file was encrypted.

![noise from a jpg](/example1.png) ![noise from the same file after it is encrypted](/example2.png)

We created this tool so we could visualize the files after encrypting them with
[goaes](https://github.com/nerdsec/goaes).
