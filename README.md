# 🖼️ pat

_Like c**at** but for **p**ictures._

Display images in the terminal without relying on special protocols or custom fonts. Works on any terminal where [true color is supported](https://github.com/termstandard/colors) (most of them).

![image](https://github.com/scottmckendry/pat/assets/39483124/4315c4a2-6915-4ed4-813b-72a49b24a725)

## 📦 Installation

**Go**

```sh
go install github.com/scottmckendry/pat@latest
```

**Windows**

```sh
winget install scottmckendry.pat
```

**Nix**

Add to your flake inputs:

```nix
{
  inputs.pat.url = "github:scottmckendry/pat";
}
```

Then you can use it in your configuration:

```nix
{
  inputs,
  pkgs,
  ...
}:

{
  # Install as a system package
  environment.systemPackages = [
    inputs.pat.packages.${pkgs.system}.default
  ];

  # OR with home-manager
  home.packages = [
    inputs.pat.packages.${pkgs.system}.default
  ];
}
```

## 🚀 Usage

Pat accepts a single argument, the path to the image you want to display. For example:

```sh
pat image.jpg
```

This can be substituted with a URL:

```sh
pat https://example.com/image.jpg
```

For more options, run `pat --help`.

## 🤝 Contributing

Contributions of all kinds are welcome! If you encounter a bug, have a question, or have an idea for a new feature, please open an issue. If you want to implement something new or fix an issue, please open a pull request.
