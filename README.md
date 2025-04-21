# eye 🔐 👀
## God sees everything, except what's encrypted.

[![Test](https://github.com/eyepipe/eye/actions/workflows/test.yml/badge.svg)](https://github.com/eyepipe/eye/actions/workflows/test.yml)

> [!WARNING]
> *eye is currently in beta testing mode*.

> [!CAUTION]
> The current limit for an encrypted message sent to the **cloud** is 30 MiB, 
> with a TTL of 7 days.

Eye is a utility for encrypting and digitally signing **files** and **text** messages.
Unlike existing programs such as [age](https://github.com/FiloSottile/age) or gpg, eye
allows you to encrypt your message and **immediately SEND it to the cloud** or your S3
bucket. The message recipient can easily decrypt the message with a single command,
knowing the link to it and your public key. The link to the message can be securely
transmitted over any public communication channel.

Eye uses the most **reliable** cryptographic primitives currently available in
their **maximum** configuration: ECDH (**P521** curve), **SHA-512**, HKDF, **AES 256-bit**.
You can change the cryptographic configuration at your discretion.

## Why?

The current political situation raises doubts about existing communication channels,
as law enforcement agencies are monitoring us. Using third-party solutions to protect
private data, even within messaging apps like Signal or SimpleX, ensures the privacy
f your most personal information.

Eye widely uses the Unix pipeline ideology, where the output of one command can be
redirected to another.

## Installation

Eye is distributed as a [single binary](https://github.com/eyepipe/eye/releases) 
file for `Windows`, `macOS`, or `Linux` on `amd64` or `arm64` platforms.
The executable file can be downloaded from the "Releases" section on GitHub.
It is also available for execution through a [Docker](#docker) container.

## The Workflow

- [keygen](#keygen) generate a key pair (signing key, private key agreement)
- [sign](#sign) your public key and send it to your correspondent along with the signature.
- [encrypt](#encrypt) and sign the message, and send it to Bob
- [decrypt](#decrypt) and verify the message received from Bob
- [tips](#tips)

## Keygen

Generate your private key. It is used to decrypt messages that other users send to you,
as well as to sign messages (or files) that you send to others.

> [!TIP]
> It is common to use separate keys for each individual correspondent and/or conversation.

> When rotating keys, generate a new pair of private and public keys, sign the public
> key with your old private key, and send the new public key to your correspondent using any
> communication channel (even an insecure one).

```bash
# generate your private key (keep it secret)
eye keygen > alice.eye

# derive the public key from the private key
# (this should be shared with your correspondent or made publicly available)
cat alice.eye | eye public > alice.eye.pub 
```

> [!NOTE]
> For educational purposes, we will "simulate" our conversation partner Bob by presenting
> our keys as Bob's keys.

```bash
cp alice.eye bob.eye
cp alice.eye.pub bob.eye.pub
```

Example of a public key

```
-----BEGIN SCHEME PROTO-----
eyJzaWduX2FsZ28iOiJFQ0RTQS1QNTIxLVNIQTUxMiIsImtleV9hZ3JlZW1lbnRf
YWxnbyI6IkVDREgtUDUyMSIsImJsb2NrX2NpcGhlcl9hbGdvIjoiQUVTX0NUUl8y
NTYiLCJoYXNoX2tleV9kZXJpdmF0aW9uX2FsZ28iOiJIS0RGLVNIQTUxMiJ9
-----END SCHEME PROTO-----
-----BEGIN SIGNER PUBLIC KEY-----
MIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQB5pWDe6uMpouEx9otgME/PrVzv4OK
pFUZSrK7hvvrYJ+mn+05+DKq6pQWicHH99TMigBjDg93N3AErS5SVQv7aoIBsto/
b6aIjc1VXRYIU1MS3r7bks0gd+BIEg1Ort2kR5aJWd+kPn6nqgqB0Fh6isNSDu4a
yfMmzjwVkrc3qptXI6o=
-----END SIGNER PUBLIC KEY-----
-----BEGIN KEY AGREEMENT PUBLIC KEY-----
MIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQAUaVQqRTC9q8FqnfeQmFl+9f0o1tV
kHD+FH0YBlLmmF72ctRyLPDqYuOe2N1FWDIG4BlPRuc8+hIjgtu35akdB5wBSTRA
6PzbCyL5JZuu4cAw9kqsWlNDAXiNyhzlY5DcjXLNojvqkWl7Xd9kxFuiiPzrA9ok
cnEs8e9f1Fe7s8AoQo8=
-----END KEY AGREEMENT PUBLIC KEY-----
```

## Sign

> [!TIP]
> You can digitally sign any files or messages without encrypting them. The recipient
> can verify the digital signature using your public key.

Sign your public key and send it along with the signature to the message recipient
(in this example, you are Alice and your correspondent is Bob).

The initial exchange of public keys should take place over a secure communication
channel — ideally during an offline, in-person meeting.

```bash
cat alice.eye.pub | eye sign -i alice.eye > alice.eye.pub.sig
```

You can verify a signature using the verify command and a public key. In this
example, you're verifying your own signature using your own public key. To verify
a digital signature from Bob, you need to use Bob's public key instead of your own.

```bash
# verify a detached signature from a file 
cat alice.eye.pub | eye verify -p alice.eye.pub -sig alice.eye.pub.sig

# verify a detached signature passed as an argument 
cat alice.eye.pub | eye verify -p alice.eye.pub -sig-hex 3081870242008011b76f30a585
```

## Encrypt

This command will **encrypt** the message (or file) and **sign** it using Bob's
public key. The recipient will be able to retrieve and decrypt the message
(and verify your digital signature) using the link provided in the response
from the encrypt send command.

```bash
echo "Hello, word" | eye encrypt -i alice.eye -p bob.eye.pub send
> https://api.eyepipe.pw/v1/downloads/01964637-d514-7fee-b0ba-730094020000
```

> [!TIP]
> You can encrypt files (or messages) of any size without uploading them to a server.
> In this case, you should send the encrypted file and its signature to Bob through
> any desired communication channel (USB drive, email, messenger).

Example of a command for encrypting a file without uploading it to a server:

```bash
# Encrypt the file and sign it:
cat README.md | eye encrypt -i alice.eye -p bob.eye.pub > README.md.enc
cat README.md.enc | eye sign -i alice.eye > README.md.enc.sig 
```

> [!TIP]
> The `encrypt` command outputs the digital signature in HEX format to `stderr`.
> You can either copy it from the console or directly redirect the signature
> output from `stderr` to a separate file.

> [!WARNING]
> When redirecting stderr, always additionally check the command's exit status,
> for example, using `echo $?`.

```bash
cat README.md | eye encrypt -i alice.eye -p bob.eye.pub > REAMDE.md.enc 2> README.md.enc.sig
echo $?
```

## Decrypt

This command will stream-download the message from Bob, decrypt it, and verify its digital signature.
```bash
eye decrypt -i bob.eye -p alice.eye.pub https://api.eyepipe.pw/v1/downloads/0196483c-2f41-7ea2-b0df-1c7cc4cb0000
```

You can also decrypt and verify the signature of a local file.
```bash
# creates a decrypted file on disk (the signature can be provided via a separate signature file).
cat README.md.enc | eye decrypt -i bob.eye -p alice.eye.pub --sig README.md.enc.sig

# prints the decrypted message to the console (the signature is provided via a command-line flag)
cat README.md.enc | eye decrypt -i bob.eye -p alice.eye.pub --sig-hex 308187024200
```

## Docker

Using the eye command-line utility via `Docker` is very simple — you just need to mount
the current directory and redirect the `stdout` stream to the container's `stdin`.
Below is an example alias for using eye in [Docker](https://github.com/eyepipe/eye/pkgs/container/eye): 
all the examples described in the README work fully with this method of running it.

```bash
docker pull ghcr.io/eyepipe/eye
```

> [!WARNING]
> This kind of usage is not optimal and works significantly slower than directly
> executing the binary on your host. However, it is perfectly suitable for infrequent
> use and small volumes of data being transferred.

```bash
alias eye="docker run --rm -i -v "$(pwd):/app" ghcr.io/eyepipe/eye"
echo "hello" | eye hex
```

To start the local eye server, use the command:

```bash
docker run --rm -p3000:3000 --entrypoint=server ghcr.io/eyepipe/eye
```

## Tips
### Progress bar

During any signing, encryption, and decryption operations, it can be useful to monitor the
progress (especially for large files). Eye solves this problem in a Unix-style manner, using pipes.

Install the popular utility pv (pipeviewer) using `brew install pv` or `apt install pv`, and simply
pipe the stream of information through it.

```bash
cat /dev/random | pv | eye encrypt -i alice.eye -p bob.eye.pub > /dev/null
# pv will display progress information in the console.
# 2,83GiB 0:00:03 [ 826MiB/s] [            <=>      ]
```

### HEX encoding
By default, the ciphertext is in binary encoding. Sometimes it is useful to get the ciphertext as
text to send it via `SMS`, `email` or any messenger as is. This is possible by using the HEX format.

This command will encrypt the message and convert it into HEX text. The stream with the digital
signature from encrypt is redirected to a separate file via `2> message.sig`.

```bash
echo "Hello, word" | eye encrypt -i alice.eye -p bob.eye.pub 2> message.sig | eye hex

# save output to the ENV variable as an example
HEX=0a0000000114000000409beea4c7a11080712bea9f0054f735b76ffe54c8d63ea6c62c5c
6b7c656a14a575a62fd3af6701dcc76db5f235371ee2095c487a3ff38a507a147e7d642b
49e41e00000010111f1c99bd90fdbc62f1f92084d3f83b631f9971e0dffc91374c0cf06d
```

To decrypt such a message, it must first be decoded from HEX.

> [!TIP]
> eye always generates the digital signature both in `stderr` and as a file in HEX format.
> There is no need to encode it further into HEX format.

```bash
echo -n $HEX | eye hex --dec | eye decrypt -i bob.eye -p alice.eye.pub --sig message.sig
```

## Feedback

You can send an encrypted feedback via email or create an [issue](https://github.com/eyepipe/eye/issues) on GitHub by
encrypting it with the [repository's public key](./repository.eye.pub).
In the case of a GitHub issue, it might be convenient to encode it in [HEX](#hex-encoding).

```bash
echo "New Issue" | eye encrypt -i alice.eye -p https://raw.githubusercontent.com/eyepipe/eye/refs/heads/master/repository.eye.pub
```
