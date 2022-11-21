# TermDict

A CLI dictionary written in Go.

![screenshot-windows-terminal](https://user-images.githubusercontent.com/70356237/203070960-8f310abc-99fd-4ad1-9c71-16891ab5a4f5.png)

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
## Special Thanks

This dictionary uses [freeDictionaryAPI](https://github.com/meetDeveloper/freeDictionaryAPI) as the source. This tool couldn't have done without the API.
