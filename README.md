# envy

`envy` is a command-line application for fetching shared environment variables.

Values are stored securely in AWS Parameter Store, and can be saved to a local `.env` file or directly sourced into your shell.


## Installation

Releases are published on [Github][releases].

### with Homebrew 🍺

A Homebrew cask is available at [`haines/tap`][homebrew-tap].

```console
$ brew cask install haines/tap/envy
```

### with Docker 🐳

[Docker images][docker-repo] are tagged with release versions.
The `latest` tag follows the `master` branch.

```console
$ docker pull ahaines/envy:$version
```

### manually 🔧

`envy` is distributed as a static binary, so installation just requires it to be downloaded from the [releases] page, then made executable:

```console
$ curl -L $url -o /usr/local/bin/envy
$ chmod +x /usr/local/bin/envy
```

Binaries can be verified by running
```console
$ gpg --keyserver ha.pool.sks-keyservers.net --recv-keys 6E225DD62262D98AAC77F9CDB16A6F178227A23E
gpg: key B16A6F178227A23E: public key "Andrew Haines <andrew@haines.org.nz>" imported

$ curl -fsSL "${url}.asc" | gpg --verify - /usr/local/bin/envy
gpg: Signature made Tue Apr 10 11:18:05 2018 UTC
gpg:                using RSA key 6E225DD62262D98AAC77F9CDB16A6F178227A23E
gpg: Good signature from "Andrew Haines <andrew@haines.org.nz>" [unknown]
```


## Usage

```console
$ envy --input /path/to/template --output /path/to/result
```

To see details of all the command-line options, run
```console
$ envy --help
```

### Templates

The input template file is a [Go text template][go-text-template], which has access to the following functions in interpolations:

* `param "path" "to" "value"` - fetches a value from AWS Parameter store
* `quote "value"` - wraps a value in single quotes, escaping embedded single quotes with `'\''` (closing the string, concatenating a literal `'`, and re-opening the string)

For example,
```shell
export FOO={{ param "secrets" "foo" | quote }}
```
would render as
```shell
export FOO='bar'
```
if the value `bar` was stored in Parameter Store under the key `/secrets/foo`.

### File permissions

When writing to a file, `envy` ensures that its permissions are `600` (only accessible by its owner).
To customize the permissions, use e.g. `--chmod 640`.
To leave the permissions alone, use `--no-chmod`.

### Sourcing output directly

`envy` reads from stdin and writes to stdout by default, so you can source things directly into your shell (assuming you trust the origin of the template!):
```console
# Bash 4, Zsh
$ source <(envy --input /path/to/template)

# Bash 3
$ source /dev/stdin <<<"$(envy --input /path/to/template)"
```

### Authenticating with AWS

#### with environment variables

Credentials are taken from the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables, if set.

#### with a shared credentials file

If the access key environment variables aren't set, `envy` can use a shared credentials file, just like the `aws` command-line interface.

If the file is not located at the default path (`~/.aws/credentials`), its location can be specified with the `AWS_SHARED_CREDENTIALS_FILE` environment variable.

The profile to use can be set by either the `--profile` command-line option, or the `AWS_PROFILE` environment variable.
If no profile is specified, the `default` profile will be used (if it exists).

#### with an EC2 instance IAM role

When running on an EC2 instance, if the access key environment variables aren't set, `envy` can use the instance's IAM role to authenticate.
You don't need to manually specify any credentials in this case.


[docker-repo]: https://hub.docker.com/r/ahaines/envy/
[go-text-template]: https://golang.org/pkg/text/template/
[homebrew-tap]: https://github.com/haines/homebrew-tap
[releases]: https://github.com/haines/envy/releases
