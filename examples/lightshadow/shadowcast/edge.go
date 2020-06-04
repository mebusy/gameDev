package shadowcast

type Edge struct {
    Sx, Sy float64
    Ex, Ey float64
}

type Cell struct {
    Edge_id [4]int
    // Edge_exist [4]bool
    Exist bool
}


const (
    NORTH = iota
    SOUTH
    EAST
    WEST
)


