module;

#include <fstream>
#include <cstring>
#include <print>
#include <string>
#include <iostream>

export module cli;
import weretype;

export {
	u8 h = 8;

	void cli(int argc, char** argv) {
		for (int i = 1; i < argc; i++) {
			if (std::strcmp(argv[i], "--setup") == 0) {
				std::string name;
				std::print("Project name: ");
				std::getline(std::cin, name);
				std::ofstream file(".weremake");
				file << "Project = \"" << name << "\"\n";
				file.close();
			}
		}
	}
}