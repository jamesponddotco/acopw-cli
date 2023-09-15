# acopw

`acopw` is a stupid simple, fast, and easy-to-use command-line utility
that allows you to generate cryptographically secure random passwords,
diceware passwords, PINs, and UUIDs.

The tool uses [crypto/rand](https://godocs.io/crypto/rand) by default
for generating random data and try to avoid selection bias. When
generating diceware passwords, it uses a [curated list with **over 23
thousand
words**](https://git.sr.ht/~jamesponddotco/acopw-go/blob/trunk/words/word-list.txt).

## Installation

### From source

First install the dependencies:

- Go 1.21 or above.
- make.
- [scdoc](https://git.sr.ht/~sircmpwn/scdoc).

Then compile and install:

```bash
make
sudo make install
```

## Usage

```bash
$ acopw --help
NAME:
   acopw - generate cryptographically secure random passwords

USAGE:
   acopw [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   diceware, d  generate a random diceware password
   random, r    generate a random password
   pin, p       generate a random numeric PIN
   uuid, u      generate a random UUID, a.k.a, UUIDv4

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

$ acopw random --length 32
Us\|0;'"::Z,K:=d0|CkdY.(kL]-k3<U

$ acopw diceware --length 6 --separator '-'
SOUTHAMPTON-isro-shocking-upstate-biography-branched
```

See _acopw(1)_ after installing for more information.

## Contributing

Anyone can help make `acopw` better. Send patches on the [mailing
list](https://lists.sr.ht/~jamesponddotco/acopw-devel) and report bugs
on the [issue tracker](https://todo.sr.ht/~jamesponddotco/acopw).

You must sign-off your work using `git commit --signoff`. Follow the
[Linux kernel developer's certificate of
origin](https://www.kernel.org/doc/html/latest/process/submitting-patches.html#sign-your-work-the-developer-s-certificate-of-origin)
for more details.

All contributions are made under [the GPL-2.0 license](LICENSE.md).

## Resources

The following resources are available:

- [Package documentation](https://godocs.io/git.sr.ht/~jamesponddotco/acopw-go).
- [Support and general discussions](https://lists.sr.ht/~jamesponddotco/acopw-discuss).
- [Patches and development related questions](https://lists.sr.ht/~jamesponddotco/acopw-devel).
- [Instructions on how to prepare patches](https://git-send-email.io/).
- [Feature requests and bug reports](https://todo.sr.ht/~jamesponddotco/acopw).

---

Released under the [GPL-2.0 license](LICENSE.md).
