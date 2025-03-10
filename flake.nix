{
  description = "Dev env for mgorepo";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem(system:
      let
        shellHook = builtins.readFile ./scripts/shell-hook.sh;
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShell = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            husky
            revive
            goimports-reviser
            mage
            go-mockery
            commitizen
            docker
            docker-compose
          ];

          shellHook = "${shellHook}";
        };
      }
    );
}
