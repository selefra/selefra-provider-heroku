#!/usr/bin/env bash

set -e

basepath=$(cd `dirname $0`; pwd)
cd $basepath/
sedi=(-i)
case "$(uname)" in
  # For macOS, use two parameters
  Darwin*) sedi=(-i "")
esac

version=v${1}
time=$(date "+%Y-%m-%d")
if [ -f "provider/heroku/metadata.yaml" ];then
  VERSION=`cat provider/heroku/metadata.yaml | grep 'latest-version' | awk -F ' ' '{print $2}'`
else
  VERSION="heroku"
  mkdir -p provider/heroku
fi
FOR=`cat selefra-provider-heroku* | awk -F '_' '{print $3,$4}' | awk -F '.' '{print $1}' |  sed "s# #_#g"`
if [ -d "provider/heroku/$version" ];then rm -rf provider/heroku/$version ; else echo "OK!"; fi && cp -r provider/template/version1 provider/heroku/$version

for f in $FOR; do
  echo "$f"
  darwin_arm64=`cat selefra-provider-heroku* | grep darwin_arm64 | awk -F ' ' '{print $1}'`
  darwin_amd64=`cat selefra-provider-heroku* | grep darwin_amd64 | awk -F ' ' '{print $1}'`
  windows_amd64=`cat selefra-provider-heroku* | grep windows_amd64 | awk -F ' ' '{print $1}'`
  linux_amd64=`cat selefra-provider-heroku* | grep linux_amd64 | awk -F ' ' '{print $1}'`
  linux_arm64=`cat selefra-provider-heroku* | grep linux_arm64 | awk -F ' ' '{print $1}'`
  windows_arm64=`cat selefra-provider-heroku* | grep windows_arm64 | awk -F ' ' '{print $1}'`
  sed "${sedi[@]}" "s#{{.PackageName}}#selefra-provider-heroku#g" provider/heroku/$version/supplement.yaml
  sed "${sedi[@]}" "s#{{.Source}}#https://github.com/selefra/selefra-provider-heroku#g" provider/heroku/$version/supplement.yaml
  sed "${sedi[@]}" "s#{{.CheckSumLinuxARM64}}#${linux_arm64}#g" provider/heroku/$version/supplement.yaml
  sed "${sedi[@]}" "s#{{.CheckSumLinuxAMD64}}#${linux_amd64}#g" provider/heroku/$version/supplement.yaml
  sed "${sedi[@]}" "s#{{.CheckSumWindowsARM64}}#${windows_arm64}#g" provider/heroku/$version/supplement.yaml
  sed "${sedi[@]}" "s#{{.CheckSumWindowsAMD64}}#${windows_amd64}#g" provider/heroku/$version/supplement.yaml
  sed "${sedi[@]}" "s#{{.CheckSumDarwinARM64}}#${darwin_arm64}#g" provider/heroku/$version/supplement.yaml
  sed "${sedi[@]}" "s#{{.CheckSumDarwinAMD64}}#${darwin_amd64}#g" provider/heroku/$version/supplement.yaml
done

if [[ "$VERSION" != "$version" ]]; then
  cp provider/template/metadata.yaml provider/template/metadata.yaml.bak
  sed "${sedi[@]}" "s#{{.ProviderName}}#heroku#g" provider/template/metadata.yaml 
  sed "${sedi[@]}" "s#{{.LatestVersion}}#${version}#g" provider/template/metadata.yaml
  sed "${sedi[@]}" "s#{{.LatestUpdated}}#${time}#g" provider/template/metadata.yaml
  sed "${sedi[@]}" "s#{{.Introduction}}#A Selefra provider for Amazon Web Services (heroku).#g" provider/template/metadata.yaml
  sed "${sedi[@]}" "s#{{.ProviderVersion}}#${version}#g" provider/template/metadata.yaml
  sed "${sedi[@]}" '6d' provider/template/metadata.yaml
  if [ -f "provider/heroku/metadata.yaml" ];then
    sed -n '/^ /p' provider/heroku/metadata.yaml >> provider/template/metadata.yaml
  fi
  echo "  - ${version}" >> provider/template/metadata.yaml
  cat provider/template/metadata.yaml > provider/heroku/metadata.yaml
  mv provider/template/metadata.yaml.bak provider/template/metadata.yaml
fi