package petl

import "testing"

func TestPipe(t *testing.T) {
	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	if 4 != <-out {
		t.Error("eror 1")
	}
	if 9 != <-out {
		t.Error("eror 1")
	}
}

func TestPipe2(t *testing.T) {
	out := sq(sq(gen(2, 3)))

	if 16 != <-out {
		t.Error("eror 2")
	}
	if 81 != <-out {
		t.Error("eror 2")
	}
}

func TestPipe3(t *testing.T) {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2.
	out := merge(c1, c2)

	if 4 != <-out {
		t.Error("eror 3")
	}
	if 9 != <-out {
		t.Error("eror 3")
	}
}

// func TestPipe4(t *testing.T) {
// 	in := gen(2, 3)

// 	// Distribute the sq work across two goroutines that both read from in.
// 	c1 := sq(in)
// 	c2 := sq(in)

// 	done := make(chan struct{}, 2)
// 	// Consume the merged output from c1 and c2.
// 	out := merge(done, c1, c2)

// 	if 4 != <-out {
// 		t.Error("eror 3")
// 	}
// 	if 9 != <-out {
// 		t.Error("eror 3")
// 	}
// }
