package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var resize_code cu.Function

type resize_args struct {
	arg_dst    unsafe.Pointer
	arg_D0     int
	arg_D1     int
	arg_D2     int
	arg_src    unsafe.Pointer
	arg_S0     int
	arg_S1     int
	arg_S2     int
	arg_layer  int
	arg_scale1 int
	arg_scale2 int
	argptr     [11]unsafe.Pointer
}

// Wrapper for resize CUDA kernel, asynchronous.
func k_resize_async(dst unsafe.Pointer, D0 int, D1 int, D2 int, src unsafe.Pointer, S0 int, S1 int, S2 int, layer int, scale1 int, scale2 int, cfg *config, str int) {
	if resize_code == 0 {
		resize_code = fatbinLoad(resize_map, "resize")
	}

	var _a_ resize_args

	_a_.arg_dst = dst
	_a_.argptr[0] = unsafe.Pointer(&_a_.arg_dst)
	_a_.arg_D0 = D0
	_a_.argptr[1] = unsafe.Pointer(&_a_.arg_D0)
	_a_.arg_D1 = D1
	_a_.argptr[2] = unsafe.Pointer(&_a_.arg_D1)
	_a_.arg_D2 = D2
	_a_.argptr[3] = unsafe.Pointer(&_a_.arg_D2)
	_a_.arg_src = src
	_a_.argptr[4] = unsafe.Pointer(&_a_.arg_src)
	_a_.arg_S0 = S0
	_a_.argptr[5] = unsafe.Pointer(&_a_.arg_S0)
	_a_.arg_S1 = S1
	_a_.argptr[6] = unsafe.Pointer(&_a_.arg_S1)
	_a_.arg_S2 = S2
	_a_.argptr[7] = unsafe.Pointer(&_a_.arg_S2)
	_a_.arg_layer = layer
	_a_.argptr[8] = unsafe.Pointer(&_a_.arg_layer)
	_a_.arg_scale1 = scale1
	_a_.argptr[9] = unsafe.Pointer(&_a_.arg_scale1)
	_a_.arg_scale2 = scale2
	_a_.argptr[10] = unsafe.Pointer(&_a_.arg_scale2)

	args := _a_.argptr[:]
	cu.LaunchKernel(resize_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream[str], args)
}

// Wrapper for resize CUDA kernel, synchronized.
func k_resize(dst unsafe.Pointer, D0 int, D1 int, D2 int, src unsafe.Pointer, S0 int, S1 int, S2 int, layer int, scale1 int, scale2 int, cfg *config) {
	const stream = 0
	k_resize_async(dst, D0, D1, D2, src, S0, S1, S2, layer, scale1, scale2, cfg, stream)
	Sync(stream)
}

var resize_map = map[int]string{0: "",
	20: resize_ptx_20,
	30: resize_ptx_30,
	35: resize_ptx_35}

const (
	resize_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry resize(
	.param .u64 resize_param_0,
	.param .u32 resize_param_1,
	.param .u32 resize_param_2,
	.param .u32 resize_param_3,
	.param .u64 resize_param_4,
	.param .u32 resize_param_5,
	.param .u32 resize_param_6,
	.param .u32 resize_param_7,
	.param .u32 resize_param_8,
	.param .u32 resize_param_9,
	.param .u32 resize_param_10
)
{
	.reg .pred 	%p<11>;
	.reg .s32 	%r<54>;
	.reg .f32 	%f<21>;
	.reg .s64 	%rd<15>;


	ld.param.u64 	%rd5, [resize_param_0];
	ld.param.u32 	%r26, [resize_param_2];
	ld.param.u32 	%r20, [resize_param_3];
	ld.param.u64 	%rd6, [resize_param_4];
	ld.param.u32 	%r21, [resize_param_6];
	ld.param.u32 	%r22, [resize_param_7];
	ld.param.u32 	%r23, [resize_param_8];
	ld.param.u32 	%r24, [resize_param_9];
	ld.param.u32 	%r25, [resize_param_10];
	.loc 2 9 1
	mov.u32 	%r1, %ctaid.y;
	mov.u32 	%r2, %ntid.y;
	mov.u32 	%r3, %tid.y;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	.loc 2 10 1
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r27, %r5, %r6, %r7;
	.loc 2 12 1
	setp.lt.s32 	%p1, %r4, %r26;
	setp.lt.s32 	%p2, %r27, %r20;
	and.pred  	%p3, %p1, %p2;
	@!%p3 bra 	BB0_10;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 17 1
	setp.lt.s32 	%p4, %r24, 1;
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	.loc 2 17 1
	@%p4 bra 	BB0_9;

	.loc 2 18 1
	mul.lo.s32 	%r8, %r4, %r24;
	.loc 2 17 1
	mul.lo.s32 	%r30, %r25, %r27;
	.loc 2 20 1
	cvt.s64.s32 	%rd1, %r30;
	.loc 2 17 1
	mul.lo.s32 	%r32, %r24, %r4;
	mad.lo.s32 	%r33, %r23, %r21, %r32;
	mul.lo.s32 	%r50, %r22, %r33;
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	mov.u32 	%r51, 0;
	cvta.to.global.u64 	%rd9, %rd6;

BB0_3:
	.loc 2 17 1
	cvt.s64.s32 	%rd7, %r50;
	.loc 2 18 1
	add.s64 	%rd8, %rd1, %rd7;
	.loc 2 17 1
	shl.b64 	%rd10, %rd8, 2;
	add.s64 	%rd14, %rd9, %rd10;
	.loc 2 18 1
	add.s32 	%r12, %r51, %r8;
	.loc 2 20 1
	setp.lt.s32 	%p5, %r25, 1;
	@%p5 bra 	BB0_8;

	mov.u32 	%r53, 0;
	mov.u32 	%r52, %r30;

BB0_5:
	.loc 2 23 1
	mov.u32 	%r14, %r52;
	setp.lt.s32 	%p6, %r14, %r22;
	setp.lt.s32 	%p7, %r12, %r21;
	.loc 2 23 1
	and.pred  	%p8, %p7, %p6;
	@!%p8 bra 	BB0_7;
	bra.uni 	BB0_6;

BB0_6:
	.loc 2 24 1
	ld.global.f32 	%f17, [%rd14];
	add.f32 	%f20, %f20, %f17;
	.loc 2 25 1
	add.f32 	%f19, %f19, 0f3F800000;

BB0_7:
	.loc 2 20 1
	add.s32 	%r16, %r14, 1;
	add.s64 	%rd14, %rd14, 4;
	.loc 2 20 22
	add.s32 	%r53, %r53, 1;
	.loc 2 20 1
	setp.lt.s32 	%p9, %r53, %r25;
	mov.u32 	%r52, %r16;
	@%p9 bra 	BB0_5;

BB0_8:
	.loc 2 17 22
	add.s32 	%r51, %r51, 1;
	.loc 2 17 1
	setp.lt.s32 	%p10, %r51, %r24;
	add.s32 	%r50, %r50, %r22;
	@%p10 bra 	BB0_3;

BB0_9:
	.loc 2 29 1
	mad.lo.s32 	%r48, %r4, %r20, %r27;
	cvta.to.global.u64 	%rd11, %rd5;
	.loc 2 29 1
	mul.wide.s32 	%rd12, %r48, 4;
	add.s64 	%rd13, %rd11, %rd12;
	.loc 3 2399 3
	div.rn.f32 	%f18, %f20, %f19;
	.loc 2 29 1
	st.global.f32 	[%rd13], %f18;

BB0_10:
	.loc 2 31 2
	ret;
}


`
	resize_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry resize(
	.param .u64 resize_param_0,
	.param .u32 resize_param_1,
	.param .u32 resize_param_2,
	.param .u32 resize_param_3,
	.param .u64 resize_param_4,
	.param .u32 resize_param_5,
	.param .u32 resize_param_6,
	.param .u32 resize_param_7,
	.param .u32 resize_param_8,
	.param .u32 resize_param_9,
	.param .u32 resize_param_10
)
{
	.reg .pred 	%p<11>;
	.reg .s32 	%r<41>;
	.reg .f32 	%f<21>;
	.reg .s64 	%rd<15>;


	ld.param.u64 	%rd4, [resize_param_0];
	ld.param.u32 	%r26, [resize_param_2];
	ld.param.u32 	%r20, [resize_param_3];
	ld.param.u64 	%rd5, [resize_param_4];
	ld.param.u32 	%r21, [resize_param_6];
	ld.param.u32 	%r22, [resize_param_7];
	ld.param.u32 	%r23, [resize_param_8];
	ld.param.u32 	%r24, [resize_param_9];
	ld.param.u32 	%r25, [resize_param_10];
	.loc 2 9 1
	mov.u32 	%r1, %ctaid.y;
	mov.u32 	%r2, %ntid.y;
	mov.u32 	%r3, %tid.y;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	.loc 2 10 1
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 2 12 1
	setp.lt.s32 	%p1, %r4, %r26;
	setp.lt.s32 	%p2, %r8, %r20;
	and.pred  	%p3, %p1, %p2;
	@!%p3 bra 	BB0_10;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 17 1
	setp.lt.s32 	%p4, %r24, 1;
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	.loc 2 17 1
	@%p4 bra 	BB0_9;

	.loc 2 18 1
	mul.lo.s32 	%r9, %r4, %r24;
	.loc 2 17 1
	mul.lo.s32 	%r10, %r25, %r8;
	mul.lo.s32 	%r30, %r24, %r4;
	mad.lo.s32 	%r31, %r23, %r21, %r30;
	mul.lo.s32 	%r37, %r22, %r31;
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	mov.u32 	%r38, 0;
	cvta.to.global.u64 	%rd9, %rd5;

BB0_3:
	.loc 2 20 1
	cvt.s64.s32 	%rd6, %r10;
	.loc 2 17 1
	cvt.s64.s32 	%rd7, %r37;
	.loc 2 18 1
	add.s64 	%rd8, %rd6, %rd7;
	.loc 2 17 1
	shl.b64 	%rd10, %rd8, 2;
	add.s64 	%rd14, %rd9, %rd10;
	.loc 2 20 1
	setp.lt.s32 	%p5, %r25, 1;
	@%p5 bra 	BB0_8;

	mov.u32 	%r40, 0;
	mov.u32 	%r39, %r10;

BB0_5:
	.loc 2 18 1
	mov.u32 	%r14, %r39;
	add.s32 	%r33, %r38, %r9;
	setp.lt.s32 	%p6, %r33, %r21;
	.loc 2 23 1
	setp.lt.s32 	%p7, %r14, %r22;
	and.pred  	%p8, %p6, %p7;
	@!%p8 bra 	BB0_7;
	bra.uni 	BB0_6;

BB0_6:
	.loc 2 24 1
	ld.global.f32 	%f17, [%rd14];
	add.f32 	%f20, %f20, %f17;
	.loc 2 25 1
	add.f32 	%f19, %f19, 0f3F800000;

BB0_7:
	.loc 2 20 1
	add.s32 	%r16, %r14, 1;
	add.s64 	%rd14, %rd14, 4;
	.loc 2 20 22
	add.s32 	%r40, %r40, 1;
	.loc 2 20 1
	setp.lt.s32 	%p9, %r40, %r25;
	mov.u32 	%r39, %r16;
	@%p9 bra 	BB0_5;

BB0_8:
	.loc 2 17 22
	add.s32 	%r38, %r38, 1;
	.loc 2 17 1
	setp.lt.s32 	%p10, %r38, %r24;
	add.s32 	%r37, %r37, %r22;
	@%p10 bra 	BB0_3;

BB0_9:
	.loc 2 29 1
	mad.lo.s32 	%r35, %r4, %r20, %r8;
	cvta.to.global.u64 	%rd11, %rd4;
	.loc 2 29 1
	mul.wide.s32 	%rd12, %r35, 4;
	add.s64 	%rd13, %rd11, %rd12;
	.loc 3 2399 3
	div.rn.f32 	%f18, %f20, %f19;
	.loc 2 29 1
	st.global.f32 	[%rd13], %f18;

BB0_10:
	.loc 2 31 2
	ret;
}


`
	resize_ptx_35 = `
.version 3.1
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry resize(
	.param .u64 resize_param_0,
	.param .u32 resize_param_1,
	.param .u32 resize_param_2,
	.param .u32 resize_param_3,
	.param .u64 resize_param_4,
	.param .u32 resize_param_5,
	.param .u32 resize_param_6,
	.param .u32 resize_param_7,
	.param .u32 resize_param_8,
	.param .u32 resize_param_9,
	.param .u32 resize_param_10
)
{
	.reg .pred 	%p<11>;
	.reg .s32 	%r<40>;
	.reg .f32 	%f<21>;
	.reg .s64 	%rd<15>;


	ld.param.u64 	%rd4, [resize_param_0];
	ld.param.u32 	%r26, [resize_param_2];
	ld.param.u32 	%r20, [resize_param_3];
	ld.param.u64 	%rd5, [resize_param_4];
	ld.param.u32 	%r21, [resize_param_6];
	ld.param.u32 	%r22, [resize_param_7];
	ld.param.u32 	%r23, [resize_param_8];
	ld.param.u32 	%r24, [resize_param_9];
	ld.param.u32 	%r25, [resize_param_10];
	.loc 3 9 1
	mov.u32 	%r1, %ctaid.y;
	mov.u32 	%r2, %ntid.y;
	mov.u32 	%r3, %tid.y;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	.loc 3 10 1
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 3 12 1
	setp.lt.s32 	%p1, %r4, %r26;
	setp.lt.s32 	%p2, %r8, %r20;
	and.pred  	%p3, %p1, %p2;
	@!%p3 bra 	BB2_10;
	bra.uni 	BB2_1;

BB2_1:
	.loc 3 17 1
	setp.lt.s32 	%p4, %r24, 1;
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	.loc 3 17 1
	@%p4 bra 	BB2_9;

	.loc 3 18 1
	mul.lo.s32 	%r9, %r4, %r24;
	.loc 3 17 1
	mul.lo.s32 	%r10, %r25, %r8;
	mul.lo.s32 	%r30, %r24, %r4;
	mad.lo.s32 	%r31, %r23, %r21, %r30;
	mul.lo.s32 	%r36, %r22, %r31;
	mov.f32 	%f20, 0f00000000;
	mov.f32 	%f19, %f20;
	mov.u32 	%r37, 0;
	cvta.to.global.u64 	%rd9, %rd5;

BB2_3:
	.loc 3 20 1
	cvt.s64.s32 	%rd6, %r10;
	.loc 3 17 1
	cvt.s64.s32 	%rd7, %r36;
	.loc 3 18 1
	add.s64 	%rd8, %rd6, %rd7;
	.loc 3 17 1
	shl.b64 	%rd10, %rd8, 2;
	add.s64 	%rd14, %rd9, %rd10;
	.loc 3 20 1
	setp.lt.s32 	%p5, %r25, 1;
	@%p5 bra 	BB2_8;

	mov.u32 	%r39, 0;
	mov.u32 	%r38, %r10;

BB2_5:
	.loc 3 18 1
	mov.u32 	%r14, %r38;
	add.s32 	%r33, %r37, %r9;
	setp.lt.s32 	%p6, %r33, %r21;
	.loc 3 23 1
	setp.lt.s32 	%p7, %r14, %r22;
	and.pred  	%p8, %p6, %p7;
	@!%p8 bra 	BB2_7;
	bra.uni 	BB2_6;

BB2_6:
	.loc 3 24 1
	ld.global.nc.f32 	%f17, [%rd14];
	add.f32 	%f20, %f20, %f17;
	.loc 3 25 1
	add.f32 	%f19, %f19, 0f3F800000;

BB2_7:
	.loc 3 20 1
	add.s32 	%r16, %r14, 1;
	add.s64 	%rd14, %rd14, 4;
	.loc 3 20 22
	add.s32 	%r39, %r39, 1;
	.loc 3 20 1
	setp.lt.s32 	%p9, %r39, %r25;
	mov.u32 	%r38, %r16;
	@%p9 bra 	BB2_5;

BB2_8:
	.loc 3 17 22
	add.s32 	%r37, %r37, 1;
	.loc 3 17 1
	setp.lt.s32 	%p10, %r37, %r24;
	add.s32 	%r36, %r36, %r22;
	@%p10 bra 	BB2_3;

BB2_9:
	.loc 3 29 1
	mad.lo.s32 	%r34, %r4, %r20, %r8;
	cvta.to.global.u64 	%rd11, %rd4;
	.loc 3 29 1
	mul.wide.s32 	%rd12, %r34, 4;
	add.s64 	%rd13, %rd11, %rd12;
	.loc 4 2399 3
	div.rn.f32 	%f18, %f20, %f19;
	.loc 3 29 1
	st.global.f32 	[%rd13], %f18;

BB2_10:
	.loc 3 31 2
	ret;
}


`
)
