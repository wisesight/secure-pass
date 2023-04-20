# ttb Secure Pass

ttb Secure Pass is a command-line tool for encrypting and decrypting passwords using AES-256-CBC encryption with a 32-byte key and 16-byte initialization vector (IV). The tool takes in plaintext passwords or base64-encoded ciphertext passwords and outputs ciphertext passwords or plaintext passwords, respectively.

## Usage

Download it from the release page according to your operating system and architecture.

```
Usage: ./ttb-secure-pass (-e|--encrypt) plaintext OR (-d|--decrypt) ciphertext
```

The tool takes in two arguments:

- `-e` or `--encrypt`: encrypts the plaintext password and outputs the base64-encoded ciphertext password
- `-d` or `--decrypt`: decrypts the base64-encoded ciphertext password and outputs the plaintext password

Examples:

Encrypting a password:

```bash
./ttb-secure-pass -e mypassword
```

Decrypting a password:

```bash
./ttb-secure-pass -d YnJvd3NlcjpteXBhc3N3b3Jk
```

## Build

**if you want to build it by yourself.**

### Requirements

To use ttb Secure Pass, you need to have the following installed on your machine:

- Go 1.16 or later
- make

1. Clone the repository or download it as a ZIP file:

```bash
git clone https://github.com/example/ttb-secure-pass.git
```

2. Build the executable:

```bash
cd ttb-secure-pass
make build
```

## Todo

- Example in NodeJs
- Download page
- Improve Docs
- Unit-test
