#!/data/data/com.termux/files/usr/bin/bash

case "$1" in
	-v | --version )
		echo -e "$0 v1.1"
		echo -e "by PhateValleyman"
		echo -e "Jonas.Ned@outlook.com"
		exit 0
		;;
	-h | --help )
		echo -e "Usage: $0 [title]"
		echo -e "Set the terminal title.\n"
		echo -e "Options:"
		echo -e "  -h, --help\t\tShow this help message."
		echo -e "  -v, --version\t\tDisplay script version."
		exit 0
		;;
esac

if [ -z "$1" ]; then
	echo -ne "\033]0;\007"
	echo -e "Title unset..."
fi

if [ "$#" -eq 1 ]; then
	echo -ne "\033]0;$1\007"
	echo -e "Title set to:\n'\e[01;33m$1\e[0m'"
fi

if [ "$#" -gt 1 ]; then
	echo -ne "\033]0;"$*"\007"
	echo -e "Title set to:\n'\e[01;33m$*\e[0m'"
fi
