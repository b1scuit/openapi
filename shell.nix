with import <nixpkgs> {};

stdenv.mkDerivation {
    name = "go";
    buildInputs = [
      cowsay
      lolcat
      go
    ];

    shellHook = ''
      echo "Open API project shell (Golang)" | cowsay | lolcat
    '';
}
