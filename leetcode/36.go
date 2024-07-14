package leetcode

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		line := board[i]
		for j, c := range line {
			if c != '.' {
				// horizental verification
				if j < 9 {
					vrow := board[i]
					for _, vc := range vrow[j+1:] {
						if vc != '.' {
							if vc == c {
								return false
							}
						}
					}
				}

				// vertical verification
				if i < 9 {
					for vi, vrow := range board {
						if vi == i {
							continue
						}
						vc := vrow[j]
						if vc != '.' {
							if vc == c {
								return false
							}
						}
					}
				}

				// square verification
				si := i / 3 * 3
				fi := (i/3 + 1) * 3
				sj := j / 3 * 3
				fj := (j/3 + 1) * 3
				for ni := si; ni < fi; ni++ {
					nrow := board[ni]
					for nj, vc := range nrow[sj:fj] {
						if ni != i && nj != j && vc != '.' {
							if vc == c {
								return false
							}
						}
					}
				}
			}
		}
	}
	return true
}
