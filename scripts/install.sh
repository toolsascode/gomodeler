#!/usr/bin/env bash

# Install the Helm S3 Publisher CLI tool.
# https://github.com/toolsascode/gomodeler
#
# Dependencies: curl, cut
#
# The version to install and the binary location can be passed in via VERSION and DESTDIR respectively.
# curl -fLSs https://raw.githubusercontent.com/toolsascode/gomodeler/main/scripts/install.sh | bash
#

PROJECT_NAME="gomodeler"
PROJECT_GH="toolsascode/$PROJECT_NAME"
# GitHub's URL for the latest release, will redirect.
GITHUB_BASE_URL="https://github.com/${PROJECT_GH}"
LATEST_URL="${GITHUB_BASE_URL}/releases/latest/"
DESTDIR="${DESTDIR:-/usr/local/bin}"

set -o errexit

function error {
	echo "An error occured installing the tool."
	echo "The contents of the directory $SCRATCH have been left in place to help to debug the issue."
}

validateCheckSum() {
	if ! grep -q "${1}" "${2}"; then
		echo "Invalid checksum" >/dev/stderr
		exit 1
	fi
	echo "Checksum is valid."
}

function getOS() {
	# Determine release filename.
	case "$(uname)" in
	Linux)
		OS='linux'
		;;
	Darwin)
		OS='darwin'
		;;
	*)
		echo "This operating system is not supported."
		exit 1
		;;
	esac
}

function getArch() {
	case "$(uname -m)" in
	aarch64 | arm64)
		ARCH='arm64'
		;;
	x86_64|amd64)
		ARCH="x86_64"
		;;
  	i386)
		ARCH="i386"
		;;
	*)
		echo "This architecture is not supported."
		exit 1
		;;
	esac
}

function run {

	echo "Starting installation."

	if [ -z "$VERSION" ]; then
		VERSION=$(curl -sLI -o /dev/null -w '%{url_effective}' "$LATEST_URL" | cut -d "v" -f 2)
	fi

	echo "Installing ${PROJECT_NAME} CLI v${VERSION}"

	# Run the script in a temporary directory that we know is empty.
	SCRATCH=$(mktemp -d || mktemp -d -t 'tmp')
	cd "$SCRATCH"

	trap error ERR

	getOS
	getArch

	RELEASE_URL="${GITHUB_BASE_URL}/releases/download/v${VERSION}/${PROJECT_NAME}_${OS}_${ARCH}.tar.gz"
	CHECKSUM_URL="${GITHUB_BASE_URL}/releases/download/v${VERSION}/${PROJECT_NAME}_${VERSION}_checksums.txt"

	mkdir "releases"
	RELEASE_FILENAME="releases/v${VERSION}.tar.gz"
	CHECKSUM_FILENAME="releases/v${VERSION}_checksums.txt"

	# Download binary and checksums files.
	(
		if command -v curl >/dev/null 2>&1; then
			curl -sSL "${RELEASE_URL}" -o "${RELEASE_FILENAME}"
			curl -sSL "${CHECKSUM_URL}" -o "${CHECKSUM_FILENAME}"
		elif command -v wget >/dev/null 2>&1; then
			wget -q "${RELEASE_URL}" -O "${RELEASE_FILENAME}"
			wget -q "${CHECKSUM_URL}" -O "${CHECKSUM_FILENAME}"
		else
			echo "ERROR: no curl or wget found to download files." >/dev/stderr
		fi
	)

	# Verify checksum.
	(
		if command -v sha256sum >/dev/null 2>&1; then
			checksum=$(sha256sum "${RELEASE_FILENAME}" | awk '{ print $1 }')
			validateCheckSum "${checksum}" "${CHECKSUM_FILENAME}"
		elif command -v openssl >/dev/null 2>&1; then
			checksum=$(openssl dgst -sha256 "${RELEASE_FILENAME}" | awk '{ print $2 }')
			validateCheckSum "${checksum}" "${CHECKSUM_FILENAME}"
		else
			echo "WARNING: no tool found to verify checksum" >/dev/stderr
		fi
	)

	# Download & unpack the release tarball.
	tar xzf "${RELEASE_FILENAME}" -C ./releases/

	echo "Installing to $DESTDIR"
	install "./releases/${PROJECT_NAME}" "$DESTDIR"

	command -v "${PROJECT_NAME}"

	# Delete the working directory when the install was successful.
	rm -r "$SCRATCH"
}

# Main
run
