package codesdk

type Compiler string

const (
	PHP5  Compiler = "php5"
	PHP7  Compiler = "php7"
	GCC32 Compiler = "gcc"
	GCC64 Compiler = "gcc64"
	JAVA  Compiler = "java"
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
			cmd: "gcc -m32",
			run: "./a.out",
		},
		GCC64: &compilerExt{
			cmd: "gcc",
			run: "./a.out",
		},
		JAVA: &compilerExt{
			file: "/usercode/Main.java",
			cmd: "java -jar /opt/ecj-4.2.2.jar -1.5 -encoding utf8",
			run: "java -Dfile.encoding=UTF-8 -classpath /usercode Main",
		},
	}
	return
}
