OS=""
ARCH=""

case $(uname -m) in
  arm64)    ARCH="arm64" ;;
  aarch64)  ARCH="arm64" ;;
  x86_64)   ARCH="amd64" ;;
  *)        ARCH=""
esac

if [ "$ARCH" = "" ] ; then
  echo "Cross-platform Fixie launcher expected arm64, aarch64, or x86_64. For other architectures, you can build from source: https://github.com/usefixie/fixie-cli"
  exit 1
fi

case "$OSTYPE" in
  darwin*)  OS="macos" ;; 
  linux*)   OS="linux" ;;
  msys*)    OS="windows" ;;
  cygwin*)  OS="windows" ;;
  *)        OS=""
esac

if [ "$OS" = "windows" ] ; then
  echo "Cross-platform Fixie launcher does not currently support Windows, but you can download a prebuilt binary for Windows: https://github.com/usefixie/fixie-cli"
  exit 1
elif [ "$OS" = "" ] ; then
  echo "Unsupported OS, but you may be able to build the Fixie CLI from source for your platform: https://github.com/usefixie/fixie-cli"
  exit 1
fi

FIXIE_BINARY="./fixie-$OS-$ARCH"
$FIXIE_BINARY $@
