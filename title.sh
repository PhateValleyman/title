#!/bin/bash

B="\e[01;34m"
G="\e[32m"
R="\e[31m"
Y="\e[01;33m"
N="\e[0m"


case "$1" in
	-v | --version )
		echo -e "${Y}$0${N} v1.1"
		echo -e "by ${Y}PhateValleyman${N}"
		echo -e "${G}Jonas.Ned@outlook.com${N}"
		exit 0
		;;
	-? | -h | --help )
		echo -e "Usage: ${Y}$0${N} '${G}TITLE${N}'"
		echo -e "Set terminal title to '${G}TITLE${N}'"
		echo -e "Options:"
		echo -e "  -h, --help\t\tShow this help message."
		echo -e "  -v, --version\t\tDisplay script version."
		exit 0
		;;
esac

if [ -z "$1" ]; then
	echo -ne "\033]0;\007"
	echo -e "${Y}Title unset${N}"
fi

if [ "$#" -eq 1 ]; then
	echo -ne "\033]0;$1\007"
	echo -e "Title set to:\n'${G}$1${N}'"
fi

if [ "$#" -gt 1 ]; then
	echo -ne "\033]0;"$*"\007"
	echo -e "Title set to:\n'${G}$*${N}'"
fi
