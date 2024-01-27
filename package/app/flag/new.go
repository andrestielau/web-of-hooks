package flag

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/samber/lo"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func castBool[T ~bool](in []T) []bool {
	out := make([]bool, len(in))
	for i, v := range in {
		out[i] = bool(v)
	}
	return out
}
func Bool[T ~bool](p *T, name string, value T, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.BoolVar((*bool)(unsafe.Pointer(p)), name, bool(value), usage) }
}
func BoolP[T ~bool](p *T, name string, value T, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.BoolVarP((*bool)(unsafe.Pointer(p)), name, sh, bool(value), usage) }
}
func BoolS[T ~bool](p *[]T, name string, value []T, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) {
		fs.BoolSliceVar((*[]bool)(unsafe.Pointer(p)), name, castBool[T](value), usage)
	}
}
func BoolSP[T ~bool](p *[]T, name string, value []T, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) {
		fs.BoolSliceVarP((*[]bool)(unsafe.Pointer(p)), name, sh, castBool[T](value), usage)
	}
}
func toString[T ~string](in []T) []string {
	out := make([]string, len(in))
	for i := range in {
		out[i] = string(in[i])
	}
	return out
}
func String[T ~string](p *T, name string, value T, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.StringVar((*string)(unsafe.Pointer(p)), name, string(value), usage) }
}
func StringP[T ~string](p *T, name string, value T, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.StringVarP((*string)(unsafe.Pointer(p)), name, sh, string(value), usage) }
}
func StringS[T ~string](p *[]T, name string, value []T, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) {
		fs.StringSliceVar((*[]string)(unsafe.Pointer(p)), name, toString(value), usage)
	}
}
func StringSP[T ~string](p *[]T, name string, value []T, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) {
		fs.StringSliceVarP((*[]string)(unsafe.Pointer(p)), name, sh, toString(value), usage)
	}
}
func StringA[T ~string](p *[]T, name string, value []T, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) {
		fs.StringArrayVar((*[]string)(unsafe.Pointer(p)), name, toString(value), usage)
	}
}
func StringAP[T ~string](p *[]T, name string, value []T, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) {
		fs.StringArrayVarP((*[]string)(unsafe.Pointer(p)), name, sh, toString(value), usage)
	}
}

func castStringString[M ~map[K]V, K, V ~string](in M) map[string]string {
	out := map[string]string{}
	for k, v := range in {
		out[string(k)] = string(v)
	}
	return out
}
func castStringInt[M ~map[K]V, K ~string, V ~int](in M) map[string]int {
	out := map[string]int{}
	for k, v := range in {
		out[string(k)] = int(v)
	}
	return out
}
func castStringInt64[M ~map[K]V, K ~string, V ~int64](in M) map[string]int64 {
	out := map[string]int64{}
	for k, v := range in {
		out[string(k)] = int64(v)
	}
	return out
}
func StringToString[M ~map[K]V, K, V ~string](p *M, name string, value M, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) {
		fs.StringToStringVar((*map[string]string)(unsafe.Pointer(p)), name, castStringString(value), usage)
	}
}
func StringToStringP[M ~map[K]V, K, V ~string](p *M, name string, value M, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) {
		fs.StringToStringVarP((*map[string]string)(unsafe.Pointer(p)), name, sh, castStringString(value), usage)
	}
}
func StringToInt[M ~map[K]V, K ~string, V ~int](p *M, name string, value M, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) {
		fs.StringToIntVar((*map[string]int)(unsafe.Pointer(p)), name, castStringInt(value), usage)
	}
}
func StringToIntP[M ~map[K]V, K ~string, V ~int](p *M, name string, value M, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) {
		fs.StringToIntVarP((*map[string]int)(unsafe.Pointer(p)), name, sh, castStringInt(value), usage)
	}
}
func StringToInt64[M ~map[K]V, K ~string, V ~int64](p *M, name string, value M, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) {
		fs.StringToInt64Var((*map[string]int64)(unsafe.Pointer(p)), name, castStringInt64(value), usage)
	}
}
func StringToInt64P[M ~map[K]V, K ~string, V ~int64](p *M, name string, value M, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) {
		fs.StringToInt64VarP((*map[string]int64)(unsafe.Pointer(p)), name, sh, castStringInt64(value), usage)
	}
}

func Int(p *int, name string, value int, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.IntVar(p, name, value, usage) }
}
func IntP(p *int, name string, value int, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.IntVarP(p, name, sh, value, usage) }
}
func IntS(p *[]int, name string, value []int, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.IntSliceVar(p, name, value, usage) }
}
func IntSP(p *[]int, name string, value []int, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.IntSliceVarP(p, name, sh, value, usage) }
}
func Int8(p *int8, name string, value int8, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Int8Var(p, name, value, usage) }
}
func Int8P(p *int8, name string, value int8, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Int8VarP(p, name, sh, value, usage) }
}
func Int16(p *int16, name string, value int16, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Int16Var(p, name, value, usage) }
}
func Int16P(p *int16, name string, value int16, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Int16VarP(p, name, sh, value, usage) }
}
func Int32(p *int32, name string, value int32, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Int32Var(p, name, value, usage) }
}
func Int32P(p *int32, name string, value int32, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Int32VarP(p, name, sh, value, usage) }
}
func Int32S(p *[]int32, name string, value []int32, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Int32SliceVar(p, name, value, usage) }
}
func Int32SP(p *[]int32, name string, value []int32, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Int32SliceVarP(p, name, sh, value, usage) }
}
func Int64(p *int64, name string, value int64, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Int64Var(p, name, value, usage) }
}
func Int64P(p *int64, name string, value int64, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Int64VarP(p, name, sh, value, usage) }
}
func Int64S(p *[]int64, name string, value []int64, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Int64SliceVar(p, name, value, usage) }
}
func Int64SP(p *[]int64, name string, value []int64, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Int64SliceVarP(p, name, sh, value, usage) }
}

func Float32(p *float32, name string, value float32, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Float32Var(p, name, value, usage) }
}
func Float32P(p *float32, name string, value float32, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Float32VarP(p, name, sh, value, usage) }
}
func Float32S(p *[]float32, name string, value []float32, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Float32SliceVar(p, name, value, usage) }
}
func Float32SP(p *[]float32, name string, value []float32, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Float32SliceVarP(p, name, sh, value, usage) }
}
func Float64(p *float64, name string, value float64, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Float64Var(p, name, value, usage) }
}
func Float64P(p *float64, name string, value float64, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Float64VarP(p, name, sh, value, usage) }
}
func Float64S(p *[]float64, name string, value []float64, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Float64SliceVar(p, name, value, usage) }
}
func Float64SP(p *[]float64, name string, value []float64, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.Float64SliceVarP(p, name, sh, value, usage) }
}

func Duration(p *time.Duration, name string, value time.Duration, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.DurationVar(p, name, value, usage) }
}
func DurationP(p *time.Duration, name string, value time.Duration, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.DurationVarP(p, name, sh, value, usage) }
}
func DurationS(p *[]time.Duration, name string, value []time.Duration, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.DurationSliceVar(p, name, value, usage) }
}
func DurationSP(p *[]time.Duration, name string, value []time.Duration, sh, usage string) func(*pflag.FlagSet) {
	return func(fs *pflag.FlagSet) { fs.DurationSliceVarP(p, name, sh, value, usage) }
}

func Bind(v *viper.Viper) func(*pflag.FlagSet) { return func(fs *pflag.FlagSet) { BindSet(fs, v) } }
func BindSet(fs *pflag.FlagSet, v *viper.Viper) {
	fs.VisitAll(func(f *pflag.Flag) {
		if !f.Changed && v.IsSet(f.Name) {
			lo.Must0(fs.Set(f.Name, fmt.Sprintf("%v", v.Get(f.Name))))
		}
	})
}
