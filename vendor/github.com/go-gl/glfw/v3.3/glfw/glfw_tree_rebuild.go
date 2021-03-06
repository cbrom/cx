package glfw

//go:generate ../../scripts/glfw_tree_rebuild.sh

// upstreamTreeSHA is a recursive hash of the full contents of the upstream
// glfw, as generated by git (doesn't need to be committed) when you run `go
// generate` on this package. This exists to invalidate the build cache (see
// https://github.com/go-gl/glfw/issues/269), which is unaffected by C source
// inputs.
const upstreamTreeSHA = "4490c2c270a92046291b021c15e33340289b33db"
