# Sencord Installer

The Sencord Installer allows you to install the Vencord fork [Sencord, another cute Discord client mod](https://github.com/sinjs/sencord)

![image](https://s3.nigga.church/images/SencordInstaller.png)

## Download

- [Windows (GUI)](https://api.nigga.church/v1/releases/sencordinstaller/latest/download/windows-gui)
- [Windows (CLI)](https://api.nigga.church/v1/releases/sencordinstaller/latest/download/windows-cli)
- [macOS (GUI)](https://api.nigga.church/v1/releases/sencordinstaller/latest/download/macos-gui)
- [Linux (CLI)](https://api.nigga.church/v1/releases/sencordinstaller/latest/download/linux-cli)

## Building from source

### Prerequisites

You need to install the [Go programming language](https://go.dev/doc/install) and GCC, the GNU Compiler Collection (MinGW on Windows)

<details>
<summary>Additionally, if you're using Linux, you have to install some additional dependencies:</summary>

#### Base dependencies

```sh
apt install -y pkg-config libsdl2-dev libglx-dev libgl1-mesa-dev
dnf install pkg-config libGL-devel libXxf86vm-devel
```

#### X11 dependencies

```sh
apt install -y xorg-dev
dnf install libXcursor-devel libXi-devel libXinerama-devel libXrandr-devel
```

#### Wayland dependencies

```sh
apt install -y libwayland-dev libxkbcommon-dev wayland-protocols extra-cmake-modules
dnf install wayland-devel libxkbcommon-devel wayland-protocols-devel extra-cmake-modules
```

</details>

### Building

#### Install dependencies

```sh
go mod tidy
```

#### Build the GUI

##### Windows / Mac / Linux X11

```sh
go build
```

##### Linux Wayland

```sh
go build --tags wayland
```

#### Build the CLI

```
go build --tags cli
```

You might want to pass some flags to this command to get a better build.
See [the GitHub workflow](https://github.com/Vendicated/VencordInstaller/blob/main/.github/workflows/release.yml) for what flags I pass or if you want more precise instructions
