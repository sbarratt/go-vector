package vector

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "math/rand"
)

func TestVectorSet(t *testing.T) {
    v := Vector{items: make([]float64, 4)}

    val, e := v.Set(0, 1)
    assert.Equal(t, v.items[0], 1.0)
    assert.Equal(t, val, 1.0)
    assert.Equal(t, e, nil)

    val, e = v.Set(-1, 1)
    assert.True(t, e != nil)

    val, e = v.Set(4, 1)
    assert.True(t, e != nil)
}

func TestVectorGet(t *testing.T) {
    v := Vector{items: make([]float64, 4)}

    v.Set(0, 1)
    v.Set(1, 2)
    v.Set(2, 3)
    v.Set(3, 4)

    val, e := v.Get(0)
    assert.Equal(t, val, 1.0)
    assert.Equal(t, e, nil)

    val, e = v.Get(1)
    assert.Equal(t, val, 2.0)
    assert.Equal(t, e, nil)

    val, e = v.Get(-1)
    assert.True(t, e != nil)

    val, e = v.Get(4)
    assert.True(t, e != nil)
}

func TestVectorOp(t *testing.T) {
    v1 := Vector{[]float64 {1, 2}}
    v2 := Vector{[]float64 {-3, 4}}

    v3, e := Add(v1, v2)

    assert.Equal(t, v3.items[0], -2.0)
    assert.Equal(t, v3.items[1], 6.0)
    assert.Equal(t, e, nil)

    v1 = Vector{[]float64 {1}}
    v2 = Vector{[]float64 {3, 4}}

    v3, e = Add(v1, v2)

    assert.True(t, e != nil)

    v1 = Vector{[]float64 {1, 2}}
    v2 = Vector{[]float64 {3}}

    v3, e = Add(v1, v2)

    assert.True(t, e != nil)
}

func TestVectorPointwiseOp(t *testing.T) {
    v1 := Vector{[]float64 {1, 2}}

    v2, e := AddScalar(v1, 2)

    assert.Equal(t, v2.items[0], 3.0)
    assert.Equal(t, v2.items[1], 4.0)
    assert.Equal(t, e, nil)
}

func TestVectorStats(t *testing.T) {
    v1 := Vector{[]float64 {3, 4}}

    min := v1.Min()
    max := v1.Max()
    mean := v1.Mean()
    variance := v1.Var()
    std := v1.Std()

    assert.Equal(t, min, 3.0)
    assert.Equal(t, max, 4.0)
    assert.Equal(t, mean, 3.5)
    assert.Equal(t, variance, .25)
    assert.Equal(t, std, .5)
}

const n int = 100000

func BenchmarkVectorAdd(b *testing.B) {
    v1 := Vector{make([]float64, n)}
    v2 := Vector{make([]float64, n)}
    for i := 0; i < n; i++ {
        v1.Set(i, (rand.Float64()-0.5)*200)
        v2.Set(i, (rand.Float64()-0.5)*200)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Add(v1, v2)
    }
}

func BenchmarkVectorSub(b *testing.B) {
    v1 := Vector{make([]float64, n)}
    v2 := Vector{make([]float64, n)}
    for i := 0; i < n; i++ {
        v1.Set(i, (rand.Float64()-0.5)*200)
        v2.Set(i, (rand.Float64()-0.5)*200)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Sub(v1, v2)
    }
}
func BenchmarkVectorMul(b *testing.B) {
    v1 := Vector{make([]float64, n)}
    v2 := Vector{make([]float64, n)}
    for i := 0; i < n; i++ {
        v1.Set(i, (rand.Float64()-0.5)*200)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Mul(v1, v2)
    }
}

func BenchmarkVectorAddScalar(b *testing.B) {
    v1 := Vector{make([]float64, n)}
    for i := 0; i < n; i++ {
        v1.Set(i, (rand.Float64()-0.5)*200)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        MulScalar(v1, 2)
    }
}
func BenchmarkVectorSubScalar(b *testing.B) {
    v1 := Vector{make([]float64, n)}
    for i := 0; i < n; i++ {
        v1.Set(i, (rand.Float64()-0.5)*200)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        SubScalar(v1, 2)
    }
}
func BenchmarkVectorMulScalar(b *testing.B) {
    v1 := Vector{make([]float64, n)}
    for i := 0; i < n; i++ {
        v1.Set(i, (rand.Float64()-0.5)*200)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        MulScalar(v1, 2)
    }
}

func BenchmarkVectorMin(b *testing.B) {
    v1 := Vector{make([]float64, n)}
    for i := 0; i < n; i++ {
        v1.Set(i, (rand.Float64()-0.5)*200)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        v1.Min()
    }
}
func BenchmarkVectorMax(b *testing.B) {
    v1 := Vector{make([]float64, n)}
    for i := 0; i < n; i++ {
        v1.Set(i, (rand.Float64()-0.5)*200)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        v1.Max()
    }
}
func BenchmarkVectorMean(b *testing.B) {
    v1 := Vector{make([]float64, n)}
    for i := 0; i < n; i++ {
        v1.Set(i, (rand.Float64()-0.5)*200)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        v1.Mean()
    }
}
func BenchmarkVectorVar(b *testing.B) {
    v1 := Vector{make([]float64, n)}
    for i := 0; i < n; i++ {
        v1.Set(i, (rand.Float64()-0.5)*200)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        v1.Var()
    }
}
func BenchmarkVectorStd(b *testing.B) {
    v1 := Vector{make([]float64, n)}
    for i := 0; i < n; i++ {
        v1.Set(i, (rand.Float64()-0.5)*200)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        v1.Std()
    }
}
