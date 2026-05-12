# `tb-override` - White Labeling for ThingsBoard CE
<p align="center">
  <img src="https://github.com/user-attachments/assets/48720da6-9b9f-4792-bc1b-fc993f3522d8" style="background: transparent;" />
</p>

---

## tb-override lets you:
- Change colors, borders, radius without rebuilding anything.
- Replace logos cleanly and control their sizes.
- Inject custom themes dynamically.
- customize Angular Material/MDC components
- Save, share, export themes!

## How it works:
NGINX proxies your ThingsBoard instance and injects:
```
<link rel="stylesheet" href="/custom.css">
<link rel="stylesheet" href="/rules.css">
```
Into the page before it reaches the UI.
This allows for fast, runtime UI modifications on the go!

## Use the CLI to manage your themes
```
../tb-override ➜  ./tb-override -h
White-labeling for ThingsBoard Community Edition

Usage:
  tb-override [command]

Available Commands:
  detect      Detects the tools necessary for tb-override to work
  help        Help about any command
  list        List of common CSS properties of ThingsBoard CE
  setup       Setup the required directories and files required for tb-override to work
  theme       Perform actions on themes

Flags:
  -h, --help   help for tb-override

Use "tb-override [command] --help" for more information about a command.
../tb-override ➜  
```
## Instructions:
> COMING SOON
