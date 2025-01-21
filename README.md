# Actionfile

This is a fast and idiot proof command runner. Easy and freindly config language. Don't judge my code. I'm new to go and learning. It's getting better i promise. But for now i really like the result.

Example. Create a `.actions` file

```txt
hello = echo hello_world
```

For now you can now have space between the command name. This will be fixed later on.

## Install

Install the easy way (with the latest version of go installed on your system or above 1.23.4)

```
go install github.com/lassejlv/action@latest
```

Then run `action --help` and to upgrade in the feature. Use `action --upgrade` 😀

#### The other way

Download a binary for your OS from the [releases](github.com/lassejlv/action/releases) page.

## Using BETA install script

### Linux and MacOS (may work for windows too. It should. It worked on my machine)

```bash
curl -fsSL https://raw.githubusercontent.com/lassejlv/action/main/scripts/install.sh | bash
```

To uninstall run

```bash
curl -fsSL https://raw.githubusercontent.com/lassejlv/action/main/scripts/uninstall.sh | bash
```

To upgrade run

```bash
curl -fsSL https://raw.githubusercontent.com/lassejlv/action/main/scripts/upgrade.sh | bash # or action --upgrade
```
