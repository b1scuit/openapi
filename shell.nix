with import <nixpkgs> {};

stdenv.mkDerivation {
    name = "go";
    buildInputs = [
      cowsay
      lolcat
      go
    ];

    shellHook = ''
      echo "Go Shell" | cowsay | lolcat
    '';
}
