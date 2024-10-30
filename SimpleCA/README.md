# Simple Certificate Authority (CA) Server

## Private key

### Documentation PK

Generate a Private Key file `ca.key`. Asks for a **Password**.
TODO: Find out more about the Private Key

#### AES Encrytion

### Commands PK which is the CA

```bash
openssl genpkey -algorithm RSA -out ca.key -aes256
```

## Create Certificate Signing Request (CSR)

### Documentation CSR

### Commands CSR

```bash
openssl req -new -key private.key -out csr.pem
```

## Create Client Certificate

### Documentation Client Certificate

### Commands Client Certificate

```bash
openssl req -x509 -days 365 -key private.key -in csr.pem -out client.crt
```

## Certificate fileextenstion

## The `.der` extension

- DER usually describes a encoded Certificate or a CMS
(Content management system) container.
- is also often stored in a `.p7` file
- Strucutre is described using the [ASN.1](https://en.wikipedia.org/wiki/ASN.1)
data representation language.

## The `.pem` extension

Private Enhanced Mail

- Methode to encode Binary data as a string (ASCII Armor)
- Contains a _header_ and a _footer_ line
- data in the middle is [Base64](https://en.wiki.org/wiki/base64)

## The `.cer` or `.crt` extension

- Normaly it is a **DER** encoded data
