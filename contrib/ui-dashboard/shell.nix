let
  pkgs = import <nixpkgs> { overlays = import ../../nix/overlays.nix; };
  inherit (pkgs) poetry2nix;

  pythonEnv = poetry2nix.mkPoetryEnv {
    projectDir = ./.;
    overrides = poetry2nix.overrides.withDefaults (
      import ./overrides.nix { inherit pkgs; }
    );
  };

  # Prevent the entirety of hydra to be in $PATH/runtime closure
  # We only want the evaluator
  hydra-eval-jobs = pkgs.runCommand "hydra-eval-jobs-${pkgs.hydra-unstable.version}" { } ''
    mkdir -p $out/bin
    cp -s ${pkgs.hydra-unstable}/bin/hydra-eval-jobs $out/bin/
  '';

  nix-nar-unpack = pkgs.callPackage ../nix-nar-unpack { };

in
pkgs.mkShell {

  buildInputs = [
    pkgs.nixpkgs-fmt
    pkgs.poetry
    pythonEnv

    nix-nar-unpack

    hydra-eval-jobs
    pkgs.sqlite

    pkgs.diffoscope

    pkgs.hivemind

    pkgs.yajl
    pkgs.pkgconfig
  ];

}
