acopw(1)

# NAME

acopw - generate cryptographically secure random passwords

# SYNOPSIS

*acopw* [GLOBAL OPTIONS] [COMMANDS] [OPTIONS]

# DESCRIPTION

*acopw* is a stupid simple, fast, and easy-to-use command-line utility that
allows you to generate cryptographically secure random passwords, diceware
passwords, PINs, and UUIDs.

# GLOBAL OPTIONS

*-h*, *--help*
	Show help message and quit.

*-v*, *--version*
	Show version number and quit.

# COMMANDS

*r*, *random* [OPTIONS]
	Generate a random password with optional character set restrictions.

	Options are:

	*-l*, *--length*
		Set the length of the password to generate. Default is 128.

	*-n*, *--numbers*
		Include numbers in the generated password. Default is true.

	*-s*, *--symbols*
		Include symbols in the generated password. Default is true.

	*-u*, *--uppercase*
		Include uppercase letters in the generated password. Default is true.

	*-L*, *--lowercase*
		Include lowercase letters in the generated password. Default is true.

	*-e*, *--exclude*
		Exclude a specific character from the generated password. Repeat to
		exclude more characters. Default is none.

*d*, *diceware* [OPTIONS]
	Generate a diceware password, also known as a passphrase.

	Options are:

	*-l*, *--length*
		Set the number of words to generate. Default is 8.

	*-s*, *--separator*
		Set the separator character between words. Default is a space.

	*-c*, *--capitalized*
		Set whether to capitalize a random word or not. Default is true.

*p*, *pin* [OPTIONS]
	Generate a numeric PIN.

	Options are:

	*-l*, *--length*
		Set the length of the PIN to generate. Default is 6.

*u*, *uuid*
	Generate a UUIDv4, the random variant of UUID.

# EXAMPLES

*Example 1. Generate a random password of length 24:*
	$ acopw random -l 24

*Example 2. Generate a random password of default length but exclude certain
characters:*
	$ acopw random -l 24 -e '@' -e '#'

*Example 3. Generate a diceware password with 12 words:*
	$ acopw diceware -l 12

*Example 4. Generate a diceware password with a custom separator:*
	$ acopw diceware -l 12 -s '-'

*Example 5. Generate a PIN of length 4:*
	$ acopw pin -l 4

*Example 6. Generate a UUIDv4:*
	$ acopw uuid

# AUTHORS

Maintained by James Pond <james@cipher.host>.
