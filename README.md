# TermDict

A CLI dictionary written in Go.

## Preview

```bash
$ dict -find happy 
Happy

----------------------------------
Meanings:
[noun]
• A happy event, thing, person, etc.
  
[noun]
• Preceded by the: happy people as a group.
  
[verb]
• Often followed by up: to become happy; to brighten up, to cheer up.
  
• Often followed by up: to make happy; to brighten, to cheer, to enliven.
...
```

## Build

For bash

```bash
cd termdict
go build -o dict
echo "alias dict='$(pwd)/dict'" >> ~/.bashrc
```

For zsh
```bash
cd termdict
go build -o dict
echo "alias dict='$(pwd)/dict'" >> ~/.zshrc
```
```