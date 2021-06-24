exec > genimports.gomacro

if [ x"$GOROOT" = "x" ]; then
  GOROOT="/usr/local/go"
fi

echo "#!/usr/bin/env gomacro"
echo


find "$GOROOT/src" -type d | \
  sed -e "s,$GOROOT/src/,," -e "s,$GOROOT/src,," | \
  grep "[a-z]" | grep -v 'builtin\|cmd\|internal\|syscall\|testdata\|vendor' | \
  sort |
while read i; do
  echo ":unload \"$i\""
  echo "import _b \"$i\""
done

