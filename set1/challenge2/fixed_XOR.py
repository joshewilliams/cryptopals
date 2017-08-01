#!/usr/bin/python3

import codecs

string1 = '1c0111001f010100061a024b53535009181c'
string2 = '686974207468652062756c6c277320657965'

output = int(string1, 16) ^ int(string2, 16)
print('{:x}'.format(output))
