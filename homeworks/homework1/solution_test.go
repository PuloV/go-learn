package main

import "testing"

func TestFirstParsePath(t *testing.T) {
	path := "D/go/code/../src/warcluster/tests/first/../../"

	if parsePath(path) != "/D/go/src/warcluster/" {
		t.Error("Result path is ", parsePath(path))
	}
}

func TestSecondParsePath(t *testing.T) {
	path := "python/movies/episode1/../../lectures/lecture1/examples/../code/../../../mostImportant/MonthyPython/quotes/../"

	if parsePath(path) != "/python/mostImportant/MonthyPython/" {
		t.Error("Result path is ", parsePath(path))
	}
}

func TestThirdParsePath(t *testing.T) {
	path := "python/movies/episode1/../../../../../../lectures/lecture1/examples/../code/../../../mostImportant/MonthyPython/quotes/../../../../"

	if parsePath(path) != "/" {
		t.Error("Result path is ", parsePath(path))
	}
}

func TestSomeHiddenPaths(t *testing.T) {
	path := "python/.movies/.class/test/one/../two/../../../"

	if parsePath(path) != "/python/.movies/" {
		t.Error("Result path is ", parsePath(path))
	}
}

func TestEmptyPath(t *testing.T) {
	path := ""
	if parsePath(path) != "/" {
		t.Error("Result path is ", parsePath(path))
	}
}

func TestEmptyPathWithDot(t *testing.T) {
	path := "."
	if parsePath(path) != "/" {
		t.Error("Result path is ", parsePath(path))	
	}
}

func TestShouldReturnTheSamePath(t *testing.T) {
	path := "/home/ndyakov/"
	if parsePath(path) != path {
		t.Error("Result path is ", parsePath(path))
	}
}

func TestNumbersInPath(t *testing.T) {
	path := "/home/test/my1dir/../"
	if parsePath(path) != "/home/test/" {
		t.Error("Result path is ", parsePath(path))
	}
}

func TestSpacesInPath(t *testing.T) {
	path := "/home/test/my 1 dir/../mysecond2 dir/"
	if parsePath(path) != "/home/test/mysecond2 dir/" {
		t.Error("Result path is ", parsePath(path))
	}
}

func TestDoubleDots(t *testing.T) {
	path := "home/test/../../home/test1/../../test2/../../test2/"
	if parsePath(path) != "/test2/" {
		t.Error("Result path is ", parsePath(path))
	}
}

func TestWithNewLineChars(t *testing.T) {
	path := "/this/is/\not/../test/"
	if parsePath(path) != "/this/is/test/" {
		t.Error("Result path is ", parsePath(path))
	}
}

func TestDotsAndSpacesInPath(t *testing.T) {
	path := "home/1 . 2/..23/../test/"
	if parsePath(path) != "/home/1 . 2/test/" {
		t.Error("Result path is ", parsePath(path))
	}
}

func TestDotBetweenChars(t *testing.T) {
	path := "/home/test/this.dir/../test2/../../"
	if parsePath(path) != "/home/" {
		t.Error("Result path is ", parsePath(path))
	}
}

func TestAllTheDots(t *testing.T) {
	path := "/home/test/this.shit.has.a.lot.of.dots./../"
	if parsePath(path) != "/home/test/" {
		t.Error("Result path is ", parsePath(path))
	}
}
