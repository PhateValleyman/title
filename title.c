#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define B "\e[01;34m"
#define G "\e[32m"
#define R "\e[31m"
#define Y "\e[01;33m"
#define N "\e[0m"

int main(int argc, char *argv[]) {
	if (argc > 1) {
		if (strcmp(argv[1], "-v") == 0 || strcmp(argv[1], "--version") == 0) {
			printf("%s%s%s v1.1\n", Y, argv[0], N);
			printf("by %sPhateValleyman%s\n", Y, N);
			printf("%sJonas.Ned@outlook.com%s\n", G, N);
			exit(0);
		} else if (strcmp(argv[1], "-?") == 0 || strcmp(argv[1], "-h") == 0 || strcmp(argv[1], "--help") == 0) {
			printf("Usage: %s%s%s '%sTITLE%s'\n", Y, argv[0], N, G, N);
			printf("Set terminal title to '%sTITLE%s'\n", G, N);
			printf("Options:\n");
			printf("  -h, --help\t\tShow this help message.\n");
			printf("  -v, --version\t\tDisplay script version.\n");
			exit(0);
		}
	}

	if (argc == 1) {
		printf("\033]0;\007");
		printf("%sTitle unset%s\n", Y, N);
	}

	if (argc == 2) {
		printf("\033]0;%s\007", argv[1]);
		printf("Title set to:\n'%s%s%s'\n", G, argv[1], N);
	}

	if (argc > 2) {
		char title[1024] = "";
		for (int i = 1; i < argc; i++) {
			strcat(title, argv[i]);
			strcat(title, " ");
		}
		title[strlen(title) - 1] = '\0'; // Removing trailing space
		printf("\033]0;%s\007", title);
		printf("Title set to:\n'%s%s%s'\n", G, title, N);
	}

	return 0;
}
