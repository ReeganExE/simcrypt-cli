# simcrypt-cli

A simple cross platform cli for encrypt/decrypt data.
It uses AES-256 CFB encryption.

## Usage

Download the pre-built binary from https://github.com/ReeganExE/simcrypt-cli/releases/latest
> Supported: Linux, macOS, Windows, Raspberry Pi family

```sh
A simple cli to encrypt/decrypt data.

Options:
  -d                    Decrypt mode
  -p, --password string Password for encrypt/decrypt
                        environment SIMCRYPT_PASSWORD
  -f <file>             Read password from a file

Example:
  export SIMCRYPT_PASSWORD='a strong password'
  echo 'test data' | simcrypt | tee /dev/stderr | simcrypt -d
  # OR
  echo 'test data' | simcrypt -p 'a strong password' | tee /dev/stderr | simcrypt -d -p 'a strong password'
  # OR
  # Encrypt then decrypt
  simcrypt -p 'abc' <<<'test data' |  simcrypt -d -p abc
```
