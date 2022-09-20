#!/bin/bash

if [ -z "${BASH_VERSION:-}" ]
then
  abort "Bash is required to execute the fixie-wrench installer"
fi

cwd=$(pwd)
install_dir="$cwd/bin"

echo "🔧 Installing fixie-wrench"

# Create ./bin if it does not exist
[ -d "$install_dir" ] && echo "📁 Installing into existing directory $install_dir" || mkdir "$install_dir"

files=(
  "fixie-wrench"
  "fixie-wrench-linux-amd64"
  "fixie-wrench-linux-arm64"
  "fixie-wrench-macos-amd64"
  "fixie-wrench-macos-arm64"
)

for file in ${files[@]}; do
  echo -ne '🟦'
  url="https://github.com/usefixie/fixie-wrench/releases/latest/download/$file"
  curl -L -s "$url" --output "$install_dir/$file"
done

echo -ne "\r✅ fixie-wrench has been installed to $install_dir\n\n"

read -p "Add fixie-wrench to Git? (Y/n): " -n 1 confirm

if ! [[  "$confirm" = "n" || "$confirm" = "N" ]]; then
  for file in ${files[@]}; do
    git add "$install_dir/$file"
  done
  if [ -f "$procfile" ]; then
    git add "$procfile"
  fi 
  echo -e '\n✅ Added fixie-wrench. To commit this change, run `git commit`'
fi

echo -e "\n🚀 Install complete! For more information, see https://usefixie.com/documentation/socks\n"
