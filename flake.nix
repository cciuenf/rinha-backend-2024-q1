{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs?rev=9a2dd8e4798be098877175e835eb8e23b85f2c33";
    flake-parts.url = "github:hercules-ci/flake-parts";
    systems.url = "github:nix-systems/default";
  };

  outputs = {
    flake-parts,
    systems,
    ...
  } @ inputs:
    flake-parts.lib.mkFlake {inherit inputs;} {
      systems = import systems;
      perSystem = {pkgs, ...}: {
        devShells.default = with pkgs;
          mkShell {
            name = "rinha-go";
            packages = with pkgs;
              [go_1_22 nginx postgresql]
              ++ lib.optional stdenv.isLinux [inotify-tools]
              ++ lib.optional stdenv.isDarwin [
                darwin.apple_sdk.frameworks.CoreServices
                darwin.apple_sdk.frameworks.CoreFoundation
              ];
          };
      };
    };
}
