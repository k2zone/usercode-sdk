package codesdk

type Compiler string

const (
	PHP5    Compiler = "php5"
	PHP7    Compiler = "php7"
	GCC32   Compiler = "gcc"
	GCC64   Compiler = "gcc64"
	JAVA    Compiler = "java"
	PYTHON2 Compiler = "python2"
	PYTHON3 Compiler = "python3"
	CPP32   Compiler = "cpp32"
	CPP64   Compiler = "cpp64"
	MCS     Compiler = "mcs"
	GOLANG  Compiler = "golang"
	LUA     Compiler = "lua"
	R       Compiler = "rscript"
	SHELL   Compiler = "shell"
	PERL    Compiler = "perl"
	NODE    Compiler = "node"
	SWIFT   Compiler = "swift"
)

type compilerExt struct {
	file string
	cmd  string
	run  string
}

type compMap map[Compiler]*compilerExt

func makeCompMap(cp *compMap) {
	*cp = compMap{
		PHP5: &compilerExt{
			cmd: "php5.6",
		},
		PHP7: &compilerExt{
			cmd: "php7.1",
		},
		GCC32: &compilerExt{
			file: "/usercode/file.c",
			cmd:  "gcc -m32",
			run:  "./a.out",
		},
		GCC64: &compilerExt{
			file: "/usercode/file.c",
			cmd:  "gcc",
			run:  "./a.out",
		},
		CPP32: &compilerExt{
			file: "/usercode/file.cpp",
			cmd:  "g++ -m32",
			run:  "./a.out",
		},
		CPP64: &compilerExt{
			file: "/usercode/file.cpp",
			cmd:  "g++",
			run:  "./a.out",
		},
		JAVA: &compilerExt{
			file: "/usercode/Main.java",
			cmd:  "java -jar /opt/ecj-4.2.2.jar -1.5 -encoding utf8",
			run:  "java -Dfile.encoding=UTF-8 -classpath /usercode Main",
		},
		PYTHON2: &compilerExt{
			cmd: "python2",
		},
		PYTHON3: &compilerExt{
			cmd: "python3",
		},
		MCS: &compilerExt{
			file: "/usercode/file.cs",
			cmd:  "mcs",
			run:  "mono /usercode/file.exe",
		},
		GOLANG: &compilerExt{
			file: "/usercode/file.go",
			cmd:  "go run",
		},
		LUA: &compilerExt{
			cmd: "lua",
		},
		R: &compilerExt{
			cmd: "Rscript",
		},
		SHELL: &compilerExt{
			cmd: "bash",
		},
		PERL: &compilerExt{
			cmd: "perl",
		},
		NODE: &compilerExt{
			cmd: "node",
		},
		SWIFT: &compilerExt{
			cmd: "/opt/swift-4.2.1-RELEASE-ubuntu18.04/usr/bin/swift",
		},
	}
	return
}
