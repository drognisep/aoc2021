It's not too often that I get the chance to use bitwise logic, so it was nice to exercise that aspect of the language.

Since the test case and the problem input have different binary digit lengths, I knew I had to make the solution work with variable size inputs.
Since all input elements have the same digit length, I checked the size of the first element to get that length.

Why is that important?

Mainly because having that number on hand informs the bit shifting and screening logic used to select and represent the scope of binary values.
All binary values are represented with uint64, which has more bits than are being used for either case.
The easiest way to select the right number of bits is with the "fScreen" which represents a 64-bit number with all bits set to 1.
With this I could work with *any* 64 bit (or less) number and then easily convert gamma to epsilon by XOR'ing against the bit screen and then using it to unset unused bits that fall outside of the digit length.
Since I'm using an unsigned type for the representation there isn't the issue of a "sticky" sign bit to deal with.
So right shifting the "fScreen" left fills with zeros and makes it easy to filter out unused bits.

With bit selection and screening done, it was pretty straightforward to build up a representation of digit statistics and solve the problem.

[Back](../README.md)
