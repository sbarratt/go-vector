package vector

import (
    "fmt"
    "math"
)

const min_to_breakup int = 7
const vector_print_format string = "%.3f"

type Vector struct {
    items []float64
}

func MakeVector(n int) (v Vector, e error){
    if n < 0 {
        return Vector{make([]float64, 0)}, fmt.Errorf("n cannot be negative: %d", n)
    }
    return Vector{make([]float64, n)}
}

func (v Vector) Set(i int, value float64) (val float64, e error) {
    n := len(v.items)
    if (i < 0 || i >= len(v.items)) {
        return value, fmt.Errorf("index %d out of range for vector of length %d", i, n)
    }
    v.items[i] = value
    return value, nil
}

func (v Vector) Get(i int) (val float64, e error) {
    n := len(v.items)
    if (i < 0 || i >= len(v.items)) {
        return 0, fmt.Errorf("index %d out of range for vector of length %d", i, n)
    }
    return v.items[i], nil
}

type operation func(float64, float64) float64

type operation_bool func(float64, float64) bool

func apply_op(v1 Vector, v2 Vector, op operation) (Vector, error) {
    v1_length := len(v1.items)
    v2_length := len(v2.items)
    if v1_length != v2_length {
        return Vector{make([]float64, 0)}, fmt.Errorf("vector length %d != %d", v1_length, v2_length)
    }

    result_v := Vector{make([]float64, v1_length)}

    for i := 0; i < v2_length; i++ {
        result_v.items[i] = op(v1.items[i], v2.items[i])
    }
    return result_v, nil
}

func Add(v1 Vector, v2 Vector) (Vector, error) {
    add_op := func(a float64, b float64) float64 { return a + b }
    v, e := apply_op(v1, v2, add_op)
    return v, e
}

func AddParallel(v1 Vector, v2 Vector) (Vector, error) {
    return Vector{make([]float64, 0)}, nil
}

func Sub(v1 Vector, v2 Vector) (Vector, error) {
    sub_op := func(a float64, b float64) float64 { return a - b }
    v, e := apply_op(v1, v2, sub_op)
    return v, e
}

func Mul(v1 Vector, v2 Vector) (Vector, error) {
    mul_op := func(a float64, b float64) float64 { return a * b }
    v, e := apply_op(v1, v2, mul_op)
    return v, e
}

type operation_pointwise func(float64) float64

func apply_op_pointwise(v1 Vector, op operation_pointwise) (Vector, error) {
    n := len(v1.items)
    result_v := Vector{make([]float64, n)}

    for i := 0; i < n; i++ {
        result_v.items[i] = op(v1.items[i])
    }
    return result_v, nil
}

func AddScalar(v1 Vector, x float64) (Vector, error) {
    add_op := func(a float64) float64 { return a + x }
    v, e := apply_op_pointwise(v1, add_op)
    return v, e
}

func SubScalar(v1 Vector, x float64) (Vector, error) {
    sub_op := func(a float64) float64 { return a - x }
    v, e := apply_op_pointwise(v1, sub_op)
    return v, e
}

func MulScalar(v1 Vector, x float64) (Vector, error) {
    mul_op := func(a float64) float64 { return a * x }
    v, e := apply_op_pointwise(v1, mul_op)
    return v, e
}

func (v Vector) Min() float64 {
    var min float64 = math.MaxFloat64
    for _, v := range v.items {
        min = math.Min(min, v)
    }
    return min
}

func (v Vector) Max() float64 {
    var max float64 = -math.MaxFloat64
    for _, v := range v.items {
        max = math.Max(max, v)
    }
    return max
}

func (v Vector) Mean() float64 {
    var mean float64 = 0.0
    for _, v := range v.items {
        mean += v
    }
    return mean/float64(len(v.items))
}

func (v Vector) Var() float64 {
    var variance float64 = 0.0
    mean := v.Mean()
    for _, v := range v.items {
        diff := v-mean
        variance += diff*diff
    }
    return variance/float64(len(v.items))
}

func (v Vector) Std() float64 {
    return math.Sqrt(v.Var())
}

func (v Vector) ToString() string {
    var str string
    n := len(v.items)
    if n > min_to_breakup {
        first_3 := v.items[:3]
        last_3 := v.items[n-3:n]

        str += "["
        for i, e := range first_3 {
            if i != 0 {
                str += ", "
            }
            str += fmt.Sprintf(vector_print_format, e)
        }
        str += " ... "
        for i, e := range last_3 {
            if i != 0 {
                str += ", "
            }
            str += fmt.Sprintf(vector_print_format, e)
        }
        str += "]"
    } else {
        str += fmt.Sprintf("[")
        for i, e := range v.items {
            if i != 0 {
                str += ", "
            }
            str += fmt.Sprintf(vector_print_format, e)
        }
        str += fmt.Sprintf("]")
    }

    return str
}


func main() {
    v1 := Vector{[]float64 {1, 2, 3, 4}}
    v2 := Vector{[]float64 {-1, -2, -3, -4}}

    fmt.Println("Min (v1): ", v1.Min())
    fmt.Println("Max (v1): ", v1.Max())
    fmt.Println("Mean (v1): ", v1.Mean())
    fmt.Println("Std (v1): ", v1.Std())
    fmt.Println("Var (v1): ", v1.Var())

    v := Vector{}
    fmt.Println(v1.ToString())
    fmt.Println(v2.ToString())
    v, _ = Add(v1, v2)
    fmt.Println("Add:", v.ToString())
    v, _ = Sub(v1, v2)
    fmt.Println("Sub:", v.ToString())
    v, _ = Mul(v1, v2)
    fmt.Println("Mult:", v.ToString())
}