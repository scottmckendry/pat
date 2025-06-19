{
  description = "pat - Display images in the terminal";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs =
    { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };
    in
    {
      packages.${system}.default = pkgs.buildGoModule {
        pname = "pat";
        version = "0.3.3"; # x-release-please-version
        src = pkgs.fetchFromGitHub {
          owner = "scottmckendry";
          repo = "pat";
          rev = "v0.3.3"; # x-release-please-version
          sha256 = "sha256-AfRA2XlDy+metUaDGJW2cIqhWLPUvUQfJRy193vdz2w";
        };
        vendorHash = "sha256-t+t0e9mqC3NV3kN9o7Vg5zggso+y862Xztv574yxroU";
        goPackagePath = "github.com/scottmckendry/pat";
        subPackages = [ "." ];
        go = pkgs.go_1_24;

        meta = with pkgs.lib; {
          description = "Like cat but for pictures - display images in the terminal";
          homepage = "https://github.com/scottmckendry/pat";
          license = licenses.mit;
          maintainers = [ "scottmckendry" ];
        };
      };
      defaultPackage.${system} = self.packages.${system}.default;
    };
}
